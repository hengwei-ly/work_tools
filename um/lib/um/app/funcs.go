package app

import (
	"fmt"

	"strings"

	"github.com/revel/revel"
	"github.com/three-plus-three/modules/environment"
	"github.com/three-plus-three/modules/permissions"
)

func initTemplateFuncs(env *environment.Environment) {
	revel.TemplateFuncs["getString"] = func(values interface{}, key string) interface{} {
		if valueMap, ok := values.(map[string]interface{}); ok {
			value := valueMap[key]
			if nil == value {
				return ""
			}
			if s, ok := value.(string); ok {
				return s
			}
			str := fmt.Sprint(value)
			str = strings.Replace(str, "[", "", -1)
			str = strings.Replace(str, "]", "", -1)
			var ips = strings.Split(str, " ")
			return strings.Join(ips, "\n")
		}
		return ""
	}
	revel.TemplateFuncs["checkInclude"] = func(new []string, old []string) bool {
		fond := false
		for _, n := range new {
			for _, o := range old {
				if n == o {
					fond = true
					break
				}
			}
		}
		return fond
	}

	revel.TemplateFuncs["checkPermissionSelect"] = func(id string, ids []string) bool {
		for _, v := range ids {
			if id == v {
				return true
			}
		}
		return false
	}

	revel.TemplateFuncs["getPermission"] = func(id string, permissionOfGroup map[string][]permissions.PermissionGroup) []permissions.PermissionGroup {
		permissionList := permissionOfGroup[id]
		if permissionList == nil {
			return nil
		}
		return permissionList
	}
}
