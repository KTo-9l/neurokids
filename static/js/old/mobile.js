

const dialog = document.getElementById('mobile-menu');
const burger = document.getElementById('burger');
const closeBtn = document.getElementById('close-menu');

burger.addEventListener('click', () => {
    dialog.showModal();
    document.body.style.overflow = 'hidden';
    document.documentElement.classList.add('box-sizing-unset');
});

closeBtn.addEventListener('click', () => {
    dialog.close();
    document.body.style.overflow = '';
    document.documentElement.classList.remove('box-sizing-unset');
});

dialog.querySelectorAll('a').forEach(link => {
    link.addEventListener('click', () => {
    dialog.close();
    document.body.style.overflow = '';
    document.documentElement.classList.remove('box-sizing-unset');
    });
});

dialog.addEventListener('click', (e) => {
  const rect = dialog.getBoundingClientRect();
  // если клик был за пределами ширины меню
  if (e.clientX > rect.left + dialog.offsetWidth) {
    dialog.close();
    document.body.style.overflow = '';
    document.documentElement.classList.remove('box-sizing-unset');
  }
});