
document.addEventListener("DOMContentLoaded", function () {
    const chatContainer = document.querySelector(".message");
    const contactItems = document.querySelectorAll(".contact-item");
    const backLink = document.querySelector(".back-link");

    contactItems.forEach(item => {
        item.addEventListener("click", function () {
            // открыть контент чата
            chatContainer.classList.add("message--content-open");
        });
    });

    if (backLink) {
        backLink.addEventListener("click", function (e) {
            e.preventDefault();
            // вернуться к списку чатов
            chatContainer.classList.remove("message--content-open");
        });
    }
});
