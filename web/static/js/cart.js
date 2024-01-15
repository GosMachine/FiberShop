document.addEventListener('DOMContentLoaded', function () {
    var couponBtn = document.getElementById('couponBtn');
    couponBtn.addEventListener('click', function() {
        couponBtn.disabled = true;
        const code = document.getElementById('code').value;
        const formData = new FormData();
        formData.append('code', code);
        sendRequest(formData, "/coupon/discount", couponBtn, "Coupon", couponHandle, couponFinallyHandle)
    });
});

var percentage = 0

function couponHandle(type, title, data) {
    var errorMessage = document.getElementById('couponError');
    if (type === "success") {
        if (data.type === "percentage") {
            document.getElementById("couponValue").innerText = data.value
            percentage = data.value
        } else if (data.type === "fixed_amount") {
            document.getElementById("couponValue").innerText = data.value
        }
        errorMessage.style.display = "none"
        document.getElementById("code").disabled = true
        updateTotalCartPrice()
        document.getElementById('couponSuccess').innerText = "Coupon applied."
    } else {
        if (data === "InvalidCoupon") {
            data = "No coupon found."
        }
        errorMessage.innerText = data
    }
}


function couponFinallyHandle(type, btn) {
    if (type === "error") {
        setTimeout(() => {
            btn.disabled = false;
        }, 1000);
    }
}

function incrementQuantity(id) {
    var input = document.getElementById('quantityInput_' + id);
    input.value = parseInt(input.value, 10) + 1;
    updateTotalPrice(id);
}

function decrementQuantity(id) {
    var input = document.getElementById('quantityInput_' + id);
    var newValue = parseInt(input.value, 10) - 1;
    input.value = newValue >= 1 ? newValue : 1;
    updateTotalPrice(id);
}

function updateTotalPrice(id) {
    var input = document.getElementById('quantityInput_' + id);
    var totalPriceElement = document.getElementById('totalPrice_' + id);

    var quantity = parseInt(input.value, 10);
    var price = parseFloat(input.getAttribute('data-price'));

    var totalPrice = quantity * price;
    totalPriceElement.textContent = totalPrice.toFixed(2);


    updateTotalCartPrice();
}


function updateTotalCartPrice() {

    var subTotalCartPriceElement = document.getElementById('subTotalCartPrice');
    var totalCartPriceElement = document.getElementById('totalCartPrice');
    var totalCartPrice = 0;

    var cartItemElements = document.querySelectorAll('[id^="totalPrice_"]');
    cartItemElements.forEach(function (item) {
        totalCartPrice += parseFloat(item.textContent);
    });

    subTotalCartPriceElement.textContent = totalCartPrice.toFixed(2);
    var couponValue = document.getElementById("couponValue");
    if (percentage !== 0) {
        couponValue.textContent = ((totalCartPrice * percentage) / 100).toFixed(2)
    }
    totalCartPrice -= couponValue.textContent;
    totalCartPriceElement.textContent = totalCartPrice.toFixed(2);
}

//TODO пофиксить NaN
