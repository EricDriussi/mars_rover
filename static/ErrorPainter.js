// TODO: movement errors should be treated differently than bad requests (sending invalid commands)
export function displayErrors(errors) {
    const errorBox = document.getElementById('error-box');
    const errorList = document.getElementById('error-list');

    errorList.innerHTML = '';

    if (errors && errors.length > 0) {
        errorBox.classList.remove('hidden');
        errors.forEach(error => {
            const listItem = document.createElement('li');
            listItem.textContent = error;
            errorList.appendChild(listItem);
        });
    } else {
        errorBox.classList.add('hidden');
    }
}
