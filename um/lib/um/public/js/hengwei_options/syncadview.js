
$(function () {
    $("#sync_ad_submit").on("click",function () {
        var innerForm = $("#sync_ad_from");
        var url=innerForm.attr("action");
        $.ajax({
            url:url,
            data:innerForm.serialize(),
            success:function (re) {
                if(re.status=="1"){
                    Msg.success(re.msg)
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
