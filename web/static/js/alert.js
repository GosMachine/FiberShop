function hideAlert(type, btn) {
    setTimeout(() => {
        var alertContainer
        if (type === "error") {
            alertContainer = document.getElementById('alert-error-container');
        } else {
            alertContainer = document.getElementById('alert-success-container')
        }
        alertContainer.removeAttribute('data-te-toast-show');
        alertContainer.setAttribute('data-te-toast-hide', '');
        btn.disabled = false;
    }, 5000);
}

function showAlert(type, title, data) {
    var alertName, alertContainer, alertMessage
    if (type === "error") {
        alertContainer = document.getElementById('alert-error-container');
        alertName = document.getElementById('alert-error-name');
        alertMessage = document.getElementById('alert-error-message');
    } else {
        alertContainer = document.getElementById('alert-success-container');
        alertName = document.getElementById('alert-success-name');
        alertMessage = document.getElementById('alert-success-message');
    }
    alertName.innerText = title;
    alertMessage.innerText = data.message;
    alertContainer.removeAttribute('data-te-toast-hide');
    alertContainer.setAttribute('data-te-toast-show', '');
}
