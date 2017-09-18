package controllers

import (
	"um/app/models"
	"um/app/routes"

	"github.com/revel/revel"
	"github.com/runner-mei/orm"
	"github.com/three-plus-three/modules/errors"
	"github.com/three-plus-three/modules/permissions"
	"github.com/three-plus-three/modules/toolbox"
	"github.com/three-plus-three/modules/util"
	"github.com/three-plus-three/modules/web_ext"
)

// HengWeiUsers - 控制器
type HengweiUserGroups struct {
	App
}

// 列出所有记录
func (c HengweiUserGroups) Index(pageIndex int, pageSize int, groupId int64) revel.Result {
	if !c.CurrentUserHasPermission("um.HengweiUserGroups", web_ext.QUERY) {
		return c.RenderError(permissions.ErrUnauthorized)
	}

	// var pageIndex, pageSize int
	// c.Params.Bind(&pageIndex, "pageIndex")
	// c.Params.Bind(&pageSize, "pageSize")

	if pageSize <= 0 {
		pageSize = toolbox.DEFAULT_SIZE_PER_PAGE
	}

	roots, err := models.LoandUserGroupTree(&c.Lifecycle.DB)
	if err != nil {
		return c.RenderError(errors.Wrap(err, "获取用户组树失败"))
	}
	if groupId == 0 && len(roots) > 0 {
		groupId = roots[0].ID
	}
	c.ViewArgs["userGroupNodes"] = roots
	userGroup := models.FindInUserGroups(roots, groupId)
	if userGroup != nil {
		userList, err := models.GetUserViewOfGroup(&c.Lifecycle.DB, userGroup)
		if err != nil {
			return c.RenderError(err)
		}
		group := map[string]interface{}{
			"userGroup": userGroup,
			"paginator": toolbox.NewPaginator(c.Request.Request, pageSize, len(userList)),
		}
		if pageIndex*pageSize < len(userList) {
			if (pageIndex+1)*pageSize > len(userList) {
				group["userList"] = userList[pageIndex*pageSize:]
			} else {
				group["userList"] = userList[pageIndex*pageSize : (pageIndex+1)*pageSize]
			}
		}
		c.ViewArgs["group"] = group
	}
	return c.Render()
}

//编辑新建记录
func (c HengweiUserGroups) New(groupId int64) revel.Result {
	if !c.CurrentUserHasPermission("um.HengweiUserGroups", web_ext.CREATE) {
		return c.RenderError(permissions.ErrUnauthorized)
	}
	//获取用户
	var allHengweiUsers []permissions.User
	err := c.Lifecycle.DB.Users().Where().OrderBy("id").All(&allHengweiUsers)
	if err != nil {
		c.Flash.Error("Get User", err.Error())
		c.FlashParams()
		return c.Redirect(routes.HengweiUserGroups.Index(0, 0, 0))
	}

	//获取用户组
	userGroupNodes, err := models.LoandUserGroupTree(&c.Lifecycle.DB)
	if err != nil {
		c.Flash.Error(err.Error())
		c.FlashParams()
		return c.Redirect(routes.HengweiUserGroups.Index(0, 0, 0))
	}

	c.ViewArgs["userGroupNodes"] = userGroupNodes
	c.ViewArgs["allHengweiUsers"] = allHengweiUsers
	c.ViewArgs["groupId"] = groupId
	return c.Render()
}

//创建记录
func (c HengweiUserGroups) Create(userGroupView models.UserGroupView) revel.Result {
	if !c.CurrentUserHasPermission("um.HengweiUserGroups", web_ext.CREATE) {
		return c.RenderError(permissions.ErrUnauthorized)
	}
	if userGroupView.UserGroup.Validate(c.Validation) {
		c.Validation.Keep()
		c.FlashParams()
		return c.Redirect(routes.HengweiUserGroups.New(userGroupView.ParentID))
	}
	tx, err := c.Lifecycle.DB.Begin()
	if err != nil {
		c.Flash.Error(err.Error())
		c.FlashParams()
		return c.Redirect(routes.HengweiUserGroups.New(userGroupView.ParentID))
	}
	defer util.CloseWith(tx)
	id, err := tx.UserGroups().Nullable("parent_id").Insert(&userGroupView)
	if err != nil {
		if oerr, ok := err.(*orm.Error); ok {
			for _, validation := range oerr.Validations {
				c.Validation.Error(validation.Message).Key(permissions.KeyForUserGroup(validation.Key))
			}
			c.Validation.Keep()
		}
		c.Flash.Error(err.Error())
		c.FlashParams()
		return c.Redirect(routes.HengweiUserGroups.New(userGroupView.ParentID))
	}
	err = models.InsertUserGroupAndUsert(tx, userGroupView, id.(int64))
	if err != nil {
		c.Flash.Error(err.Error())
		c.FlashParams()
		return c.Redirect(routes.HengweiUserGroups.New(userGroupView.ParentID))
	}

	if err = tx.Commit(); err != nil {
		c.Flash.Error(err.Error())
		c.FlashParams()
		return c.Redirect(routes.HengweiUserGroups.New(userGroupView.ParentID))
	}
	c.Flash.Success("创建成功")
	c.FlashParams()
	return c.Redirect(routes.HengweiUserGroups.Index(0, 0, id.(int64)))
}

//编辑指定 id 的记录
func (c HengweiUserGroups) Edit(id int64) revel.Result {
	if !c.CurrentUserHasPermission("um.HengweiUserGroups", web_ext.UPDATE) {
		return c.RenderError(permissions.ErrUnauthorized)
	}

	//获取所有用户
	var allHengweiUsers []permissions.User
	err := c.Lifecycle.DB.Users().Where().OrderBy("id").All(&allHengweiUsers)
	if err != nil {
		c.Flash.Error("获取用户失败" + err.Error())
		c.FlashParams()
		return c.Redirect(routes.HengweiUserGroups.Index(0, 0, id))
	}
	//获取用户组树
	userGroupNodes, err := models.LoandUserGroupTree(&c.Lifecycle.DB)
	if err != nil {
		c.Flash.Error(err.Error())
		c.FlashParams()
		return c.Redirect(routes.HengweiUserGroups.Index(0, 0, 0))
	}
	//获取用户组用户信息
	userGroupView, err := models.GetUserGroupView(&c.Lifecycle.DB, id)
	if err != nil {
		return c.RenderError(err)
	}
	c.ViewArgs["userGroupNodes"] = userGroupNodes
	c.ViewArgs["allHengweiUsers"] = allHengweiUsers
	c.ViewArgs["userGroupView"] = userGroupView
	return c.Render()
}

// 按 id 更新记录
func (c HengweiUserGroups) Update(userGroupView models.UserGroupView) revel.Result {
	if !c.CurrentUserHasPermission("um.HengweiUserGroups", web_ext.UPDATE) {
		return c.RenderError(permissions.ErrUnauthorized)
	}
	if userGroupView.UserGroup.Validate(c.Validation) {
		c.Validation.Keep()
		c.FlashParams()
		return c.Redirect(routes.HengweiUserGroups.Edit(userGroupView.ID))
	}
	tx, err := c.Lifecycle.DB.Begin()
	if err != nil {
		c.Flash.Error(err.Error())
		c.FlashParams()
		return c.Redirect(routes.HengweiUsers.Edit(userGroupView.ID))
	}
	defer util.CloseWith(tx)

	err = tx.UserGroups().Id(userGroupView.ID).Nullable("parent_id").Update(&userGroupView)
	if err != nil {
		if err == orm.ErrNotFound {
			c.Flash.Error(revel.Message(c.Request.Locale, "update.record_not_found"))
		} else {
			if oerr, ok := err.(*orm.Error); ok {
				for _, validation := range oerr.Validations {
					c.Validation.Error(validation.Message).Key(permissions.KeyForUserGroup(validation.Key))
				}
				c.Validation.Keep()
			}
			c.Flash.Error(err.Error())
		}
		c.FlashParams()
		return c.Redirect(routes.HengweiUserGroups.Edit(userGroupView.ID))
	}

	err = models.CheckUserForUserGroup(tx, userGroupView.ID, userGroupView.UserID)
	if err != nil {
		c.Flash.Error(err.Error())
		c.FlashParams()
		return c.Redirect(routes.HengweiUserGroups.Edit(userGroupView.ID))
	}
	if err = tx.Commit(); err != nil {
		c.Flash.Error(err.Error())
		c.FlashParams()
		return c.Redirect(routes.HengweiUserGroups.Edit(userGroupView.ID))
	}
	c.Flash.Success(revel.Message(c.Request.Locale, "update.success"))
	c.FlashParams()
	return c.Redirect(routes.HengweiUserGroups.Index(0, 0, userGroupView.ID))
}

// 按 id 删除记录
func (c HengweiUserGroups) Delete(id string, groupId int64) revel.Result {
	if !c.CurrentUserHasPermission("um.HengweiUserGroups", web_ext.DELETE) {
		return c.RenderError(permissions.ErrUnauthorized)
	}
	err := c.Lifecycle.DB.UserGroups().Id(id).Delete()
	if err != nil {
		c.Flash.Error("删除失败", err.Error())
	} else {
		c.Flash.Success("删除成功")
	}
	c.FlashParams()
	return c.Redirect(routes.HengweiUserGroups.Index(0, 0, groupId))
}

//导入组
func (c HengweiUserGroups) ImportGroup(id string) revel.Result {
	var result = map[string]interface{}{}
	userlist, err := models.GetUserOfGroup(&c.Lifecycle.DB, id)
	if err != nil {
		c.Response.Status = 500
		result["msg"] = err.Error()
		return c.RenderJSON(result)
	}
	if len(userlist) == 0 {
		c.Response.Status = 500
		result["msg"] = "选择组为没有关联用户请重新选择"
		return c.RenderJSON(result)
	}
	result["msg"] = "获取成功"
	result["data"] = userlist
	return c.RenderJSON(result)
}
