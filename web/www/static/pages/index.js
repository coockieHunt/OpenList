const refreshList = async () => {
    const startTime = performance.now();
    const data = await GetList();
    const list = data.data;
    const responseTime = performance.now() - startTime;
    const time = `${(responseTime * 1000).toFixed(0)} µs`;

	const Render = {
		"list-container": () => {
			const container = document.getElementById("list-container");
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
								<div style="display:flex; flex-direction: row; gap:8px;">
									<span class="item-count">${item.items ? item.items.length : 0}</span>
								</div>
							</div>
						</li>
					`).join('')}
				</ul>
			`;
		},
	}

	Render["list-container"]();

	document.getElementById("response-time").textContent = time;
	document.getElementById("response-count").textContent = data['data'].length;

};

onload = async () => {
	refreshList();
};