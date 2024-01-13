document.addEventListener('DOMContentLoaded', function () {
    initForm("changeEmailForm", "submitBtn", "Change email", "/account/settings/change_email", changeEmailHandle, changeEmailFinallyHandle)
});


function changeEmailHandle(type, title, data) {
    if (type === "success") {
        document.open();
        document.write(data);
        document.close();
    } else {
        if (data === "EmailAlreadyUsed") {
            data = "This email address is already in use."
        } else if (data === "InternalError") {
            data = "Internal error. Please try again."
        }
        var errorMessage = document.getElementById('settingsError');
        errorMessage.innerText = data
    }
}


function changeEmailFinallyHandle(type, btn) {
    if (type === "error") {
        setTimeout(() => {
            btn.disabled = false;
        }, 1000);
    }
}