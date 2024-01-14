function incrementQuantity(id) {
    var input = document.getElementById('quantityInput_' + id);
    input.value = parseInt(input.value, 10) + 1;
}

function decrementQuantity(id) {
    var input = document.getElementById('quantityInput_' + id);
    var newValue = parseInt(input.value, 10) - 1;
    input.value = newValue >= 1 ? newValue : 1;
}