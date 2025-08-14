// user ID укажем вручную
let courseUserId = null;
let currentStage = 0;

// Получаем courseId из URL
function getCourseIdFromURL() {
    const urlParams = new URLSearchParams(window.location.search);
    return urlParams.get("id");
}

const courseId = getCourseIdFromURL();

// Загружаем courseUser при загрузке страницы
async function fetchOrCreateCourseUser() {
    const url = `/getCourseUserByCourseId?courseId=${courseId}`;
    console.log("Запрос на:", url);
    try {
        const response = await fetch(url, {
            method: "GET",
            credentials: 'include'
        });

        const text = await response.text();
        console.log("Ответ:", text);

        if (response.status === 404 || text === "null") {
            console.log("Создание пользователя...");
            // const createRes = await fetch("/createCourseUser", {
            //     method: "POST",
            //     headers: {
            //         "Content-Type": "application/json",
            //     },
            //     body: JSON.stringify({
            //         userId: USER_ID,
            //         courseId: courseId,
            //         purchased: true
            //     }),
            //     credentials: 'include'
            // });
            // const data = await createRes.json();
            // console.log("Создано:", data);
            // courseUserId = data.id;
            // if (data.progress && typeof data.progress.stage === "number") {
            //     currentStage = data.progress.stage;
            // }
        } else {
            const data = JSON.parse(text);
            courseUserId = data.id;
            console.log("Найдено:", data);
            if (data.progress && typeof data.progress.stage === "number") {
                currentStage = data.progress.stage;
            }
        }
    } catch (err) {
        console.error("Ошибка запроса:", err);
    }
}

// Обновление прогресса при открытии урока/материала
function markLessonOpened(lessonNumber, isLast = false) {
    if (!courseUserId) {
        console.warn("courseUserId ещё не установлен, пропуск обновления прогресса");
        return;
    }

    // ✅ Проверка: не обновлять, если стадия уже >= lessonNumber
    if (lessonNumber <= currentStage) {
        console.log(`📌 Урок ${lessonNumber} уже открыт (текущий прогресс: ${currentStage}), не обновляем`);
        return;
    }

    const payload = {
        id: courseUserId,
        purchased: true,
        progress: {
            opened: true,
            stage: lessonNumber,
            finished: isLast
        }
    };

    console.log("🔄 Отправка прогресса курса:", payload);

    fetch("/updateCourseUser", {
        method: "POST",
        headers: {
            "Content-Type": "application/json",
        },
        body: JSON.stringify(payload),
        credentials: 'include'
    })
        .then(res => res.text())
        .then(text => {
            console.log("✅ Ответ от updateCourseUser:", text);
            currentStage = lessonNumber; // обновляем локально stage после успешного запроса
        })
        .catch(err => {
            console.error("❌ Ошибка при отправке прогресса курса:", err);
        });
}





// Отправка прогресса теста
async function markTestFinished(testId, correctCount) {
    try {
        const res = await fetch(`/getTestUserByTestId?testId=${testId}`, {
            method: "GET",
            credentials: 'include',
        });

        if (!res.ok) {
            console.error("Ошибка при получении testUser:", res.status);
            return;
        }

        const data = await res.json();
        const testUserId = data.id;

        await fetch("/updateTestUser", {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
            },
            body: JSON.stringify({
                id: testUserId,
                progress: {
                    opened: true,
                    finished: true,
                    correct: correctCount
                }
            }),
        });
    } catch (err) {
        console.error("Ошибка при обновлении прогресса теста:", err);
    }
}


// Инициализация
fetchOrCreateCourseUser();

function getTotalLessons() {
    return document.querySelectorAll('.tabs[data-tabgroup] li').length;
}

function onLessonClick(li) {
    const submenu = li.querySelector('.submenu');
    if (!submenu) return;

    const isOpen = submenu.style.display === 'block';
    submenu.style.display = isOpen ? 'none' : 'block';

    if (!isOpen) {
        const lessonNum = parseInt(li.dataset.lesson);
        const totalLessons = getTotalLessons();
        const isLastLesson = (lessonNum === totalLessons);

        console.log(`Открыт урок ${lessonNum}, последний: ${isLastLesson}`);

        markLessonOpened(lessonNum, isLastLesson);
    }
}


window.addEventListener("load", () => {
    document.querySelectorAll(".opener").forEach((opener) => {
        opener.addEventListener("click", (event) => {
            event.stopPropagation(); // <-- Это ключевой момент!

            const parentLi = opener.closest("li[data-lesson][data-step]");
            if (parentLi) {
                const lessonNum = parseInt(parentLi.getAttribute("data-lesson"));
                const totalLessons = getTotalLessons();
                const isLastLesson = (lessonNum === totalLessons);

                console.log(`Открыт урок ${lessonNum}, последний: ${isLastLesson}`);

                markLessonOpened(lessonNum, isLastLesson);
            }
        });
    });
    document.querySelectorAll(".checkTest").forEach((button) => {
        button.addEventListener("click", () => {
            const container = button.closest("li");
            const testId = button.dataset.testId;
            const summaryBlock = container.querySelector(".test-summary");

            let total = 0;
            let correct = 0;

            const questionNames = new Set();
            container.querySelectorAll("[name]").forEach(input => {
                questionNames.add(input.name);
            });

            questionNames.forEach(name => {
                const key = `${testId}_${name}`;
                const correctAnswer = correctAnswers[key];

                const questionBlock = container.querySelector(`[name="${name}"]`).closest("li");
                const existingMark = questionBlock.querySelector(".answer-mark");
                if (existingMark) existingMark.remove();

                questionBlock.classList.remove("correct", "incorrect");
                questionBlock.querySelectorAll("label").forEach(lbl => {
                    lbl.classList.remove("correct", "incorrect");
                });

                const field = container.querySelector(`[name="${name}"]`);
                total++;

                if (field && (field.tagName === "TEXTAREA" || field.type === "text")) {
                    const userAnswer = field.value.trim();

                    if (!userAnswer) {
                        questionBlock.classList.add("incorrect");
                        addMark(questionBlock, "❌ Ответ не дан");
                    } else if (
                        (Array.isArray(correctAnswer) && correctAnswer.includes(userAnswer)) ||
                        (userAnswer.toLowerCase() === String(correctAnswer).toLowerCase())
                    ) {
                        correct++;
                        questionBlock.classList.add("correct");
                        addMark(questionBlock, `✔  Верно!`);
                    } else {
                        questionBlock.classList.add("incorrect");
                        addMark(questionBlock, `❌ Правильно: ${correctAnswer}`);
                    }
                } else {
                    const selectorName = name.includes('[') ? name.replace('[', '\\[').replace(']', '\\]') : name;
                    const inputs = container.querySelectorAll(`input[name="${selectorName}"]`);
                    const checkedInputs = [...inputs].filter(i => i.checked);

                    if (checkedInputs.length === 0) {
                        questionBlock.classList.add("incorrect");
                        addMark(questionBlock, "❌ Ответ не дан");
                    } else {
                        const userAnswers = checkedInputs.map(i => i.value);
                        const isMultipleCorrect = Array.isArray(correctAnswer) && correctAnswer.length > 1;

                        let isCorrect = false;
                        if (isMultipleCorrect) {
                            isCorrect = correctAnswer.length === userAnswers.length &&
                                correctAnswer.every(ans => userAnswers.includes(ans));
                        } else {
                            isCorrect = userAnswers.length === 1 &&
                                userAnswers[0].toLowerCase() === String(correctAnswer).toLowerCase();
                        }

                        if (isCorrect) {
                            correct++;
                            checkedInputs.forEach(i => i.closest("label").classList.add("correct"));
                            addMark(questionBlock, `✔  Верно!`);
                        } else {
                            checkedInputs.forEach(i => i.closest("label").classList.add("incorrect"));
                            addMark(questionBlock, `❌ Правильно: ${isMultipleCorrect ? correctAnswer.join(", ") : correctAnswer}`);
                        }
                    }
                }
            });

            summaryBlock.textContent = `Результат: ${correct} из ${total} верно`;
            summaryBlock.style.margin = "1em 0";
            summaryBlock.style.fontWeight = "bold";

            markTestFinished(testId, correct);

            function addMark(el, text) {
                const span = document.createElement("div");
                span.className = "answer-mark";

                if (text.includes("✔")) {
                    span.classList.add("corrects");
                } else if (text.includes("❌")) {
                    span.classList.add("incorrects");
                }

                span.innerText = text;
                el.appendChild(span);
            }
        });
    });
});
