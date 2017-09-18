package models

import (
	"github.com/runner-mei/orm"
	"github.com/three-plus-three/modules/errors"
	"github.com/three-plus-three/modules/permissions"
)

func CheckAdminUser(db *permissions.DB) {
	var users []permissions.User
	err := db.Users().Where(orm.Cond{"name": "admin"}).All(&users)
	if err != nil {
		panic(errors.Wrap(err, "find admin user"))
	}

	if len(users) == 0 {
		var user permissions.User
		user.Name = "admin"
		user.Password = "Admin"
		user.Description = "超级用户"
		user.Source = "sys"
		_, err := db.Users().Insert(&user)
		if err != nil {
			panic(errors.Wrap(err, "insert admin user"))
		}
	}
}

func  CheckAdministratorRole(db *permissions.DB) {
	var roles []permissions.Role
	err := db.Roles().Where(orm.Cond{"name": "administrator"}).All(&roles)
	if err != nil {
		panic(errors.Wrap(err, "find administrator role"))
	}
	if len(roles) == 0 {
		var role permissions.Role
		role.Name = "administrator"
		role.Description = "内置管理员角色"
		_, err := db.Roles().Insert(&role)
		if err != nil {
			panic(errors.Wrap(err, "insert administrator role"))
		}
	}
}

func  CheckVisitorRole (db *permissions.DB) {
	var roles []permissions.Role
	err := db.Roles().Where(orm.Cond{"name": "visitor"}).All(&roles)
	if err != nil {
		panic(errors.Wrap(err, "find visitor role"))
	}
	if len(roles) == 0 {
		var role permissions.Role
		role.Name = "visitor"
		role.Description = "内置角色拥有所有查看权限"
		_, err := db.Roles().Insert(&role)
		if err != nil {
			panic(errors.Wrap(err, "insert visitor role"))
		}
	}
}
