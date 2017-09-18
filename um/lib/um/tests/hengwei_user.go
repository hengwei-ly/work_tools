package tests

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"um/app"
	"um/app/models"
)

//  HengWeiUsersTest 测试
type HengweiUserTest struct {
	BaseTest
}

func (t HengweiUserTest) TestIndex() {
	t.ClearTable("hengwei_user")
	t.LoadFiles("tests/fixtures/heng_wei_users.yaml")
	//conds := EQU{"name": "这是一个规则名,请替换成正确的值"}
	conds := EQU{}
	ruleId := t.GetIDFromTable("hengwei_user", conds)

	t.Get(t.ReverseUrl("HengWeiUsers.Index"))
	t.AssertOk()
	t.AssertContentType("text/html; charset=utf-8")
	//t.AssertContains("这是一个规则名,请替换成正确的值")

	var hengWeiUser models.User
	err := app.Lifecycle.DB.Users().Id(ruleId).Get(&hengWeiUser)
	if err != nil {
		t.Assertf(false, err.Error())
	}

	t.AssertContains(fmt.Sprint(hengWeiUser.Name))
	t.AssertContains(fmt.Sprint(hengWeiUser.Description))
}

func (t HengweiUserTest) TestNew() {
	t.ClearTable("hengwei_user")
	t.Get(t.ReverseUrl("HengWeiUsers.New"))
	t.AssertOk()
	t.AssertContentType("text/html; charset=utf-8")
}

func (t HengweiUserTest) TestCreate() {
	t.ClearTable("hengwei_user")
	v := url.Values{}

	v.Set("hengWeiUser.Name", "gux")

	v.Set("hengWeiUser.Password", "2a5yj1593")

	v.Set("hengWeiUser.Description", "Fugit asperiores doloremque ut.")

	v.Set("hengWeiUser.CreatedAt", "2009-11-15T01:55:34+08:00")

	v.Set("hengWeiUser.UpdatedAt", "1971-05-02T20:46:49+08:00")

	t.Post(t.ReverseUrl("HengWeiUsers.Create"), "application/x-www-form-urlencoded", strings.NewReader(v.Encode()))
	t.AssertOk()

	//conds := EQU{"name": "这是一个规则名,请替换成正确的值"}
	conds := EQU{}
	ruleId := t.GetIDFromTable("hengwei_user", conds)

	var hengWeiUser models.User
	err := app.Lifecycle.DB.Users().Id(ruleId).Get(&hengWeiUser)
	if err != nil {
		t.Assertf(false, err.Error())
	}

	t.AssertEqual(fmt.Sprint(hengWeiUser.Name), v.Get("hengWeiUser.Name"))
	t.AssertEqual(fmt.Sprint(hengWeiUser.Password), v.Get("hengWeiUser.Password"))
	t.AssertEqual(fmt.Sprint(hengWeiUser.Description), v.Get("hengWeiUser.Description"))
}

func (t HengweiUserTest) TestEdit() {
	t.ClearTable("hengwei_user")
	t.LoadFiles("tests/fixtures/heng_wei_users.yaml")
	//conds := EQU{"name": "这是一个规则名,请替换成正确的值"}
	conds := EQU{}
	ruleId := t.GetIDFromTable("hengwei_user", conds)
	t.Get(t.ReverseUrl("HengWeiUsers.Edit", ruleId))
	t.AssertOk()
	t.AssertContentType("text/html; charset=utf-8")

	var hengWeiUser models.User
	err := app.Lifecycle.DB.Users().Id(ruleId).Get(&hengWeiUser)
	if err != nil {
		t.Assertf(false, err.Error())
	}
	fmt.Println(string(t.ResponseBody))

	t.AssertContains(fmt.Sprint(hengWeiUser.Name))
	t.AssertContains(fmt.Sprint(hengWeiUser.Description))
}

func (t HengweiUserTest) TestUpdate() {
	t.ClearTable("hengwei_user")
	t.LoadFiles("tests/fixtures/heng_wei_users.yaml")
	//conds := EQU{"name": "这是一个规则名,请替换成正确的值"}
	conds := EQU{}
	ruleId := t.GetIDFromTable("hengwei_user", conds)
	v := url.Values{}
	v.Set("_method", "PUT")
	v.Set("hengWeiUser.ID", strconv.FormatInt(ruleId, 10))

	v.Set("hengWeiUser.Name", "4m1")

	v.Set("hengWeiUser.Password", "w00il3zq0")

	v.Set("hengWeiUser.Description", "Voluptatem ut ut molestiae qui voluptatem.")

	v.Set("hengWeiUser.CreatedAt", "2001-10-22T10:55:54+08:00")

	v.Set("hengWeiUser.UpdatedAt", "2003-01-02T02:18:46+08:00")

	t.Post(t.ReverseUrl("HengWeiUsers.Update"), "application/x-www-form-urlencoded", strings.NewReader(v.Encode()))
	t.AssertOk()

	var hengWeiUser models.User
	err := app.Lifecycle.DB.Users().Id(ruleId).Get(&hengWeiUser)
	if err != nil {
		t.Assertf(false, err.Error())
	}

	t.AssertEqual(fmt.Sprint(hengWeiUser.Name), v.Get("hengWeiUser.Name"))

	t.AssertEqual(fmt.Sprint(hengWeiUser.Password), v.Get("hengWeiUser.Password"))

	t.AssertEqual(fmt.Sprint(hengWeiUser.Description), v.Get("hengWeiUser.Description"))

}

func (t HengweiUserTest) TestDelete() {
	t.ClearTable("hengwei_user")
	t.LoadFiles("tests/fixtures/heng_wei_users.yaml")
	//conds := EQU{"name": "这是一个规则名,请替换成正确的值"}
	conds := EQU{}
	ruleId := t.GetIDFromTable("hengwei_user", conds)
	t.Delete(t.ReverseUrl("HengWeiUsers.Delete", ruleId))
	t.AssertStatus(http.StatusOK)
	//t.AssertContentType("application/json; charset=utf-8")
	count := t.GetCountFromTable("hengwei_user", nil)
	t.Assertf(count == 0, "count != 0, actual is %v", count)
}

func (t HengweiUserTest) TestDeleteByIDs() {
	t.ClearTable("hengwei_user")
	t.LoadFiles("tests/fixtures/heng_wei_users.yaml")
	//conds := EQU{"name": "这是一个规则名,请替换成正确的值"}
	conds := EQU{}
	ruleId := t.GetIDFromTable("hengwei_user", conds)
	t.Delete(t.ReverseUrl("HengWeiUsers.DeleteByIDs", []interface{}{ruleId}))
	t.AssertStatus(http.StatusOK)
	//t.AssertContentType("application/json; charset=utf-8")
	count := t.GetCountFromTable("hengwei_user", nil)
	t.Assertf(count == 0, "count != 0, actual is %v", count)
}
