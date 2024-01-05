function setupPasswordMatchValidation(formId, passwordFieldId, confirmPasswordFieldId, messageElementId) {
    var form = document.getElementById(formId);
    var passwordField = document.getElementById(passwordFieldId);
    var confirmPasswordField = document.getElementById(confirmPasswordFieldId);
    var messageElement = document.getElementById(messageElementId);

    form.addEventListener('submit', function(event) {
        var password = passwordField.value;
        var confirmPassword = confirmPasswordField.value;

        if (password !== confirmPassword) {
            event.preventDefault();
            messageElement.innerText = 'Password mismatch.';
        } else {
            messageElement.innerText = '';
        }
    });
}