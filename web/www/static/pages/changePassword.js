const form = document.getElementById('change-password-form');
const errorBox = document.getElementById('change-password-error');
const successBox = document.getElementById('change-password-success');

const showError = (message) => {
    errorBox.textContent = message;
    errorBox.style.display = 'block';
    successBox.style.display = 'none';
};

const showSuccess = (message) => {
    successBox.textContent = message;
    successBox.style.display = 'block';
    errorBox.style.display = 'none';
};

window.addEventListener('DOMContentLoaded', async () => {
    const status = await AuthStatus();
    if (status.status !== 'success') {
        window.location.href = '/login';
        return;
    }

    form.addEventListener('submit', async (event) => {
        event.preventDefault();

        const formData = new FormData(form);
        const currentPassword = String(formData.get('current_password') || '');
        const newPassword = String(formData.get('new_password') || '');
        const confirmPassword = String(formData.get('confirm_password') || '');

        if (!currentPassword || !newPassword || !confirmPassword) {
            showError('Tous les champs sont obligatoires.');
            return;
        }

        if (newPassword.length < 8) {
            showError('Le mot de passe doit contenir au moins 8 caractères.');
            return;
        }

        if (newPassword !== confirmPassword) {
            showError('Les deux mot de passe ne correspondent pas.');
            return;
        }

        const response = await ChangePassword(currentPassword, newPassword);
        if (response.status !== 'success') {
            showError(response.message || 'Impossible de changer le mot de passe.');
            return;
        }

        showSuccess('Mot de passe mis à jour');
        setTimeout(() => {
            window.location.href = '/login';
        }, 900);
    });
});
