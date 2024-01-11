function sendRequest(form, data, btn, title, handler, finallyHandler) {
    fetch(form.action, {
        method: form.method,
        body: data
    })
        .then(response => {
            if (!response.ok) {
                throw new Error('Network response was not ok');
            }
            return response.json();
        })
        .then(data => {
            handler('success', title, data.message);
            finallyHandler("success", btn)
        })
        .catch(error => {
            handler("error", title, error.message);
            finallyHandler("error", btn)
        })
}