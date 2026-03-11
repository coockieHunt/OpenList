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

function OpenConfirmModal(title, content, onConfirm) {
    const modal = document.createElement('dialog');
    modal.classList.add('modal');
    modal.innerHTML = `
        <article style="min-width: min(32rem, 90vw); margin: 0;">
            <h3>${title}</h3>
            <p>${content}</p>
            <footer style="display: flex; justify-content: flex-end; gap: 0.75rem;">
                <button type="button" class="secondary" data-action="cancel">Annuler</button>
                <button type="button" class="contrast" style="background-color: var(--pico-del-color); color:white" data-action="confirm">Supprimer</button>
            </footer>
        </article>
    `;

    document.body.appendChild(modal);

    const cancelButton = modal.querySelector('[data-action="cancel"]');
    const confirmButton = modal.querySelector('[data-action="confirm"]');

    cancelButton.addEventListener('click', () => {
        modal.close();
    });

    confirmButton.addEventListener('click', async () => {
        try {
            await onConfirm();
        } finally {
            if (modal.open) {
                modal.close();
            }
        }
    });

    modal.addEventListener('click', (event) => {
        if (event.target === modal) {
            modal.close();
        }
    });

    modal.addEventListener('close', () => {
        document.body.removeChild(modal);
    });

    modal.showModal();
}