document.getElementById('resendCodeButton').addEventListener('click', function() {
    var button = this;
    button.disabled = true;
    const email = document.getElementById('email').value;
    const formData = new FormData();
    formData.append('email', email);

    fetch('/email/resend', {
        method: 'POST',
        body: formData
    })
        .then(response => {
            if (!response.ok) {
                throw new Error(`HTTP error! Status: ${response.status}`);
            }
        })
        .catch(error => {
            console.error('Error resending code:', error);
        })
        .finally(() => {
            setTimeout(() => {
                button.disabled = false;
            }, 10000);
        });
});