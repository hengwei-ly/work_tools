package tests

//  PermissionsGroupsTest 测试
type PermissionsGroupsTest struct {
	BaseTest
}

	/*
	func (t PermissionsGroupsTest) TestIndex() {
		t.ClearTable("hengwei_permissions_group")
		t.LoadFiles("tests/fixtures/permissions_groups.yaml")
		//conds := EQU{"name": "这是一个规则名,请替换成正确的值"}
		conds := EQU{}
		ruleId := t.GetIDFromTable("hengwei_permissions_group", conds)

		t.Get(t.ReverseUrl("PermissionsGroups.Index"))
		t.AssertOk()
		t.AssertContentType("text/html; charset=utf-8")
		//t.AssertContains("这是一个规则名,请替换成正确的值")

		var permissionsGroup models.PermissionGroup
		err := app.Lifecycle.DB.PermissionGroup().Id(ruleId).Get(&permissionsGroup)
		if err != nil {
			t.Assertf(false, err.Error())
		}

		t.AssertContains(fmt.Sprint(permissionsGroup.Name))
		t.AssertContains(fmt.Sprint(permissionsGroup.Description))
	}

	func (t PermissionsGroupsTest) TestNew() {
		t.ClearTable("hengwei_permissions_group")
		t.Get(t.ReverseUrl("PermissionsGroups.New"))
		t.AssertOk()
		t.AssertContentType("text/html; charset=utf-8")
	}

	func (t PermissionsGroupsTest) TestCreate() {
		t.ClearTable("hengwei_permissions_group")
		v := url.Values{}

		v.Set("permissionsGroup.Name", "9lh")

		v.Set("permissionsGroup.Description", "Quo adipisci unde dicta sunt.")

		v.Set("permissionsGroup.CreatedAt", "1983-08-31T05:30:18+08:00")

		v.Set("permissionsGroup.UpdatedAt", "2014-07-25T01:16:57+08:00")

		t.Post(t.ReverseUrl("PermissionsGroups.Create"), "application/x-www-form-urlencoded", strings.NewReader(v.Encode()))
		t.AssertOk()

		//conds := EQU{"name": "这是一个规则名,请替换成正确的值"}
		conds := EQU{}
		ruleId := t.GetIDFromTable("hengwei_permissions_group", conds)

		var permissionsGroup models.PermissionGroup
		err := app.Lifecycle.DB.PermissionGroup().Id(ruleId).Get(&permissionsGroup)
		if err != nil {
			t.Assertf(false, err.Error())
		}

		t.AssertEqual(fmt.Sprint(permissionsGroup.Name), v.Get("permissionsGroup.Name"))
		t.AssertEqual(fmt.Sprint(permissionsGroup.Description), v.Get("permissionsGroup.Description"))
	}

	func (t PermissionsGroupsTest) TestEdit() {
		t.ClearTable("hengwei_permissions_group")
		t.LoadFiles("tests/fixtures/permissions_groups.yaml")
		//conds := EQU{"name": "这是一个规则名,请替换成正确的值"}
		conds := EQU{}
		ruleId := t.GetIDFromTable("hengwei_permissions_group", conds)
		t.Get(t.ReverseUrl("PermissionsGroups.Edit", ruleId))
		t.AssertOk()
		t.AssertContentType("text/html; charset=utf-8")

		var permissionsGroup models.PermissionGroup
		err := app.Lifecycle.DB.PermissionGroup().Id(ruleId).Get(&permissionsGroup)
		if err != nil {
			t.Assertf(false, err.Error())
		}
		fmt.Println(string(t.ResponseBody))

		t.AssertContains(fmt.Sprint(permissionsGroup.Name))
		t.AssertContains(fmt.Sprint(permissionsGroup.Description))
	}

	func (t PermissionsGroupsTest) TestUpdate() {
		t.ClearTable("hengwei_permissions_group")
		t.LoadFiles("tests/fixtures/permissions_groups.yaml")
		//conds := EQU{"name": "这是一个规则名,请替换成正确的值"}
		conds := EQU{}
		ruleId := t.GetIDFromTable("hengwei_permissions_group", conds)
		v := url.Values{}
		v.Set("_method", "PUT")
		v.Set("permissionsGroup.ID", strconv.FormatInt(ruleId, 10))

		v.Set("permissionsGroup.Name", "310")

		v.Set("permissionsGroup.Description", "Autem cupiditate velit voluptas.")

		v.Set("permissionsGroup.CreatedAt", "1988-07-20T13:19:06+08:00")

		v.Set("permissionsGroup.UpdatedAt", "1970-01-17T20:29:58+08:00")

		t.Post(t.ReverseUrl("PermissionsGroups.Update"), "application/x-www-form-urlencoded", strings.NewReader(v.Encode()))
		t.AssertOk()

		var permissionsGroup models.PermissionGroup
		err := app.Lifecycle.DB.PermissionGroup().Id(ruleId).Get(&permissionsGroup)
		if err != nil {
			t.Assertf(false, err.Error())
		}

		t.AssertEqual(fmt.Sprint(permissionsGroup.Name), v.Get("permissionsGroup.Name"))

		t.AssertEqual(fmt.Sprint(permissionsGroup.Description), v.Get("permissionsGroup.Description"))

	}

	func (t PermissionsGroupsTest) TestDelete() {
		t.ClearTable("hengwei_permissions_group")
		t.LoadFiles("tests/fixtures/permissions_groups.yaml")
		//conds := EQU{"name": "这是一个规则名,请替换成正确的值"}
		conds := EQU{}
		ruleId := t.GetIDFromTable("hengwei_permissions_group", conds)
		t.Delete(t.ReverseUrl("PermissionsGroups.Delete", ruleId))
		t.AssertStatus(http.StatusOK)
		//t.AssertContentType("application/json; charset=utf-8")
		count := t.GetCountFromTable("hengwei_permissions_group", nil)
		t.Assertf(count == 0, "count != 0, actual is %v", count)
	}

	func (t PermissionsGroupsTest) TestDeleteByIDs() {
		t.ClearTable("hengwei_permissions_group")
		t.LoadFiles("tests/fixtures/permissions_groups.yaml")
		//conds := EQU{"name": "这是一个规则名,请替换成正确的值"}
		conds := EQU{}
		ruleId := t.GetIDFromTable("hengwei_permissions_group", conds)
		t.Delete(t.ReverseUrl("PermissionsGroups.DeleteByIDs", []interface{}{ruleId}))
		t.AssertStatus(http.StatusOK)
		//t.AssertContentType("application/json; charset=utf-8")
		count := t.GetCountFromTable("hengwei_permissions_group", nil)
		t.Assertf(count == 0, "count != 0, actual is %v", count)
	}
	*/
