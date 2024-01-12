function sendRequest(data, endpoint, btn, title, handler, finallyHandler) {
    fetch(endpoint, {
        method: "POST",
        body: data
    })
        .then(response => {
            if (!response.ok) {
                return response.json().then(errorData => {
                    throw new Error(errorData.error || 'Ошибка при выполнении запроса');
                });
            }
            if (response.redirected) {
                window.location.href = response.url;
            } else {
                const contentType = response.headers.get("content-type");
                return contentType.includes("application/json") ? response.json() : response.text();
            }

        })
        .then(data => {
            if (data) {
                handler('success', title, data);
                finallyHandler("success", btn)
            }
        })
        .catch(error => {
            handler("error", title, error.message);
            finallyHandler("error", btn)
        })
}

function initForm(formId, btnId, title, endpoint, handler, finallyHandler) {
    document.getElementById(formId).addEventListener('submit', function(event) {
        event.preventDefault();
        var submitButton = document.getElementById(btnId);
        submitButton.disabled = true;
        const form = document.getElementById(formId);
        const formData = new FormData(form);
        sendRequest(formData, endpoint, submitButton, title, handler, finallyHandler)
    })
}
