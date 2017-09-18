function delectPermission(a) {
    $(a).parent().remove()
    var id=$(a).attr("key")
    var claShow=".select_permission_"+id
    var claRem=".selected_permission_"+id
    $(claShow).show()
    $(claRem).remove()
}

function groupDeleteTaboyAllTr(){
    bootbox.confirm("确认删除选定信息？", function(result){
        if (!result) {
            return;
        }
        $(".hengwei_permissions_groups-checker:checked").each(function (i) {
           var calss=".permission_selected_"+$(this).attr("key")
           $(calss).remove()
            $(this).parent().parent().remove()
        });
        $("#hengwei_permissions_groups-all-checker").prop("checked","");
    })
    return false
}

function groupsubmitFrom(){
    var f=document.getElementById("hengwei_permissions_groups-edit");
    var pewimssionids=new Array()
    $(".hengwei_permissions_groups-checker").each(function (i) {
        var inputField = document.createElement("input");
        inputField.type = "hidden";
        inputField.name = "id_list[]";
        inputField.value = $(this).attr("key");
        pewimssionids.push($(this).attr("key"))
        f.appendChild(inputField);
    });
    $("#hengwei_permissions_groups_data").val(pewimssionids)
    document.body.appendChild(f);
    f.submit();
}

$(document).ready(function(){

    if($("#select_permission_tags_value").val().length>0){
        var va = $("#select_permission_tags_value").val().split(",")
        $(".select_permission_tags").val(va)
        $(".select_permission_tags").chosen();
    }else{
        $(".select_permission_tags").chosen();
    }

    $('#jstree1').jstree({
        'core' : {
            multiple:false,
            check_callback: true,
            dblclick_toggle:true
        },
        'plugins' : [ "themes",'types', "contextmenu",'dnd','check_callback'],
        'types' : {
            'default' : {
                'icon' : 'fa fa-folder'
            }
        },
        "contextmenu":{
            "items":{
                "create":null,
                "rename":null,
                "remove":null,
                "copy":null
            }
        }
    });

    $('#jstree1').bind('click.jstree', function(event) {
        var eventNodeName = event.target.nodeName;
        if (eventNodeName == 'A') {
            var idStr =  $(event.target).parents('li').attr('id')
            if (idStr=="group:0"){
                return false
            }
            var url=$("#get_permission_url").val()
            var tag = $(".select_permission_tags").val()
            var ids =[]
            $(".permission_selected").each(function () {
                  ids.push($(this).attr("key"))
            })
            var id=idStr.split(":")[1]
            $.ajax({
                url:url,
                data:{"id":id,"tag":tag.toString(),"ids":ids.toString()},
                success:function (data) {
                    console.log(data)
                    $("#permission_list").html("")
                    if(data==""){
                        alert("该组没有添加权限")
                        return false
                    }
                    $("#permission_list").append(data)

                    $(".permission_select").on("click",function () {
                        var data =$(this).next().val()
                        var datas=data.split(":")
                        var calss=".permission_selected_"+datas[0]
                        if ($(calss).length==0){
                            var li= '<li class="dd-handle permission_selected_'+datas[0]+'"  >'+datas[1]+
                                       '<a class="permission_selected" key="'+datas[0]+'" onclick="delectPermission(this)"><i class="fa fa-arrow-circle-o-left pull-right" ></i></a>'+
                                       '<input hidden type="text" value="'+data+'">'+
                                   '</li>'
                            $("#permission_list_seleceted").append(li)
                            $(this).parent().hide()
                        }
                    })
                },
                error:function (data) {
                    alert("获取权限错误 请稍后再试")
                }
            });
        }
    });

    $("#permission-all-checker").on("click", function () {
        $(".permission_select").each(function () {
            if(!$(this).is(":hidden")) {
                var data =$(this).next().val()
                var datas=data.split(":")
                var calss=".permission_selected_"+datas[0]
                if ($(calss).length==0){
                    var li= '<li class="dd-handle permission_selected_'+datas[0]+'" >'+datas[1]+
                        '<a class="permission_selected" key="'+datas[0]+'" onclick="delectPermission(this)"><i class="fa fa-arrow-circle-o-left pull-right"></i></a>'+
                        '<input hidden type="text" value="'+data+'">'+
                        '</li>'
                    $("#permission_list_seleceted").append(li)
                    $(this).parent().hide()
                }
            }
        })
    });
    
    $("#permission-select-delete-all").on("click", function () {
        $(".permission_selected").each(function () {
            var id=$(this).attr("key")
            var claShow=".select_permission_"+id
            var claRem=".selected_permission_"+id
            $(claShow).show()
            $(claRem).remove()
            $(this).parent().remove()
        })
    });

    $(".hengwei_group_add_permission").on("click",function () {

        $(".permission_selected").each(function () {
            var data=$(this).next().val()
            var datas=data.split(":")
            var tr='<tr>'+
                        '<td>' +
                        '<input type="checkbox" class="hengwei_permissions_groups-checker " key="'+datas[0]+'">' +
                        '<input hidden type="text" name="permissionGroupView.SelectedID[]" value="'+datas[0]+'">'+
                        '<input type="text" hidden class="group-row-data-input" key="'+data+'">' +
                        '</td>'+
                        '<td>'+datas[1]+'</td>'+
                        '<td>'+datas[2]+'</td>'+
                    '</tr>'
            $("tbody").append(tr)
        })
        $("#myModal_add_permission").modal('hide');
    });

    $(".hengwei_group_Import_GroupsPermission").on("click",function () {
        var permsissionsid =new Array()
        var checkedObj = $('input:checkbox[name="selectGroup"]:checked');
        checkedObj.each(function() {
            var isCheck = this.value;
            permsissionsid.push(isCheck)
        });

        if (permsissionsid.length==0){
            Msg.warn("请选择权限组")
            return false
        }
        var url =$("#import_permission_group_url").val()
        $.ajax({
            url:url,
            data:{"id":permsissionsid.toString()},
            success:function (data) {
                checkedObj.each(function() {
                    $(this).prop("checked","")
                });
                var per= data.data
                var pers=[]
                for (var i=0;i<per.length;i++ ){
                    var fond=true
                    $(".hengwei_permissions_groups-checker").each(function () {
                         if (per[i].id == $(this).attr("key")){
                             fond=false
                         }
                    })
                    if(fond){
                        pers.push(per[i])
                    }
                }
                var max=pers.length
                for (var i=0;i<max;i++ ){
                    var a=pers[i];
                    var v =a.id+":"+a.name+":"+a.description
                    var tr='<tr>'+
                                '<td>' +
                                '<input type="checkbox" class="hengwei_permissions_groups-checker" key="'+a.id+'">' +
                                '<input hidden type="text" name="permissionGroupView.SelectedID[]" value="'+a.id+'">'+
                                '<input type="text" hidden class="group-row-data-input" key="'+v+'">' +
                                '</td>'+
                                '<td>'+a.name+'</td>'+
                                '<td>'+a.description+'</td>'+
                            '</tr>'
                    $("tbody").append(tr)

                    var li= '<li class="dd-handle permission_selected_'+a.id+'" >'+a.name+
                        '<a class="permission_selected" key="'+a.id+'" onclick="delectPermission(this)"><i class="fa fa-arrow-circle-o-left pull-right"></i></a>'+
                        '<input hidden type="text" value="'+v+'">'+
                        '</li>'
                    $("#permission_list_seleceted").append(li)
                }
            },
            error:function (data) {
                var d=data.responseJSON
                alert(d.msg)
            }
        })
        $("#myModa_input_group").modal("hide")
    })

    $("#hengwei_permissions_groups-all-checker").on("click", function () {
        var all_checked =  this.checked
        $(".hengwei_permissions_groups-checker").each(function(){
            this.checked = all_checked
            return true;
        });
        return true;
    });

    $("#myModal_add_permission").on('hide.bs.modal', function (e) {
            $("#permission_list").html("")
    })

    $("#myModa_input_group").on('hide.bs.modal', function (e) {
        var checkedObj = $('input:checkbox[name="selectGroup"]:checked');
        checkedObj.each(function() {
            $(this).prop("checked","")
        });
    })

})



