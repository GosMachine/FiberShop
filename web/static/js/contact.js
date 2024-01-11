document.addEventListener('DOMContentLoaded', function () {
    var submitButton = document.getElementById('submit');
    submitButton.addEventListener('click', function() {
        submitButton.disabled = true;
        const contactForm = document.getElementById('contactForm');
        const formData = new FormData(contactForm);
        sendRequest(formData, "/contact", submitButton, "Contact", showAlert, hideAlert)
    });
});