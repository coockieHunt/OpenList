async function appendToList(domId) {
    const container = document.getElementById(domId);
    const data = await GetList();
    const list = data.data;

    container.addEventListener('click', (e) => {
        const item = e.target.closest('.list-item');
        if (item) {
            window.location.href = `/list/${item.dataset.id}`;
        }
    });

    const htmlContent = `
        ${list.length === 0 ? '<p style="text-align: center; color: var(--pico-muted-color);">Aucune liste trouvée.</p>' : ''}
        <ul>
            ${list.map(item => `
                <li>
                    <div class="list-item" 
                        style="display: flex; justify-content: space-between; align-items: center;"
                        data-id="${item.id}">
                        <span>${item.title}</span>
                        <span class="item-count">${item.items ? item.items.length : 0}</span>
                    </div>
                </li>
            `).join('')}
        </ul>
    `;

    container.innerHTML = htmlContent;
}

    const ValidateItemHandler = async (itemId, listId, domId) => {
        const response = await ValidateItem(listId, itemId);
        console.log(response);
        if (response.status === 'success') {
            console.log(`Item with ID: ${itemId} validated successfully`);
            appendToItemList(domId, listId);
        } else {
            console.error(`Failed to validate item with ID: ${itemId}`);
        }
    }

    const DeleteItemHandler = async (itemId, listId, domId) => {
        const response = await DeleteItem(listId, itemId);
        console.log(response);
        if (response.status === 'success') {
            console.log(`Item with ID: ${itemId} deleted successfully`);
            appendToItemList(domId, listId);
        } else {
            console.error(`Failed to delete item with ID: ${itemId}`);
        }
    }


async function appendToItemList(domId, listId) {
    const container = document.getElementById(domId);
    const data = await getListItems(listId);
    const items = data.data.items || [];


    const htmlContent = `
        <ul>
            ${items.length === 0 ? '<p style="text-align: center; color: var(--pico-muted-color);">Aucun item dans cette liste.</p>' : ''}
            ${items.map(item => `
                <li>
                    <div class="list-item" 
                        style="display: flex; justify-content: space-between; align-items: center;"
                        data-id="${item.id}">
                        <span>${item.name} | ${item.quantity}</span>
                        <span style="display:flex; align-items:center; gap:8px;">
                            <span onclick="ValidateItemHandler(${item.id}, ${listId}, '${domId}')"><input type="checkbox" ${item.validated ? 'checked' : ''} disabled></span>
                            <button type="button" class="secondary" style="padding: 0.2rem 0.5rem; margin: 0;" onclick="DeleteItemHandler(${item.id}, ${listId}, '${domId}')">✕</button>
                        </span>
                    </div>
                </li>
            `).join('')}
        </ul>
    `;

    container.innerHTML = htmlContent;
}