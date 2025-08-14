const form = document.getElementById('password-form');
const newPassword = document.getElementById('new-password');
const confirmPassword = document.getElementById('confirm-password');
const errorMessage = document.getElementById('error-message');

form.addEventListener('submit', function (event) {
    event.preventDefault(); // Отменяем стандартную отправку формы

    if (newPassword.value !== confirmPassword.value) {
        errorMessage.textContent = 'Новый пароль и подтверждение не совпадают!';
        errorMessage.style.display = 'block';
        return; // Прекращаем отправку
    }

    errorMessage.style.display = 'none';

    // Если нужно отправить данные через fetch/ajax, пример:
    const formData = new URLSearchParams(new FormData(form));

    // Пример отправки (подкорректируй URL и метод под свой бекенд)
    fetch('/updatePassword', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/x-www-form-urlencoded;charset=UTF-8'
        },
        body: formData.toString()
    })
        .then(res => {
            if (res.ok) {
                errorMessage.style.color = 'green';
                errorMessage.textContent = 'Пароль успешно изменён';
                errorMessage.style.display = 'block';
                form.reset();
            } else {
                errorMessage.style.color = 'red';
                errorMessage.textContent = 'Ошибка при изменении пароля';
                errorMessage.style.display = 'block';
            }
        })
        .catch(() => {
            errorMessage.style.color = 'red';
            errorMessage.textContent = 'Ошибка сети';
            errorMessage.style.display = 'block';
        });


    // Если же форма отправляется традиционно, вместо fetch вызови form.submit() после проверки.
});

const formN = document.getElementById('profile-form');
const formMessage = document.getElementById('form-message');

formN.addEventListener('submit', async (event) => {
    event.preventDefault(); // отменяем стандартную отправку

    // собираем данные из формы
    const formData = new FormData(formN);
    const urlEncodedData = new URLSearchParams(formData).toString();

    try {
        const response = await fetch('/updateInfo', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/x-www-form-urlencoded;charset=UTF-8'
            },
            body: urlEncodedData
        });

        if (response.ok) {
            formMessage.style.color = 'green';
            formMessage.textContent = 'Данные успешно обновлены';
            formMessage.style.display = 'block';
        } else {
            formMessage.style.color = 'red';
            formMessage.textContent = 'Ошибка при обновлении данных';
            formMessage.style.display = 'block';
        }
    } catch (error) {
        formMessage.style.color = 'red';
        formMessage.textContent = 'Ошибка сети';
        formMessage.style.display = 'block';
    }
});