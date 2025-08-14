
let chatContent = null;
let currentChatId = null;
let lastDate = null;

let messageWithFiles = false;
let filesToChange = null;

// Множество ID чатов с непрочитанными сообщениями
const unreadChatIds = new Set();

const socket = new WebSocket(`ws://${location.host}/openChatWs`);

socket.onopen = () => {
    console.log('WebSocket connected');
};

socket.onerror = (err) => {
    console.error("WebSocket error:", err);
};

socket.onmessage = (event) => {
    const data = JSON.parse(event.data);
    console.log("WebSocket message:", data);

    switch (data.event) {
        case "addMessage":
            const msg = data.data;
            if (data.from === currentUser.Id && msg.attachments !== null) {
                updateFilesPath(msg.attachments, msg.id);
            }
            if (msg.chatId === currentChatId) {
                addMessageToChat(msg);
                chatContent.scrollTop = chatContent.scrollHeight;
                removeNotificationDot(msg.chatId);
                unreadChatIds.delete(msg.chatId);
            } else {
                unreadChatIds.add(msg.chatId);
                addNotificationDot(msg.chatId);
            }
            updateNotificationDotInHeader();
            addMessageToList(msg);
            break;

        case "createChat":
            const chat = data.data;
            allChats.unshift(chat);

            let isGroup = chat.isGroup;
            let title = isGroup ? chat.title : getOtherUserName(chat.members);

            addChatToList(chat, title);

            // Открываем чат только если это создал текущий пользователь
            if (data.from === currentUser.Id) {
                openChat(chat.id, title);
            }

            // Сброс формы, если модалка открыта (только для группового чата)
            if (createGroupDialog.open) {
                createGroupDialog.close();
                document.getElementById("groupTitle").value = "";
                searchInput.value = "";
                searchResults.innerHTML = "";
                selectedUsers.innerHTML = "";
                selectedUserIds = [];
            }
            break;

        case "updateChat":
            const updatedChat = data.data;
            // Обновить allChats
            const index = allChats.findIndex(c => c.id === updatedChat.id);
            if (index !== -1) {
                allChats[index] = updatedChat;
            }

            if (addUsersModal.open) {
                addUsersModal.close();
                addSelectedUserIds = [];
                addSelectedUsers.innerHTML = '';
                addUserSearch.value = '';
                addUserSearchResults.innerHTML = '';
                if (data.from === currentUser.Id) {
                    openChat(chat.id, title);
                }
            }
            break;

        default:
            console.log("Unknown event");
            break;
    }
};

// Получить имя второго участника
function getOtherUserName(members) {
    const otherId = members.find(id => id !== currentUser.Id);
    const user = findUserById(otherId);
    return user ? user.Name : "Чат";
}


let userId = null;
let currentUser = null;
let allUsers = [];
let otherUsers = [];
let allChats = [];


function initUserData() {
    fetch('/getMe')
        .then(res => res.json())
        .then(me => {
            currentUser = me;
            userId = me.Id;
            return fetch('/getAllUsers');
        })
        .then(res => res.json())
        .then(users => {
            allUsers = users;
            otherUsers = users.filter(u => u.Id !== currentUser.Id);
            console.log("Текущий пользователь:", currentUser);
            console.log("Текущий пользователь:", userId);
            console.log("Все пользователи:", allUsers);
            console.log("Все пользователи кроме меня:", otherUsers);
            return fetch('/getChats'); // <--- получаем все чаты
        })
        .then(res => res.json())
        .then(Chats => {
            if (Chats === null) {
                allChats = [];
                console.log("Нет чатов");
            } else {
                allChats = Chats;
                console.log("Все чаты:", allChats);
            }
            // .filter(chat =>
            //     chat.members && chat.members.some(m => m.Id === currentUser.Id)
            // );
            setupUserSearch(); // инициализация поиска
        })
        .catch(console.error);
}

function setupUserSearch() {
    const input = document.querySelector('.message__search input');
    const resultBox = document.createElement('div');
    resultBox.className = 'search-user-results';
    input.parentNode.style.position = 'relative';
    input.parentNode.appendChild(resultBox);

    input.addEventListener('input', () => {
        const query = input.value.toLowerCase().trim();
        resultBox.innerHTML = '';
        if (!query) {
            resultBox.style.display = 'none';
            return;
        }

        const results = otherUsers.filter(user =>
            user.Name.toLowerCase().includes(query)
        );

        results.forEach(user => {
            const item = document.createElement('div');
            item.textContent = user.Name;
            item.addEventListener('click', () => {
                createPrivateChat(user);
                resultBox.innerHTML = '';
                input.value = '';
                resultBox.style.display = 'none';
            });
            resultBox.appendChild(item);
        });

        resultBox.style.display = results.length ? 'block' : 'none';
    });

    document.addEventListener('click', (e) => {
        if (!e.target.closest('.message__search')) {
            resultBox.style.display = 'none';
        }
    });
}


function createPrivateChat(otherUser) {
    const existingChat = allChats.find(chat => {
        return (
            !chat.isGroup &&
            chat.members.includes(currentUser.Id) &&
            chat.members.includes(otherUser.Id) &&
            chat.members.length === 2
        );
    });

    if (existingChat) {
        // Уже существует — просто открыть
        openChat(existingChat.id, getOtherUserName(existingChat.members));
        return;
    }

    const title = `${currentUser.Name}&${otherUser.Name}`;
    const members = [currentUser.Id, otherUser.Id];

    const message = {
        event: "createChat",
        from: currentUser.Id,  // добавляем from
        to: members,           // добавляем to — все участники
        data: {
            title,
            members,
            isGroup: false
        }
    };

    socket.send(JSON.stringify(message));
}

function addChatToList(chat, nameOverride) {
    const listContainer = document.querySelector('.message__contacts .mCSB_container');

    const item = document.createElement('div');
    item.className = 'contact-item';

    const nameToShow = nameOverride || chat.title || "Новая беседа";

    item.innerHTML = `
        <div class="contact-item__body" data-chat-id="${chat.id}">
            <div class="contact-item__number">${chat.members?.length || 2} участника</div>
            <div class="contact-item__name">${nameToShow}</div>
            <div class="contact-item__text"></div>
        </div>
    `;

    listContainer.prepend(item);

    const chatBody = item.querySelector('.contact-item__body');
    chatBody.addEventListener('click', () => {
        openChat(chat.id, nameToShow);
    });
}




function updateNotificationDotInHeader() {
    const headerDot = document.querySelector(".header-user__message");
    if (!headerDot) return;

    if (unreadChatIds.size > 0) {
        headerDot.classList.add("new");
    } else {
        headerDot.classList.remove("new");
    }
}

function formatDateTime(isoString) {
    const date = new Date(isoString);
    return date.toLocaleTimeString([], { hour: '2-digit', minute: '2-digit' });
}

function openChat(chatId) {
    currentChatId = chatId;

    removeNotificationDot(chatId);
    unreadChatIds.delete(chatId);
    updateNotificationDotInHeader();

    fetch(`/getMessages?chatId=${chatId}`)
        .then(res => res.json())
        .then(messages => {
            renderChatWindow(chatId, messages);
        })
        .catch(err => {
            console.error("Ошибка загрузки чата:", err);
        });
}

let attachedFiles = [];
let attachedMedia = [];
function renderChatWindow(chatId, messages, chatTitle) {
    const container = document.querySelector(".message__content");
    container.innerHTML = "";

    // Заголовок
    const headerTemplate = document.getElementById("chat-header-template");
    const headerClone = headerTemplate.content.cloneNode(true);

    console.log(findChatById(chatId).isGroup);

    if (!findChatById(chatId).isGroup) {
        const newArray = findChatById(chatId).members.filter(item => item !== userId)
        headerClone.querySelector(".chat-title").textContent = findUserById(newArray[0]).Name;
    } else {
        headerClone.querySelector(".chat-title").textContent = findChatById(chatId).title || "Чат";
    }

    container.appendChild(headerClone);

    const chaters = findChatById(chatId);
    const addLink = container.querySelector('.add-link');
    const menuLink = container.querySelector('.menu-link');

    if (chaters && !chaters.isGroup) {
        if (addLink) addLink.style.display = 'none';
        if (menuLink) menuLink.style.display = 'none';
    } else {
        if (addLink) addLink.style.display = '';
        if (menuLink) menuLink.style.display = '';
    }

    // Контейнер чата
    const chat = document.createElement("div");
    chat.classList.add("chat");
    container.appendChild(chat);

    chatContent = document.createElement("div");
    chatContent.classList.add("chat__content", "mCustomScrollbar");
    chatContent.id = "content2";
    chatContent.setAttribute("data-mcs-theme", "dark");
    chat.appendChild(chatContent);

    if (!Array.isArray(messages)) {
        messages = [];
    }
    // Добавляем сообщения
    lastDate = null; // сбрасываем предыдущую дату
    messages.reverse().forEach(msg => {
        addMessageToChat(msg);
    });

    chatContent.scrollTop = chatContent.scrollHeight;

    // Поле ввода
    const inputTemplate = document.getElementById("chat-input-template");
    const inputClone = inputTemplate.content.cloneNode(true);
    container.appendChild(inputClone);

    const from = userId;

    const sendBtn = container.querySelector("#send");
    const inputField = container.querySelector("#msg-input");

    const attachMedia = container.querySelector("#attach-media");
    attachedMedia = [];
    attachMedia.onclick = (event) => {
        event.preventDefault();
        let inputMedia = document.createElement('input');
        inputMedia.type = 'file';
        inputMedia.accept = 'image/*,video/*';
        inputMedia.multiple = true;
        inputMedia.style.display = 'none';

        inputMedia.click();
        attachMedia.appendChild(inputMedia);
        inputMedia.addEventListener('change', () => {
            attachedMedia = Array.from(inputMedia.files);
            renderFilePreview()
        })
    }

    const attachFiles = container.querySelector("#attach-files");
    attachedFiles = [];
    attachFiles.onclick = (event) => {
        event.preventDefault();
        let inputFiles = document.createElement('input');
        inputFiles.type = 'file';
        inputFiles.multiple = true;
        inputFiles.style.display = 'none';

        inputFiles.click();
        attachFiles.appendChild(inputFiles);
        inputFiles.addEventListener('change', () => {
            attachedFiles = Array.from(inputFiles.files);
            renderFilePreview()
        })
    }

    sendBtn.onclick = () => {
        sendMessageBtn();
    };

    inputField.addEventListener("keydown", (e) => {
        if (e.key === "Enter" && !e.shiftKey) {
            e.preventDefault();
            sendMessageBtn();
        }
    });

    async function sendMessageBtn() {
        const attachments = new FormData();

        attachedMedia.forEach(oneMedia => {
            attachments.append('files', oneMedia)
        });
        attachedMedia = []

        attachedFiles.forEach(oneFile => {
            attachments.append('files', oneFile)
        });
        attachedFiles = []

        if (Array.from(attachments.entries()).length > 0) {
            await sendTmpFiles(attachments);
            messageWithFiles = true;
        } else {
            messageWithFiles = false;
            filesToChange = null;
        }


        let text = inputField.value.trim();

        if (!messageWithFiles && !text) {
            return;
        }
        if (!text) text = null;

        sendMessageWebSocket(chatId, from, text);
        inputField.value = "";
        document.getElementById('file-preview').innerHTML = '';
    }

    setupToggles();
    setupCloseButtons();
    setupMessageSearch();

    const addLinkBtn = container.querySelector('.add-link');
    if (addLinkBtn) {
        addLinkBtn.addEventListener('click', () => {
            setupAddUsersModal(chatId);
        });
    }

    const deleteChatBtn = container.querySelector('.delete-chat-btn');
    const leaveGroupBtn = container.querySelector('.leave-group-btn');

    

    if (deleteChatBtn) {
        deleteChatBtn.addEventListener('click', (e) => {
            e.preventDefault();
            leaveChat(chatId, currentUser.Id);
        });
    }

    if (leaveGroupBtn) {
        leaveGroupBtn.addEventListener('click', (e) => {
            e.preventDefault();
            leaveChat(chatId, currentUser.Id);
        });
    }
    
}

function renderFilePreview() {
    const filePreviewContainer = document.getElementById('file-preview');
    filePreviewContainer.innerHTML = "";

    [...attachedMedia, ...attachedFiles].forEach((file, index) => {
        const fileItem = document.createElement("div");
        fileItem.className = "file-preview__item";

        // Если это изображение — отрисовать как превью
        if (file.type.startsWith("image/")) {
            const img = document.createElement("img");
            img.src = URL.createObjectURL(file);
            img.style.maxHeight = "100px";
            img.style.marginRight = "10px";

            fileItem.appendChild(img);
        }

        const fileInfo = document.createElement("span");
        fileInfo.textContent = `${file.name} (${Math.ceil(file.size / 1024)} KB)`;
        fileItem.appendChild(fileInfo);

        const removeBtn = document.createElement("button");
        removeBtn.textContent = "✕";
        removeBtn.className = "file-preview__remove";
        removeBtn.onclick = () => {
            if (attachedMedia.includes(file)) {
                attachedMedia = attachedMedia.filter(f => f !== file);
            } else {
                attachedFiles = attachedFiles.filter(f => f !== file);
            }
            renderFilePreview();
        };
        fileItem.appendChild(removeBtn);

        filePreviewContainer.appendChild(fileItem);
    });
}


async function sendTmpFiles(files) {
    files.append('path', currentChatId);
    files.append('path', 'tmp');
    filesToChange = await uploadFiles(files);
}

function uploadFiles(formData) {
    return fetch("/uploadFiles", {
        method: 'POST',
        body: formData
    }).then(response => {
        if (!response.ok) {
            console.error("error uploading files");
            return null
        }
        return response.json();
    })
}

function updateFilesPath(filesIds, messageId) {
    filesIds.forEach(async fileId => {
        fileInfo = await getFileInfoById(fileId);
        filePath = fileInfo.Path;

        let updatedFileFormData = new FormData();

        updatedFileFormData.append('id', fileId);

        updatedFileFormData.append('path', filePath[0]);
        updatedFileFormData.append('path', messageId);
        updatedFileFormData.append('path', fileInfo.Filename);

        ok = updateFilePathById(updatedFileFormData);
    })
}

function updateFilePathById(formData) {
    return fetch("/updateFilePathById", {
        method: 'POST',
        body: formData
    }).then(response => {
        if (!response.ok) {
            console.error("error update file path by id");
            return false;
        }
        return true;
    })
}

function getFileInfoById(fileId) {
    return fetch(`/getFileInfoById?id=${fileId}`, {
    }).then(response => {
        if (!response.ok) {
            log.error("error getFileInfoById");
            return null;
        }
        return response.json();
    })
}

function getFileById(fileId) {
    return fetch(`/getFileById?id=${fileId}`, {
    }).then(response => {
        if (!response.ok) {
            log.error("error getFileById");
            return null;
        }
        return response.blob();
    }) // .then(blob => {}) здесь могут быть реализованы любые действия с потоком байтов из файла (я octet-stream в ответе бахаю и сам файл)
}

async function addMessageToChat(msg) {
    const messageTemplate = document.getElementById("chat-message-template");

    const msgDate = new Date(msg.time);
    const msgDay = msgDate.toDateString();

    if (lastDate !== msgDay) {
        lastDate = msgDay;

        const dateSpan = document.createElement("span");
        dateSpan.className = "chat__date";

        const options = { day: 'numeric', month: 'long', year: 'numeric' };
        dateSpan.textContent = msgDate.toLocaleDateString('ru-RU', options);

        chatContent.appendChild(dateSpan);
    }

    const msgClone = messageTemplate.content.cloneNode(true);
    msgClone.querySelector(".chat-item__name").innerHTML = `${findUserById(msg.from).Name}<span>${formatDateTime(msg.time)}</span>`;
    const messageText = msgClone.querySelector(".message-text");
    messageText.innerHTML = ""; // очищаем, если шаблон содержит текст по умолчанию

    // Добавляем текст, если он есть
    if (msg.text && msg.text.trim() !== "") {
        const textPara = document.createElement("p");
        textPara.textContent = msg.text;
        messageText.appendChild(textPara);
    }

    // === 🔥 Обработка вложений ===
    if (msg.attachments && msg.attachments.length > 0) {
        const filesContainer = document.createElement("div");
        filesContainer.className = "message-files";

        for (const fileId of msg.attachments) {
            try {
                const fileInfo = await getFileInfoById(fileId);
                if (!fileInfo) continue;

                const fileLink = document.createElement("a");
                fileLink.href = `/getFileById?id=${fileId}`;
                fileLink.download = fileInfo.Filename;
                fileLink.className = "file-link";
                fileLink.target = "_blank";
                fileLink.rel = "noopener noreferrer";

                const fileSizeKb = Math.ceil(fileInfo.Length / 1024);
                fileLink.textContent = `${fileInfo.Filename} — ${fileSizeKb} KB`;

                filesContainer.appendChild(fileLink);
            } catch (err) {
                console.error("Ошибка при получении информации о файле:", err);
            }
        }

        messageText.appendChild(filesContainer);
    }


    chatContent.appendChild(msgClone);
}


function findUserById(id) {
    return allUsers.find(user => user.Id === id);
}

function findChatById(id) {
    return allChats.find(chat => chat.id === id);
}





function addMessageToList(msg) {
    const chatElem = document.querySelector(`.contact-item__body[data-chat-id="${msg.chatId}"]`);
    if (!chatElem) {
        console.log("Чат в списке не найден для chatId:", msg.chatId);
        return;
    }

    const lastMsgElem = chatElem.querySelector('.contact-item__text');
    if (lastMsgElem) {
        lastMsgElem.textContent = msg.text;
    }
}



function sendMessageWebSocket(chatId, from, text) {
    const message = {
        event: "addMessage",
        from: from,
        to: findChatById(chatId).members,
        data: {
            chatId: chatId,
            from: from,
            text: text,
            time: new Date().toISOString()
        }
    };

    if (messageWithFiles) {
        message.data.attachments = filesToChange.map(item => item.id);
    }
    socket.send(JSON.stringify(message));
}

function addNotificationDot(chatId) {
    const chatElem = document.querySelector(`.contact-item__body[data-chat-id="${chatId}"]`);
    if (!chatElem) return;

    if (!chatElem.querySelector(".new-message-dot")) {
        const dot = document.createElement("span");
        dot.className = "new-message-dot";
        chatElem.appendChild(dot);
    }
}

function removeNotificationDot(chatId) {
    const chatElem = document.querySelector(`.contact-item__body[data-chat-id="${chatId}"]`);
    if (!chatElem) return;

    const dot = chatElem.querySelector(".new-message-dot");
    if (dot) dot.remove();
}

// --- Управление меню и поиска ---

function setupToggles() {
    setupToggle('.search-link', ['.search-link', '.search-chat', '.chat__content']);
    setupToggle('.search-chat .close', ['.search-link', '.search-chat', '.chat__content']);
    setupToggle('.add-link', ['.add-link', '.add-users']);
    setupToggle('.add-users .close', ['.add-link', '.add-users']);
    setupToggle('.menu-link', ['.menu-link', '.message-menu']);
    setupToggle('.message-menu .close', ['.menu-link', '.message-menu']);
    setupToggle('.clip-link', ['.clip-link', '.clip-menu']);
    setupOutsideClickClose(['.menu-link', '.message-menu']);
    setupOutsideClickClose(['.clip-link', '.clip-menu']);
    setupOutsideClickClose(['.search-link', '.search-chat', '.chat__content']);
    setupOutsideClickClose(['.add-link', '.add-users']);
}

function setupToggle(triggerSelector, toggleSelectors) {
    const trigger = document.querySelector(triggerSelector);
    if (!trigger) return;

    trigger.addEventListener('click', (e) => {
        e.stopPropagation();
        const isOpen = toggleSelectors.some(sel => {
            const el = document.querySelector(sel);
            return el?.classList.contains('open');
        });

        toggleSelectors.forEach(sel => {
            const el = document.querySelector(sel);
            if (el) {
                el.classList.toggle('open', !isOpen);
            }
        });
    });
}

function setupOutsideClickClose(selectors) {
    document.addEventListener('click', (e) => {
        const clickedInside = selectors.some(sel => e.target.closest(sel));
        if (!clickedInside) {
            selectors.forEach(sel => {
                const el = document.querySelector(sel);
                if (el) el.classList.remove('open');
            });

            if (selectors.includes('.search-chat')) {
                const input = document.querySelector('.search-input input');
                if (input) {
                    input.value = '';
                    const messages = document.querySelectorAll('.chat-item');
                    messages.forEach(msg => msg.style.display = '');
                }
            }
        }
    });
}

function setupCloseButtons() {
    document.querySelectorAll('a.close').forEach(closeBtnGroupDialog => {
        closeBtnGroupDialog.addEventListener('click', () => {
            let parent = closeBtnGroupDialog.closest('.open');
            if (parent) {
                parent.classList.remove('open');
            }
        });
    });
}

function setupMessageSearch() {
    const input = document.querySelector('.search-input input');
    const cancelBtn = document.querySelector('.search-chat__cancel');
    const clearBtn = document.querySelector('.search-chat__clear');

    if (!input) return;

    const filterMessages = () => {
        const search = input.value.trim().toLowerCase();
        const messages = document.querySelectorAll('.chat-item');

        messages.forEach(msg => {
            const text = msg.querySelector('.message-text')?.textContent.toLowerCase() || '';
            msg.style.display = text.includes(search) ? '' : 'none';
        });
        cleanChatDates();
    };

    input.addEventListener('input', filterMessages);

    if (cancelBtn) {
        cancelBtn.addEventListener('click', () => {
            input.value = '';
            filterMessages();
        });
    }

    if (clearBtn) {
        clearBtn.addEventListener('click', () => {
            input.value = '';
            filterMessages();
        });
    }
}

function cleanChatDates() {
    const chatContent = document.querySelector('.chat__content');
    if (!chatContent) return;

    const nodes = Array.from(chatContent.children);
    let hideDate = true;

    for (let i = nodes.length - 1; i >= 0; i--) {
        const node = nodes[i];

        if (node.classList.contains('chat-item')) {
            const style = window.getComputedStyle(node);
            const isVisible = style.display !== 'none' && style.visibility !== 'hidden' && style.opacity !== '0';
            if (isVisible) {
                hideDate = false;
            }
        }

        if (node.classList.contains('chat__date')) {
            if (hideDate) {
                node.style.display = 'none';
            } else {
                node.style.display = '';
                hideDate = true;
            }
        }
    }
}
document.addEventListener("DOMContentLoaded", () => {
    initUserData();
});



const createGroupDialog = document.getElementById("createGroupModal");
const openBtn = document.getElementById("openCreateGroupModal");
const closeBtnGroupDialog = document.getElementById("closeCreateGroupModal");

const searchInput = document.getElementById("userSearch");
const searchResults = document.getElementById("searchResults");
const selectedUsers = document.getElementById("selectedUsers");

let selectedUserIds = [];

// Открытие
openBtn.addEventListener("click", (e) => {
    e.preventDefault();
    createGroupDialog.showModal();
});

// Закрытие
closeBtnGroupDialog.addEventListener("click", () => {
    createGroupDialog.close();
});

createGroupDialog.addEventListener("click", (event) => {
    if (event.target === createGroupDialog) {
        createGroupDialog.close();
    }
});

searchInput.addEventListener("input", () => {
    const query = searchInput.value.trim().toLowerCase();
    searchResults.innerHTML = "";

    if (!query) {
        searchResults.style.display = "none";
        return;
    }

    const results = otherUsers.filter(user =>
        user.Name.toLowerCase().includes(query) &&
        !selectedUserIds.includes(user.Id)
    );

    if (results.length === 0) {
        searchResults.style.display = "none";
        return;
    }

    // Показываем список
    searchResults.style.display = "block";

    results.forEach(user => {
        const option = document.createElement("div");
        option.textContent = user.Name;
        option.classList.add("search-result-item");

        option.addEventListener("click", () => {
            selectedUserIds.push(user.Id);

            const li = document.createElement("li");
            li.textContent = user.Name;
            li.classList.add("userSel");


            li.addEventListener("click", () => {
                selectedUsers.removeChild(li);
                selectedUserIds = selectedUserIds.filter(id => id !== user.Id);
            });

            selectedUsers.appendChild(li);
            searchResults.innerHTML = "";
            searchResults.style.display = "none";
            searchInput.value = "";
        });

        searchResults.appendChild(option);
    });
});


// Создание группового чата
document.getElementById("createGroupBtn").addEventListener("click", () => {
    const title = document.getElementById("groupTitle").value.trim();

    if (!title || selectedUserIds.length === 0) {
        alert("Введите название и выберите участников");
        return;
    }

    const membersChat = [...selectedUserIds];
    if (!membersChat.includes(currentUser.Id)) {
        membersChat.push(currentUser.Id);
    }

    socket.send(JSON.stringify({
        event: "createChat",
        from: currentUser.Id,      // добавляем поле from
        to: membersChat,           // добавляем поле to — список участников
        data: {
            title: title,
            members: membersChat,
            isGroup: true
        }
    }));
});

function setupAddUsersModal(chatId) {
    const modal = document.getElementById('addUsersModal');
    if (!modal) return;

    // Открываем модалку
    modal.showModal();

    // Сбросим выбранных пользователей и очистим интерфейс
    addSelectedUserIds = [];
    document.getElementById('addSelectedUsers').innerHTML = '';
    document.getElementById('addUserSearch').value = '';
    document.getElementById('addUserSearchResults').innerHTML = '';

    // Получаем список участников текущего чата
    const chat = findChatById(chatId);
    if (!chat) {
        alert('Чат не найден');
        modal.close();
        return;
    }
    const currentMembers = new Set(chat.members);

    // Все пользователи (должен быть массив всех пользователей в переменной allUsers)
    // Фильтруем — показываем только тех, кто еще не в чате
    const availableUsers = allUsers.filter(user => !currentMembers.has(user.Id));

    const searchInput = document.getElementById('addUserSearch');
    const searchResults = document.getElementById('addUserSearchResults');
    const selectedUsersList = document.getElementById('addSelectedUsers');

    // Отрисовка результатов поиска по пользователям
    function renderSearchResults(users) {
        searchResults.innerHTML = '';
        users.forEach(user => {
            const div = document.createElement('div');
            div.className = 'search-result-item';
            div.textContent = user.Name;
            div.dataset.userid = user.Id;

            div.addEventListener('click', () => {
                if (!addSelectedUserIds.includes(user.Id)) {
                    addSelectedUserIds.push(user.Id);

                    // Добавить в список выбранных
                    const li = document.createElement('li');
                    li.classList.add("userSel");
                    li.textContent = user.Name;
                    li.dataset.userid = user.Id;

                    li.addEventListener("click", () => {
                        selectedUsersList.removeChild(li);
                        addSelectedUserIds = addSelectedUserIds.filter(id => id !== user.Id);
                    });
                    // Кнопка удаления из выбранных
                    // const removeBtn = document.createElement('button');
                    // removeBtn.textContent = '×';
                    // removeBtn.onclick = () => {
                    //     selectedUsersList.removeChild(li);
                    //     addSelectedUserIds = addSelectedUserIds.filter(id => id !== user.Id);
                    // };
                    // li.appendChild(removeBtn);

                    selectedUsersList.appendChild(li);
                }
                searchInput.value = '';
                searchResults.innerHTML = '';
            });

            searchResults.appendChild(div);
        });

        if (users.length === 0) {
            searchResults.innerHTML = '<div class="search-no-results">Пользователи не найдены</div>';
        }
    }

    // Поиск пользователей при вводе
    searchInput.oninput = () => {
        const query = searchInput.value.trim().toLowerCase();
        if (!query) {
            searchResults.innerHTML = '';
            return;
        }

        const filtered = availableUsers.filter(user =>
            user.Name.toLowerCase().includes(query) &&
            !addSelectedUserIds.includes(user.Id)
        );
        renderSearchResults(filtered);
    };

    // Обработчик кнопки "Добавить"
    const addBtn = document.getElementById('addUsersBtn');
    addBtn.onclick = () => {
        if (addSelectedUserIds.length === 0) {
            alert('Выберите хотя бы одного пользователя');
            return;
        }

        const updatedMembers = [...new Set([...chat.members, ...addSelectedUserIds])];

        socket.send(JSON.stringify({
            event: 'updateChat',
            from: currentUser.Id,
            to: updatedMembers,
            data: {
                id: chatId,
                title: findChatById(chatId).title,
                members: updatedMembers,
                isGroup: true
            }
        }));
    }
}

function leaveChat(chatId, userId) {

    const chat = findChatById(chatId);
    if (!chat) return;

    // Создаем новый список участников, исключая текущего пользователя
    const updatedMembers = chat.members.filter(id => id !== userId);

    // Отправляем через WebSocket сообщение серверу об обновлении чата
    socket.send(JSON.stringify({
        event: "updateChat",
        from: userId,
        to: updatedMembers,
        data: {
            id: chatId,
            title: findChatById(chatId).title,
            members: updatedMembers,
            isGroup: chat.isGroup
        }
    }));

    removeChatFromSidebar(chatId);
    closeChatWindow();
    currentChatId = null;
}
function closeChatWindow() {
    const container = document.querySelector(".message__content");
    container.innerHTML = "";
}

function removeChatFromSidebar(chatId) {
    const chatItem = document.querySelector(`.contact-item__body[data-chat-id="${chatId}"]`);
    if (chatItem) {
        const parent = chatItem.closest('.contact-item');
        if (parent) parent.remove();
    }
}