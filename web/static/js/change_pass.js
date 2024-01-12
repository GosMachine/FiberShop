document.addEventListener('DOMContentLoaded', function () {
    const validPassword = setupPasswordMatchValidation('password', 'confirmPassword', 'changePassError');
    if (validPassword) {
        initForm("changePassForm", "submitBtn", "Change pass", "/account/change_pass", changePassHandle, changePassFinallyHandle)
    }
});


function changePassHandle(type, title, data) {
    if (type === "error") {
        if (data === "InternalError") {
            data = "Internal error. Please try again."
        }
        var errorMessage = document.getElementById('recoveryError');
        errorMessage.innerText = data
    }

}


function changePassFinallyHandle(type, btn) {
    if (type === "error") {
        setTimeout(() => {
            btn.disabled = false;
        }, 1000);
    }
}