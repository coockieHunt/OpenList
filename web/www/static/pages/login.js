const loginForm = document.getElementById('login-form');
const errorBox = document.getElementById('login-error');

const showError = (message) => {
    errorBox.textContent = message;
    errorBox.style.display = 'block';
};

const hideError = () => {
    errorBox.textContent = '';
    errorBox.style.display = 'none';
};

window.addEventListener('DOMContentLoaded', async () => {
    try {
        const status = await AuthStatus();
        if (status.status === 'success') {
            if (status.data?.first_login) {
                window.location.href = '/change-password';
                return;
            }

            window.location.href = '/';
            return;
        }
    } catch (_) {}

    loginForm.addEventListener('submit', async (event) => {
        event.preventDefault();
        hideError();

        const formData = new FormData(loginForm);
        const username = String(formData.get('username') || '').trim();
        const password = String(formData.get('password') || '');

        if (!username || !password) {
            showError('Identifiant et mot de passe requis.');
            return;
        }

        const response = await Login(username, password);
        if (response.status !== 'success') {
            showError(response.message || 'Connexion impossible.');
            return;
        }

        if (response.data?.first_login) {
            window.location.href = '/change-password';
            return;
        }

        window.location.href = '/';
    });
});
