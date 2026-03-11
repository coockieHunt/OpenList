async function appendAllToList(domId) {
    const container = document.getElementById(domId);

    const startTime = performance.now();
    const data = await GetList();
    const list = data.data;
    const responseTime = performance.now() - startTime;
    const time = `${(responseTime * 1000).toFixed(0)} µs`;

    const fresh = container.cloneNode(false);
    container.parentNode.replaceChild(fresh, container);
    fresh.addEventListener('click', (e) => {
        const item = e.target.closest('.list-item');
        if (item) window.location.href = `/list/${item.dataset.id}`;
    });

    fresh.innerHTML = `
        ${list.length === 0 ? '<p style="text-align: center; color: var(--pico-muted-color);">Aucune liste trouvée.</p>' : ''}
        <ul>
            ${list.map(item => `
                <li>
                    <div class="list-item" style="display: flex; justify-content: space-between; align-items: center;" data-id="${item.id}">
                        <span>${item.title}</span>
                        <span class="item-count">${item.items ? item.items.length : 0}</span>
                    </div>
                </li>
            `).join('')}
        </ul>
    `;

    return {time, count: data['data'].length};
}

const ValidateItemHandler = async (itemId, listId, domId) => {
    const response = await ValidateItem(listId, itemId);
    if (response.status === 'success') {
        appendToItemList(domId, listId);
    } else {
        console.error(`Failed to validate item with ID: ${itemId}`);
    }
}


