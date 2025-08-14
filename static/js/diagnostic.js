document.querySelectorAll('.feedback-item').forEach(item => {
    const minusBtn = item.querySelector('.minus');
    const plusBtn = item.querySelector('.plus');
    const input = item.querySelector('input[type="number"]');
    const min = parseInt(input.min) || 0;

    minusBtn.addEventListener('click', () => {
        let value = parseInt(input.value) || 0;
        if (value > min) {
            input.value = value - 1;
        }
    });

    plusBtn.addEventListener('click', () => {
        let value = parseInt(input.value) || 0;
        input.value = value + 1;
    });
});

// Инпут прятать там сям тоси боси

const searchWrapper = document.querySelector(".search-input");
const searchIcon = searchWrapper.querySelector(".icon-search");
const closeIcon = searchWrapper.querySelector(".icon-close");
const input = searchWrapper.querySelector("input");

searchIcon.addEventListener("click", () => {
    searchWrapper.classList.add("active");
    input.focus();
});

closeIcon.addEventListener("click", () => {
    searchWrapper.classList.remove("active");
    input.value = "";
});

// Посик по списку детей 

const childrenCards = document.querySelectorAll(".diagnostics__children");

// Открыть поиск
searchIcon.addEventListener("click", () => {
    searchWrapper.classList.add("active");
    input.focus();
});

// Закрыть и сбросить
closeIcon.addEventListener("click", () => {
    searchWrapper.classList.remove("active");
    input.value = "";
    showAllChildren();
});

// Фильтрация
input.addEventListener("input", () => {
    const query = input.value.toLowerCase().trim();

    if (query === "") {
        showAllChildren();
        return;
    }

    const searchWords = query.split(/\s+/); // массив слов из запроса

    childrenCards.forEach(card => {
        const firstName = card.querySelector("#name")?.textContent.toLowerCase().trim() || "";
        const middleName = card.querySelector("#fatherland")?.textContent.toLowerCase().trim() || "";
        const lastName = card.querySelector("#second-name")?.textContent.toLowerCase().trim() || "";

        // объединяем всё в одну строку
        const fullText = firstName + " " + middleName + " " + lastName;

        // проверяем, что для каждого слова из запроса fullText содержит его
        const matchesAll = searchWords.every(word => fullText.includes(word));

        if (matchesAll) {
            card.style.display = "grid";
        } else {
            card.style.display = "none";
        }
    });
});

// Показать всех детей
function showAllChildren() {
    childrenCards.forEach(card => {
        card.style.display = "grid";
    });
}