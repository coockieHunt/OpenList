function ChangeTheme() {
    const htmlElement = document.documentElement;
    const currentTheme = htmlElement.getAttribute('data-theme');
    const newTheme = currentTheme === 'dark' ? 'light' : 'dark';
    htmlElement.setAttribute('data-theme', newTheme);

    const themeToggle = document.getElementById('theme-toggle');
    if (newTheme === 'dark') {
        themeToggle.classList.remove('icon-brightness-down');
        themeToggle.classList.add('icon-brightness-up');
    } else {
        themeToggle.classList.remove('icon-brightness-up');
        themeToggle.classList.add('icon-brightness-down');
    }
}