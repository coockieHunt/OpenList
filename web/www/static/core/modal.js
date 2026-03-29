function OpenModal(title, content) {
    const modal = document.createElement('dialog');
    modal.classList.add('modal');
    modal.innerHTML = `
        <form method="dialog" style="display: flex; flex-direction: column; gap: 1rem;">
            <h3>${title}</h3>
            <div>${content}</div>
            <button type="submit" class="primary">Fermer</button>
        </form>
    `;

    document.body.appendChild(modal);
    modal.showModal();

    modal.addEventListener('close', () => {
        document.body.removeChild(modal);
    });
}

function OpenConfirmModal(title, content, onConfirm, footerButtons = [
    { label: 'Annuler', class: 'secondary', action: 'cancel' },
    { label: 'Confirmer', class: 'contrast', style: 'background-color: var(--pico-ins-color);', action: 'confirm' }
]) {
    const modal = document.createElement('dialog');
    modal.classList.add('modal');
    modal.innerHTML = `
        <article style="min-width: min(32rem, 90vw); margin: 0;">
            <h3>${title}</h3>
            <div>${content}</div>
            <footer style="display: flex; justify-content: flex-end; gap: 0.75rem;">
                ${footerButtons.map(btn => `
                    <button
                        type="button"
                        class="${btn.class || ''}"
                        style="${btn.style || ''}"
                        data-action="${btn.action || ''}"
                    >${btn.label || btn.text || ''}</button>
                `).join('')}
            </footer>
        </article>
    `;

    document.body.appendChild(modal);

    const cancelButton = modal.querySelector('[data-action="cancel"]');
    if (cancelButton) {
        cancelButton.addEventListener('click', () => modal.close());
    }

    const confirmButton = modal.querySelector('[data-action="confirm"]');
    if (confirmButton) {
        confirmButton.addEventListener('click', async () => {
            try {
                await onConfirm();
                if (modal.open) modal.close();
            } catch {}
        });
    }

    modal.addEventListener('click', (event) => {
        if (event.target === modal) modal.close();
    });

    modal.addEventListener('close', () => {
        document.body.removeChild(modal);
    });

    modal.showModal();
}

function CloseModal() {
    const modal = document.querySelector('dialog[open]');
    if (modal) modal.close();
}