document.addEventListener('DOMContentLoaded', function () {
    const path = window.location.pathname
    if (path === "/login") {
        document.getElementById('loginForm').addEventListener('submit', function(event) {
            event.preventDefault();
            authenticateUser("/login")
        });
    } else if (path === "/register") {
        document.getElementById('registerForm').addEventListener('submit', function(event) {
            event.preventDefault();
            const validPassword = setupPasswordMatchValidation('password', 'confirmPassword', 'authError');
            if (validPassword) {
                authenticateUser("/register")
            }
        });
    }
});


function authenticateUser(endpoint) {
    var submitButton = document.getElementById('submitBtn');
    submitButton.disabled = true;
    const email = document.getElementById('email').value;
    const password = document.getElementById('password').value;
    const remember = document.getElementById("remember").value;

    const formData = new FormData();
    formData.append('email', email);
    formData.append('password', password);
    formData.append("remember", remember)
    sendRequest(formData, endpoint, submitButton, null, authHandle, authFinallyHandle)
}

function authHandle(type, title, data) {
    if (type === "error") {
        if (data === "InvalidCredentials") {
            data = "Invalid email or password."
        } else if (data === "AlreadyExists") {
            data = "User already exists."
        } else if (data === "InternalError") {
            data = "Internal error. Please try again."
        }
        var errorMessage = document.getElementById('authError');
        errorMessage.innerText = data
    } else if (type === "success") {
        document.open();
        document.write(data);
        document.close();
    }
}

function authFinallyHandle(type, btn) {
    if (type === "error") {
        setTimeout(() => {
            btn.disabled = false;
        }, 1000);
    }
}