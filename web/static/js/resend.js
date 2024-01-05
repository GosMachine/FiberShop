document.getElementById('resendCodeButton').addEventListener('click', function() {
    var button = this;
    button.disabled = true;

    var email = document.querySelector('#email [name="email"]').value;
    var action = document.querySelector('#email [name="action"]').value;

    fetch('/email/resend', {
        method: 'POST',
        headers: {'Content-Type': 'application/x-www-form-urlencoded'},
        body: 'email=' + encodeURIComponent(email) + '&action=' + encodeURIComponent(action)
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