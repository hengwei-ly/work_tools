package controllers

import (
	"time"
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

// HengweiUsers - 控制器
type HengweiUsers struct {
	App
}

// 列出所有记录
func (c HengweiUsers) Index(pageIndex int, pageSize int) revel.Result {
	if !c.CurrentUserHasPermission("um.HengweiUsers", web_ext.QUERY) {
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
	total, err := c.Lifecycle.DB.Users().Where().And(cond).Count()
	if err != nil {
		return c.RenderError(errors.Wrap(err, "读用户失败"))
	}

	var hengweiUsers []models.User
	err = c.Lifecycle.DB.Users().Where().And(cond).OrderBy("id").
		Offset(pageIndex * pageSize).
		Limit(pageSize).
		All(&hengweiUsers)
	if err != nil {
		return c.RenderError(errors.Wrap(err, "读用户失败"))
	}

	//查询每个用户的角色
	for i := range hengweiUsers {
		roles, err := models.GetRolesFromUser(&c.Lifecycle.DB, hengweiUsers[i].ID)
		if err != nil {
			return c.RenderError(errors.Wrap(err, "读用户的角色信息失败"))
		}
		hengweiUsers[i].Roles = roles
	}

	fields, err := readFieldDefinitionsFromFile(c.Lifecycle.Env)
	if err != nil {
		return c.RenderError(errors.Wrap(err, "读用户自定义字段失败"))
	}

	var onlineUsers []permissions.OnlineUser
	err = c.Lifecycle.DB.OnlineUsers().Where().All(&onlineUsers)
	if err != nil {
		return c.RenderError(errors.Wrap(err, "获取在线用户失败"))
	}

	paginator := toolbox.NewPaginator(c.Request.Request, pageSize, total)
	return c.Render(hengweiUsers, paginator, fields, onlineUsers)
}

// 编辑新建记录
func (c HengweiUsers) New() revel.Result {
	if !c.CurrentUserHasPermission("um.HengweiUsers", web_ext.CREATE) {
		return c.RenderError(permissions.ErrUnauthorized)
	}

	fields, err := readFieldDefinitionsFromFile(c.Lifecycle.Env)
	if err != nil {
		return c.RenderError(err)
	}
	c.ViewArgs["allRoles"] = choices.Role(&c.Lifecycle.DB, false)
	return c.Render(fields)
}

//创建记录
func (c HengweiUsers) Create(hengweiUser models.User, role_id_list []int64) revel.Result {
	if !c.CurrentUserHasPermission("um.HengweiUsers", web_ext.CREATE) {
		return c.RenderError(permissions.ErrUnauthorized)
	}

	if hengweiUser.Validate(c.Validation) {
		c.Validation.Keep()
		c.FlashParams()
		return c.Redirect(routes.HengweiUsers.New())
	}

	tx, err := c.Lifecycle.DB.Begin()
	if err != nil {
		c.Flash.Error(err.Error())
		c.FlashParams()
		return c.Redirect(routes.HengweiUsers.New())
	}

	defer util.CloseWith(tx)

	user := hengweiUser.ToUser()
	userID, err := tx.Users().Insert(&user)
	if err != nil {
		if oerr, ok := err.(*orm.Error); ok {
			for _, validation := range oerr.Validations {
				c.Validation.Error(validation.Message).Key(permissions.KeyForUsers(validation.Key))
			}
			c.Validation.Keep()
		}
		c.Flash.Error(err.Error())
		c.FlashParams()
		return c.Redirect(routes.HengweiUsers.New())
	}

	//将角色与用户关联
	if len(role_id_list) > 0 {
		userID64 := userID.(int64)
		for _, v := range role_id_list {
			_, err = tx.UsersAndRoles().Insert(&permissions.UserAndRole{
				RoleID: v, UserID: userID64})
			if err != nil {
				c.Flash.Error(err.Error())
				c.FlashParams()
				return c.Redirect(routes.HengweiUsers.New())
			}
		}
	}

	if err := tx.Commit(); err != nil {
		c.Flash.Error(err.Error()) //报错： Transaction has already been committed or rolled back
		c.FlashParams()
		return c.Redirect(routes.HengweiUsers.New())
	}
	c.Flash.Success(revel.Message(c.Request.Locale, "insert.success"))
	return c.Redirect(routes.HengweiUsers.Index(0, 0))
}

//编辑指定 id 的记录
func (c HengweiUsers) Edit(id int64) revel.Result {
	if !c.CurrentUserHasPermission("um.HengweiUsers", web_ext.UPDATE) {
		return c.RenderError(permissions.ErrUnauthorized)
	}
	var hengweiUser models.User
	err := c.Lifecycle.DB.Users().Id(id).Get(&hengweiUser)
	if err != nil {
		if err == orm.ErrNotFound {
			c.Flash.Error(revel.Message(c.Request.Locale, "update.record_not_found"))
		} else {
			c.Flash.Error(err.Error())
		}
		c.FlashParams()
		return c.Redirect(routes.HengweiUsers.Index(0, 0))
	}

	//查询角色
	roles, err := models.GetRolesFromUser(&c.Lifecycle.DB, id)
	hengweiUser.Roles = roles
	if err != nil {
		c.Flash.Error(err.Error())
		c.FlashParams()
		return c.Redirect(routes.HengweiUsers.Index(0, 0))
	}

	fields, err := readFieldDefinitionsFromFile(c.Lifecycle.Env)
	if err != nil {
		c.Flash.Error(err.Error())
		c.FlashParams()
		return c.Redirect(routes.HengweiUsers.Index(0, 0))
	}
	c.ViewArgs["allRoles"] = choices.Remove(&c.Lifecycle.DB, false, roles)
	return c.Render(hengweiUser, fields)
}

// 按 id 更新记录
func (c HengweiUsers) Update(hengweiUser models.User, role_id_list []int64) revel.Result {
	if !c.CurrentUserHasPermission("um.HengweiUsers", web_ext.UPDATE) {
		return c.RenderError(permissions.ErrUnauthorized)
	}
	hasPassword := hengweiUser.Password != ""
	if !hasPassword {
		hengweiUser.Password = "Validate_$2dfg&123_Is_Ok"
	}

	validation := hengweiUser.Validate(c.Validation)
	if !hasPassword {
		hengweiUser.Password = ""
	}
	if validation {
		c.Validation.Keep()
		c.FlashParams()
		return c.Redirect(routes.HengweiUsers.Edit(hengweiUser.ID))
	}

	tx, err := c.Lifecycle.DB.Begin()
	if err != nil {
		c.Flash.Error(err.Error())
		c.FlashParams()
		return c.Redirect(routes.HengweiUsers.Edit(hengweiUser.ID))
	}
	defer util.CloseWith(tx)

	var updater orm.Updater = tx.Users().Id(hengweiUser.ID)
	if !hasPassword {
		updater = updater.Omit("password")
	}

	//若是DB的数据
	user := hengweiUser.ToUser()
	err = updater.Update(&user)
	if err != nil {
		if err == orm.ErrNotFound {
			c.Flash.Error(revel.Message(c.Request.Locale, "update.record_not_found"))
		} else {
			if oerr, ok := err.(*orm.Error); ok {
				for _, validation := range oerr.Validations {
					c.Validation.Error(validation.Message).Key(permissions.KeyForUsers(validation.Key))
				}
				c.Validation.Keep()
			}
			c.Flash.Error(err.Error())
		}
		c.FlashParams()
		return c.Redirect(routes.HengweiUsers.Edit(hengweiUser.ID))
	}

	//更新用户与角色关系
	roles, err := models.GetRolesFromUser(&c.Lifecycle.DB, hengweiUser.ID)
	if err != nil {
		c.Flash.Error(err.Error())
		c.FlashParams()
		return c.Redirect(routes.HengweiUsers.Edit(hengweiUser.ID))
	}

	err = models.CheckUserForRole(tx, hengweiUser.ID, roles, role_id_list)
	if err != nil {
		c.Flash.Error(err.Error())
		c.FlashParams()
		return c.Redirect(routes.HengweiUsers.Edit(hengweiUser.ID))
	}

	if err = tx.Commit(); err != nil {
		c.Flash.Error(err.Error())
		c.FlashParams()
		return c.Redirect(routes.HengweiUsers.Edit(hengweiUser.ID))
	}

	c.Flash.Success(revel.Message(c.Request.Locale, "update.success"))
	return c.Redirect(routes.HengweiUsers.Index(0, 0))
}

// 按 id 删除记录
func (c HengweiUsers) Delete(id int64) revel.Result {
	if !c.CurrentUserHasPermission("um.HengweiUsers", web_ext.DELETE) {
		return c.RenderError(permissions.ErrUnauthorized)
	}
	err := c.Lifecycle.DB.Users().Id(id).Delete()
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
	return c.Redirect(HengweiUsers.Index)
}

// 按 id_list 删除记录
func (c HengweiUsers) DeleteByIDs(user_id_list []int64) revel.Result {
	if !c.CurrentUserHasPermission("um.HengweiUsers", web_ext.DELETE) {
		return c.RenderError(permissions.ErrUnauthorized)
	}
	_, err := c.Lifecycle.DB.Users().Where().And(orm.Cond{"id IN": user_id_list}).Delete()
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
	return c.Redirect(HengweiUsers.Index)
}

//用户解锁
func (c HengweiUsers) UserUnlock(id int64) revel.Result {
	result := map[string]interface{}{}
	var user permissions.User
	err := c.Lifecycle.DB.Users().Id(id).Get(&user)
	if err != nil {
		result["status"] = 0
		result["msg"] = "操作失败 :" + "UnlockGet fail" + err.Error()
		return c.RenderJSON(result)
	}
	var msg string
	if user.LockedAt != nil {
		user.LockedAt = nil
		_, err = c.Lifecycle.DB.Engine.Cols("locked_at").ID(id).Update(&user)
		if err != nil {
			result["status"] = 0
			result["msg"] = "解锁失败 :" + "UnlockUpdate fail" + err.Error()
			return c.RenderJSON(result)
		}
		msg = "解锁成功"
	} else {
		ti := time.Now()
		user.LockedAt = &ti
		err = c.Lifecycle.DB.Users().Id(id).Update(&user)
		if err != nil {
			result["status"] = 0
			result["msg"] = "锁定失败 :" + "UnlockUpdate fail" + err.Error()
			return c.RenderJSON(result)
		}
		msg = "锁定成功"
	}
	result["status"] = 1
	result["msg"] = msg
	c.RenderJSON(result)
	return c.RenderJSON(result)
}

//同步AD域 的用户信息
func (c HengweiUsers) SyncAD() revel.Result {
	var errCount, successCount, addCountInt = 0, 0, 0
	result := map[string]interface{}{}
	newUser, err := permissions.ReadUserFromLDAP(c.Lifecycle.Env)
	if err != nil {
		c.Response.Status = 500
		result["status"] = 0
		result["msg"] = "SyncAD.GetADUser:  " + err.Error()
		return c.RenderJSON(result)
	}

	var oldUsers []permissions.User
	err = c.Lifecycle.DB.Users().Where(orm.Cond{"source": "AD"}).All(&oldUsers)
	if err != nil {
		c.Response.Status = 500
		result["status"] = 0
		result["msg"] = "Marshal.GetDBUser:  " + err.Error()
		return c.RenderJSON(result)
	}
	var users []permissions.User
	for _, r := range newUser {
		found := true
		for i, v := range oldUsers {
			if r.Name == v.Name {
				k := i + 1
				oldUsers = append(oldUsers[:i], oldUsers[k:]...)
				found = false
				err := c.Lifecycle.DB.Users().Id(v.ID).Update(r)
				if err != nil {
					errCount += 1
					revel.ERROR.Println("Update.User:", err)
				}
				successCount += 1
				break
			}
		}
		if found {
			users = append(users, r)
		}
	}
	for _, user := range users {
		_, err := c.Lifecycle.DB.Users().Insert(&user)
		if err != nil {
			errCount += 1
			revel.ERROR.Println("Insert.User:", err)
		} else {
			addCountInt += 1
		}
	}
	result["status"] = 1
	result["errCount"] = errCount
	result["successCount"] = successCount
	result["addCountint"] = addCountInt
	result["oldUsers"] = oldUsers
	return c.RenderJSON(result)
}
