// GENERATED CODE - DO NOT EDIT
package routes

import "github.com/revel/revel"


type tApp struct {}
var App tApp



type tStatic struct {}
var Static tStatic


func (_ tStatic) Serve(
		prefix string,
		filepath string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "prefix", prefix)
	revel.Unbind(args, "filepath", filepath)
	return revel.MainRouter.Reverse("Static.Serve", args).URL
}

func (_ tStatic) ServeModule(
		moduleName string,
		prefix string,
		filepath string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "moduleName", moduleName)
	revel.Unbind(args, "prefix", prefix)
	revel.Unbind(args, "filepath", filepath)
	return revel.MainRouter.Reverse("Static.ServeModule", args).URL
}


type tTestRunner struct {}
var TestRunner tTestRunner


func (_ tTestRunner) Index(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("TestRunner.Index", args).URL
}

func (_ tTestRunner) Suite(
		suite string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "suite", suite)
	return revel.MainRouter.Reverse("TestRunner.Suite", args).URL
}

func (_ tTestRunner) Run(
		suite string,
		test string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "suite", suite)
	revel.Unbind(args, "test", test)
	return revel.MainRouter.Reverse("TestRunner.Run", args).URL
}

func (_ tTestRunner) List(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("TestRunner.List", args).URL
}


type tHengweiUserGroups struct {}
var HengweiUserGroups tHengweiUserGroups


func (_ tHengweiUserGroups) Index(
		pageIndex int,
		pageSize int,
		groupId int64,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "pageIndex", pageIndex)
	revel.Unbind(args, "pageSize", pageSize)
	revel.Unbind(args, "groupId", groupId)
	return revel.MainRouter.Reverse("HengweiUserGroups.Index", args).URL
}

func (_ tHengweiUserGroups) New(
		groupId int64,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "groupId", groupId)
	return revel.MainRouter.Reverse("HengweiUserGroups.New", args).URL
}

func (_ tHengweiUserGroups) Create(
		userGroupView interface{},
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "userGroupView", userGroupView)
	return revel.MainRouter.Reverse("HengweiUserGroups.Create", args).URL
}

func (_ tHengweiUserGroups) Edit(
		id int64,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "id", id)
	return revel.MainRouter.Reverse("HengweiUserGroups.Edit", args).URL
}

func (_ tHengweiUserGroups) Update(
		userGroupView interface{},
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "userGroupView", userGroupView)
	return revel.MainRouter.Reverse("HengweiUserGroups.Update", args).URL
}

func (_ tHengweiUserGroups) Delete(
		id string,
		groupId int64,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "id", id)
	revel.Unbind(args, "groupId", groupId)
	return revel.MainRouter.Reverse("HengweiUserGroups.Delete", args).URL
}

func (_ tHengweiUserGroups) ImportGroup(
		id string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "id", id)
	return revel.MainRouter.Reverse("HengweiUserGroups.ImportGroup", args).URL
}


type tHengweiOptions struct {}
var HengweiOptions tHengweiOptions


func (_ tHengweiOptions) SetUserField(
		active string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "active", active)
	return revel.MainRouter.Reverse("HengweiOptions.SetUserField", args).URL
}

func (_ tHengweiOptions) UpdateUserField(
		fields interface{},
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "fields", fields)
	return revel.MainRouter.Reverse("HengweiOptions.UpdateUserField", args).URL
}

func (_ tHengweiOptions) SyncADView(
		active string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "active", active)
	return revel.MainRouter.Reverse("HengweiOptions.SyncADView", args).URL
}

func (_ tHengweiOptions) SyncADRule(
		umFields interface{},
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "umFields", umFields)
	return revel.MainRouter.Reverse("HengweiOptions.SyncADRule", args).URL
}

func (_ tHengweiOptions) GetPermission(
		permission string,
		operation string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "permission", permission)
	revel.Unbind(args, "operation", operation)
	return revel.MainRouter.Reverse("HengweiOptions.GetPermission", args).URL
}


type tHengweiPermissionGroups struct {}
var HengweiPermissionGroups tHengweiPermissionGroups


func (_ tHengweiPermissionGroups) Index(
		pageIndex int,
		pageSize int,
		groupId int64,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "pageIndex", pageIndex)
	revel.Unbind(args, "pageSize", pageSize)
	revel.Unbind(args, "groupId", groupId)
	return revel.MainRouter.Reverse("HengweiPermissionGroups.Index", args).URL
}

func (_ tHengweiPermissionGroups) New(
		groupId int64,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "groupId", groupId)
	return revel.MainRouter.Reverse("HengweiPermissionGroups.New", args).URL
}

func (_ tHengweiPermissionGroups) Create(
		permissionGroupView interface{},
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "permissionGroupView", permissionGroupView)
	return revel.MainRouter.Reverse("HengweiPermissionGroups.Create", args).URL
}

func (_ tHengweiPermissionGroups) Edit(
		id int64,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "id", id)
	return revel.MainRouter.Reverse("HengweiPermissionGroups.Edit", args).URL
}

func (_ tHengweiPermissionGroups) Update(
		permissionGroupView interface{},
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "permissionGroupView", permissionGroupView)
	return revel.MainRouter.Reverse("HengweiPermissionGroups.Update", args).URL
}

func (_ tHengweiPermissionGroups) Delete(
		id int64,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "id", id)
	return revel.MainRouter.Reverse("HengweiPermissionGroups.Delete", args).URL
}

func (_ tHengweiPermissionGroups) PermissionChoices(
		id int64,
		tag string,
		ids string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "id", id)
	revel.Unbind(args, "tag", tag)
	revel.Unbind(args, "ids", ids)
	return revel.MainRouter.Reverse("HengweiPermissionGroups.PermissionChoices", args).URL
}

func (_ tHengweiPermissionGroups) Copy(
		id int64,
		parentId int64,
		name string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "id", id)
	revel.Unbind(args, "parentId", parentId)
	revel.Unbind(args, "name", name)
	return revel.MainRouter.Reverse("HengweiPermissionGroups.Copy", args).URL
}


type tHengweiRoles struct {}
var HengweiRoles tHengweiRoles


func (_ tHengweiRoles) Index(
		pageIndex int,
		pageSize int,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "pageIndex", pageIndex)
	revel.Unbind(args, "pageSize", pageSize)
	return revel.MainRouter.Reverse("HengweiRoles.Index", args).URL
}

func (_ tHengweiRoles) New(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("HengweiRoles.New", args).URL
}

func (_ tHengweiRoles) Create(
		hengweiRole interface{},
		group_id_list []string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "hengweiRole", hengweiRole)
	revel.Unbind(args, "group_id_list", group_id_list)
	return revel.MainRouter.Reverse("HengweiRoles.Create", args).URL
}

func (_ tHengweiRoles) Edit(
		id int64,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "id", id)
	return revel.MainRouter.Reverse("HengweiRoles.Edit", args).URL
}

func (_ tHengweiRoles) Update(
		hengweiRole interface{},
		group_id_list []string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "hengweiRole", hengweiRole)
	revel.Unbind(args, "group_id_list", group_id_list)
	return revel.MainRouter.Reverse("HengweiRoles.Update", args).URL
}

func (_ tHengweiRoles) Delete(
		id int64,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "id", id)
	return revel.MainRouter.Reverse("HengweiRoles.Delete", args).URL
}

func (_ tHengweiRoles) DeleteByIDs(
		id_list []int64,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "id_list", id_list)
	return revel.MainRouter.Reverse("HengweiRoles.DeleteByIDs", args).URL
}


type tHengweiUsers struct {}
var HengweiUsers tHengweiUsers


func (_ tHengweiUsers) Index(
		pageIndex int,
		pageSize int,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "pageIndex", pageIndex)
	revel.Unbind(args, "pageSize", pageSize)
	return revel.MainRouter.Reverse("HengweiUsers.Index", args).URL
}

func (_ tHengweiUsers) New(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("HengweiUsers.New", args).URL
}

func (_ tHengweiUsers) Create(
		hengweiUser interface{},
		role_id_list []int64,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "hengweiUser", hengweiUser)
	revel.Unbind(args, "role_id_list", role_id_list)
	return revel.MainRouter.Reverse("HengweiUsers.Create", args).URL
}

func (_ tHengweiUsers) Edit(
		id int64,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "id", id)
	return revel.MainRouter.Reverse("HengweiUsers.Edit", args).URL
}

func (_ tHengweiUsers) Update(
		hengweiUser interface{},
		role_id_list []int64,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "hengweiUser", hengweiUser)
	revel.Unbind(args, "role_id_list", role_id_list)
	return revel.MainRouter.Reverse("HengweiUsers.Update", args).URL
}

func (_ tHengweiUsers) Delete(
		id int64,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "id", id)
	return revel.MainRouter.Reverse("HengweiUsers.Delete", args).URL
}

func (_ tHengweiUsers) DeleteByIDs(
		user_id_list []int64,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "user_id_list", user_id_list)
	return revel.MainRouter.Reverse("HengweiUsers.DeleteByIDs", args).URL
}

func (_ tHengweiUsers) UserUnlock(
		id int64,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "id", id)
	return revel.MainRouter.Reverse("HengweiUsers.UserUnlock", args).URL
}

func (_ tHengweiUsers) SyncAD(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("HengweiUsers.SyncAD", args).URL
}


