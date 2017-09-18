package controllers

import (
	"strings"
	"um/app/libs/choices"
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

// HengweiPermissionGroups - 控制器
type HengweiPermissionGroups struct {
	App
}

// Index 列出所有记录
func (c HengweiPermissionGroups) Index(pageIndex int, pageSize int, groupId int64) revel.Result {
	if !c.CurrentUserHasPermission("um.HengweiPermissionGroups", web_ext.QUERY) {
		return c.RenderError(permissions.ErrUnauthorized)
	}

	//var pageIndex, pageSize int
	//c.Params.Bind(&pageIndex, "pageIndex")
	//c.Params.Bind(&pageSize, "pageSize")

	if pageSize <= 0 {
		pageSize = toolbox.DEFAULT_SIZE_PER_PAGE
	}

	//查询所有权限组， 然后生成树
	roots, err := models.LoadPermissionGroupTreeFromDB(&c.Lifecycle.DB)
	if err != nil {
		return c.RenderError(err)
	}
	tags, err := choices.Tags(false)
	if err != nil {
		return c.RenderError(errors.Wrap(err, "获取权限标签失败"))
	}
	c.ViewArgs["tags"] = tags
	c.ViewArgs["permissionGroupNodes"] = roots

	if groupId == 0 && len(roots) > 0 {
		groupId = roots[0].ID
	}

	permissionGroup := models.FindInGroups(roots, groupId)
	if permissionGroup != nil {
		permissionList, err := models.GetPermissionViewOfGroup(&c.Lifecycle.DB, permissionGroup)
		if err != nil {
			return c.RenderError(err)
		}

		group := map[string]interface{}{
			"permissionGroup": permissionGroup,
		}
		if pageIndex*pageSize < len(permissionList) {
			if (pageIndex+1)*pageSize > len(permissionList) {
				group["permissionList"] = permissionList[pageIndex*pageSize:]
			} else {
				group["permissionList"] = permissionList[pageIndex*pageSize : (pageIndex+1)*pageSize]
			}
		}
		c.ViewArgs["group"] = group
		c.ViewArgs["paginator"] = toolbox.NewPaginator(c.Request.Request, pageSize, len(permissionList))
	}

	return c.Render()
}

// 编辑新建记录
func (c HengweiPermissionGroups) New(groupId int64) revel.Result {
	if !c.CurrentUserHasPermission("um.HengweiPermissionGroups", web_ext.CREATE) {
		return c.RenderError(permissions.ErrUnauthorized)
	}

	//查询所有权限组， 然后生成树
	roots, err := models.LoadPermissionGroupTreeFromDB(&c.Lifecycle.DB)
	if err != nil {
		return c.RenderError(err)
	}

	//获取标签
	tags, err := choices.Tags(true)
	if err != nil {
		return c.RenderError(errors.Wrap(err, "获取权限标签失败"))
	}
	c.ViewArgs["permissionGroupNodes"] = roots
	c.ViewArgs["tags"] = tags
	c.ViewArgs["groupId"] = groupId
	return c.Render()
}

//创建记录
func (c HengweiPermissionGroups) Create(permissionGroupView models.PermissionGroupView) revel.Result {
	if !c.CurrentUserHasPermission("um.HengweiPermissionGroups", web_ext.CREATE) {
		return c.RenderError(permissions.ErrUnauthorized)
	}
	if permissionGroupView.PermissionGroup.Validate(c.Validation) {
		c.Validation.Keep()
		c.FlashParams()
		return c.Redirect(routes.HengweiPermissionGroups.New(permissionGroupView.ParentID))
	}
	tx, err := c.Lifecycle.DB.Begin()
	if err != nil {
		c.Flash.Error(err.Error())
		c.FlashParams()
		return c.Redirect(routes.HengweiPermissionGroups.New(permissionGroupView.ParentID))
	}
	defer util.CloseWith(tx)

	//添加组
	id, err := tx.PermissionGroups().Nullable("parent_id").Insert(&permissionGroupView)
	if err != nil {
		if oerr, ok := err.(*orm.Error); ok {
			for _, validation := range oerr.Validations {
				c.Validation.Error(validation.Message).Key(permissions.KeyForPermissionsGroups(validation.Key))
			}
			c.Validation.Keep()
		}
		c.Flash.Error(err.Error())
		c.FlashParams()
		return c.Redirect(routes.HengweiPermissionGroups.New(permissionGroupView.ParentID))
	}
	var groupId = id.(int64)
	//添加组和权限的关系
	err = models.InsertPerssionsAndGroup(tx, permissionGroupView, groupId)
	if err != nil {
		c.Flash.Error(err.Error())
		c.FlashParams()
		return c.Redirect(routes.HengweiPermissionGroups.New(permissionGroupView.ParentID))
	}

	if err = tx.Commit(); err != nil {
		c.Flash.Error(err.Error())
		c.FlashParams()
		return c.Redirect(routes.HengweiPermissionGroups.New(permissionGroupView.ParentID))
	}
	c.Flash.Success(revel.Message(c.Request.Locale, "insert.success"))
	return c.Redirect(routes.HengweiPermissionGroups.Index(0, 0, 0))
}

// 编辑指定 id 的记录
func (c HengweiPermissionGroups) Edit(id int64) revel.Result {
	if !c.CurrentUserHasPermission("um.HengweiPermissionGroups", web_ext.UPDATE) {
		return c.RenderError(permissions.ErrUnauthorized)
	}

	permissionGroupView, err := models.GetPermissionGroupView(&c.Lifecycle.DB, id)
	if err != nil {
		c.Flash.Error(err.Error())
		c.FlashParams()
		return c.Render(routes.HengweiPermissionGroups.Index(0, 0, 0))
	}

	roots, err := models.LoadPermissionGroupTreeFromDB(&c.Lifecycle.DB)
	if err != nil {
		return c.RenderError(err)
	}

	//获取标签
	tags, err := choices.Tags(true)
	if err != nil {
		return c.RenderError(errors.Wrap(err, "获取权限标签失败"))
	}
	c.ViewArgs["permissionGroupView"] = permissionGroupView
	c.ViewArgs["permissionGroupNodes"] = roots
	c.ViewArgs["tags"] = tags
	return c.Render()
}

// 按 id 更新记录
func (c HengweiPermissionGroups) Update(permissionGroupView models.PermissionGroupView) revel.Result {
	if !c.CurrentUserHasPermission("um.HengweiPermissionGroups", web_ext.UPDATE) {
		return c.RenderError(permissions.ErrUnauthorized)
	}
	if permissionGroupView.PermissionGroup.Validate(c.Validation) {
		c.Validation.Keep()
		c.FlashParams()
		return c.Redirect(routes.HengweiPermissionGroups.Edit(permissionGroupView.ID))
	}
	tx, err := c.Lifecycle.DB.Begin()
	if err != nil {
		c.Flash.Error(err.Error())
		c.FlashParams()
		return c.Redirect(routes.HengweiPermissionGroups.Edit(permissionGroupView.ID))
	}

	defer util.CloseWith(tx)

	err = tx.PermissionGroups().Id(permissionGroupView.ID).Nullable("parent_id").Update(&permissionGroupView)
	if err != nil {
		if err == orm.ErrNotFound {
			c.Flash.Error(revel.Message(c.Request.Locale, "update.record_not_found"))
		} else {
			if oerr, ok := err.(*orm.Error); ok {
				for _, validation := range oerr.Validations {
					c.Validation.Error(validation.Message).Key(permissions.KeyForPermissionsGroups(validation.Key))
				}
				c.Validation.Keep()
			}
			c.Flash.Error(err.Error())
		}
		c.FlashParams()
		return c.Redirect(routes.HengweiPermissionGroups.Edit(permissionGroupView.ID))
	}
	err = models.CheckGroupForPermissionGroup(tx, permissionGroupView.ID, permissionGroupView.SelectedID, permissionGroupView.Tags)
	if err != nil {
		c.Flash.Error(err.Error())
		c.FlashParams()
		return c.Redirect(routes.HengweiPermissionGroups.Edit(permissionGroupView.ID))
	}
	if err = tx.Commit(); err != nil {
		c.Flash.Error(err.Error())
		c.FlashParams()
		return c.Redirect(routes.HengweiPermissionGroups.Edit(permissionGroupView.ID))
	}
	c.Flash.Success(revel.Message(c.Request.Locale, "update.success"))
	c.FlashParams()
	return c.Redirect(routes.HengweiPermissionGroups.Index(0, 0, permissionGroupView.ID))
}

// Delete 按 id 删除记录
func (c HengweiPermissionGroups) Delete(id int64) revel.Result {
	if !c.CurrentUserHasPermission("um.HengweiPermissionGroups", web_ext.DELETE) {
		return c.RenderError(permissions.ErrUnauthorized)
	}
	revel.INFO.Println("--------", id)
	err := c.Lifecycle.DB.PermissionGroups().Id(id).Delete()
	if err != nil {
		c.Flash.Error("删除失败", err)
	} else {
		c.Flash.Success("删除成功")
	}
	c.FlashParams()
	return c.Redirect(routes.HengweiPermissionGroups.Index(0, 0, 0))
}

func (c HengweiPermissionGroups) PermissionChoices(id int64, tag string, ids string) revel.Result {
	var result = map[string]interface{}{}
	var tags []string
	var permissionIds []string
	if tag != "" {
		tags = strings.Split(tag, ",")
	}
	if ids != "" {
		permissionIds = strings.Split(ids, ",")
	}
	permissionView, err := models.GetSelectPermissionView(&c.Lifecycle.DB, id, tags, permissionIds)
	if err != nil {
		c.Response.Status = 500
		result["status"] = 0
		result["msg"] = "获取权限失败" + err.Error()
		return c.RenderJSON(result)
	}
	c.ViewArgs["permissions"] = permissionView
	return c.Render()
}

func (c HengweiPermissionGroups) Copy(id int64, parentId int64, name string) revel.Result {
	if !c.CurrentUserHasPermission("um.HengweiPermissionGroups", web_ext.CREATE) {
		return c.RenderError(permissions.ErrUnauthorized)
	}

	nodes, err := models.OnePermissionGroupTreeFromDB(&c.Lifecycle.DB, id)
	if err != nil {
		return c.RenderError(err)
	}
	err = models.InsertPerssionGROUP(&c.Lifecycle.DB, nodes[0], parentId, name)
	if err != nil {
		return c.RenderError(err)
	}
	return c.Redirect(routes.HengweiPermissionGroups.Index(0, 0, parentId))
}
