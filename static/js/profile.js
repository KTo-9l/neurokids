
const menuItems = document.querySelectorAll(".profile-menu__item");
const panels = document.querySelectorAll(".content-panel");

menuItems.forEach(item => {
    item.addEventListener("click", () => {
        // убрать активный класс у меню
        menuItems.forEach(i => i.classList.remove("profile-menu__item--active"));
        item.classList.add("profile-menu__item--active");

        const target = item.getAttribute("data-target");

        // показать нужную панель
        panels.forEach(panel => {
            if (panel.getAttribute("data-panel") === target) {
                panel.classList.add("content-panel--active");
            } else {
                panel.classList.remove("content-panel--active");
            }
        });
    });
});
document.addEventListener("DOMContentLoaded", () => {

    const urlParams = new URLSearchParams(window.location.search);
    const tab = urlParams.get('tab');
    if (tab) {
        const menuItem = document.querySelector(`.profile-menu__item[data-target="${tab}"]`);
        if (menuItem) menuItem.click();
    }


    const menuItems = document.querySelectorAll(".profile-menu__item");
    const panels = document.querySelectorAll(".content-panel");
    const contentWrapper = document.querySelector(".profile__content");
    const backButtons = document.querySelectorAll(".back-button");

    menuItems.forEach(item => {
        item.addEventListener("click", () => {
            // Активный пункт меню
            menuItems.forEach(i => i.classList.remove("profile-menu__item--active"));
            item.classList.add("profile-menu__item--active");

            const target = item.getAttribute("data-target");

            // Переключить панели
            panels.forEach(panel => {
                if (panel.getAttribute("data-panel") === target) {
                    panel.classList.add("content-panel--active");
                } else {
                    panel.classList.remove("content-panel--active");
                }
            });

            // Если мобилка — показать контент
            if (window.innerWidth <= 600) {
                contentWrapper.classList.add("active-mobile");
                document.querySelector(".profile__side").style.display = "none";
            }
        });
    });

    // Назад — возвращаем меню
    backButtons.forEach(btn => {
        btn.addEventListener("click", () => {
            contentWrapper.classList.remove("active-mobile");
            document.querySelector(".profile__side").style.display = "block";
        });
    });
});