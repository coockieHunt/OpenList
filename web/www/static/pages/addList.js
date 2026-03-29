export const PutNewList = async (name) => {
    const response = await newList(name);
    if (response.status === "success") {
        window.location.href = `/list/${response.data.id}`;
    } else {
        alert("Erreur lors de la création de la liste.");
    }
};

document.addEventListener('DOMContentLoaded', () => {
    const form = document.getElementById('add-list-form');
    
    if (form) {
        form.addEventListener('submit', async (event) => {
            event.preventDefault(); 
            
            const input = form.querySelector('input[name="title"]');
            const name = input ? input.value.trim() : '';

            if (!name) {
                alert("Le titre de la liste est requis.");
                return;
            }
            
            await PutNewList(name);
        });
    }
});