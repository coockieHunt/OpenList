function ChangeTheme() {
    const htmlElement = document.documentElement;
    const currentTheme = htmlElement.getAttribute('data-theme');
    const newTheme = currentTheme === 'dark' ? 'light' : 'dark';
    htmlElement.setAttribute('data-theme', newTheme);
    localStorage.setItem('theme', newTheme); 

    const themeToggle = document.getElementById('theme-toggle');
    if (newTheme === 'dark') {
        themeToggle.classList.remove('icon-brightness-down');
        themeToggle.classList.add('icon-brightness-up');
    } else {
        themeToggle.classList.remove('icon-brightness-up');
        themeToggle.classList.add('icon-brightness-down');
    }
}

const savedTheme = localStorage.getItem('theme');
if (savedTheme) {
    document.documentElement.setAttribute('data-theme', savedTheme);
    const themeToggle = document.getElementById('theme-toggle');
    if (themeToggle) {
        themeToggle.classList.toggle('icon-brightness-up', savedTheme === 'dark');
        themeToggle.classList.toggle('icon-brightness-down', savedTheme === 'light');
    }
}