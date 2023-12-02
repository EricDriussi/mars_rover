export class ErrorPainter {
    #dom;

    constructor(dom) {
        this.#dom = dom;
    }

    // TODO: movement errors should be treated differently than bad requests (sending invalid commands)
    displayErrors(errors) {
        const errorBox = this.#dom.getElementById('error-box');
        const errorList = this.#dom.getElementById('error-list');

        errorList.innerHTML = '';

        if (errors && errors.length > 0) {
            errorBox.classList.remove('hidden');
            errors.forEach(error => {
                const listItem = this.#dom.createElement('li');
                listItem.textContent = error;
                errorList.appendChild(listItem);
            });
        } else {
            errorBox.classList.add('hidden');
        }
    }
}