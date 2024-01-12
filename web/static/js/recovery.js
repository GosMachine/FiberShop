document.addEventListener('DOMContentLoaded', function () {
    initForm("recoveryForm", "submitBtn", "Recovery", "/account_recovery", recoveryHandle, recoveryFinallyHandle)
});


function recoveryHandle(type, title, data) {
    if (type === "success") {
        document.open();
        document.write(data);
        document.close();
    } else {
        if (data === "UserIsNotFound") {
            data = "User is not found."
        } else if (data === "InternalError") {
            data = "Internal error. Please try again."
        }
        var errorMessage = document.getElementById('recoveryError');
        errorMessage.innerText = data
    }
}


function recoveryFinallyHandle(type, btn) {
    if (type === "error") {
        setTimeout(() => {
            btn.disabled = false;
        }, 1000);
    }
}