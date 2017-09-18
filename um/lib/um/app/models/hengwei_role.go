package models

import (
	"strings"

	"strconv"

	"github.com/runner-mei/orm"
	"github.com/three-plus-three/modules/errors"
	"github.com/three-plus-three/modules/permissions"
)

func GetPermissionsGroupFormRole(db *permissions.DB, roleID int64) ([]PermissionGroup, error) {
	var groups []PermissionGroup
	err := db.PermissionGroups().Where(orm.Cond{
		"EXISTS (SELECT * FROM " + db.PermissionGroupsAndRoles().Name() + " par WHERE par.group_id = " + db.PermissionGroups().Name() + ".id AND par.role_id=?)": roleID,
	}).All(&groups)
	if err != nil {
		return nil, errors.Wrap(err, "获取权限组失败:")
	}
	return groups, nil
}

func CreateGroupAndRole(groupID int64, roleID int64, op string) permissions.PermissionGroupAndRole {
	var groupAndRole permissions.PermissionGroupAndRole
	groupAndRole.GroupID = groupID
	groupAndRole.RoleID = roleID
	operation := strings.Split(strings.Split(op, ":")[1], ",")
	for _, oper := range operation {
		switch oper {
		case permissions.CREATE:
			groupAndRole.CreateOperation = true
		case permissions.DELETE:
			groupAndRole.DeleteOperation = true
		case permissions.UPDATE:
			groupAndRole.UpdateOperation = true
		case permissions.QUERY:
			groupAndRole.QueryOperation = true
		}
	}
	return groupAndRole
}

func CheckRoleForPermissionGroup(db *permissions.DB, roleID int64, groupInDB []PermissionGroup, newIDs []string) error {
	var deleted, updated []int64
	var created []string
	for _, v := range groupInDB {
		found := false
		for _, new := range newIDs {
			newId, _ := strconv.ParseInt(strings.Split(new, ":")[0], 10, 64)
			if newId == v.ID {
				found = true
				newId, _ := strconv.ParseInt(strings.Split(new, ":")[0], 10, 64)
				var groupAndRole permissions.PermissionGroupAndRole
				err := db.PermissionGroupsAndRoles().Where(orm.Cond{"role_id": roleID, "group_id": newId}).One(&groupAndRole)
				if err != nil {
					return errors.Wrap(err, "获取权限组与角色关系出错")
				}
				var newGroupAndRole = CreateGroupAndRole(groupAndRole.GroupID, groupAndRole.RoleID, new)
				err = db.PermissionGroupsAndRoles().Id(groupAndRole.ID).Update(&newGroupAndRole)
				if err != nil {
					return errors.Wrap(err, "更新权限组和角色关系出错")
				}
				break
			}
		}
		if !found {
			deleted = append(deleted, v.ID)
		} else {
			updated = append(updated, v.ID)
		}
	}
	for _, id := range newIDs {
		found := false
		newId, _ := strconv.ParseInt(strings.Split(id, ":")[0], 10, 64)
		for _, r := range groupInDB {
			if r.ID == newId {
				found = true
				break
			}
		}
		if !found {
			created = append(created, id)
		}
	}
	_, err := db.PermissionGroupsAndRoles().Where(orm.Cond{"role_id": roleID}).And(orm.Cond{"group_id IN": deleted}).Delete()
	if err != nil {
		return errors.Wrap(err, "删除权限组与角色的关系")
	}
	for _, new := range created {
		newId, _ := strconv.ParseInt(strings.Split(new, ":")[0], 10, 64)
		var groupAndRole = CreateGroupAndRole(newId, roleID, new)
		_, err = db.PermissionGroupsAndRoles().Insert(&groupAndRole)
		if err != nil {
			return errors.Wrap(err, "添加权限组与角色的关系")
		}
	}
	return nil
}
