// ==============================
// 1. Модалка с галереей от ящиков
// ==============================

const modalCell = document.getElementById('box');
const gallery = document.getElementById('galleryImages');
const textElement = document.getElementById('imageText');
const prevBtn = document.getElementById('prevBtn');
const nextBtn = document.getElementById('nextBtn');

let currentIndex = 0;
let images = [];
let texts = [];

function updateGallery() {
  const allImgs = gallery.querySelectorAll('img');
  allImgs.forEach((img, i) => {
    img.classList.toggle('active', i === currentIndex);
  });
  textElement.textContent = texts[currentIndex];
}

document.querySelectorAll('.box-grid__cell').forEach(cell => {
  cell.addEventListener('click', () => {
    images = JSON.parse(cell.dataset.images);
    texts = JSON.parse(cell.dataset.texts);
    currentIndex = 0;

    gallery.innerHTML = images.map((src, i) =>
      `<img src="${src}" class="modal-gallery__img ${i === 0 ? 'active' : ''}">`
    ).join('');

    updateGallery();
    modalCell.showModal();
    document.body.style.overflow = 'hidden';
  });
});

prevBtn.addEventListener('click', () => {
  currentIndex = (currentIndex - 1 + images.length) % images.length;
  updateGallery();
});

nextBtn.addEventListener('click', () => {
  currentIndex = (currentIndex + 1) % images.length;
  updateGallery();
});

modalCell.addEventListener('click', (e) => {
  if (e.target === modalCell) {
    modalCell.close();
    document.body.style.overflow = '';
  }
});

modalCell.addEventListener('close', () => {
  document.body.style.overflow = '';
});


// ==============================
// 2. Независимые галереи на странице (вне модалки)
// ==============================

const galleries = document.querySelectorAll(".section__gallery");
const galleryStates = new Map();

function updateGallerySection(gallery, index) {
  const images = gallery.querySelectorAll(".gallery__img");
  images.forEach((img, i) => {
    img.classList.toggle("active", i === index);
  });
}

for (const gallery of galleries) {
  const images = gallery.querySelectorAll(".gallery__img");
  let index = 0;
  galleryStates.set(gallery, index);

  const left = gallery.querySelector(".icon--left");
  const right = gallery.querySelector(".icon--right");

  if (left) {
    left.addEventListener("click", () => {
      let currentIndex = galleryStates.get(gallery);
      currentIndex = (currentIndex - 1 + images.length) % images.length;
      galleryStates.set(gallery, currentIndex);
      updateGallerySection(gallery, currentIndex);
    });
  }

  if (right) {
    right.addEventListener("click", () => {
      let currentIndex = galleryStates.get(gallery);
      currentIndex = (currentIndex + 1) % images.length;
      galleryStates.set(gallery, currentIndex);
      updateGallerySection(gallery, currentIndex);
    });
  }

  updateGallerySection(gallery, index);
}


// ==============================
// 3. Модалки: открытие, закрытие (универсально)
// ==============================

// ✅ myModal (открывается по .openModal)
const modal = document.getElementById('myModal');
document.querySelectorAll('.openModal').forEach((btn) => {
  btn.addEventListener('click', () => {
    modal.showModal();
    document.body.style.overflow = 'hidden';
  });
});
modal.addEventListener('click', (e) => {
  if (e.target.classList.contains('modal')) {
    modal.close();
  }
});
modal.addEventListener('close', () => {
  document.body.style.overflow = '';
});

// rekvModal (открывается по #openRekv)
const modal2 = document.getElementById('rekvModal');
const openBtn2 = document.getElementById('openRekv');

openBtn2.addEventListener('click', (e) => {
  e.preventDefault();
  modal2.showModal();
  document.body.style.overflow = 'hidden';
});

modal2.addEventListener('click', (e) => {
  if (e.target.classList.contains('modal') || e.target === modal2) {
    modal2.close();
  }
});
modal2.addEventListener('close', () => {
  document.body.style.overflow = '';
});


// ==============================
// 4. Логин-модалка (loginModal)
// ==============================

const loginModal = document.getElementById('loginModal');
const openLoginBtn = document.querySelector('.open-login-modal');

openLoginBtn.addEventListener('click', (e) => {
  e.preventDefault();
  loginModal.showModal();
  document.body.style.overflow = 'hidden';
});

loginModal.addEventListener('click', (e) => {
  const content = loginModal.querySelector('.modal__content');
  if (e.target === loginModal) {
    loginModal.close();
    document.body.style.overflow = '';
  }
});

loginModal.addEventListener('close', () => {
  document.body.style.overflow = '';
});


// ==============================
// 5. Маска телефона (Inputmask)
// ==============================

const phoneInput = document.getElementById("phone");
const phoneInputModal = document.getElementById("phone-modal");
const im = new Inputmask("+7 (999) 999-99-99");
im.mask(phoneInput);
im.mask(phoneInputModal);


// ==============================
// 6. Бургер-меню (моб.)
// ==============================
const dialog = document.getElementById('mobile-menu');
const burger = document.getElementById('burger');
const closeBtn = document.getElementById('close-menu');

burger.addEventListener('click', () => {
  dialog.showModal();
  document.body.style.overflow = 'hidden';
});

closeBtn.addEventListener('click', () => {
  dialog.close();
  document.body.style.overflow = '';
});

dialog.querySelectorAll('a').forEach(link => {
  link.addEventListener('click', () => {
    dialog.close();
    document.body.style.overflow = '';
  });
});

dialog.addEventListener('click', (e) => {
  const rect = dialog.getBoundingClientRect();
  // если клик был за пределами ширины меню
  if (e.clientX > rect.left + dialog.offsetWidth) {
    dialog.close();
    document.body.style.overflow = '';
  }
});



// ==============================
// 7. Аккордеоны
// ==============================

document.querySelectorAll('.accordion').forEach(accordion => {
  accordion.addEventListener('click', () => {
    const contentId = accordion.dataset.target;
    const content = document.querySelector(contentId);
    accordion.classList.toggle('active');
    content.classList.toggle('active');
  });
});

// ==============================
// 8. Кнопка "крестик" для закрытия модалок 
// ==============================

document.querySelectorAll('.modal__close').forEach(btn => {
  btn.addEventListener('click', () => {
    const dialog = btn.closest('dialog');
    if (dialog && dialog.open) {
      dialog.close();
      document.body.style.overflow = '';
    }
  });
});

document.querySelector('.login-form').addEventListener('submit', function (e) {
  e.preventDefault();

  const form = e.target;
  const errorBlock = form.querySelector('.login-error');
  errorBlock.style.display = 'none';
  errorBlock.textContent = '';

  const login = form.login.value;
  const pass = form.pass.value;
  const body = `login=${encodeURIComponent(login)}&pass=${encodeURIComponent(pass)}`;

  fetch('/login', {
    method: 'POST',
    headers: { 'Content-Type': 'application/x-www-form-urlencoded' },
    body: body,
    credentials: 'include'
  })
    .then(response => {
      console.log(response.status)
      if (response.status === 302 || response.status === 200) {
        window.location.reload();
      } else {
        errorBlock.textContent = 'Неверный логин или пароль';
        errorBlock.style.display = 'block';
      }
    })
    .catch(error => {
      errorBlock.textContent = 'Сетевая ошибка. Проверьте подключение.';
      errorBlock.style.display = 'block';
      console.error(error);
    });
});