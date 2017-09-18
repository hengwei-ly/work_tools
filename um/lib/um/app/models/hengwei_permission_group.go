package models

import (
	"strings"

	"github.com/runner-mei/orm"
	"github.com/three-plus-three/modules/errors"
	"github.com/three-plus-three/modules/permissions"
)

type Group struct {
	permissions.PermissionGroup `xorm:"extends"` // nolint

	Children []*Group `json:"children,omitempty" xorm:"-"`
}

func (group *Group) GetChild(childID int64) *Group {
	return FindInGroups(group.Children, childID)
}

func FindInGroups(groups []*Group, id int64) *Group {
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

func LoadPermissionGroupTreeFromDB(db *permissions.DB) ([]*Group, error) {
	//查询所有权限组， 然后生成json数据
	var allPermissionsGroups []Group
	err := db.PermissionGroups().Where().OrderBy("id").All(&allPermissionsGroups)
	if err != nil {
		return nil, err
	}

	if len(allPermissionsGroups) == 0 {
		return nil, err
	}

	byID := map[int64]*Group{}
	for idx := range allPermissionsGroups {
		byID[allPermissionsGroups[idx].ID] = &allPermissionsGroups[idx]
	}

	var roots []*Group
	for idx := range allPermissionsGroups {
		group := &allPermissionsGroups[idx]
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

//复制时查找选择的组的树
func OnePermissionGroupTreeFromDB(db *permissions.DB, groupID int64) ([]*Group, error) {
	//查询所有权限组， 然后生成json数据
	var allPermissionsGroups []Group
	err := db.PermissionGroups().Query(GET_Permission_Group, groupID).All(&allPermissionsGroups)
	if err != nil {
		return nil, err
	}

	if len(allPermissionsGroups) == 0 {
		return nil, err
	}

	byID := map[int64]*Group{}
	for idx := range allPermissionsGroups {
		byID[allPermissionsGroups[idx].ID] = &allPermissionsGroups[idx]
	}

	var roots []*Group
	for idx := range allPermissionsGroups {
		group := &allPermissionsGroups[idx]
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

type Permission struct {
	permissions.Permission `xorm:"extends"`
	Type                   int64  `json:"type" xorm:"-"`
	Tag                    string `json:"tag" xorm:"-"`
}

type PermissionView struct {
	Type        int64   `json:"type"`
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description,emitempty"`
	Groups      []Group `json:"groups,emitempty"`
}

func GetPermissionViewOfGroup(db *permissions.DB, group *Group) ([]PermissionView, error) {
	//获取所有组与权限的关系
	var allGroupAndPermissions []permissions.PermissionAndGroup
	err := db.PermissionsAndGroups().Query(GET_Permission_And_Group, group.ID).All(&allGroupAndPermissions)
	if err != nil {
		return nil, errors.Wrap(err, "加载权限失败")
	}

	//获取当前所选的权限组含有的所有权限
	var permissionsByKey = map[string]PermissionView{}
	for _, v := range allGroupAndPermissions {
		var key string

		if v.Type == permissions.PERMISSION_TAG {
			key = "tag:" + v.PermissionObject
		} else {
			key = ":" + v.PermissionObject
		}

		old, ok := permissionsByKey[key]
		if !ok {
			old = PermissionView{
				Type: v.Type,
				ID:   v.PermissionObject,
			}
			if v.Type == permissions.PERMISSION_ID {
				permission, err := permissions.GetPermissionByID(old.ID)
				if err != nil {
					if err == permissions.ErrPermissionNotFound {
						continue
					}
					return nil, errors.Wrap(err, "加载权限失败")
				}
				old.Name = permission.Name
				old.Description = permission.Description
			} else {
				old.Name = v.PermissionObject
				old.Description = v.PermissionObject
			}
		}
		if child := group.GetChild(v.ID); child != nil {
			old.Groups = append(old.Groups, *child)
		}
		if v.GroupID == group.ID {
			old.Groups = append(old.Groups, *group)
		}
		permissionsByKey[key] = old
	}

	var permissionList []PermissionView
	for _, view := range permissionsByKey {
		permissionList = append(permissionList, view)
	}
	return permissionList, nil
}

type PermissionGroupView struct {
	permissions.PermissionGroup `xorm:"extends"`
	Permissions                 []*permissions.Permission `json:"permissions" xorm:"-"`
	Tags                        []string                  `json:"tags" xorm:"-"`
	SelectedID                  []string                  `json:"selected_id" xorm:"-"`
	SelectedTags                string                    `json:"selected_tags" xorm:"-"`
}

func GetPermissionGroupView(db *permissions.DB, groupID int64) (PermissionGroupView, error) {
	//获取当前组
	var permissionGroupView PermissionGroupView
	err := db.PermissionGroups().Id(groupID).Get(&permissionGroupView)
	if err != nil {
		return PermissionGroupView{}, errors.Wrap(err, "获取权限组信息失败")
	}
	//获取组与权限的关系
	var permissionAndGroups []permissions.PermissionAndGroup
	err = db.PermissionsAndGroups().Where(orm.Cond{"group_id": groupID}).All(&permissionAndGroups)
	if err != nil {
		return PermissionGroupView{}, errors.Wrap(err, "获取权限组与权限关系失败")
	}

	for _, v := range permissionAndGroups {
		if v.Type == permissions.PERMISSION_ID {
			permission, err := permissions.GetPermissionByID(v.PermissionObject)
			if err != nil {
				return PermissionGroupView{}, errors.Wrap(err, "获取该组权限失败")
			}
			permissionGroupView.Permissions = append(permissionGroupView.Permissions, permission)
		} else {
			permissionGroupView.Tags = append(permissionGroupView.Tags, v.PermissionObject)
		}
	}
	permissionGroupView.SelectedTags = strings.Join(permissionGroupView.Tags, ",")
	return permissionGroupView, nil
}

type SelectPermissionView struct {
	permissions.Permission `xorm:"extends"`
	IsSelected             bool `json:"is_select"`
	IsSelectedTag          bool `json:"is_selected_tag"`
}

func GetSelectPermissionView(db *permissions.DB, groupID int64, tags []string, permissionIds []string) ([]SelectPermissionView, error) {
	//获取所有组与权限的关系
	var allGroupAndPermissions []permissions.PermissionAndGroup
	err := db.PermissionsAndGroups().Query(GET_Permission_And_Group, groupID).All(&allGroupAndPermissions)
	if err != nil {
		return nil, errors.Wrap(err, "加载权限失败")
	}
	//获取当前所选的权限组含有的所有权限
	var permissionsByKey = map[string]SelectPermissionView{}
	for _, v := range allGroupAndPermissions {
		var key string

		if v.Type == permissions.PERMISSION_TAG {
			key = "tag:" + v.PermissionObject
		} else {
			key = ":" + v.PermissionObject
		}

		old, ok := permissionsByKey[key]
		if !ok {
			old = SelectPermissionView{
				Permission: permissions.Permission{ID: v.PermissionObject},
			}
			if v.Type == permissions.PERMISSION_ID {
				permission, err := permissions.GetPermissionByID(old.ID)
				if err != nil {
					if err == permissions.ErrPermissionNotFound {
						continue
					}
					return nil, errors.Wrap(err, "加载权限失败")
				}
				old.Name = permission.Name
				old.Description = permission.Description
				old.Tags = permission.Tags
				permissionsByKey[key] = old
			}
		}
	}

	var permissionList []SelectPermissionView
	for _, view := range permissionsByKey {
		permissionList = append(permissionList, checkIsSelected(view, tags, permissionIds))
	}
	return permissionList, nil
}

func checkIsSelected(view SelectPermissionView, tags []string, permissionIds []string) SelectPermissionView {
	for _, tag := range tags {
		for _, t := range view.Tags {
			if tag == t {
				view.IsSelectedTag = true
				return view
			}
		}
	}

	for _, id := range permissionIds {
		if view.ID == id {
			view.IsSelected = true
			return view
		}
	}

	return view
}

func InsertPerssionGROUP(db *permissions.DB, group *Group, id int64, name string) error {
	var oldPermissionGroup permissions.PermissionGroup
	err := db.PermissionGroups().Id(group.ID).Get(&oldPermissionGroup)
	if err != nil {
		return errors.Wrap(err, "GetPermissionGroup fail"+group.Name)
	}
	var permssionAndGroups []permissions.PermissionAndGroup
	err = db.PermissionsAndGroups().Where(orm.Cond{"group_id": group.ID}).All(&permssionAndGroups)
	if err != nil {
		return errors.Wrap(err, "GetPermssionAndGroups fail"+group.Name)
	}
	var permissionGroup permissions.PermissionGroup
	if name != "" {
		permissionGroup.Name = name
	} else {
		permissionGroup.Name = oldPermissionGroup.Name
	}
	permissionGroup.IsDefault = false
	if id != 0 {
		permissionGroup.ParentID = id
	}
	permissionGroup.Description = oldPermissionGroup.Description
	grouId, err := db.PermissionGroups().Nullable("parent_id").Insert(&permissionGroup)
	if err != nil {
		return errors.Wrap(err, "InsertPermssionGroups fail"+group.Name)
	}
	for _, pag := range permssionAndGroups {
		var newPAG permissions.PermissionAndGroup
		newPAG.PermissionObject = pag.PermissionObject
		newPAG.Type = pag.Type
		newPAG.GroupID = grouId.(int64)
		_, err := db.PermissionsAndGroups().Insert(&newPAG)
		if err != nil {
			return errors.Wrap(err, "InsertPermissionsAndGroups fail")
		}
	}

	if len(group.Children) != 0 {
		for _, g := range group.Children {
			err = InsertPerssionGROUP(db, g, grouId.(int64), "")
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func InsertPerssionsAndGroup(tx *permissions.DB, permissionGroupView PermissionGroupView, grouId int64) error {
	var tags = permissionGroupView.Tags
	var permissionIDS = permissionGroupView.SelectedID
	if len(tags) != 0 {
		var perssionsandgroup permissions.PermissionAndGroup
		perssionsandgroup.GroupID = grouId
		for _, v := range tags {
			perssionsandgroup.PermissionObject = v
			perssionsandgroup.Type = permissions.PERMISSION_TAG
			_, err := tx.PermissionsAndGroups().Insert(perssionsandgroup)
			if err != nil {
				return errors.Wrap(err, "创建组与权限标签关系失败")
			}
		}
	}
	if len(permissionIDS) != 0 {
		var perssionsandgroup permissions.PermissionAndGroup
		perssionsandgroup.GroupID = grouId
		for _, v := range permissionIDS {
			perssionsandgroup.PermissionObject = v
			perssionsandgroup.Type = permissions.PERMISSION_ID
			_, err := tx.PermissionsAndGroups().Insert(perssionsandgroup)
			if err != nil {
				return errors.Wrap(err, "创建组与权限关系失败")
			}
		}
	}
	return nil
}

func CheckGroupForPermissionGroup(tx *permissions.DB, id int64, new []string, tags []string) error {
	var permissionAndGroups []permissions.PermissionAndGroup
	err := tx.PermissionsAndGroups().Where(orm.Cond{"group_id": id}).All(&permissionAndGroups)
	if err != nil {
		return errors.Wrap(err, "GetPermissionsAndGroups")
	}
	var ids []string
	for _, h := range permissionAndGroups {
		found := false
		if h.Type == permissions.PERMISSION_ID {
			for i, v := range new {
				if h.PermissionObject == v {
					k := i + 1
					new = append(new[:i], new[k:]...)
					found = true
					break
				}
			}
		} else {
			for i, v := range tags {
				if h.PermissionObject == v {
					k := i + 1
					tags = append(tags[:i], tags[k:]...)
					found = true
					break
				}
			}
		}
		if !found {
			ids = append(ids, h.PermissionObject)
		}
	}
	_, err = tx.PermissionsAndGroups().Where(orm.Cond{"group_id": id}).And(orm.Cond{"permission_object IN": ids}).Delete()
	if err != nil {
		return errors.Wrap(err, "DeletePermissionsAndGroups")
	}
	for _, v := range new {
		var permissionAndGroup permissions.PermissionAndGroup
		permissionAndGroup.GroupID = id
		permissionAndGroup.PermissionObject = v
		permissionAndGroup.Type = permissions.PERMISSION_ID
		_, err = tx.PermissionsAndGroups().Insert(&permissionAndGroup)
		if err != nil {
			return errors.Wrap(err, "InsertPermissionsAndGroups")
		}
	}

	for _, v := range tags {
		var permissionAndGroup permissions.PermissionAndGroup
		permissionAndGroup.GroupID = id
		permissionAndGroup.PermissionObject = v
		permissionAndGroup.Type = permissions.PERMISSION_TAG
		_, err = tx.PermissionsAndGroups().Insert(&permissionAndGroup)
		if err != nil {
			return errors.Wrap(err, "InsertPermissionsAndGroups")
		}
	}
	return nil
}
