document.addEventListener('DOMContentLoaded', function () {
    document.getElementById('contactForm').addEventListener('submit', function(event) {
        event.preventDefault();
        var submitButton = document.getElementById('submitBtn');
        submitButton.disabled = true;
        const contactForm = document.getElementById('contactForm');
        const formData = new FormData(contactForm);
        sendRequest(formData, "/contact", submitButton, "Contact", showAlert, hideAlert)
    });
});