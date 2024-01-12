function setupPasswordMatchValidation(passwordFieldId, confirmPasswordFieldId, messageElementId) {
    var passwordField = document.getElementById(passwordFieldId);
    var confirmPasswordField = document.getElementById(confirmPasswordFieldId);
    var messageElement = document.getElementById(messageElementId);
    var password = passwordField.value;
    var confirmPassword = confirmPasswordField.value;
    if (password !== confirmPassword) {
        messageElement.innerText = 'Password mismatch.';
        return false
    } else {
        messageElement.innerText = '';
        return true
    }
}