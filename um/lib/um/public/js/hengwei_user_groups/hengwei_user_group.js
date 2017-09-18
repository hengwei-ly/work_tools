
$(document).ready(function(){

    $('#jstree1').jstree({
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
                "ccp":null,
                "添加":{
                    "label": "添加",
                    "action": function(nodedata){
                        var inst = jQuery.jstree.reference(nodedata.reference),
                            obj = inst.get_node(nodedata.reference);
                        var f = document.createElement("form");
                        f.action =$("#HengweiUserGroup_new").val()
                        f.method="GET";
                        var inputFieldid = document.createElement("input");
                        inputFieldid.type = "hidden";
                        inputFieldid.name = "groupId";
                        inputFieldid.value =obj.id.split(":")[1];
                        f.appendChild(inputFieldid);
                        document.body.appendChild(f);
                        f.submit();
                    }
                }
            }
        }
    });

    $('#jstree1').bind('click.jstree', function(event) {
        var eventNodeName = event.target.nodeName;
        if (eventNodeName == 'A') {
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
        }
    });

    $('#jstree1').on('changed.jstree', function (e, data) {
        var i, j, r = [];
        for(i = 0, j = data.selected.length; i < j; i++) {
            r.push(data.instance.get_node(data.selected[i]).text);
        }
        var menu =[{
            "label":"删除",
            "action":function(nodedata){
                var inst = jQuery.jstree.reference(nodedata.reference),
                    obj = inst.get_node(nodedata.reference);
                var ids=obj.id.split(":")
                var f = document.createElement("form");
                f.action =$("#HengweiUserGroup_delete").val()
                f.method="GET";
                var inputFieldid = document.createElement("input");
                inputFieldid.type = "hidden";
                inputFieldid.name = "id";
                inputFieldid.value =ids[1];
                f.appendChild(inputFieldid);
                var inputFieldgroupid = document.createElement("input");
                inputFieldgroupid.type = "hidden";
                inputFieldgroupid.name = "groupId";
                inputFieldgroupid.value =obj.parent.split(":")[1];
                f.appendChild(inputFieldgroupid);
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
                f.action =$("#HengweiUserGroup_edit").val();
                f.method="GET";
                var inputFieldid = document.createElement("input");
                inputFieldid.type = "hidden";
                inputFieldid.name = "id";
                inputFieldid.value =ids[1];
                f.appendChild(inputFieldid);
                document.body.appendChild(f);
                f.submit();
            }
        }]

        //生成右键菜单
        if (data.selected[0] == "group:0"){/*进入权限组管理*/
            $('#jstree1').jstree(true).settings.contextmenu.items["删除"]= false
            $('#jstree1').jstree(true).settings.contextmenu.items["编辑"]= false
        }else if (data.selected[0].split(":")[0].indexOf("group")>-1) {/*单个分组管理*/
            $('#jstree1').jstree(true).settings.contextmenu.items["删除"]= menu[0];
            $('#jstree1').jstree(true).settings.contextmenu.items["编辑"]= menu[1]
        }else if (data.selected[0]=="permissions:0") {/*进入未分组权限管理*/
            $('#jstree1').jstree(true).settings.contextmenu.items["删除"]= false
            $('#jstree1').jstree(true).settings.contextmenu.items["编辑"]= false
        }else if(data.selected[0].split(":")[0].indexOf("permissions")>-1){/*点击位分组的权限*/
            $('#jstree1').jstree(true).settings.contextmenu.items["删除"]= menu[0];
            $('#jstree1').jstree(true).settings.contextmenu.items["编辑"]= menu[1]
        }
    });

});



