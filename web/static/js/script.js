function initializeForm(form, title) {
    var submitButton = document.getElementById('submit');

    form.addEventListener('submit', function (event) {
        event.preventDefault();
        const formData = new FormData(form);
        submitButton.disabled = true;

        sendRequest(form, formData, submitButton, title, showAlert, hideAlert)
    });
}