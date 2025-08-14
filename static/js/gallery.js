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
      `  <picture>
                <source srcset="/images/399${src}" media="(max-width: 399px)">
                <source srcset="/images/412${src}" media="(max-width: 599px)">
                <source srcset="/images/600${src}" media="(max-width: 992px)">
                <source srcset="/images/900${src}" media="(max-width: 1200px)">
                <source srcset="/images/900${src}" media="(min-width: 1201px)">
                <img src="/images/900${src}" loading="lazy" decoding="async" class="modal-gallery__img ${i === 0 ? 'active' : ''}">
          </picture>
          `
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

 // ==============================
  // 9. Сабхедер
  // ==============================
  if (window.innerWidth >= 768) {
    const header = document.getElementById('main-header');
    const subheader = document.getElementById('subheader');
    const aboutLink = document.querySelector('a[href="#about"]');
    const sectionIds = ['compound', 'school', 'dealers', 'delivery', 'diagnostics'];
    let lastScroll = window.pageYOffset || document.documentElement.scrollTop;

    aboutLink?.addEventListener('click', function (e) {
      if (!subheader) return;
      e.preventDefault();
      subheader.classList.toggle('open');
      subheader.setAttribute('aria-hidden', subheader.classList.contains('open') ? 'false' : 'true');
      if (subheader.classList.contains('open')) {
        subheader.classList.add('sticky');
        subheader.style.top = (header ? header.offsetHeight : 0) + 'px';
      } else {
        if (window.scrollY <= (header ? header.offsetHeight : 0)) {
          subheader.classList.remove('sticky');
          subheader.style.top = '';
        }
      }
    });

    const observerOptions = { root: null, rootMargin: '0px 0px -55% 0px', threshold: 0 };
    const observer = new IntersectionObserver((entries) => {
      entries.forEach(entry => {
        const id = entry.target.id;
        const link = document.querySelector('.subheader__item[href="#' + id + '"]');
        if (!link) return;
        if (entry.isIntersecting) {
          document.querySelectorAll('.subheader__item').forEach(l => l.classList.remove('active'));
          link.classList.add('active');
        }
      });
    }, observerOptions);

    sectionIds.forEach(id => {
      const el = document.getElementById(id);
      if (el) observer.observe(el);
    });

    window.addEventListener('scroll', function () {
      const current = window.pageYOffset || document.documentElement.scrollTop;
      if (header) {
        if (current > lastScroll && current > header.offsetHeight + 20) {
          header.classList.add('hide');
        } else {
          header.classList.remove('hide');
        }
      }
      lastScroll = current <= 0 ? 0 : current;

      if (subheader && header) {
        if (window.scrollY > header.offsetHeight) {
          subheader.classList.add('sticky');
          subheader.style.top = (header.classList.contains('hide') ? 0 : header.offsetHeight) + 'px';
          subheader.style.height = (header.classList.contains('hide') ? 50 : 35) + 'px';
        } else {
          if (!subheader.classList.contains('open')) {
            subheader.classList.remove('sticky');
            subheader.style.top = '';
          }
        }
      }

      const scrollTopBtn = document.getElementById('scrollTopBtn');
      if (scrollTopBtn) {
        if (window.scrollY > 400) scrollTopBtn.classList.add('show');
        else scrollTopBtn.classList.remove('show');
      }
    }, { passive: true });
  }


  
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
