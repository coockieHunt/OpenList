const MainUrl = (window.API_URL || 'http://localhost:8080/api').replace(/\/$/, '');

const GetList = () => {
    url = `${MainUrl}/list`;
    const response = fetch(url, {
        method: 'GET',
        headers: {
            'Content-Type': 'application/json'
        }
    }).then(response => response.json())
    .then(data => {
        return data;
    })
    .catch(error => {
        console.error('Error:', error);
    });
    return response;
}

const getListItems = (listId) => {
    url = `${MainUrl}/list/${listId}`;
    const response = fetch(url, {
        method: 'GET',
        headers: {
            'Content-Type': 'application/json'
        }
    }).then(response => response.json())
    .then(data => {
        return data;
    })
    .catch(error => {
        console.error('Error:', error);
    });
    return response;
}

const ValidateItem = (listId, itemId) => {
    url = `${MainUrl}/item/${listId}/${itemId}`;

    console.log(listId, itemId);
    const response = fetch(url, {
        method: 'PUT',
        headers: {
            'Content-Type': 'application/json'
        }
    }).then(response => response.json())
    .then(data => {
        return data;
    })
    .catch(error => {
        console.error('Error:', error);
    });
    return response;
}

const DeleteItem = (listId, itemId) => {
    url = `${MainUrl}/item/${listId}/${itemId}`;

    const response = fetch(url, {
        method: 'DELETE',
        headers: {
            'Content-Type': 'application/json'
        }
    }).then(response => response.json())
    .then(data => {
        return data;
    })
    .catch(error => {
        console.error('Error:', error);
    });
    return response;
}

const newList = (name) => {
    url = `${MainUrl}/list`;

    const response = fetch(url, {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify({
            title: name,
            items: []
        })
    }).then(response => response.json())
    .then(data => {
        return data;
    })
    .catch(error => {
        console.error('Error:', error);
    });
    return response;
}

const newItem = (name, quantity, listId) => {
    url = `${MainUrl}/item/${listId}`;
    console.log(name, quantity, listId);
    const response = fetch(url, {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify({
            name,
            quantity: parseInt(quantity),
            validated: false
        })
    }).then(response => response.json())
    .then(data => {
        return data;
    })
    .catch(error => {
        console.error('Error:', error);
    });
    return response;
}

const DeleteList = (listId) => {
    url = `${MainUrl}/list/${listId}`;

    const response = fetch(url, {
        method: 'DELETE',
        headers: {
            'Content-Type': 'application/json'
        }
    }).then(response => response.json())
    .then(data => {
        return data;
    })
    .catch(error => {
        console.error('Error:', error);
    });
    return response;
}