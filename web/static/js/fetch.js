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