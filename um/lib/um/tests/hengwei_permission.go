package tests

//  PermissionsTest 测试
type HengweiPermissionsTest struct {
	BaseTest
}

/*func (t HengweiPermissionsTest) TestIndex() {
	t.ClearTable("hengwei_permission")
	t.LoadFiles("tests/fixtures/permissions.yaml")
	//conds := EQU{"name": "这是一个规则名,请替换成正确的值"}
	conds := EQU{}
	ruleId := t.GetIDFromTable("hengwei_permission", conds)

	t.Get(t.ReverseUrl("HengweiPermission.Index"))
	t.AssertOk()
	t.AssertContentType("text/html; charset=utf-8")
	//t.AssertContains("这是一个规则名,请替换成正确的值")

	var permissions models.HengweiPermissionObject
	err := app.Lifecycle.DB.HengweiPermissionObject().Id(ruleId).Get(&permissions)
	if err != nil {
		t.Assertf(false, err.Error())
	}

	t.AssertContains(fmt.Sprint(permissions.Name))
}*/

/*
func (t HengweiPermissionsTest) TestNew() {
	t.ClearTable("hengwei_permission")
	t.Get(t.ReverseUrl("HengweiPermission.New"))
	t.AssertOk()
	t.AssertContentType("text/html; charset=utf-8")
}

func (t HengweiPermissionsTest) TestCreate() {
	t.ClearTable("hengwei_permission")
	v := url.Values{}

	v.Set("hengweiPermission.Name", "Et rerum est suscipit et.")

	v.Set("hengweiPermission.CreatedAt", "1996-03-15T20:48:01+08:00")

	v.Set("hengweiPermission.UpdatedAt", "1976-06-12T13:39:11+08:00")

	t.Post(t.ReverseUrl("HengweiPermission.Create"), "application/x-www-form-urlencoded", strings.NewReader(v.Encode()))
	t.AssertOk()

	//conds := EQU{"name": "这是一个规则名,请替换成正确的值"}
	conds := EQU{}
	ruleId := t.GetIDFromTable("hengwei_permission", conds)

	var hengweiPermission models.HengweiPermissionObject
	err := app.Lifecycle.DB.HengweiPermissionObject().Id(ruleId).Get(&hengweiPermission)
	if err != nil {
		t.Assertf(false, err.Error())
	}

	t.AssertEqual(fmt.Sprint(hengweiPermission.Name), v.Get("hengweiPermission.Name"))
}

func (t HengweiPermissionsTest) TestEdit() {
	t.ClearTable("hengwei_permission")
	t.LoadFiles("tests/fixtures/permissions.yaml")
	//conds := EQU{"name": "这是一个规则名,请替换成正确的值"}
	conds := EQU{}
	ruleId := t.GetIDFromTable("hengwei_permission", conds)
	t.Get(t.ReverseUrl("HengweiPermission.Edit", ruleId))
	t.AssertOk()
	t.AssertContentType("text/html; charset=utf-8")

	var hengweiPermission models.HengweiPermissionObject
	err := app.Lifecycle.DB.HengweiPermissionObject().Id(ruleId).Get(&hengweiPermission)
	if err != nil {
		t.Assertf(false, err.Error())
	}
	fmt.Println(string(t.ResponseBody))

	t.AssertContains(fmt.Sprint(hengweiPermission.Name))
}

func (t HengweiPermissionsTest) TestUpdate() {
	t.ClearTable("hengwei_permission")
	t.LoadFiles("tests/fixtures/permissions.yaml")
	//conds := EQU{"name": "这是一个规则名,请替换成正确的值"}
	conds := EQU{}
	ruleId := t.GetIDFromTable("hengwei_permission", conds)
	v := url.Values{}
	v.Set("_method", "PUT")
	v.Set("hengweiPermission.ID", strconv.FormatInt(ruleId, 10))

	v.Set("hengweiPermission.Name", "Repellendus itaque aliquam consequuntur.")

	v.Set("hengweiPermission.CreatedAt", "2000-03-17T11:08:04+08:00")

	v.Set("hengweiPermission.UpdatedAt", "1976-04-16T18:52:18+08:00")

	t.Post(t.ReverseUrl("HengweiPermission.Update"), "application/x-www-form-urlencoded", strings.NewReader(v.Encode()))
	t.AssertOk()

	var hengweiPermission models.HengweiPermissionObject
	err := app.Lifecycle.DB.HengweiPermissionObject().Id(ruleId).Get(&hengweiPermission)
	if err != nil {
		t.Assertf(false, err.Error())
	}

	t.AssertEqual(fmt.Sprint(hengweiPermission.Name), v.Get("permissions.Name"))

}

func (t HengweiPermissionsTest) TestDelete() {
	t.ClearTable("hengwei_permission")
	t.LoadFiles("tests/fixtures/permissions.yaml")
	//conds := EQU{"name": "这是一个规则名,请替换成正确的值"}
	conds := EQU{}
	ruleId := t.GetIDFromTable("hengwei_permission", conds)
	t.Delete(t.ReverseUrl("HengweiPermission.Delete", ruleId))
	t.AssertStatus(http.StatusOK)
	//t.AssertContentType("application/json; charset=utf-8")
	count := t.GetCountFromTable("hengwei_permission", nil)
	t.Assertf(count == 0, "count != 0, actual is %v", count)
}

func (t HengweiPermissionsTest) TestDeleteByIDs() {
	t.ClearTable("hengwei_permission")
	t.LoadFiles("tests/fixtures/permissions.yaml")
	//conds := EQU{"name": "这是一个规则名,请替换成正确的值"}
	conds := EQU{}
	ruleId := t.GetIDFromTable("hengwei_permission", conds)
	t.Delete(t.ReverseUrl("HengweiPermission.DeleteByIDs", []interface{}{ruleId}))
	t.AssertStatus(http.StatusOK)
	//t.AssertContentType("application/json; charset=utf-8")
	count := t.GetCountFromTable("hengwei_permission", nil)
	t.Assertf(count == 0, "count != 0, actual is %v", count)
}
*/
