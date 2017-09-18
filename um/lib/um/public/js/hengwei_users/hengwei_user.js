var tableBody =new Array()

$(".user-row-data-input").each(function (i) {
    tableBody.push($(this).attr("key"))
})

$(document).ready(function() {
    $(".user_unlock").on("click",function () {
       var op= $(this).attr("key").split(":")
        var url= $("#user_unlock_url").val()
        var msg
        if(op[0] == "true"){
            msg="是否解锁该用户"
        }else{
            msg="是否锁定该用户"
        }
        var r=confirm(msg)
        if (r==true) {
            $.ajax({
                url:url,
                data:{"id":op[1]},
                success:function (data) {
                    if(data.ststus==1){
                        alert(data.msg)
                    }else {
                        alert(data.msg)
                    }
                    location.replace(location)
                },
                error:function () {
                     alert("未知错误 请稍后再试")
                }
            })
        }else {
            return false
        }
    })


    if($("input[name='hengweiUser.Source']").val() == "AD" ){
        $("input[name^='hengweiUser']").attr("readonly","readonly")
        $("textarea").attr("readonly","readonly")
    }
    if($("input[name='hengweiUser.Source']").val() == "sys" ){
        $("input[name='hengweiUser.Name']").attr("readonly","readonly")
    }

    if ($("#heng_wei_users_roles_data").val()){
        var rolesList=$("#heng_wei_users_roles_data").val().split(",")
        for (var i=0;i<rolesList.length;i++){
            $(".heng_wei_user_select").each(function () {
                if ($(this).val().indexOf(rolesList[i])==0){
                    tableBody.push($(this).val());
                    var data=$(this).val()
                    $(this).remove()
                    var a=$(this).val().split(":");
                    var tr='<tr>'+
                        '<td>' +
                        '<input type="checkbox" class="heng_wei_users—edit-row-checker" key="'+a[0]+'">' +
                        '<input type="text" hidden class="user-row-data-input" key="'+data+'">' +
                        '</td>'+
                        '<td><a href="/hengwei/aaa/HengweiRole/Edit?id='+a[0]+'">'+a[1]+'</a></td>'+
                        '<td>'+a[2]+'</td>'+
                        '<td>'+
                        '<a onclick="userdeleteTaboyOneTr($(this),'+a[0]+')">删除</a>'+
                        '</td>'+
                        '</tr>'
                    $("tbody").append(tr)
                }
            });
        }
    }
    if ($('.footable').length !=0){
        $('.footable').footable({pageNavigation:".pagination1"})
    }
});

$("#heng_wei_users-edit-all-checker").on("click", function () {
    var all_checked =  this.checked;
    $(".heng_wei_users—edit-row-checker").each(function(){
        this.checked = all_checked;
        return true;
    });
    return true;
});

$("#myModal").on('hide.bs.modal', function (e) {
    $(".heng_wei_user_select").removeAttr("selected");
});

$(function () {
    $(".heng_wei_roles_add_group").on("click",function () {
        var tableBody1 =new Array()
        var optaion =new Array()
        $(".form-control :selected").each(function(){
            tableBody.push($(this).val());
            tableBody1.push($(this).val());
            $(this).remove()
        });
        $("#myModal").modal('hide')
        for (var i=0;i<tableBody1.length;i++ ){
            var a=tableBody1[i].split(":");
            var tr='<tr>'+
                        '<td>' +
                            '<input type="checkbox" class="heng_wei_users—edit-row-checker" key="'+a[0]+'">' +
                            '<input type="text" hidden class="user-row-data-input" key="'+tableBody1[i]+'">' +
                        '</td>'+
                        '<td><a href="/hengwei/aaa/HengweiRole/Edit?id='+a[0]+'">'+a[1]+'</a></td>'+
                        '<td>'+a[2]+'</td>'+
                        '<td>'+
                            '<a onclick="userdeleteTaboyOneTr($(this),'+a[0]+')">删除</a>'+
                        '</td>'+
                    '</tr>'
            $("tbody").append(tr)
        }
    });

    $("#sync_AD_user").on("click",function () {
        var url =$("#sync_AD_user_url").val()
        $.ajax({
            url:url,
            success:function (data) {
                alert("成功添加条数 : "+data.addCountint+"\n"+"成功更新条数 : "+data.successCount+"\n"+"失败条数 : "+data.errCount)
                if (data.oldUsers.length>0){
                    $("#sync_AD_user_chose tbody").html("")
                    var u=data.oldUsers
                    for (var i=0 ;i<data.oldUsers.length;i++){
                        var tr='<option value="'+u[i].id+'" class="sync_AD_user_select">'+u[i].name+'</option>'
                        $("#sync_AD_user_chose").append(tr)
                    }
                    $("#myModalUser").modal('show');
                }else{
                    window.location.reload();
                }
            },
            error:function (data) {
                var re=data.responseJSON
                if (re.status==0){
                    alert("同步失败 ："+re.msg)
                }
            }
        })
    })

    $("#sync_AD_user_delete_commit").on("click",function () {
        var f = document.createElement("form");
        f.action = $("#heng_wei_users-delete").attr("url");
        f.method="POST";
        var inputField1 = document.createElement("input");
        inputField1.type = "hidden";
        inputField1.name = "_method";
        inputField1.value = "DELETE";
        f.appendChild(inputField1);
        var ids =[]
        $(".sync_AD_user_select:selected").each(function (i) {
            ids.push($(this).val())
        });
        var inputField = document.createElement("input");
        inputField.type = "hidden";
        inputField.name = "id";
        inputField.value = ids.toString();
        f.appendChild(inputField);
        document.body.appendChild(f);
        f.submit();
        $("#myModalUser").modal('hide');
    })
})

function submitFrom(){
    var f=document.getElementById("hengwei_users-edit");
    var rolesList =new Array();
    $(".heng_wei_users—edit-row-checker").each(function (i) {
        var inputField = document.createElement("input");
        inputField.type = "hidden";
        inputField.name = "role_id_list[]";
        inputField.value = $(this).attr("key");
        rolesList.push($(this).attr("key"))
        f.appendChild(inputField);
    });
    $("#heng_wei_users_roles_data").val(rolesList)
    document.body.appendChild(f);
    f.submit();
}

function userdeleteTaboyAllTr() {
    bootbox.confirm("确认删除选定信息？", function(result){
        if (!result) {
            return;
        }
        $(".heng_wei_users—edit-row-checker:checked").each(function (i) {
            var id=$(this).attr("key").split(":")[0]
            for (var i=0;i<tableBody.length;i++ ){
                if(tableBody[i].indexOf(id)>-1){
                    var op ='<option value="'+tableBody[i]+'" class="heng_wei_user_select">'+tableBody[i].split(":")[1]+'</option>'
                    $(".form-control").append(op)
                    tableBody.splice(i,1)
                }
            }
            $(this).parent().parent().remove()
        });
        $("#heng_wei_users-edit-all-checker").prop("checked",false);
    });
    return false
}

function userdeleteTaboyOneTr(a,id) {
    for (var i=0;i<tableBody.length;i++ ){
        if(tableBody[i].indexOf(id)>-1){
            var op ='<option value="'+tableBody[i]+'" class="heng_wei_user_select">'+tableBody[i].split(":")[1]+'</option>'
            $(".form-control").append(op)
            tableBody.splice(i,1)
        }
    }
    a.parent().parent().remove()
    bootbox.alert('删除成功')
}

$(function () {
    var urlPrefix = $("#urlPrefix").val();
    $("#heng_wei_users-all-checker").on("click", function () {
        var all_checked =  this.checked
        $(".heng_wei_users-row-checker").each(function(){
            this.checked = all_checked
            return true;
        });
        return true;
    });

    $("#heng_wei_users-delete").on("click", function () {
        bootbox.confirm("确认删除选定信息？", function(result){
            if (!result) {
                return;
            }
            var f = document.createElement("form");
            f.action = $("#heng_wei_users-delete").attr("url");
            f.method="POST";
            var inputField1 = document.createElement("input");
            inputField1.type = "hidden";
            inputField1.name = "_method";
            inputField1.value = "DELETE";
            f.appendChild(inputField1);
            $(".heng_wei_users-row-checker:checked").each(function (i) {
                var inputField = document.createElement("input");
                inputField.type = "hidden";
                inputField.name = "user_id_list[]";
                inputField.value = $(this).attr("key");
                f.appendChild(inputField);
            });
            document.body.appendChild(f);
            f.submit();
        })
        return false
    });

    $("#heng_wei_users-edit").on("click", function () {
        var elements = $(".heng_wei_users-row-checker:checked");
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
