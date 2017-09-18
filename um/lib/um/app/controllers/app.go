package controllers

import (
	"um/app"
	"um/app/libs"

	"github.com/revel/revel"
	"github.com/three-plus-three/modules/web_ext"
)

type App struct {
	*revel.Controller
	Lifecycle *libs.Lifecycle
}

func (c *App) IsAJAX() bool {
	return c.Request.Header.Get("X-Requested-With") == "XMLHttpRequest"
}

func (c *App) CurrentUser() web_ext.User {
	return c.Lifecycle.CurrentUser(c.Controller)
}

func (c *App) initLifecycle() revel.Result {
	c.Lifecycle = app.Lifecycle
	c.ViewArgs["menuList"] = c.Lifecycle.MenuList

	if active, ok := c.Params.Values["active"]; ok {
		c.ViewArgs["active"] = active[0]
	} else {
		c.ViewArgs["active"] = c.Name
	}

	user := c.CurrentUser()
	if user != nil {
		c.ViewArgs["currentUser"] = user
	}
	return nil
}

func (c *App) CurrentUserHasPermission(permissionObject, action string) bool {
	user := c.CurrentUser()
	if user == nil {
		return false
	}
	return user.HasPermission(permissionObject, action)
}

func (c App) checkUser() revel.Result {
	if c.Name == "Permission" {
		return nil
	}
	return c.Lifecycle.CheckUser(c.Controller)
}

func init() {
	revel.InterceptMethod((*App).initLifecycle, revel.BEFORE)
	revel.InterceptMethod(func(c interface{}) revel.Result {
		if check, ok := c.(interface {
			CheckUser() revel.Result
		}); ok {
			return check.CheckUser()
		}
		return nil
	}, revel.BEFORE)

	revel.InterceptMethod((*App).checkUser, revel.BEFORE)
}
