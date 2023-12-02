export class InfoPainter {
    #dom;

    constructor(dom) {
        this.#dom = dom;
    }

    // TODO: movement errors should be treated differently than bad requests (sending invalid commands)
    errors(errors) {
        const errorList = this.#dom.getElementById('error-list');
        this.#clear(errorList);

        const errorBox = this.#dom.getElementById('error-box');
        if (!errors || errors.length === 0) {
            this.#hide(errorBox);
            return
        }

        this.#reveal(errorBox);
        this.#paintErrors(errors, errorList);
    }

    #paintErrors(errors, errorList) {
        errors.forEach(error => {
            const listItem = this.#dom.createElement('li');
            listItem.textContent = error;
            errorList.appendChild(listItem);
        });
    }

    #clear(errorList) {
        errorList.innerHTML = '';
    }

    #reveal(errorBox) {
        errorBox.classList.remove('hidden');
    }

    #hide(errorBox) {
        errorBox.classList.add('hidden');
    }
}