// document.addEventListener('DOMContentLoaded', function () {
//     const submitBtn = document.getElementById('submitBtn');
//
//     submitBtn.addEventListener('click', function () {
//         showAlert();
//         setTimeout(function () {
//             hideAlert();
//         }, 10000); // 10 секунд (или другое значение в миллисекундах)
//     });
// });
//
// function showAlert() {
//     const alertContainer = document.getElementById('alert-container');
//     alertContainer.style.display = 'block';
// }

function hideAlert() {
    const alertContainer = document.getElementById('alert-container');
    alertContainer.removeAttribute("data-te-toast-show");
    alertContainer.setAttribute("data-te-toast-hide", "");
}

document.addEventListener('DOMContentLoaded', function () {
    const contactForm = document.getElementById('contact');

    contactForm.addEventListener('submit', function (event) {
        event.preventDefault();
        const formData = new FormData(contactForm);
        var button = document.getElementById("submit");
        button.disabled = true

        fetch('/contact', {
            method: 'POST',
            body: formData
        })
            .then(response => {
                if (!response.ok) {
                    throw new Error('Network response was not ok');
                }
            })
            .then(data => {
                showAlert('success', data.message);
            })
            .catch(error => {
                showAlert("error", error.message);
            })
        function showAlert(type, message) {
            //TODO add error alert
            const alertContainer = document.getElementById('alert-container');
            const alertName = document.getElementById('alertName');
            const alertMessage = document.getElementById('alertMessage');
            alertName.innerText = "Contact"
            alertMessage.innerText = "Ticket create successfully."
            alertContainer.removeAttribute("data-te-toast-hide");
            alertContainer.setAttribute("data-te-toast-show", "");
            setTimeout(function () {
                button.disabled = false;
                hideAlert();
            }, 5000);
        }
    });

});