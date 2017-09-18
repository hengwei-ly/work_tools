package app

import (
	"um/app/libs"
	"um/app/routes"

	"um/app/models"

	"github.com/pkg/errors"
	_ "github.com/three-plus-three/modules/bind"
	"github.com/three-plus-three/modules/environment"
	"github.com/three-plus-three/modules/permissions"
	"github.com/three-plus-three/modules/toolbox"
	"github.com/three-plus-three/modules/web_ext"
)

var Lifecycle *libs.Lifecycle

var umPermissions = []permissions.Permission{permissions.Permission{ID: "um.HengweiUsers", Name: "用户管理权限", Description: "用户", Tags: []string{"um"}},
	permissions.Permission{ID: "um.HengweiUserGroups", Name: "用户组管理权限", Description: "用户组", Tags: []string{"um"}},
	permissions.Permission{ID: "um.HengweiRoles", Name: "角色管理权限", Description: "角色", Tags: []string{"um"}},
	permissions.Permission{ID: "um.HengweiPermissionGroups", Name: "权限组管理权限", Description: "权限组", Tags: []string{"um"}},
	permissions.Permission{ID: "um_set_field", Name: "设置用户字段权限", Description: "用户字段权限", Tags: []string{"um"}},
	permissions.Permission{ID: "um_sync_AD", Name: "设置活动目录同步规则权限", Description: "用户活动目录同步规则权限", Tags: []string{"um"}},
}

var umGroups = []permissions.Group{permissions.Group{Name: "用户权限管理模块", Description: "用户权限管理模块",
	Children: []permissions.Group{permissions.Group{Name: "用户权限管理模块", Description: "用户权限管理模块", PermissionIDs: []string{"um.HengweiUsers", "um.HengweiUserGroups"}},
		permissions.Group{Name: "角色管理组", Description: "角色管理组", PermissionIDs: []string{"um.HengweiRoles"}},
		permissions.Group{Name: "权限组", Description: "权限组", PermissionIDs: []string{"um.HengweiPermissionGroups"}},
		permissions.Group{Name: "设置权限组", Description: "设置权限组", PermissionIDs: []string{"um_set_field", "um_sync_AD"}},
	}}}

func init() {

	web_ext.InitUser = permissions.InitUser

	web_ext.Init(environment.ENV_UM_PROXY_ID, "用户管理",
		func(data *web_ext.Lifecycle) error {
			data.Variables["footer_copyright_text"] = "用户管理"

			Lifecycle = &libs.Lifecycle{
				Lifecycle: data,
				DB:        permissions.DB{Engine: data.ModelEngine},
				DataDB:    permissions.DB{Engine: data.DataEngine},
			}

			//if err := permissions.DropTables(data.ModelEngine); err != nil {
			//	return err
			//}

			//if err := permissions.InitTables(data.ModelEngine); err != nil {
			//	return err
			//}

			initTemplateFuncs(data.Env)

			err := permissions.LoadHTTP(data.Env.Fs.FromLib("permissions", "http"),
				map[string]interface{}{
					"httpAddress":        data.Env.GetServiceConfig(environment.ENV_WSERVER_PROXY_ID).UrlFor(),
					"urlPrefix":          data.URLPrefix,
					"urlRoot":            data.URLRoot,
					"applicationRoot":    data.ApplicationRoot,
					"applicationContext": data.ApplicationContext,
				})
			if err != nil {
				return errors.New("load permission config fail: " + err.Error())
			}

			permissions.RegisterPermissions("um_bultin",
				permissions.PermissionProviderFunc(func() (*permissions.PermissionData, error) {
					return &permissions.PermissionData{
						Permissions: umPermissions,
						Groups:      umGroups,
					}, nil
				}))

			err = permissions.SaveDefaultPermissionGroups(&Lifecycle.DB)
			if err != nil {
				return errors.Wrap(err, "SaveDefaultPermissionGroups")
			}

			models.CheckAdminUser(&Lifecycle.DB)
			models.CheckAdministratorRole(&Lifecycle.DB)
			models.CheckVisitorRole(&Lifecycle.DB)
			return nil
		},

		func(data *web_ext.Lifecycle) ([]toolbox.Menu, error) {
			return []toolbox.Menu{
				{Title: "用户管理", Icon: "fa-user", Name: "HengweiUser", Permission: "", URL: routes.HengweiUsers.Index(0, 0)},
				{Title: "用户组管理", Icon: "fa-users", Name: "HengweiUserGroup", Permission: "", URL: routes.HengweiUserGroups.Index(0, 0, 0)},
				{Title: "角色管理", Icon: "fa-user-circle", Name: "HengweiRole", Permission: "", URL: routes.HengweiRoles.Index(0, 0)},
				{Title: "权限管理", Icon: "fa-archive", Name: "HengweiPermissionGroup", Permission: "", URL: routes.HengweiPermissionGroups.Index(0, 0, 0)},
				{Title: "设置", Icon: "fa-cogs", URL: "#", Name: "HengweiOptions", Permission: "", Children: []toolbox.Menu{
					{Title: "设置用户字段", Name: "setUserField", Permission: "", URL: routes.HengweiOptions.SetUserField("setUserField")},
					{Title: "设置活动目录同步规则", Name: "syncAD", Permission: "", URL: routes.HengweiOptions.SyncADView("syncAD")},
				}},
			}, nil
		})
}
