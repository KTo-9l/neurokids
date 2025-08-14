const USER_ID = "6873cace5475b657e758f6f3";
const currentPage = window.location.pathname;
let currentChatId = null;

// Хранилище новых чатов с непрочитанными сообщениями
let unreadChatIds = new Set();

const socket = new WebSocket(`ws://${location.host}/openChatWs`);

socket.onopen = () => console.log("WebSocket подключен");

socket.onmessage = (event) => {
    const data = JSON.parse(event.data);

    if (data.event === "addMessage") {
        const msg = data.data;

        const isMessengerOpen = currentPage.includes("message.html");
        const isChatOpen = currentChatId === msg.chatId;

        if (!isChatOpen) {
            unreadChatIds.add(msg.chatId);
            updateNotificationDot();
        }
    }
};

function updateNotificationDot() {
    const dot = document.querySelector(".header-user__message");
    if (!dot) return;

    if (unreadChatIds.size > 0) {
        dot.classList.add("new");
    } else {
        dot.classList.remove("new");
    }
}

// Вызывается при открытии чата
function setCurrentChatId(chatId) {
    currentChatId = chatId;

    // Если пользователь зашел в чат с уведомлением — убираем его
    if (unreadChatIds.has(chatId)) {
        unreadChatIds.delete(chatId);
        updateNotificationDot();
    }
}
