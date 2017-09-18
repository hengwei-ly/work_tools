

$(document).ready(function () {
  $.validator.defaults.ignore = ":hidden:not(.force-validate)";
  $("form.validation-required").validate();

  // $('.tagsinput').tagsinput({
  //  tagClass: 'label label-primary'
  // });
});

function adjustSecondMenu($menu) {
  if($menu.length > 0 && $(".fixed-sidebar").hasClass("mini-navbar")) {
    var secTop = $menu.offset().top;
    var secHeight = $menu.height();

    if (secHeight > $("#wrapper").height()) {
      $menu.css("top", -secTop);
    } else {
      if (secTop + secHeight > $("#wrapper").height()) {
        var moveOffset = $("#wrapper").height() - (secTop + secHeight);

        $menu.css("top", parseInt($menu.css("top")) + moveOffset);
      }
    }
  }
}

$(document).ready(function(){
  $("#side-menu").delegate(".nav.nav-second-level", "shown.metisMenu", function(){
    adjustSecondMenu($(this));
  });

  $("#side-menu").delegate(".nav.nav-third-level", "shown.metisMenu", function(){
    adjustSecondMenu($(this).parents(".nav-second-level:first"));
  });

  $("#side-menu > li").mouseenter(function(){
    adjustSecondMenu($(this).find(".nav-second-level"));
  });
});

$(document).ready(function () {
  if($(".fixed-sidebar").hasClass("mini-navbar")) {
    $('.sidebar-collapse').slimScroll({destroy:true});
    $('.sidebar-collapse').css("overflow", "visible");
  } else {
    $('.sidebar-collapse').slimScroll({
      height: '100%',
      railOpacity: 0.9
    });
  }


  function memuWheel(e) {
    var delta = 0;
    if (e.wheelDelta) { delta = -e.wheelDelta/120; }
    if (e.detail) { delta = e.detail / 3; }

    var newTop = parseFloat($('.sidebar-collapse').css("top"))-delta*4;
    if(newTop < 0) {
      

      newTop = 0;
    }

    $('.sidebar-collapse').css("top", newTop)
  }

  function attachWheel(target)
  {
    $('.sidebar-collapse').css("top", 0);

    if (window.addEventListener)
    {
      target.addEventListener('DOMMouseScroll', memuWheel, false );
      target.addEventListener('mousewheel', memuWheel, false );
    }
    else
    {
      document.attachEvent("onmousewheel", memuWheel)
    }
  }

  // attachWheel($('.sidebar-collapse').get(0));
});

function SmoothlyMenu() {
  if (!$('body').hasClass('mini-navbar') || $('body').hasClass('body-small')) {
    // Hide menu in order to smoothly turn on when maximize menu
    $('#side-menu').hide();
    // For smoothly turn on menu
    setTimeout(
        function () {
          $('#side-menu').fadeIn(400, function(){
            $('#side-menu').trigger("fadeInComplete");
          });
        }, 200);
  } else if ($('body').hasClass('fixed-sidebar')) {
    $('#side-menu').hide();
    setTimeout(
        function () {
          $('#side-menu').fadeIn(400, function(){
            $('#side-menu').trigger("fadeInComplete");
          });
        }, 100);
  } else {
    // Remove all inline style from jquery fadeIn function to reset menu state
    $('#side-menu').removeAttr('style');
  }

  if ($('body').hasClass('fixed-sidebar')) {
    if($(".fixed-sidebar").hasClass("mini-navbar")) {
      $('.sidebar-collapse').slimScroll({destroy:true});
      $('.sidebar-collapse').css("overflow", "visible");
    } else {
      $('.sidebar-collapse').slimScroll({
        height: '100%',
        railOpacity: 0.9
      });
    }
  }
}

$(function(){
  window.setTimeout(function(){
    $("#side-menu").focus();
    $("#side-menu").trigger("mouseover");
  }, 1000);

  $('#side-menu').bind("fadeInComplete", function() {
    if($("body").hasClass("mini-navbar")) {
      App.setData("as-sidebar-visible", "false");
    } else {
      $(".nav-second-level, .nav-third-level").css("top", "");
      App.setData("as-sidebar-visible", "true");
    }
  });
})
