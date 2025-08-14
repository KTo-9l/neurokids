document.addEventListener('DOMContentLoaded', function () {
  // Все ссылки внутри ul.tabs с data-tabgroup
  const tabs = document.querySelectorAll('ul.tabs[data-tabgroup] a');

  tabs.forEach(tab => {
    tab.addEventListener('click', function (e) {
      e.preventDefault();

      // Так как data-tabgroup на ul.tabs, ищем ближайший ul.tabs с data-tabgroup
      const tabGroupEl = this.closest('ul.tabs[data-tabgroup]');
      const tabGroupName = tabGroupEl.getAttribute('data-tabgroup');

      // Контейнер с содержимым табов
      const group = document.getElementById(tabGroupName);

      // Все табы (ссылки) внутри этого ul.tabs
      const allTabsInGroup = tabGroupEl.querySelectorAll('a');

      allTabsInGroup.forEach(t => t.classList.remove('active'));
      this.classList.add('active');

      // Скрываем все контентные блоки
      const tabItems = group.querySelectorAll('.tab-item');
      tabItems.forEach(item => item.style.display = 'none');

      // Показываем выбранный
      const target = document.querySelector(this.getAttribute('href'));
      if (target) target.style.display = '';
    });
  });
});
