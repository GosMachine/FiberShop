var percentage = 0;

function updateTotalPrice(id) {
    var input = document.getElementById('quantityInput_' + id);
    var totalPriceElement = document.getElementById('totalPrice_' + id);

    var quantity = parseInt(input.value, 10);
    var price = parseFloat(input.getAttribute('data-price'));

    var totalPrice = quantity * price;
    totalPriceElement.textContent = totalPrice.toFixed(2);

    updateTotalCartPrice()
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