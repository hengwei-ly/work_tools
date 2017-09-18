package controllers

import (
	"encoding/json"
	"os"
	"sort"
	"um/app/models"

	"github.com/revel/revel"
	"github.com/three-plus-three/modules/environment"
	"github.com/three-plus-three/modules/errors"
	"github.com/three-plus-three/modules/permissions"
	"github.com/three-plus-three/modules/tid"
	"github.com/three-plus-three/modules/web_ext"
)

type HengweiOptions struct {
	App
}

type ADFields struct {
	ID   string
	Name string
}

func (c HengweiOptions) SetUserField(active string) revel.Result {
	if !c.CurrentUserHasPermission("um_set_field", web_ext.QUERY) {
		return c.RenderError(permissions.ErrUnauthorized)
	}
	fields, err := readFieldDefinitionsFromFile(c.Lifecycle.Env)
	if err != nil {
		return c.RenderError(errors.Wrap(err, "GetFields"))
	}
	c.ViewArgs["fields"] = fields
	return c.Render()
}

func (c HengweiOptions) UpdateUserField(fields []models.Field) revel.Result {
	sort.Slice(fields, func(i, j int) bool {
		return fields[i].ID < fields[j].ID
	})
	result := map[string]string{}
	filName := c.Lifecycle.Env.Fs.FromDataConfig(umExtendFields)
	if len(fields) > 0 {
		for i, f := range fields {
			if f.ID == "" {
				fields[i].ID = tid.GenerateID()
			}
		}
	}
	err := writeToFile(fields, filName)
	if err != nil {
		result["status"] = "0"
		result["msg"] = "writeToFile" + err.Error()
	} else {
		result["status"] = "1"
		result["msg"] = "保存成功"
	}
	return c.RenderJSON(result)
}

func (c HengweiOptions) SyncADView(active string) revel.Result {
	if !c.CurrentUserHasPermission("um_sync_AD", web_ext.QUERY) {
		return c.RenderError(permissions.ErrUnauthorized)
	}

	fields, err := readFieldDefinitionsFromFile(c.Lifecycle.Env)
	if err != nil {
		return c.RenderError(errors.Wrap(err, "GetFields:"))
	}

	fileName := c.Lifecycle.Env.Fs.FromDataConfig("syncADRule.json")
	fieldBind, err := readLDAPToUserBindsFromFile(fileName)
	if err != nil {
		return c.RenderError(errors.Wrap(err, "GetFields:"))
	}
	c.ViewArgs["fields"] = fields
	c.ViewArgs["fieldBind"] = fieldBind
	c.ViewArgs["adfields"] = getLDAPFields()
	return c.Render()
}

func (c HengweiOptions) SyncADRule(umFields []models.LDAPToUserBind) revel.Result {
	filName := c.Lifecycle.Env.Fs.FromDataConfig("syncADRule.json")

	err := writeToFile(umFields, filName)
	result := map[string]string{}
	if err != nil {
		result["status"] = "0"
		result["msg"] = errors.Wrap(err, "writeToFile:").Error()
	}
	result["status"] = "1"
	result["msg"] = "保存成功"
	return c.RenderJSON(result)
}

func (c HengweiOptions) GetPermission(permission string, operation string) revel.Result {
	var op string
	switch operation {
	case "edit":
		op = web_ext.UPDATE
	case "del":
		op = web_ext.DELETE
	}
	result := map[string]interface{}{}
	if !c.CurrentUserHasPermission(permission, op) {
		result["status"] = false
	} else {
		result["status"] = true
	}
	return c.RenderJSON(result)
}

const umExtendFields = "um_extend_fields.json"

func readFieldDefinitionsFromFile(env *environment.Environment) ([]models.Field, error) {
	fileName := env.Fs.FromDataConfig(umExtendFields)
	in, err := os.Open(fileName)
	if err != nil {
		if os.IsNotExist(err) {
			return []models.Field{
				{ID: "white_address_list",
					Name:      "登录IP",
					IsDefault: "true",
					Type:      "text"},
			}, nil
		}
		return nil, errors.New("readFieldDefinitionsFromFile: " + err.Error())
	}
	defer in.Close() // nolint

	var fields []models.Field
	err = json.NewDecoder(in).Decode(&fields)
	if err != nil {
		return nil, errors.New("read '" + fileName + "' fail: " + err.Error())
	}
	return fields, nil
}

func readLDAPToUserBindsFromFile(filename string) ([]models.LDAPToUserBind, error) {
	in, err := os.Open(filename)
	if err != nil {
		if os.IsNotExist(err) {
			return []models.LDAPToUserBind{}, nil
		}
		return nil, errors.New("ReadFieldDefinitionsFromFile: " + err.Error())
	}
	defer in.Close() // nolint

	var fields []models.LDAPToUserBind
	err = json.NewDecoder(in).Decode(&fields)
	if err != nil {
		return nil, errors.New("read '" + filename + "' fail: " + err.Error())
	}
	return fields, nil
}

func writeToFile(fields interface{}, filename string) error {
	fd, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644) // nolint
	if err != nil {
		return errors.Wrap(err, "打开文件失败 ")
	}
	defer fd.Close() // nolint

	buf, err := json.Marshal(fields)
	if err != nil {
		return errors.Wrap(err, "序列化失败 ")
	}
	_, err = fd.Write(buf)
	if err != nil {
		return errors.Wrap(err, "写文件失败")
	}
	return err
}

//获取活动目录的字段
func getLDAPFields() []ADFields {
	var field = []ADFields{
		{"name", "名称"},
		{"age", "年龄"},
		{"number", "电话"}}
	return field
}
