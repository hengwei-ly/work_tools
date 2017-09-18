
function rolesubmitFrom(){
    var opt =new Array()
    var f=document.getElementById("hengwei_roles-new-edit");
    $(".hengweiRoles-row-checker").each(function (i) {
        var inputField = document.createElement("input");
        inputField.type = "hidden";
        inputField.name = "group_id_list[]";
        inputField.value = $(this).attr("key");
        opt.push(("/"+$(this).attr("key")))
        f.appendChild(inputField);
    });
    $("#heng_wei_roles_all_data").val(opt)
    document.body.appendChild(f);
    f.submit();
}

function rolesdeleteTaboyAllTr() {
    bootbox.confirm("确认删除选定信息？", function(result){
        if (!result) {
            return;
        }
        $(".hengweiRoles-row-checker:checked").each(function (i) {
            var id=$(this).attr("key").split(":")[0]
            $('input:checkbox[name="selectGroup"]').each(function () {
                if(id==this.value){
                    this.removeAttribute('disabled');
                    $(this).attr("title","")
                }
            })

            $(this).parent().parent().remove()
        });
        $("#hengweiRoles-all-checker").prop("checked",false);
    });
    return false
}

function deleteTaboyOneTr(a,id) {
    $('input:checkbox[name="selectGroup"]').each(function () {
        if(id==this.value){
            this.removeAttribute('disabled');
            $(this).attr("title","")
        }
    })
    a.parent().parent().remove()
    bootbox.alert('删除成功')
}

function openEdit(a){
    $(":input[type='checkbox']").prop("checked",false);
    var opt=new Array();
    var opt =$($(a.parent().siblings()[0]).children()[0]).attr("key").split(":")[1].split(",")
    $("#myModal2").modal('show')
    for(var i=0;i<opt.length;i++){
        if(opt[i]=="create"){
            $("#inlineCheckbox5").prop("checked",true);
        }else if (opt[i]=="delete"){
            $("#inlineCheckbox6").prop("checked",true);
        }else if(opt[i]=="update"){
            $("#inlineCheckbox7").prop("checked",true);
        }else if(opt[i]=="query"){
            $("#inlineCheckbox8").prop("checked",true);
        }
    }
    $("#roles-row-data-edit_optaion").val($($(a.parent().siblings()[0]).children()[0]).attr("key").split(":")[0])
}


$(function () {

    if ($('.footable').length>0){
        $('.footable').footable({pageNavigation:".pagination1"})
    }

    $(".roles-row-data-input").each(function () {
        var id =$(this).attr("key").split(":")[0]
        $('input:checkbox[name="selectGroup"]').each(function () {
            if(id==this.value){
                this.setAttribute('disabled','disabled')
                $(this).attr("title","该组已选")
                $(this).addClass("easyui-tooltip")
            }
        })
    })

    $("#myModal").on('hidden.bs.modal', function (e) {
        $(":input[type='checkbox']").prop("checked",false);
    });

    $("#myModal2").on('hidden.bs.modal', function (e) {
        $(":input[type='checkbox']").prop("checked",false);
    });

    $(".roles-row-data-input").each(function (i) {
        var optaion=$(this).prev().attr("key").split(":")[1].split(",")
        var op="";
        for(var i=0;i<optaion.length;i++){
            switch (optaion[i]){
                case "create":
                    op+=" 添加"
                    break
                case "delete":
                    op+=" 删除"
                    break
                case "update":
                    op+=" 修改"
                    break
                case "query":
                    op+=" 查看"
                    break
            }
        }
        $($(this).parent().siblings()[2]).html(op)
    })


    $(".heng_wei_roles_add_group").on("click",function () {
        var optaion =new Array()
        var tableBody1 = []
        $(":input[type='checkbox']:checked").each(function (i) {
            optaion.push($(this).val());
        });

        var checkedObj = $('input:checkbox[name="selectGroup"]:checked');
        checkedObj.each(function() {
            tableBody1.push($(this).attr("key"));
            this.setAttribute('disabled','disabled')
            $(this).attr("title","该组已选")
            $(this).addClass("easyui-tooltip")

        });

        var op="";
        for(var i=0;i<optaion.length;i++){
            switch (optaion[i]){
                case "create":
                    op+=" 添加"
                    break
                case "delete":
                    op+=" 删除"
                    break
                case "update":
                    op+=" 修改"
                    break
                case "query":
                    op+=" 查看"
                    break
            }
        }

        $("#myModal").modal('hide')
        for (var i=0;i<tableBody1.length;i++ ){
            var a=tableBody1[i].split(":");
            var tr='<tr>'+
                '<td>' +
                '<input type="checkbox" id="roles_group_id'+a[0]+'"  class="hengweiRoles-row-checker" key='+a[0]+':'+optaion+'>' +
                '<input type="text" hidden class="roles-row-data-input" key="'+tableBody1[i]+'">'+
                '</td>' +
                '<td><a href="/hengwei/aaa/PermissionsGroups/Edit?id='+a[0]+'">'+a[1]+'</a></td>'+
                '<td>'+a[2]+'</td>'+
                '<td>'+op+'</td>'+
                '<td>'+
                '<a onclick="openEdit($(this))">编辑</a>  ' +
                '<a onclick="deleteTaboyOneTr($(this),'+a[0]+')">删除</a>'+
                '</td>'+
                '</tr>';
            $("tbody").append(tr)
        }
    })

    $("#heng_wei_roles_edit_options").on("click",function () {
        debugger
        var optaion = []
        $(".hengwei_role_optaion_check:checked").each(function (i) {
            optaion.push($(this).val());
        });

        var groupid=$("#roles-row-data-edit_optaion").val()
        $("#roles_group_id"+groupid).attr("key",groupid+":"+optaion)
        var op="";

        for(var i=0;i<optaion.length;i++){
            switch (optaion[i]){
                case "create":
                    op+=" 添加"
                    break
                case "delete":
                    op+=" 删除"
                    break
                case "update":
                    op+=" 修改"
                    break
                case "query":
                    op+=" 查看"
                    break
            }
        }

        console.log("0000000000",optaion)
        $($("#roles_group_id"+groupid).parent().siblings()[2]).text(op);
        $("#myModal2").modal('hide');
        return false
    });


    $("#hengweiRoles-all-checker").on("click", function () {
        var all_checked =  this.checked;
        $(".hengweiRoles-row-checker").each(function(){
            this.checked = all_checked;
            return true;
        });
        return true;
    });

});



$(function () {
    var urlPrefix = $("#urlPrefix").val();
    $("#roles-all-checker").on("click", function () {
        var all_checked =  this.checked
        $(".roles-row-checker").each(function(){
            this.checked = all_checked
            return true;
        });
        return true;
    });
    $("#roles-delete").on("click", function () {
        bootbox.confirm("确认删除选定信息？", function(result){
            if (!result) {
                return;
            }
            var f = document.createElement("form");
            f.action = $("#roles-delete").attr("url");
            f.method="POST";
            var inputField = document.createElement("input");
            inputField.type = "hidden";
            inputField.name = "_method";
            inputField.value = "DELETE";

            $(".roles-row-checker:checked").each(function (i) {
                var inputField = document.createElement("input");
                inputField.type = "hidden";
                inputField.name = "id_list[]";
                inputField.value = $(this).attr("key").split(":")[0];
                f.appendChild(inputField);
            });

            document.body.appendChild(f);
            f.submit();
        })
        return false
    });

    $("#roles-edit").on("click", function () {
        var elements = $(".roles-row-checker:checked");
        if (elements.length == 1) {
            window.location.href= elements.first().attr("url");
        } else if (elements.length == 0) {
            bootbox.alert('请选择一条记录！')
        } else {
            bootbox.alert('你选择了多条记录，请选择一条记录！')
        }
        return false
    });
});

