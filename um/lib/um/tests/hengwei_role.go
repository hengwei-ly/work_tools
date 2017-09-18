package tests

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"um/app"

	Permissions "github.com/three-plus-three/modules/permissions"
)

//  RolesTest 测试
type HengweiRoleTest struct {
	BaseTest
}

func (t HengweiRoleTest) TestIndex() {
	t.ClearTable("hengwei_roles")
	t.LoadFiles("tests/fixtures/roles.yaml")
	//conds := EQU{"name": "这是一个规则名,请替换成正确的值"}
	conds := EQU{}
	ruleId := t.GetIDFromTable("hengwei_roles", conds)

	t.Get(t.ReverseUrl("Roles.Index"))
	t.AssertOk()
	t.AssertContentType("text/html; charset=utf-8")
	//t.AssertContains("这是一个规则名,请替换成正确的值")

	var roles Permissions.Role
	err := app.Lifecycle.DB.Roles().Id(ruleId).Get(&roles)
	if err != nil {
		t.Assertf(false, err.Error())
	}

	t.AssertContains(fmt.Sprint(roles.Name))
}

func (t HengweiRoleTest) TestNew() {
	t.ClearTable("hengwei_roles")
	t.Get(t.ReverseUrl("Roles.New"))
	t.AssertOk()
	t.AssertContentType("text/html; charset=utf-8")
}

func (t HengweiRoleTest) TestCreate() {
	t.ClearTable("hengwei_roles")
	v := url.Values{}

	v.Set("roles.Name", "Fugiat fuga cumque.")

	v.Set("roles.CreatedAt", "1999-08-06T23:20:19+08:00")

	v.Set("roles.UpdatedAt", "1979-04-21T17:08:30+08:00")

	t.Post(t.ReverseUrl("Roles.Create"), "application/x-www-form-urlencoded", strings.NewReader(v.Encode()))
	t.AssertOk()

	//conds := EQU{"name": "这是一个规则名,请替换成正确的值"}
	conds := EQU{}
	ruleId := t.GetIDFromTable("hengwei_roles", conds)

	var roles Permissions.Role
	err := app.Lifecycle.DB.Roles().Id(ruleId).Get(&roles)
	if err != nil {
		t.Assertf(false, err.Error())
	}

	t.AssertEqual(fmt.Sprint(roles.Name), v.Get("roles.Name"))
}

func (t HengweiRoleTest) TestEdit() {
	t.ClearTable("hengwei_roles")
	t.LoadFiles("tests/fixtures/roles.yaml")
	//conds := EQU{"name": "这是一个规则名,请替换成正确的值"}
	conds := EQU{}
	ruleId := t.GetIDFromTable("hengwei_roles", conds)
	t.Get(t.ReverseUrl("Roles.Edit", ruleId))
	t.AssertOk()
	t.AssertContentType("text/html; charset=utf-8")

	var roles Permissions.Role
	err := app.Lifecycle.DB.Roles().Id(ruleId).Get(&roles)
	if err != nil {
		t.Assertf(false, err.Error())
	}
	fmt.Println(string(t.ResponseBody))

	t.AssertContains(fmt.Sprint(roles.Name))
}

func (t HengweiRoleTest) TestUpdate() {
	t.ClearTable("hengwei_roles")
	t.LoadFiles("tests/fixtures/roles.yaml")
	//conds := EQU{"name": "这是一个规则名,请替换成正确的值"}
	conds := EQU{}
	ruleId := t.GetIDFromTable("hengwei_roles", conds)
	v := url.Values{}
	v.Set("_method", "PUT")
	v.Set("roles.ID", strconv.FormatInt(ruleId, 10))

	v.Set("roles.Name", "Delectus vel maiores eaque modi adipisci sapiente et.")

	v.Set("roles.CreatedAt", "1970-11-20T00:59:11+08:00")

	v.Set("roles.UpdatedAt", "1983-12-12T07:44:14+08:00")

	t.Post(t.ReverseUrl("Roles.Update"), "application/x-www-form-urlencoded", strings.NewReader(v.Encode()))
	t.AssertOk()

	var roles Permissions.Role
	err := app.Lifecycle.DB.Roles().Id(ruleId).Get(&roles)
	if err != nil {
		t.Assertf(false, err.Error())
	}

	t.AssertEqual(fmt.Sprint(roles.Name), v.Get("roles.Name"))
}

func (t HengweiRoleTest) TestDelete() {
	t.ClearTable("hengwei_roles")
	t.LoadFiles("tests/fixtures/roles.yaml")
	//conds := EQU{"name": "这是一个规则名,请替换成正确的值"}
	conds := EQU{}
	ruleId := t.GetIDFromTable("hengwei_roles", conds)
	t.Delete(t.ReverseUrl("Roles.Delete", ruleId))
	t.AssertStatus(http.StatusOK)
	//t.AssertContentType("application/json; charset=utf-8")
	count := t.GetCountFromTable("hengwei_roles", nil)
	t.Assertf(count == 0, "count != 0, actual is %v", count)
}

func (t HengweiRoleTest) TestDeleteByIDs() {
	t.ClearTable("hengwei_roles")
	t.LoadFiles("tests/fixtures/roles.yaml")
	//conds := EQU{"name": "这是一个规则名,请替换成正确的值"}
	conds := EQU{}
	ruleId := t.GetIDFromTable("hengwei_roles", conds)
	t.Delete(t.ReverseUrl("Roles.DeleteByIDs", []interface{}{ruleId}))
	t.AssertStatus(http.StatusOK)
	//t.AssertContentType("application/json; charset=utf-8")
	count := t.GetCountFromTable("hengwei_roles", nil)
	t.Assertf(count == 0, "count != 0, actual is %v", count)
}
