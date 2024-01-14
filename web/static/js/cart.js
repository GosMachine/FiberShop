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


    updateTotalCartPrice(); // activate coupon
}


function updateTotalCartPrice() {
    var couponValue = 5.00;

    var subTotalCartPriceElement = document.getElementById('subTotalCartPrice');
    var totalCartPriceElement = document.getElementById('totalCartPrice');
    var totalCartPrice = 0;

    var cartItemElements = document.querySelectorAll('[id^="totalPrice_"]');
    cartItemElements.forEach(function (item) {
        totalCartPrice += parseFloat(item.textContent);
    });

    subTotalCartPriceElement.textContent = totalCartPrice.toFixed(2);

    totalCartPrice -= couponValue;
    totalCartPriceElement.textContent = totalCartPrice.toFixed(2);
}