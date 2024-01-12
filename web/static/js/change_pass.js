document.getElementById('changePassForm').addEventListener('submit', function(event) {
    event.preventDefault();
    const validPassword = setupPasswordMatchValidation('password', 'confirmPassword', 'changePassError');
    if (validPassword) {
        var submitButton = document.getElementById("submitBtn");
        submitButton.disabled = true;
        const form = document.getElementById("changePassForm");
        const formData = new FormData(form);
        sendRequest(formData, "/account/change_pass", submitButton, "Change pass", changePassHandle, changePassFinallyHandle)
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