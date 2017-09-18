// GENERATED CODE - DO NOT EDIT
package main

import (
	"flag"
	"reflect"
	"github.com/revel/revel"
	controllers0 "github.com/revel/modules/static/app/controllers"
	_ "github.com/revel/modules/testrunner/app"
	controllers1 "github.com/revel/modules/testrunner/app/controllers"
	permissions "github.com/three-plus-three/modules/permissions"
	_ "um/app"
	controllers "um/app/controllers"
	models "um/app/models"
	tests "um/tests"
	"github.com/revel/revel/testing"
)

var (
	runMode    *string = flag.String("runMode", "", "Run mode.")
	port       *int    = flag.Int("port", 0, "By default, read from app.conf")
	importPath *string = flag.String("importPath", "", "Go Import Path for the app.")
	srcPath    *string = flag.String("srcPath", "", "Path to the source root.")

	// So compiler won't complain if the generated code doesn't reference reflect package...
	_ = reflect.Invalid
)

func main() {
	flag.Parse()
	revel.Init(*runMode, *importPath, *srcPath)
	revel.INFO.Println("Running revel server")
	
	revel.RegisterController((*controllers.App)(nil),
		[]*revel.MethodType{
			
		})
	
	revel.RegisterController((*controllers0.Static)(nil),
		[]*revel.MethodType{
			&revel.MethodType{
				Name: "Serve",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "prefix", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "filepath", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "ServeModule",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "moduleName", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "prefix", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "filepath", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			
		})
	
	revel.RegisterController((*controllers1.TestRunner)(nil),
		[]*revel.MethodType{
			&revel.MethodType{
				Name: "Index",
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
					76: []string{ 
						"testSuites",
					},
				},
			},
			&revel.MethodType{
				Name: "Suite",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "suite", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "Run",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "suite", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "test", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
					129: []string{ 
					},
				},
			},
			&revel.MethodType{
				Name: "List",
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			
		})
	
	revel.RegisterController((*controllers.HengweiUserGroups)(nil),
		[]*revel.MethodType{
			&revel.MethodType{
				Name: "Index",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "pageIndex", Type: reflect.TypeOf((*int)(nil)) },
					&revel.MethodArg{Name: "pageSize", Type: reflect.TypeOf((*int)(nil)) },
					&revel.MethodArg{Name: "groupId", Type: reflect.TypeOf((*int64)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
					62: []string{ 
					},
				},
			},
			&revel.MethodType{
				Name: "New",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "groupId", Type: reflect.TypeOf((*int64)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
					90: []string{ 
					},
				},
			},
			&revel.MethodType{
				Name: "Create",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "userGroupView", Type: reflect.TypeOf((*models.UserGroupView)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "Edit",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "id", Type: reflect.TypeOf((*int64)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
					168: []string{ 
					},
				},
			},
			&revel.MethodType{
				Name: "Update",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "userGroupView", Type: reflect.TypeOf((*models.UserGroupView)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "Delete",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "id", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "groupId", Type: reflect.TypeOf((*int64)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "ImportGroup",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "id", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			
		})
	
	revel.RegisterController((*controllers.HengweiOptions)(nil),
		[]*revel.MethodType{
			&revel.MethodType{
				Name: "SetUserField",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "active", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
					35: []string{ 
					},
				},
			},
			&revel.MethodType{
				Name: "UpdateUserField",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "fields", Type: reflect.TypeOf((*[]models.Field)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "SyncADView",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "active", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
					80: []string{ 
					},
				},
			},
			&revel.MethodType{
				Name: "SyncADRule",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "umFields", Type: reflect.TypeOf((*[]models.LDAPToUserBind)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "GetPermission",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "permission", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "operation", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			
		})
	
	revel.RegisterController((*controllers.HengweiPermissionGroups)(nil),
		[]*revel.MethodType{
			&revel.MethodType{
				Name: "Index",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "pageIndex", Type: reflect.TypeOf((*int)(nil)) },
					&revel.MethodArg{Name: "pageSize", Type: reflect.TypeOf((*int)(nil)) },
					&revel.MethodArg{Name: "groupId", Type: reflect.TypeOf((*int64)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
					74: []string{ 
					},
				},
			},
			&revel.MethodType{
				Name: "New",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "groupId", Type: reflect.TypeOf((*int64)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
					97: []string{ 
					},
				},
			},
			&revel.MethodType{
				Name: "Create",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "permissionGroupView", Type: reflect.TypeOf((*models.PermissionGroupView)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "Edit",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "id", Type: reflect.TypeOf((*int64)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
					159: []string{ 
					},
					175: []string{ 
					},
				},
			},
			&revel.MethodType{
				Name: "Update",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "permissionGroupView", Type: reflect.TypeOf((*models.PermissionGroupView)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "Delete",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "id", Type: reflect.TypeOf((*int64)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "PermissionChoices",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "id", Type: reflect.TypeOf((*int64)(nil)) },
					&revel.MethodArg{Name: "tag", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "ids", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
					263: []string{ 
					},
				},
			},
			&revel.MethodType{
				Name: "Copy",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "id", Type: reflect.TypeOf((*int64)(nil)) },
					&revel.MethodArg{Name: "parentId", Type: reflect.TypeOf((*int64)(nil)) },
					&revel.MethodArg{Name: "name", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			
		})
	
	revel.RegisterController((*controllers.HengweiRoles)(nil),
		[]*revel.MethodType{
			&revel.MethodType{
				Name: "Index",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "pageIndex", Type: reflect.TypeOf((*int)(nil)) },
					&revel.MethodArg{Name: "pageSize", Type: reflect.TypeOf((*int)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
					65: []string{ 
						"hengweiRoles",
						"paginator",
					},
				},
			},
			&revel.MethodType{
				Name: "New",
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
					80: []string{ 
					},
				},
			},
			&revel.MethodType{
				Name: "Create",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "hengweiRole", Type: reflect.TypeOf((*permissions.Role)(nil)) },
					&revel.MethodArg{Name: "group_id_list", Type: reflect.TypeOf((*[]string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "Edit",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "id", Type: reflect.TypeOf((*int64)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
					204: []string{ 
					},
				},
			},
			&revel.MethodType{
				Name: "Update",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "hengweiRole", Type: reflect.TypeOf((*permissions.Role)(nil)) },
					&revel.MethodArg{Name: "group_id_list", Type: reflect.TypeOf((*[]string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "Delete",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "id", Type: reflect.TypeOf((*int64)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "DeleteByIDs",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "id_list", Type: reflect.TypeOf((*[]int64)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			
		})
	
	revel.RegisterController((*controllers.HengweiUsers)(nil),
		[]*revel.MethodType{
			&revel.MethodType{
				Name: "Index",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "pageIndex", Type: reflect.TypeOf((*int)(nil)) },
					&revel.MethodArg{Name: "pageSize", Type: reflect.TypeOf((*int)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
					76: []string{ 
						"hengweiUsers",
						"paginator",
						"fields",
						"onlineUsers",
					},
				},
			},
			&revel.MethodType{
				Name: "New",
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
					90: []string{ 
						"fields",
					},
				},
			},
			&revel.MethodType{
				Name: "Create",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "hengweiUser", Type: reflect.TypeOf((*models.User)(nil)) },
					&revel.MethodArg{Name: "role_id_list", Type: reflect.TypeOf((*[]int64)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "Edit",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "id", Type: reflect.TypeOf((*int64)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
					184: []string{ 
						"hengweiUser",
						"fields",
					},
				},
			},
			&revel.MethodType{
				Name: "Update",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "hengweiUser", Type: reflect.TypeOf((*models.User)(nil)) },
					&revel.MethodArg{Name: "role_id_list", Type: reflect.TypeOf((*[]int64)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "Delete",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "id", Type: reflect.TypeOf((*int64)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "DeleteByIDs",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "user_id_list", Type: reflect.TypeOf((*[]int64)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "UserUnlock",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "id", Type: reflect.TypeOf((*int64)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "SyncAD",
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			
		})
	
	revel.DefaultValidationKeys = map[string]map[int]string{ 
		"um/app/controllers.HengweiPermissionGroups.Create": { 
			123: "validation.Message",
		},
		"um/app/controllers.HengweiPermissionGroups.Update": { 
			204: "validation.Message",
		},
		"um/app/controllers.HengweiRoles.Create": { 
			106: "validation.Message",
		},
		"um/app/controllers.HengweiRoles.Update": { 
			249: "validation.Message",
		},
		"um/app/controllers.HengweiUserGroups.Create": { 
			114: "validation.Message",
		},
		"um/app/controllers.HengweiUserGroups.Update": { 
			196: "validation.Message",
		},
		"um/app/controllers.HengweiUsers.Create": { 
			119: "validation.Message",
		},
		"um/app/controllers.HengweiUsers.Update": { 
			229: "validation.Message",
		},
	}
	testing.TestSuites = []interface{}{ 
		(*tests.AppTest)(nil),
		(*tests.BaseTest)(nil),
		(*tests.HengweiPermissionsTest)(nil),
		(*tests.PermissionsGroupsTest)(nil),
		(*tests.HengweiRoleTest)(nil),
		(*tests.HengweiUserTest)(nil),
	}

	revel.Run(*port)
}
