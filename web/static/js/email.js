document.addEventListener('DOMContentLoaded', function () {
    initForm("emailForm", "submitBtn", "Email", "/email", emailHandle, emailFinallyHandle)
    var resendButton = document.getElementById('resendBtn');
    resendButton.addEventListener('click', function() {
        resendButton.disabled = true;
        const email = document.getElementById('email').value;
        const formData = new FormData();
        formData.append('email', email);
        sendRequest(formData, "/email/resend", resendButton, "Resend code", showAlert, hideAlert)
    });
});


function emailHandle(type, title, data) {
    if (type === "success") {
        document.open();
        document.write(data);
        document.close();
    } else {
        if (data === "WrongCode") {
            data = "Incorrect code."
        } else if (data === "CodeTimeError") {
            data = "Confirmation code has expired. Please request a new one"
        } else if (data === "InternalError") {
            data = "Internal error. Please try again."
        }
        var errorMessage = document.getElementById('emailError');
        errorMessage.innerText = data
    }
}


function emailFinallyHandle(type, btn) {
    if (type === "error") {
        setTimeout(() => {
            btn.disabled = false;
        }, 1000);
    }
}