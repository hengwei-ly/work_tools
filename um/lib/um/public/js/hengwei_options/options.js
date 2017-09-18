var id=  $("#TableId>tbody").children("tr").length;

function userFieldEdit(tr) {
    var url=$("#current_user_has_permission").val()
    $.ajax({
        url:url,
        data:{"name":"um_set_field","operation":"edit"},
        success:function (re) {
            if(re.status=="false"){
                alert("你没有该权限")
                return false
            }
            var a=[]
            $(tr).parent().find("input").each(function () {
                a.push($(this).val())
            })
            $("#original_field_name").val(a[1])
            $("[name='field_name']").val(a[1])
            $("[name='field_type']").val(a[2])
            $('#addTextModal').modal('show')
        },
        error:function(re){
            alert("获取权限失败 请稍后再试")
        }
    })
}

function userFieldDelect(tr) {
    var url=$("#current_user_has_permission").val()
    $.ajax({
        url:url,
        data:{"permission":"um_set_field","operation":"del"},
        success:function (re) {
            if(re.status=="false"){
                alert("你没有该权限")
                return false
            }
            var r=confirm("是否确认删除")
            if (r==true) {
                var data=[]
                $(tr).parent().find("input").each(function () {
                    data.push($(this).val())
                })
                data[3]=$(tr).parent().prev().text()
                $($(tr).parent()).parent().remove()
                if (data[0]!=""){
                    var innerForm = $("#field_from");
                    var url=innerForm.attr("action");
                    $.ajax({
                        url:url,
                        data:innerForm.serialize(),
                        success:function (re) {
                            if(re.status=="1"){
                                Msg.success("删除成功")
                            }else{
                                alert("删除失败 \n"+re.msg)
                                var li='<tr >'+
                                    '<td>'+data[1]+'</td>'+
                                    '<td>'+data[3]+'</td>'+
                                    '<td>'+
                                    '<a href="#" class="field_edit"  onclick="userFieldEdit(this)" ><i class="fa fa-edit"></i></a>'+
                                    '<a href="#" class="field_delete" onclick="userFieldDelect(this)" ><i class="fa fa-trash"></i></a>'+
                                    '<input hidden name="fields['+id+'].ID" value="'+data[0]+'">'+
                                    '<input hidden class="all_fields_name" name="fields['+id+'].Name" value="'+data[1]+'">'+
                                    '<input hidden name="fields['+id+'].Type" value="'+data[2]+'">'+
                                    '</td>'+
                                    '</tr>'
                                $("#completed").append(li)
                            }
                        },
                        error:function (a) {
                            alert("未知错误 删除失败")
                            var li='<tr >'+
                                '<td>'+data[1]+'</td>'+
                                '<td>'+data[3]+'</td>'+
                                '<td>'+
                                '<a href="#" class="field_edit"  onclick="userFieldEdit(this)" ><i class="fa fa-edit"></i></a>'+
                                '<a href="#" class="field_delete" onclick="userFieldDelect(this)" ><i class="fa fa-trash"></i></a>'+
                                '<input hidden name="fields['+id+'].ID" value="'+data[0]+'">'+
                                '<input hidden class="all_fields_name" name="fields['+id+'].Name" value="'+data[1]+'">'+
                                '<input hidden name="fields['+id+'].Type" value="'+data[2]+'">'+
                                '</td>'+
                                '</tr>'
                            $("#completed").append(li)
                        }
                    })
                }
            }else{
                return false
            }
        },
        error:function(re){
            alert("获取权限失败 请稍后再试")
        }
    })
}

$(function () {
    $("#btnOk_add").on("click", function () {
      var name =$("[name='field_name']").val()
      var type =$("[name='field_type']").val()
      var txt=$("[name='field_type']").find("option:selected").text();
      var original= $("#original_field_name").val()
      if (name==""){
          alert("字段名称不能为空")
      }
      if (original!=""){
          console.log(name,type,txt)
        $(".all_fields_name").each(function () {
            if ($(this).val() == original){
                $(this).val(name)
                $(this).next().val(type)
                $(this).parent().prev().text(txt)
                $(this).parent().prev().prev().text(name)
                $('#addTextModal').modal('hide')
            }
        })
      }else {
          id=id+1
          var li='<tr >'+
                      '<td>'+name+'</td>'+
                      '<td>'+txt+'</td>'+
                      '<td>'+
                          '<a href="#" class="field_edit"  onclick="userFieldEdit(this)" ><i class="fa fa-edit"></i></a>'+
                          '<a href="#" class="field_delete" onclick="userFieldDelect(this)" ><i class="fa fa-trash"></i></a>'+
                          '<input hidden name="fields['+id+'].ID" value="">'+
                          '<input hidden class="all_fields_name" name="fields['+id+'].Name" value="'+name+'">'+
                          '<input hidden name="fields['+id+'].Type" value="'+type+'">'+
                      '</td>'+
                '</tr>'
          $("#completed").append(li)
          $('#addTextModal').modal('hide')
          $("[name='field_name']").val("")

          // $(".field_delete").on("click", function () {
          //     $($(this).parent()).parent().remove()
          // });
      }
    });

    $("#field_add").on("click", function () {
        $("#original_field_name").val("")
        $("[name='field_name']").val("")
        $('#addTextModal').modal('show')
    });

    // $(".field_delete").on("click", function () {
    //     $($(this).parent()).parent().remove()
    // });

    $("#field_from_submit").on("click",function () {
         var innerForm = $("#field_from");
        var url=innerForm.attr("action");
        $.ajax({
            url:url,
            data:innerForm.serialize(),
            success:function (re) {
                if(re.status=="1"){
                  Msg.success("保存成功")
                }else{
                  alert("保存失败 \n"+re.msg)
                }
            },
            error:function (a) {
              alert("未知错误 保存失败")
            }
        })
    })

});
