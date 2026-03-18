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

    container.innerHTML = `
        <ul>
            ${items.length === 0 ? '<p style="text-align: center; color: var(--pico-muted-color);">Aucun item dans cette liste.</p>' : ''}
            ${items.map(item => `
                <li>
                    <div 
                        class="list-item" 
                        style="display: flex; justify-content: space-between; align-items: center; ${item.validated ? 'background-color: var(--pico-primary-border);' : ''}" 
                        data-id="${item.id}" 
                        onclick="ValidateItemHandler(${item.id}, ${listId}, '${domId}')"
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
                                onclick="DeleteItemHandler(${item.id}, ${listId}, '${domId}')"
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