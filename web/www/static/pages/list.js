onload = () => {
	const listID = window.location.pathname.split('/').pop()
	appendToItemList('item-container', listID)

	document.getElementById('btn-add-item')?.addEventListener('click', OpenModalAddItem);
	document.getElementById('btn-delete-list')?.addEventListener('click', () => HandelDeleteList(listID));

}


const HandelAddItemToList = async (itemName, itemQuantity, listID) => {
    await newItem(itemName, itemQuantity, listID)
    appendToItemList('item-container', listID)
}

const HandelDeleteList = (listID) => {
    OpenConfirmModal(
        'Supprimer cette liste ?',
        'Tous les articles de cette liste seront egalement supprimes.',
        async () => {
            await DeleteList(listID)
            window.location.href = '/'
        },
        [
            {
                label: 'Annuler',
                class: 'secondary',
                action: 'cancel',
            },
            {
                label: 'Supprimer',
                class: 'contrast',
                action: 'confirm',
                style: 'background-color: var(--pico-del-color);',
            },
        ]
    )
}

const OpenModalAddItem = () => {
    const modalContent = `
		<form id="add-item-form">
			<input type="text" name="itemName" placeholder="Banane" required>
			<label>Quantité :</label>
			<div style="display: grid; grid-template-columns: 2.5rem 2.5rem 2.5rem; align-items: center; justify-content: center; column-gap: 1rem; width: 100%; margin-bottom: 1rem;">
				<button type="button" onclick="
					const q = document.getElementById('quantity-input');
					if (parseInt(q.value) > 1) { q.value = parseInt(q.value) - 1; document.getElementById('quantity-value').textContent = q.value; }
				" style="width: 2.5rem; height: 2.5rem; margin: 0; padding: 0; display: inline-flex; align-items: center; justify-content: center; font-size: 1.2rem;">−</button>
				<span id="quantity-value" style="display: inline-flex; align-items: center; justify-content: center; width: 2.5rem; height: 2.5rem; text-align: center; font-size: 2rem;">1</span>
				<button type="button" onclick="
					const q = document.getElementById('quantity-input');
					if (parseInt(q.value) < 10) { q.value = parseInt(q.value) + 1; document.getElementById('quantity-value').textContent = q.value; }
				" style="width: 2.5rem; height: 2.5rem; margin: 0; padding: 0; display: inline-flex; align-items: center; justify-content: center; font-size: 1.2rem;">+</button>
				<input type="hidden" id="quantity-input" name="itemQuantity" value="1">
			</div>
		</form>
	`

    OpenConfirmModal(
        'Ajouter un article',
        modalContent,
        () => {
            const form = document.getElementById('add-item-form')
            if (!form.reportValidity()) return Promise.reject()
            const listID = window.location.pathname.split('/').pop()
            HandelAddItemToList(
                form.itemName.value,
                form.itemQuantity.value,
                listID
            )
        },
        [
            {
                label: 'Annuler',
                class: 'secondary',
                action: 'cancel',
            },
            {
                label: 'Ajouter',
                class: 'contrast',
                action: 'confirm',
                style: 'background-color: var(--pico-ins-color);',
            },
        ]
    )
}




const DeleteItemHandler = (itemId, listId, domId) => {
    OpenConfirmModal('Supprimer cet article ?', 'Cette action est irreversible.', async () => {
        const response = await DeleteItem(listId, itemId);
        if (response.status === 'success') {
            appendToItemList(domId, listId);
            return;
        }

        console.error(`Failed to delete item with ID: ${itemId}`);
    });
}

const ValidateItemHandler = async (itemId, listId, domId) => {
    const response = await ValidateItem(listId, itemId);
    if (response.status === 'success') {
        appendToItemList(domId, listId);
    } else {
        console.error(`Failed to validate item with ID: ${itemId}`);
    }
}

async function appendToItemList(domId, listId) {
    const container = document.getElementById(domId);
    if (!container) {
        console.error(`Container not found: ${domId}`);
        return;
    }

    const data = await getListItems(listId);
    const items = data.data.items || [];

    container.dataset.domId = domId;
    container.dataset.listId = String(listId);

    if (!container.dataset.clickBound) {
        container.addEventListener('click', (event) => {
            const deleteBtn = event.target.closest('button[data-action="delete-item"]');
            if (deleteBtn) {
                event.stopPropagation();
                const itemId = Number(deleteBtn.dataset.itemId);
                const currentListId = Number(container.dataset.listId);
                const currentDomId = container.dataset.domId;
                DeleteItemHandler(itemId, currentListId, currentDomId);
                return;
            }

            const row = event.target.closest('.list-item[data-item-id]');
            if (!row || !container.contains(row)) {
                return;
            }

            const itemId = Number(row.dataset.itemId);
            const currentListId = Number(container.dataset.listId);
            const currentDomId = container.dataset.domId;
            ValidateItemHandler(itemId, currentListId, currentDomId);
        });

        container.dataset.clickBound = '1';
    }




    container.innerHTML = `
        <ul>
            ${items.length === 0 ? '<p style="text-align: center; color: var(--pico-muted-color);">Aucun item dans cette liste.</p>' : ''}
            ${items.map(item => `
                <li>
                    <div 
                        class="list-item" 
                        style="display: flex; justify-content: space-between; align-items: center; ${item.validated ? 'background-color: var(--pico-primary-border);' : ''}" 
                        data-item-id="${item.id}"
                    >
                        <span>${item.name} | ${item.quantity}</span>
                        <span style="display:flex; align-items:center; gap:8px;">
                            <span >
                                <input type="checkbox" ${item.validated ? 'checked' : ''} disabled>
                            </span>
                            <button type="button" 
                                class="secondary" 
                                style="padding: 0.2rem 0.5rem; 
                                margin: 0; 
                                background-color: var(--pico-del-color);"
                                data-action="delete-item"
                                data-item-id="${item.id}"
                            >
                                ✕
                            </button>
                        </span>
                    </div>
                </li>
            `).join('')}
        </ul>
    `;
}