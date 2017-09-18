package models

import (
	"strings"

	"github.com/runner-mei/orm"
	"github.com/three-plus-three/modules/errors"
	"github.com/three-plus-three/modules/permissions"
)

const GET_USER_And_USER_Group = "select * from hengwei_users_and_user_groups where group_id in (WITH RECURSIVE T (ID)  AS (SELECT ID, name, PARENT_ID, ARRAY[ID] AS PATH, 1 AS DEPTH " +
	" FROM hengwei_user_groups WHERE id=? " +
	" UNION ALL SELECT  D.ID, D.NAME, D.PARENT_ID, T.PATH || D.ID, T.DEPTH + 1 AS DEPTH " +
	" FROM hengwei_user_groups D JOIN T ON D.PARENT_ID = T.ID) " +
	" SELECT ID FROM T ORDER BY PATH)"

type UserGroupForUserview struct {
	permissions.User `xorm:"extends"`
	userGroups       []permissions.UserGroup `json:"groups,emitempty"`
}

type UserGroupTree struct {
	permissions.UserGroup `xorm:"extends"`
	Children              []*UserGroupTree `json:"user_groups" xorm:"-"`
}

type userView struct {
	permissions.User `xorm:"extends"`
	Groups           []UserGroupTree `json:"groups,emitempty"`
}

type UserGroupView struct {
	permissions.UserGroup `xorm:"extends"`
	Users                 []permissions.User `json:"user" xorm:"-"`
	UserID                []int64            `json:"user_id" xorm:"-"`
}

func (group *UserGroupTree) GetChild(childID int64) *UserGroupTree {
	return FindInUserGroups(group.Children, childID)
}

func FindInUserGroups(groups []*UserGroupTree, id int64) *UserGroupTree {
	for idx := range groups {
		if groups[idx].ID == id {
			return groups[idx]
		}
	}

	for idx := range groups {
		if child := groups[idx].GetChild(id); child != nil {
			return child
		}
	}
	return nil
}

func LoandUserGroupTree(db *permissions.DB) ([]*UserGroupTree, error) {
	var allUserGroups []UserGroupTree
	err := db.UserGroups().Where().All(&allUserGroups)
	if err != nil {
		return nil, errors.Wrap(err, "获取用户组失败")
	}
	if len(allUserGroups) == 0 {
		return nil, err
	}
	byID := map[int64]*UserGroupTree{}
	for idx := range allUserGroups {
		byID[allUserGroups[idx].ID] = &allUserGroups[idx]
	}

	var roots []*UserGroupTree
	for idx := range allUserGroups {
		group := &allUserGroups[idx]
		if group.ParentID == 0 {
			roots = append(roots, group)
			continue
		}
		parent := byID[group.ParentID]
		if parent != nil {
			parent.Children = append(parent.Children, group)
		} else {
			roots = append(roots, group)
		}
	}
	return roots, nil
}

func GetUserViewOfGroup(db *permissions.DB, group *UserGroupTree) ([]userView, error) {
	//获取所有组与权限的关系
	var allGroupAndUser []permissions.UserAndUserGroup
	err := db.UsersAndUserGroups().Query(GET_USER_And_USER_Group, group.ID).All(&allGroupAndUser)
	if err != nil {
		return nil, errors.Wrap(err, "加载权限失败")
	}

	//获取当前所选的权限组含有的所有权限
	var userByKey = map[int64]userView{}
	for _, v := range allGroupAndUser {
		key := v.UserID
		old, ok := userByKey[key]
		if !ok {
			old = userView{User: permissions.User{ID: v.UserID}}
			user, err := getUserByID(db, v.UserID)
			if err != nil {
				return []userView{}, err
			}
			if user.Name == "" {
				return []userView{}, errors.New("加载用户失败")
			}
			old.Name = user.Name
			old.Description = user.Description
		}
		if v.GroupID == group.ID {
			old.Groups = append(old.Groups, *group)
		}
		if child := group.GetChild(v.GroupID); child != nil {
			old.Groups = append(old.Groups, *child)
		}
		userByKey[key] = old
	}
	var permissionList []userView
	for _, view := range userByKey {
		permissionList = append(permissionList, view)
	}
	return permissionList, nil
}

func getUserByID(db *permissions.DB, ID int64) (permissions.User, error) {
	var allUsers []permissions.User
	err := db.Users().Where().All(&allUsers)
	if err != nil {
		return permissions.User{}, errors.Wrap(err, "加载用户失败")
	}

	for _, user := range allUsers {
		if user.ID == ID {
			return user, nil
		}
	}
	return permissions.User{}, nil
}

func GetUserGroupView(db *permissions.DB, id int64) (UserGroupView, error) {
	var userGroupView UserGroupView
	err := db.UserGroups().Where(orm.Cond{"id": id}).One(&userGroupView)
	if err != nil {
		return UserGroupView{}, errors.Wrap(err, "获取用户组失败")
	}

	var userAndUserGroups []permissions.UserAndUserGroup
	err = db.UsersAndUserGroups().Where(orm.Cond{"group_id": id}).All(&userAndUserGroups)
	if err != nil {
		return UserGroupView{}, errors.Wrap(err, "获取用户组与用户关系失败")
	}
	for _, userAndUserGroup := range userAndUserGroups {
		user, err := getUserByID(db, userAndUserGroup.UserID)
		if err != nil {
			return UserGroupView{}, err
		}
		if user.Name != "" {
			userGroupView.Users = append(userGroupView.Users, user)
		}
	}
	return userGroupView, nil
}

func InsertUserGroupAndUsert(tx *permissions.DB, userGroupView UserGroupView, groudId int64) error {
	for _, id := range userGroupView.UserID {
		var userAndUserGroup permissions.UserAndUserGroup
		userAndUserGroup.GroupID = groudId
		userAndUserGroup.UserID = id
		_, err := tx.UsersAndUserGroups().Insert(&userAndUserGroup)
		if err != nil {
			return err
		}
	}
	return nil
}

func CheckUserForUserGroup(tx *permissions.DB, groupId int64, userIds []int64) error {
	var hengweiUserAndUserGroups []permissions.UserAndUserGroup
	err := tx.UsersAndUserGroups().Where(orm.Cond{"group_id": groupId}).All(&hengweiUserAndUserGroups)
	if err != nil {
		return errors.New("获取用户和用户组失败" + err.Error())
	}

	var created, deleted, updated []int64
	for _, h := range hengweiUserAndUserGroups {
		found := false
		for i, id := range userIds {
			if h.UserID == id {
				k := i + 1
				userIds = append(userIds[:i], userIds[k:]...)
				found = true
				break
			}
		}
		if !found {
			deleted = append(deleted, h.UserID)
		} else {
			updated = append(updated, h.UserID)
		}
	}

	for _, id := range userIds {
		found := false
		for _, v := range updated {
			if id == v {
				found = true
				break
			}
		}
		if !found {
			created = append(created, id)
		}
	}

	_, err = tx.UsersAndUserGroups().Where(orm.Cond{"group_id": groupId}).And(orm.Cond{"user_id IN": deleted}).Delete()
	if err != nil {
		return err
	}

	for _, v := range created {
		var userAndGroup permissions.UserAndUserGroup
		userAndGroup.GroupID = groupId
		userAndGroup.UserID = v
		_, err = tx.UsersAndUserGroups().Insert(&userAndGroup)
		if err != nil {
			return err
		}
	}
	return nil
}

func GetUserOfGroup(db *permissions.DB, id string) ([]permissions.User, error) {
	ids := strings.Split(id, ",")
	var userAndGroup []permissions.UserAndUserGroup
	err := db.UsersAndUserGroups().Where(orm.Cond{"group_id IN": ids}).All(&userAndGroup)
	if err != nil {
		return []permissions.User{}, errors.Wrap(err, "获取用户与用户组信息失败")
	}
	var allUsers []permissions.User
	err = db.Users().Where().All(&allUsers)
	if err != nil {
		return []permissions.User{}, errors.Wrap(err, "获取用户失败 稍后再试 GetUsers")
	}
	var userMap = map[int64]permissions.User{}
	for _, u := range userAndGroup {
		_, ok := userMap[u.UserID]
		if !ok {
			for _, alluser := range allUsers {
				if alluser.ID == u.UserID {
					userMap[u.UserID] = alluser
				}
			}
		}
	}
	var users []permissions.User
	for _, v := range userMap {
		users = append(users, v)
	}
	return users, nil

}
