package controllers

import (
	"strconv"
	"strings"
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

// HengweiRoles - 控制器
type HengweiRoles struct {
	App
}

// 列出所有记录
func (c HengweiRoles) Index(pageIndex int, pageSize int) revel.Result {
	if !c.CurrentUserHasPermission("um.HengweiRoles", web_ext.QUERY) {
		return c.RenderError(permissions.ErrUnauthorized)
	}

	// var pageIndex, pageSize int
	// c.Params.Bind(&pageIndex, "pageIndex")
	// c.Params.Bind(&pageSize, "pageSize")
	if pageSize <= 0 {
		pageSize = toolbox.DEFAULT_SIZE_PER_PAGE
	}

	var cond orm.Cond
	if name := c.Params.Get("query"); name != "" {
		cond = orm.Cond{"name LIKE": "%" + name + "%"}
	}

	total, err := c.Lifecycle.DB.Roles().Where().And(cond).Count()
	if err != nil {
		return c.RenderError(errors.Wrap(err, "读角色失败"))
	}

	var hengweiRoles []models.Role
	err = c.Lifecycle.DB.Roles().Where().
		And(cond).
		Offset(pageIndex * pageSize).
		Limit(pageSize).OrderBy("id").
		All(&hengweiRoles)
	if err != nil {
		return c.RenderError(errors.Wrap(err, "读角色失败"))
	}

	//查询每个角色的组
	for i := range hengweiRoles {
		permissionGroup, err := models.GetPermissionsGroupFormRole(&c.Lifecycle.DB, hengweiRoles[i].ID)
		if err != nil {
			return c.RenderError(errors.Wrap(err, "读角色关联权限组失败"))
		}
		hengweiRoles[i].Groups = permissionGroup
	}
	paginator := toolbox.NewPaginator(c.Request.Request, pageSize, total)
	return c.Render(hengweiRoles, paginator)
}

// 编辑新建记录
func (c HengweiRoles) New() revel.Result {
	if !c.CurrentUserHasPermission("um.HengweiRoles", web_ext.CREATE) {
		return c.RenderError(permissions.ErrUnauthorized)
	}
	roots, err := models.LoadPermissionGroupTreeFromDB(&c.Lifecycle.DB)
	if err != nil {
		c.Flash.Error(err.Error())
		c.FlashParams()
		return c.Redirect(routes.HengweiRoles.Index(0, 0))
	}
	c.ViewArgs["allGroupNodes"] = roots
	return c.Render()
}

//创建记录
func (c HengweiRoles) Create(hengweiRole permissions.Role, group_id_list []string) revel.Result {
	if !c.CurrentUserHasPermission("um.HengweiRoles", web_ext.CREATE) {
		return c.RenderError(permissions.ErrUnauthorized)
	}
	if hengweiRole.Validate(c.Validation) {
		c.Validation.Keep()
		c.FlashParams()
		return c.Redirect(routes.HengweiRoles.New())
	}

	tx, err := c.Lifecycle.DB.Begin()
	if err != nil {
		c.Flash.Error(err.Error())
		c.FlashParams()
		return c.Redirect(routes.HengweiRoles.New())
	}
	defer util.CloseWith(tx)

	rolesId, err := tx.Roles().Insert(&hengweiRole)
	if err != nil {
		if oerr, ok := err.(*orm.Error); ok {
			for _, validation := range oerr.Validations {
				c.Validation.Error(validation.Message).Key(permissions.KeyForRoles(validation.Key))
			}
			c.Validation.Keep()
		}
		c.Flash.Error(err.Error())
		c.FlashParams()
		return c.Redirect(routes.HengweiRoles.New())
	}
	//将角色与组关联
	if len(group_id_list) != 0 {
		for _, v := range group_id_list {
			groupid, err := strconv.ParseInt(strings.Split(v, ":")[0], 10, 64)
			if err != nil {
				c.Flash.Error(err.Error())
				c.FlashParams()
				return c.Redirect(routes.HengweiRoles.New())
			}
			var groupAndRole = models.CreateGroupAndRole(groupid, rolesId.(int64), v)
			_, err = tx.PermissionGroupsAndRoles().Insert(&groupAndRole)
			if err != nil {
				c.Flash.Error(err.Error())
				c.FlashParams()
				return c.Redirect(routes.HengweiRoles.New())
			}
		}
	}
	err = tx.Commit()
	if err != nil {
		c.Flash.Error(err.Error())
		c.FlashParams()
		return c.Redirect(routes.HengweiRoles.New())
	}
	c.Flash.Success(revel.Message(c.Request.Locale, "insert.success"))
	return c.Redirect(routes.HengweiRoles.Index(0, 0))
}

//编辑指定 id 的记录
func (c HengweiRoles) Edit(id int64) revel.Result {
	if !c.CurrentUserHasPermission("um.HengweiRoles", web_ext.UPDATE) {
		return c.RenderError(permissions.ErrUnauthorized)
	}
	var hengweiRole models.Role
	err := c.Lifecycle.DB.Roles().Id(id).Get(&hengweiRole)
	if err != nil {
		if err == orm.ErrNotFound {
			c.Flash.Error(revel.Message(c.Request.Locale, "update.record_not_found"))
		} else {
			c.Flash.Error(err.Error())
		}
		c.FlashParams()
		return c.Redirect(routes.HengweiRoles.Index(0, 0))
	}

	//根据id 查询权限分组
	permissionGroupsForRoleOnly, err := models.GetPermissionsGroupFormRole(&c.Lifecycle.DB, id)
	if err != nil {
		c.Flash.Error(err.Error())
		c.FlashParams()
		return c.Redirect(routes.HengweiRoles.Index(0, 0))
	}

	hengweiRole.Groups = permissionGroupsForRoleOnly

	for i := range permissionGroupsForRoleOnly {
		var groupAndRoles []permissions.PermissionGroupAndRole
		err := c.Lifecycle.DB.PermissionGroupsAndRoles().Where(orm.Cond{"role_id": id}).And(orm.Cond{"group_id": permissionGroupsForRoleOnly[i].ID}).All(&groupAndRoles)
		if err != nil {
			c.Flash.Error(err.Error())
			c.FlashParams()
			return c.Redirect(routes.HengweiRoles.Index(0, 0))
		}
		var operation []string
		if len(groupAndRoles) > 0 {
			if groupAndRoles[0].CreateOperation {
				operation = append(operation, permissions.CREATE)
			}
			if groupAndRoles[0].DeleteOperation {
				operation = append(operation, permissions.DELETE)
			}
			if groupAndRoles[0].UpdateOperation {
				operation = append(operation, permissions.UPDATE)
			}
			if groupAndRoles[0].QueryOperation {
				operation = append(operation, permissions.QUERY)
			}
		}
		str := strings.Join(operation, ",")
		permissionGroupsForRoleOnly[i].Operation = str
	}

	roots, err := models.LoadPermissionGroupTreeFromDB(&c.Lifecycle.DB)
	if err != nil {
		c.Flash.Error(err.Error())
		c.FlashParams()
		return c.Redirect(routes.HengweiRoles.Index(0, 0))
	}
	c.ViewArgs["allGroupNodes"] = roots
	c.ViewArgs["hengweiRole"] = hengweiRole
	return c.Render()
}

// 按 id 更新记录
func (c HengweiRoles) Update(hengweiRole permissions.Role, group_id_list []string) revel.Result {
	if !c.CurrentUserHasPermission("um.HengweiRoles", web_ext.UPDATE) {
		return c.RenderError(permissions.ErrUnauthorized)
	}

	if hengweiRole.Validate(c.Validation) {
		c.Validation.Keep()
		c.FlashParams()
		return c.Redirect(routes.HengweiRoles.Edit(hengweiRole.ID))
	}

	tx, err := c.Lifecycle.DB.Begin()
	if err != nil {
		c.Flash.Error(err.Error())
		c.FlashParams()
		return c.Redirect(routes.HengweiRoles.Edit(hengweiRole.ID))
	}

	defer util.CloseWith(tx)

	//获取原角色权限组
	onlyRoleGroups, err := models.GetPermissionsGroupFormRole(tx, hengweiRole.ID)
	if err != nil {
		c.Flash.Error(err.Error())
		c.FlashParams()
		return c.Redirect(routes.HengweiRoles.Edit(hengweiRole.ID))
	}
	err = models.CheckRoleForPermissionGroup(tx, hengweiRole.ID, onlyRoleGroups, group_id_list)
	if err != nil {
		c.Flash.Error(err.Error())
		c.FlashParams()
		return c.Redirect(routes.HengweiRoles.Edit(hengweiRole.ID))
	}
	//更新角色
	err = tx.Roles().Id(hengweiRole.ID).Update(hengweiRole)
	if err != nil {
		if err == orm.ErrNotFound {
			c.Flash.Error(revel.Message(c.Request.Locale, "update.record_not_found"))
		} else {
			if oerr, ok := err.(*orm.Error); ok {
				for _, validation := range oerr.Validations {
					c.Validation.Error(validation.Message).Key(permissions.KeyForRoles(validation.Key))
				}
				c.Validation.Keep()
			}
			c.Flash.Error(err.Error())
		}
		c.FlashParams()
		return c.Redirect(routes.HengweiRoles.Edit(hengweiRole.ID))
	}
	if err = tx.Commit(); err != nil {
		c.Flash.Error(err.Error())
		c.FlashParams()
		return c.Redirect(routes.HengweiRoles.Edit(hengweiRole.ID))
	}
	c.Flash.Success(revel.Message(c.Request.Locale, "update.success"))
	c.FlashParams()
	return c.Redirect(routes.HengweiRoles.Index(0, 0))
}

// 按 id 删除记录
func (c HengweiRoles) Delete(id int64) revel.Result {
	if !c.CurrentUserHasPermission("um.HengweiRoles", web_ext.DELETE) {
		return c.RenderError(permissions.ErrUnauthorized)
	}
	err := c.Lifecycle.DB.Roles().Id(id).Delete()
	if nil != err {
		if err == orm.ErrNotFound {
			c.Flash.Error(revel.Message(c.Request.Locale, "delete.record_not_found"))
		} else {
			c.Flash.Error(err.Error())
		}
	} else {
		c.Flash.Success(revel.Message(c.Request.Locale, "delete.success"))
	}
	c.FlashParams()
	return c.Redirect(HengweiRoles.Index)
}

// 按 id 列表删除记录
func (c HengweiRoles) DeleteByIDs(id_list []int64) revel.Result {
	if !c.CurrentUserHasPermission("um.HengweiRoles", web_ext.DELETE) {
		return c.RenderError(permissions.ErrUnauthorized)
	}

	if len(id_list) == 0 {
		c.Flash.Error("请至少选择一条记录！")
		return c.Redirect(HengweiRoles.Index)
	}
	_, err := c.Lifecycle.DB.Roles().Where().And(orm.Cond{"id IN": id_list}).Delete()
	if nil != err {
		c.Flash.Error(err.Error())
	} else {
		c.Flash.Success(revel.Message(c.Request.Locale, "delete.success"))
	}

	c.FlashParams()
	return c.Redirect(HengweiRoles.Index)
}
