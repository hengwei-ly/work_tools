$(document).ready(function(){
    $('#jstree').jstree({
        'core' : {
            multiple:false,
            check_callback: true,
            dblclick_toggle:false,
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

    //节点点击事件
    $('#jstree').bind('click.jstree', function(event) {
        var eventNodeName = event.target.nodeName;
        if (eventNodeName == 'A') {
            if($(event.target).parents('li').attr('id').split(":")[0]!="tags") {
                var groupId = $(event.target).parents('li').attr('id').split(":")[1]
                var f = document.createElement("form");
                f.action = $("#PermissionGroup_index_url").val()
                f.method = "GET";
                var inputField = document.createElement("input");
                inputField.type = "hidden";
                inputField.name = "groupId";
                inputField.value = groupId;
                f.appendChild(inputField);
                document.body.appendChild(f);
                f.submit()
            }
        }
    });

    $('#jstree').on('changed.jstree', function (e, data) {
        var menu =[{
            "label": "添加",
            "action": function(nodedata){
                var inst = jQuery.jstree.reference(nodedata.reference),
                    obj = inst.get_node(nodedata.reference);
                var f = document.createElement("form");
                f.action =$("#PermissionGroup_new_url").val()
                f.method="GET";
                var inputFieldid = document.createElement("input");
                inputFieldid.type = "hidden";
                inputFieldid.name = "groupId";
                inputFieldid.value =obj.id.split(":")[1];
                f.appendChild(inputFieldid);
                document.body.appendChild(f);
                f.submit();
            }
        },{
            "label":"删除",
            "action":function(nodedata){
                var inst = jQuery.jstree.reference(nodedata.reference),
                    obj = inst.get_node(nodedata.reference);
                var ids=obj.id.split(":")
                var f = document.createElement("form");
                f.action =$("#PermissionGroup_delete_url").val()
                f.method="GET";
                var inputFieldid = document.createElement("input");
                inputFieldid.type = "hidden";
                inputFieldid.name = "id";
                inputFieldid.value =ids[1];
                f.appendChild(inputFieldid);
                var inputFieldgroupId = document.createElement("input");
                inputFieldgroupId.type = "hidden";
                inputFieldgroupId.name = "groupId";
                inputFieldgroupId.value =obj.parent.split(":")[1];
                f.appendChild(inputFieldgroupId);
                document.body.appendChild(f);
                f.submit();
            }
        },{
            "label":"编辑",
            "action":function(nodedata){
                var inst = jQuery.jstree.reference(nodedata.reference),
                    obj = inst.get_node(nodedata.reference);
                var ids=obj.id.split(":")
                var f = document.createElement("form");
                f.action =$("#PermissionGroup_edit_url").val();
                f.method="GET";
                var inputFieldid = document.createElement("input");
                inputFieldid.type = "hidden";
                inputFieldid.name = "id";
                inputFieldid.value =ids[1];
                var inputFieldgroupId = document.createElement("input");
                inputFieldgroupId.type = "hidden";
                inputFieldgroupId.name = "groupId";
                inputFieldgroupId.value =obj.parent.split(":")[1]
                f.appendChild(inputFieldid);
                f.appendChild(inputFieldgroupId);
                document.body.appendChild(f);
                f.submit();
            }
        },{
            "label":"查看",
            "action":function(nodedata){
                alert("okok")
            }
        },{
            "label": "复制",
            "action": function(nodedata){
                $("#myModal_Copy").modal("show")
                var inst = jQuery.jstree.reference(nodedata.reference),
                    obj = inst.get_node(nodedata.reference);
                $("#permission_group_copy_id").val(obj.id.split(":")[1])
            }
        }]
        //生成右键菜单
        if (data.selected[0] == "group:0"){/*进入权限组管理*/
            $('#jstree').jstree(true).settings.contextmenu.items["添加"]= menu[0]
            $('#jstree').jstree(true).settings.contextmenu.items["删除"]= false
            $('#jstree').jstree(true).settings.contextmenu.items["编辑"]= false
            $('#jstree').jstree(true).settings.contextmenu.items["复制"]= false;
        }else if (data.selected[0].split(":")[0].indexOf("group")>-1) {/*单个分组管理*/
            $('#jstree').jstree(true).settings.contextmenu.items["添加"]= menu[0];
            $('#jstree').jstree(true).settings.contextmenu.items["删除"]= menu[1];
            $('#jstree').jstree(true).settings.contextmenu.items["编辑"]= menu[2];
            $('#jstree').jstree(true).settings.contextmenu.items["复制"]= false;
        }else if(data.selected[0].split(":")[0].indexOf("Default")>-1){/*点击位分组的权限*/
            $('#jstree').jstree(true).settings.contextmenu.items["添加"]= false;
            $('#jstree').jstree(true).settings.contextmenu.items["删除"]= false;
            $('#jstree').jstree(true).settings.contextmenu.items["编辑"]= false;
            $('#jstree').jstree(true).settings.contextmenu.items["复制"]= menu[4];
        }
        else if(data.selected[0].split(":")[0].indexOf("tags")>-1){/*点击位分组的权限*/
            $('#jstree').jstree(true).settings.contextmenu.items["添加"]= false;
            $('#jstree').jstree(true).settings.contextmenu.items["删除"]= false;
            $('#jstree').jstree(true).settings.contextmenu.items["编辑"]= false;
            $('#jstree').jstree(true).settings.contextmenu.items["复制"]= false;
        }
    });

    //双击事件
    $('#jstree').bind('dblclick.jstree', function(event) {
        if($(event.target).parents('li').attr('id').split(":")[0]=="group") {
            var groupId = $(event.target).parents('li').attr('id').split(":")[1]
            var f = document.createElement("form");
            f.action = event.currentTarget.baseURI.split("&")[0];
            f.method = "GET";
            var inputField = document.createElement("input");
            inputField.type = "hidden";
            inputField.name = "groupId";
            inputField.value = groupId;
            f.appendChild(inputField);
            document.body.appendChild(f);
            f.submit()
        }
    });

    $("#hengwei_permission_group_copy").on("click",function () {
        var groupName=$("[name='newGroupName']").val()
        var parentid=$('input:radio[name="name"]:checked').val()
        var id=$("#permission_group_copy_id").val()
        var f = document.createElement("form");
        f.action =$("#Copy_DefaultPermissionGroups_url").val()
        f.method="GET";
        var inputFieldid = document.createElement("input");
        inputFieldid.type = "hidden";
        inputFieldid.name = "id";
        inputFieldid.value =id;
        f.appendChild(inputFieldid);
        var inputFieldid1 = document.createElement("input");
        inputFieldid1.type = "hidden";
        inputFieldid1.name = "name";
        inputFieldid1.value =groupName;
        f.appendChild(inputFieldid1);
        var inputFieldid2 = document.createElement("input");
        inputFieldid2.type = "hidden";
        inputFieldid2.name = "parentId";
        inputFieldid2.value =parentid;
        f.appendChild(inputFieldid2);
        document.body.appendChild(f);
        f.submit();
    })

});