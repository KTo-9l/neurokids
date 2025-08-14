// function moveElement(selector, target) {
//   var homeSelector = selector + '-home';
//   var home = document.querySelector(homeSelector);
//   if (home) {
//     return;
//   }

//   var element = document.querySelector(selector);
//   var oldParent = element.parentElement;
//   var newParent = document.querySelector(target);
//   if (element && newParent) {
//     newParent.appendChild(element);
//     oldParent.classList.add(homeSelector.slice(1));
//   }
// }

// function moveElementBack(selector) {
//   var homeSelector = selector + '-home';
//   var home = document.querySelector(homeSelector);
//   if (!home) {
//     return;
//   }

//   var element = document.querySelector(selector);
//   if (element) {
//     home.appendChild(element);
//     home.classList.remove(homeSelector.slice(1));
//   }
// }

// function handleResize() {
//   var width = window.innerWidth;


//   if (width < 768) {
//     moveElement(".move-1", ".move-2");
//     moveElement(".move-5", ".move-6");
//   } else {
//     moveElementBack(".move-1");
//     moveElementBack(".move-5");
//   }


//   if (width < 993) {
//     moveElement(".move-3", ".move-4");
//   } else {
//     moveElementBack(".move-3");
//   }
// }

// window.addEventListener("resize", handleResize);
// handleResize();



$(document).ready(function () {

  smoothScroll();

  // Animate menu using animate.css
  $('.navbar').addClass('animated bounceInDown');

  // Scroll event listen
  $(window).on('scroll', function () {
    updateNavigation();
  });

  $('.button-container, .switch-container').bind('touchstart mousedown', function (e) {});

  // Update nav selected when click
  $('.courses-nav__item a').on('click', function () {
    console.log("haha");
    $('.nav-item').removeClass('active');
    $(this).parent().addClass('active');
    var width = window.innerWidth;
    if (width < 768) {
      $('.courses-nav').animate({
        scrollLeft: this.getBoundingClientRect().left
      });
    }
  });

  slideSwitch();

});

// Smooth the scroll action
function smoothScroll() {

  $('.courses-nav__item a[href*="#"]:not([href="#"])').click(function () {
    if (location.pathname.replace(/^\//, '') == this.pathname.replace(/^\//, '') && location.hostname == this.hostname) {
      var target = $(this.hash);
      target = target.length ? target : $('[name=' + this.hash.slice(1) + ']');
      if (target.length) {
        $('html, body').animate({
          scrollTop: target.offset().top
        }, 500);
        return false;
      }
    }
  });
}

// Update nav selected
function updateNavigation() {
  var lastId,
    topMenu = $(".navbar"),
    topMenuHeight = topMenu.outerHeight() + 15,
    // All list items
    menuItems = topMenu.find(".courses-nav__item a"),
    // Anchors corresponding to menu items
    scrollItems = menuItems.map(function () {
      var item = $($(this).attr("href"));
      if (item.length) {
        return item;
      }
    });

  // Get container scroll position
  var fromTop = $(this).scrollTop() + topMenuHeight;

  // Get id of current scroll item
  var cur = scrollItems.map(function () {
    if ($(this).offset().top < fromTop)
      return this;
  });
  // Get the id of the current element
  cur = cur[cur.length - 1];
  var id = cur && cur.length ? cur[0].id : "";

  if (lastId !== id) {
    lastId = id;
    // Set/remove active class
    menuItems
      .parent().removeClass("active")
      .end().filter("[href='#" + id + "']").parent().addClass("active");
  }
}

// Update slide switch highlight
function slideSwitch() {
  $('.switch-slide').on('click', function () {

    var activeSpan = $('.switch-toggle-slide .active');

    if (activeSpan.css('left') == '0px') {
      activeSpan.css('left', '50%');
    }

    if (activeSpan.css('left') == '125px') {
      activeSpan.css('left', '0');
    }

    if ($(this).hasClass('active-switch')) {
      $('.switch-slide').addClass('active-switch')
      $(this).removeClass('active-switch');
    } else {
      $('.switch-slide').removeClass('active-switch')
      $(this).addClass('active-switch');
    }
  });
}

$('.tabgroup > div').hide();
$('.tabgroup > div:first-of-type').show();
$('.tabs a').click(function (e) {
  e.preventDefault();
  var $this = $(this),
    tabgroup = '#' + $this.parents('.tabs').data('tabgroup'),
    others = $this.closest('li').siblings().children('a'),
    target = $this.attr('href');
  others.removeClass('active');
  $this.addClass('active');
  $(tabgroup).children('div').hide();
  $(target).show();

});

$('.catalog-menu__title').click(function () {
  $(this).parent().find(".catalog-menu__col").slideToggle();
  $(this).parent().toggleClass('open');
  return false;
});




$('.link-setting').click(function () {
  $(this).parent().find(".list-section").slideToggle();
  $(this).parent().toggleClass('open');
  return false;
});

$('.list-section__bottom .close').click(function () {
  if ($('.list-section').hasClass('open')) {
    $('.list-section').removeClass('open');
  } else {
    $('.list-section').addClass('open');
  }
});

$('.user-link').click(function () {
  if ($('.user-menu').hasClass('open')) {
    $('.user-menu').removeClass('open');
  } else {
    $('.user-menu').addClass('open');
  }
});

$('.bookmark').on('click', function (event) {
  event.preventDefault();
  $(this).toggleClass('active');
});

$('.publish').on('click', function (event) {
  event.preventDefault();
  $(this).toggleClass('active');
});

$('.attention').on('click', function (event) {
  event.preventDefault();
  $(this).toggleClass('delete');
});

$('.photo-slider .owl-carousel').owlCarousel({
  items: 1,
  loop: true,
  margin: 6,
  nav: true,
  dots: true,
  autoHeight: true,
  responsive: {
    0: {
      nav: false,
      dots: true,
      stagePadding: 16,
    },
    768: {
      nav: true,
      dots: true,
    }
  }
});
$('.author-slider .owl-carousel').owlCarousel({
  items: 1,
  loop: true,
  margin: 6,
  nav: true,
  dots: true,
  autoHeight: true,
  responsive: {
    0: {
      nav: false,
      dots: true,
    },
    992: {
      nav: true,
      dots: true,
    }
  }
});



$(function () {
  var owl = $('.cost-slider .owl-carousel');
  owl.owlCarousel({
    items: 1,
    loop: false,
    margin: 10,
    nav: true,
    dots: false,
    autoHeight: true,
    onInitialized: counter, //When the plugin has initialized.
    onTranslated: counter //When the translation of the stage has finished.
  });

  function counter(event) {
    var element = event.target; // DOM element, in this example .owl-carousel
    var items = event.item.count; // Number of items
    var item = event.item.index + 1; // Position of the current item

    // it loop is true then reset counter from 1
    if (item > items) {
      item = item - items
    }
    $('#counter').html(item + " / " + items)
  }
});


$(function () {
  var owl = $('.review-slider .owl-carousel');
  owl.owlCarousel({
    items: 1,
    loop: false,
    margin: 10,
    nav: true,
    dots: false,
    autoHeight: true,
    onInitialized: counter, //When the plugin has initialized.
    onTranslated: counter //When the translation of the stage has finished.
  });

  function counter(event) {
    var element = event.target; // DOM element, in this example .owl-carousel
    var items = event.item.count; // Number of items
    var item = event.item.index + 1; // Position of the current item

    // it loop is true then reset counter from 1
    if (item > items) {
      item = item - items
    }
    $('#counter2').html(item + " / " + items)
  }
});

$('.carousel-section .owl-carousel').owlCarousel({
  nav: true,
  loop: false,
  dots: false,
  pagination: false,
  margin: 15,
  autoHeight: false,
  stagePadding: 1,
  responsive: {
    0: {
      items: 1,
      stagePadding: 0,
    },
    566: {
      items: 2,
    },
    1000: {
      items: 3,
    }
  }
});

if ($('.carousel').hasClass('active'))
  $('div.owl-item').css('opacity', '1')
else $('div.owl-item').css('opacity', '1');


$('.more-text__link').click(function () {
  if ($(this).parent().hasClass('active')) {
    $('.more-text,.more-text__link').parent().removeClass('active');
  } else {
    $(this).parent().addClass('active');
  }
});





$('.menu-link,.message-menu .close').click(function () {
  if ($('.menu-link, .message-menu').hasClass('open')) {
    $('.menu-link,.message-menu').removeClass('open');
  } else {
    $('.menu-link,.message-menu').addClass('open');
  }
});

// Закрытие при клике вне меню
$(document).on('click', function (e) {
  if (!$(e.target).closest('.menu-link, .message-menu').length) {
    $('.menu-link,.message-menu').removeClass('open');
  }
});

$('.add-link,.add-users .close').click(function () {
  if ($('.add-link, .add-users').hasClass('open')) {
    $('.add-link,.add-users').removeClass('open');
  } else {
    $('.add-link,.add-users').addClass('open');
  }
});

$(document).on('click', function (e) {
  if (!$(e.target).closest('.add-link, .add-users').length) {
    $('.add-link, .add-users').removeClass('open');
  }
});

//$('.catalog-link').click(function () {
//  if ($('.catalog-menu,.catalog-link').hasClass('open')) {
//    $('.catalog-menu,.catalog-link').removeClass('open');
//  } else {
//    $('.catalog-menu,.catalog-link').addClass('open');
//  }
//});
//
$(".catalog-link").click(function () {
  $('.catalog-menu').toggle();
});
$(document).on('click', function (e) {
  if (!$(e.target).closest(".catalog-menu,.catalog-link").length) {
    $('.catalog-menu,.catalog-link').removeClass('open');
    $('.catalog-menu').hide();
  }
  e.stopPropagation();
});



$('.clip-link').click(function () {
  if ($('.clip-link,.clip-menu').hasClass('open')) {
    $('.clip-link,.clip-menu').removeClass('open');
  } else {
    $('.clip-link,.clip-menu').addClass('open');
  }
});



$('.search-link,.search-chat .close').click(function () {
  if ($('.search-link,.search-chat,.chat__content').hasClass('open')) {
    $('.search-link,.search-chat,.chat__content').removeClass('open');
  } else {
    $('.search-link,.search-chat,.chat__content').addClass('open');
  }
});

$('.volume').on('click', function (event) {
  event.preventDefault();
  $(this).toggleClass('active');
});



$('.filter-type-link').click(function () {
  if ($('.filter-type-link,.filter').hasClass('open')) {
    $('.filter-type-link,.filter').removeClass('open');
  } else {
    $('.filter-type-link,.filter').addClass('open');
  }
});



$(".custom-select").each(function () {
  var classes = $(this).attr("class"),
    id = $(this).attr("id"),
    name = $(this).attr("name");
  var template = '<div class="' + classes + '">';
  template +=
    '<span class="custom-select-trigger">' +
    $(this).attr("placeholder") +
    "</span>";
  template += '<div class="custom-options">';
  $(this)
    .find("option")
    .each(function () {
      template +=
        '<span class="custom-option ' +
        $(this).attr("class") +
        '" data-value="' +
        $(this).attr("value") +
        '">' +
        $(this).html() +
        "</span>";
    });
  template += "</div></div>";

  $(this).wrap('<div class="custom-select-wrapper"></div>');
  $(this).hide();
  $(this).after(template);
});
$(".custom-option:first-of-type").hover(
  function () {
    $(this)
      .parents(".custom-options")
      .addClass("option-hover");
  },
  function () {
    $(this)
      .parents(".custom-options")
      .removeClass("option-hover");
  }
);
$(".custom-select-trigger").on("click", function () {
  $("html").one("click", function () {
    $(".custom-select").removeClass("opened");
  });
  $(this)
    .parents(".custom-select")
    .toggleClass("opened");
  event.stopPropagation();
});
$(".custom-option").on("click", function () {
  $(this)
    .parents(".custom-select-wrapper")
    .find("select")
    .val($(this).data("value"));
  $(this)
    .parents(".custom-options")
    .find(".custom-option")
    .removeClass("selection");
  $(this).addClass("selection");
  $(this)
    .parents(".custom-select")
    .removeClass("opened");
  $(this)
    .parents(".custom-select")
    .find(".custom-select-trigger")
    .text($(this).text());
});





$('.text-show__link').click(function () {
  if ($('.text-show__link, .text-show p').hasClass('active')) {
    $('.text-show__link, .text-show p').removeClass('active');
  } else {
    $('.text-show__link, .text-show p').addClass('active');
  }
});


$(document).ready(function () {
  $('ul.accordion .opener').click(function () {
    $(this).parent().find("ul:first").slideToggle();
    $(this).parent().toggleClass('active');
    return false;
  });
});

$(document).ready(function () {
  var $inpt = $('.form-group input');

  // for already filled input
  $inpt.each(function () {
    if ($(this).val() !== '') {
      $(this).addClass('filled')
    }
  });

  //for newly filles input
  $inpt.on('change', function () {
    if ($(this).val() !== '') {
      $(this).addClass('filled');
    } else {
      $(this).removeClass('filled');
    }
  })
});

$(document).ready(function () {
  $('.edit-link').click(function () {
    if ($('.form-group--data .form-control').hasClass('edit')) {
      $('.form-group--data .form-control').removeClass('edit');
    } else {
      $('.form-group--data .form-control').addClass('edit');
    }
  });
});


(function () {
  "use strict";
  var jQueryPlugin = (window.jQueryPlugin = function (ident, func) {
    return function (arg) {
      if (this.length > 1) {
        this.each(function () {
          var $this = $(this);

          if (!$this.data(ident)) {
            $this.data(ident, func($this, arg));
          }
        });

        return this;
      } else if (this.length === 1) {
        if (!this.data(ident)) {
          this.data(ident, func(this, arg));
        }

        return this.data(ident);
      }
    };
  });
})();

let lastScrollTop = 0;

function getScrollDirection() {
  let winTop = $(window).scrollTop();
  let direction = 'up';

  if (winTop > lastScrollTop) {
    direction = 'down';
  }

  lastScrollTop = winTop;
  return direction;
}

$(function () {
  $(window).scroll(function () {
    var winTop = $(window).scrollTop();
    let scrollDirection = getScrollDirection();
    let isScrollDown = scrollDirection === 'down';
    if (winTop >= 30 && isScrollDown) {
      $("body").addClass("sticky-header");
    } else {
      $("body").removeClass("sticky-header");
      $("body").removeClass("sticky-up");
    } //if-else

    switch (scrollDirection) {
      case 'down':
        $("body").removeClass("sticky-up");
        break;
      case 'up':
        if (winTop >= 30) {
          $("body").addClass("sticky-up");
        }
        break;
      default:
    }
  }); //win func.
}); //ready func.








$(document).ready(function () {

  $("#datepicker,#datepicker2,#datepicker3,#datepicker4").datepicker({
    buttonText: "ДД.ММ.ГГГГ",
    showOtherMonths: true

  });

  /* Локализация datepicker */
  $.datepicker.regional['ru'] = {
    closeText: 'Закрыть',
    prevText: 'Предыдущий',
    nextText: 'Следующий',
    currentText: 'Сегодня',
    monthNames: ['Январь', 'Февраль', 'Март', 'Апрель', 'Май', 'Июнь', 'Июль', 'Август', 'Сентябрь', 'Октябрь', 'Ноябрь', 'Декабрь'],
    monthNamesShort: ['Янв', 'Фев', 'Мар', 'Апр', 'Май', 'Июн', 'Июл', 'Авг', 'Сен', 'Окт', 'Ноя', 'Дек'],
    dayNames: ['воскресенье', 'понедельник', 'вторник', 'среда', 'четверг', 'пятница', 'суббота'],
    dayNamesShort: ['вск', 'пнд', 'втр', 'срд', 'чтв', 'птн', 'сбт'],
    dayNamesMin: ['Вс', 'Пн', 'Вт', 'Ср', 'Чт', 'Пт', 'Сб'],
    weekHeader: 'Не',
    dateFormat: 'dd.mm.yy',
    firstDay: 1,
    isRTL: false,
    showMonthAfterYear: false,
    yearSuffix: ''
  };
  $.datepicker.setDefaults($.datepicker.regional['ru']);



  const imgUpload = document.getElementById('img-upload');
  const realInput = document.getElementById('real-input');

  imgUpload.addEventListener('click', () => {
    realInput.click();
  });


  function readURL(input) {
    if (input.files && input.files[0]) {
      var reader = new FileReader();

      reader.onload = function (e) {
        imgUpload.setAttribute("src", e.target.result);
      };

      reader.readAsDataURL(input.files[0]);
    }
  }

});
(function () {
  "use strict";

  function Pass_Show_Hide($root) {
    const element = $root;
    const pass_target = $root.first("data-password");
    const pass_elemet = $root.find("[data-pass-target]");
    const pass_show_hide_btn = $root.find("[data-pass-show-hide]");
    const pass_show = $root.find("[data-pass-show]");
    const pass_hide = $root.find("[data-pass-hide]");
    $(pass_hide).hide();
    $(pass_show_hide_btn).click(function () {
      if (pass_elemet.attr("type") === "password") {
        pass_elemet.attr("type", "text");
        $(pass_show).hide();
        $(pass_hide).show();
      } else {
        pass_elemet.attr("type", "password");
        $(pass_hide).hide();
        $(pass_show).show();
      }
    });
  }
  $.fn.Pass_Show_Hide = jQueryPlugin("Pass_Show_Hide", Pass_Show_Hide);
  $("[data-password]").Pass_Show_Hide();
})();






$('[data-fancybox="modal"]').fancybox({
  afterClose: function () {
    $('#show').show()
  }
});

$('#modal-form').on('submit', function () {
  $.fancybox.close();

  return false;
});
