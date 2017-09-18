
$(document).ready(function() {
    $(".hengwei_user_groups-checker").each(function (i) {
        var id=$(this).attr("key")
        var className=".hengwei_user_group_select-"+id
        $(className).hide()
    });

    $(".hengwei_group_add_user").on("click",function () {
        var tableBody =new Array()
        $(".hengwei_user_group_allUser :selected").each(function(){
            tableBody.push($(this).val());
            $(this).hide();
        });

        $("#myModal").modal('hide');
        for (var i=0;i<tableBody.length;i++ ){
            var a=tableBody[i].split(":");
            var tr='<tr>'+
                '<td>' +
                '<input type="checkbox" class="hengwei_user_groups-checker" key="'+a[0]+'">' +
                '<input type="text" hidden name="userGroupView.UserID[]" value="'+a[0]+'">' +
                '</td>'+
                '<td>'+a[1]+'</td>'+
                '<td>'+a[2]+'</td>'+
                '</tr>'
            $("tbody").append(tr)
        }
    });

    $(".hengwei_group_Import_GroupsUser").on("click",function () {
        var groupsId =new Array()
        $(".user_group_import:checked").each(function (i) {
            var isCheck =$(this).attr("key");
            groupsId.push(isCheck)
        })

        if (groupsId.length==0){
            Msg.warn("请选择用户组")
            return false
        }

        var url =$("#ImportGroupUrl").val()
        $.ajax({
            url:url,
            data:{"id":groupsId.toString()},
            success:function (data) {
                $("#myModa2").modal('hide');
                var users= data.data
                if($(".hengwei_user_groups-checker").length ==0){
                    for (var i=0;i<users.length;i++ ){
                        var user=users[i];
                        var tr='<tr>'+
                            '<td>' +
                            '<input type="checkbox" class="hengwei_user_groups-checker" key="'+user["id"]+'">' +
                            '<input type="text" hidden name="userGroupView.UserID[]" value="'+user["id"]+'">' +
                            '</td>'+
                            '<td>'+user.name+'</td>'+
                            '<td>'+user.description+'</td>'+
                            '</tr>'
                        $("tbody").append(tr)
                        var className=".hengwei_user_group_select-"+a["id"]
                        $(className).hide()
                    }
                }else {
                    for (var i = 0; i < users.length; i++) {
                        var user = users[i];
                        $(".hengwei_user_groups-checker").each(function () {
                            if (user["id"] != $(this).attr("key")) {
                                var tr = '<tr>' +
                                    '<td>' +
                                    '<input type="checkbox" class="hengwei_user_groups-checker" key="' + user["id"] + '">' +
                                    '<input type="text" hidden name="userGroupView.UserID[]" value="' + user["id"] + '">' +
                                    '</td>' +
                                    '<td>' + user.name + '</td>' +
                                    '<td>' + user.description + '</td>' +
                                    '</tr>'
                                $("tbody").append(tr)
                                var className = ".hengwei_user_group_select-" + a["id"]
                                $(className).hide()
                            }
                        })
                    }
                }
            },
            error:function (data) {
                var d=data.responseJSON
                alert(d.msg)
            }
        })
    })

    $("#hengwei_user_groups-all-checker").on("click", function () {
        var all_checked =  this.checked
        $(".hengwei_user_groups-checker").each(function(){
            this.checked = all_checked
            return true;
        });
        return true;
    });

    $("#hengwei_user_groups-delect_user").on("click",function () {
        bootbox.confirm("确认删除选定信息？", function(result){
            if (!result) {
                return;
            }
            $(".hengwei_user_groups-checker:checked").each(function (i) {
                var id=$(this).attr("key")
                $(this).parent().parent().remove()
                var className=".hengwei_user_group_select-"+id
                $(className).show()
            });
            $("#hengwei_user_groups-all-checker").prop("checked",false);
        })
        return false
    })

});









