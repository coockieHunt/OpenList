const MainUrl = (window.API_URL || 'http://localhost:8080/api').replace(/\/$/, '');

const apiFetch = async (url, options = {}) => {
    const response = await fetch(url, {
        credentials: 'include',
        ...options,
        headers: {
            'Content-Type': 'application/json',
            ...(options.headers || {}),
        },
    });
    return response.json();
};

const GetList = async () => {
    try {
        return await apiFetch(`${MainUrl}/list`);
    } catch (error) {
        console.error('Error:', error);
    }
};

const getListItems = async (listId) => {
    try {
        return await apiFetch(`${MainUrl}/list/${listId}`);
    } catch (error) {
        console.error('Error:', error);
    }
};

const ValidateItem = async (listId, itemId) => {
    try {
        return await apiFetch(`${MainUrl}/item/${listId}/${itemId}`, {
            method: 'PUT'
        });
    } catch (error) {
        console.error('Error:', error);
    }
};

const DeleteItem = async (listId, itemId) => {
    try {
        return await apiFetch(`${MainUrl}/item/${listId}/${itemId}`, {
            method: 'DELETE'
        });
    } catch (error) {
        console.error('Error:', error);
    }
};

const newList = async (name) => {
    try {
        return await apiFetch(`${MainUrl}/list`, {
            method: 'POST',
            body: JSON.stringify({ title: name, items: [] })
        });
    } catch (error) {
        console.error('Error:', error);
    }
};

const newItem = async (name, quantity, listId) => {
    try {
        return await apiFetch(`${MainUrl}/item/${listId}`, {
            method: 'POST',
            body: JSON.stringify({ name, quantity: parseInt(quantity), validated: false })
        });
    } catch (error) {
        console.error('Error:', error);
    }
};

const DeleteList = async (listId) => {
    try {
        return await apiFetch(`${MainUrl}/list/${listId}`, {
            method: 'DELETE'
        });
    } catch (error) {
        console.error('Error:', error);
    }
};

const Login = async (username, password) => {
    try {
        return await apiFetch(`${MainUrl}/auth/login`, {
            method: 'POST',
            body: JSON.stringify({ username, password })
        });
    } catch (error) {
        console.error('Error:', error);
    }
};

const Logout = async () => {
    try {
        return await apiFetch(`${MainUrl}/auth/logout`, {
            method: 'POST'
        });
    } catch (error) {
        console.error('Error:', error);
    }
};

const ChangePassword = async (current_password, new_password) => {
    try {
        return await apiFetch(`${MainUrl}/auth/change-password`, {
            method: 'POST',
            body: JSON.stringify({ current_password, new_password })
        });
    } catch (error) {
        console.error('Error:', error);
    }
};

const AuthStatus = async () => {
    try {
        return await apiFetch(`${MainUrl}/auth/status`);
    } catch (error) {
        console.error('Error:', error);
    }
};