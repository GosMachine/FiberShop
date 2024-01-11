function authenticateUser(endpoint) {
    const username = document.getElementById('username').value;
    const password = document.getElementById('password').value;

    const formData = new FormData();
    formData.append('username', username);
    formData.append('password', password);

    fetch(endpoint, {
        method: 'POST',
        body: formData,
    })
        .then(response => response.json())
        .then(data => {
            if (data.success) {
                alert('Authentication successful!');
                // Дополнительные действия при успешной аутентификации
            } else {
                alert('Authentication failed. Please try again.');
                // Дополнительные действия при неудачной аутентификации
            }
        })
        .catch(error => {
            console.error('Error during authentication:', error);
        });
}