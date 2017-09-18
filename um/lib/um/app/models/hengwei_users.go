package models

import (
	"cn/com/hengwei/commons"
	"errors"
	"strings"

	"github.com/revel/revel"
	"github.com/runner-mei/orm"
	"github.com/three-plus-three/modules/permissions"
)

const GET_Permission_And_Group = "select * from hengwei_permissions_and_groups where group_id in (WITH RECURSIVE T (ID, NAME)  AS (" +
	" SELECT ID, name, PARENT_ID, ARRAY[ID] AS PATH, 1 AS DEPTH " +
	" FROM hengwei_permission_groups WHERE id=?" +
	" UNION ALL SELECT  D.ID, D.NAME, D.PARENT_ID, T.PATH || D.ID, T.DEPTH + 1 AS DEPTH " +
	" FROM hengwei_permission_groups D JOIN T ON D.PARENT_ID = T.ID)" +
	" SELECT ID FROM T ORDER BY PATH);"

const GET_Permission_Group = "WITH RECURSIVE T (ID, NAME ,PARENT_ID)  AS (" +
	" SELECT ID, name, PARENT_ID,ARRAY[ID] AS PATH, 1 AS DEPTH " +
	" FROM hengwei_permission_groups WHERE id=?" +
	" UNION ALL SELECT  D.ID, D.NAME, D.PARENT_ID, T.PATH || D.ID, T.DEPTH + 1 AS DEPTH " +
	" FROM hengwei_permission_groups D JOIN T ON D.PARENT_ID = T.ID)" +
	" SELECT ID, name, PARENT_ID FROM T ORDER BY PATH;"

type User struct {
	permissions.User `xorm:"extends"`
	Fields           []ExtendField      `xorm:"-"`
	Roles            []permissions.Role `json:"role" xorm:"-"`
}
type ExtendField struct {
	FieldID string
	Value   string
}

type Field struct {
	ID        string
	Name      string
	Type      string
	IsDefault string
}

type LDAPToUserBind struct {
	Field  string
	Column string
}

func (u User) ToUser() permissions.User {

	var user permissions.User
	user.Name = u.Name
	if u.Password != "" {
		encryptPassword := commons.CopyToBlock(u.Password)
		user.Password = encryptPassword
	}
	user.Description = u.Description
	revel.INFO.Println("--------------------", u.Fields)
	att := map[string]interface{}{}
	for _, field := range u.Fields {
		if field.FieldID == "white_address_list" {
			strs := strings.Split(field.Value, "\n")
			var ips []string
			for _, str := range strs {
				str = strings.Replace(str, "\n", "", -1)
				str = strings.Replace(str, "\r", "", -1)
				str = strings.Replace(str, " ", "", -1)
				if str != "" {
					ips = append(ips, str)
				}
			}
			att[field.FieldID] = ips
		} else {
			att[field.FieldID] = field.Value
		}
	}

	user.Attributes = att
	user.Source = u.Source
	return user
}

type UserGroup struct {
	permissions.UserGroup `xorm:"extends"`
	Users                 []permissions.User `json:"users" xorm:"-"`
}

type Role struct {
	permissions.Role `xorm:"extends"`
	Groups           []PermissionGroup `json:"groups" xorm:"-"`
}

type PermissionGroup struct {
	permissions.PermissionGroup `xorm:"extends"`
	Permissions                 []permissions.Permission `json:"permissionss" xorm:"-"`
	Operation                   string                   `json:"operation" xorm:"-"`
}

func CheckUserForRole(tx *permissions.DB, userID int64, rolesInDB []permissions.Role, newIDs []int64) error {
	var created, deleted, updated []int64
	for _, r := range rolesInDB {
		found := false
		for _, v := range newIDs {
			if r.ID == v {
				found = true
				break
			}
		}

		if found {
			updated = append(updated, r.ID)
		} else {
			deleted = append(deleted, r.ID)
		}
	}
	for _, id := range newIDs {
		found := false
		for _, r := range rolesInDB {
			if r.ID == id {
				found = true
				break
			}
		}

		if !found {
			created = append(created, id)
		}
	}

	_, err := tx.UsersAndRoles().Where(orm.Cond{"user_id": userID}).
		And(orm.Cond{"role_id IN": deleted}).Delete()
	if err != nil {
		return errors.New("Delete UsersAndRoles :" + err.Error())
	}

	for _, id := range created {
		_, err = tx.UsersAndRoles().Insert(permissions.UserAndRole{
			UserID: userID,
			RoleID: id,
		})
		if err != nil {
			return errors.New("Insert UsersAndRoles :" + err.Error())
		}
	}
	return nil
}

func GetRolesFromUser(db *permissions.DB, userID int64) ([]permissions.Role, error) {
	var roles []permissions.Role
	err := db.Roles().Where(orm.Cond{
		"EXISTS (SELECT * FROM " + db.UsersAndRoles().Name() + " uar WHERE uar.role_id = " + db.Roles().Name() + ".id AND uar.user_id=?)": userID,
	}).All(&roles)
	return roles, err
}
