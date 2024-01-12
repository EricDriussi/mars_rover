export class Logger {
    #dom;

    constructor(dom) {
        this.#dom = dom;
    }

    error(errors) {
        this.#render('error', errors)
    }

    warning(warnings) {
        this.#render('warn', warnings)
    }

    // TODO.LM: This is kind of a hack, it should not be here
    populateRoverId(roverId) {
        this.#dom.getElementById('currentRoverId').innerHTML = roverId;
    }

    #render(type, messages) {
        if (!messages || messages.length === 0) {
            return;
        }

        const messageList = this.#dom.getElementById(`${type}-list`);
        const messageBox = this.#dom.getElementById(`${type}-box`);
        this.#paintMessages(messages, messageList);
        this.#reveal(messageBox);
    }


    resetLogMessages() {
        this.#resetMessages('error')
        this.#resetMessages('warn')
    }

    #resetMessages(type) {
        const listId = `${type}-list`;
        const boxId = `${type}-box`;
        this.#clear(this.#dom.getElementById(listId));
        this.#hide(this.#dom.getElementById(boxId));
    }

    #paintMessages(messages, list) {
        messages = Array.isArray(messages) ? messages : [messages];
        messages.forEach(message => {
            const listItem = this.#dom.createElement('li');
            listItem.textContent = message;
            list.appendChild(listItem);
        });
    }

    #clear(list) {
        list.innerHTML = '';
    }

    #reveal(box) {
        box.classList.remove('hidden');
    }

    #hide(box) {
        box.classList.add('hidden');
    }
}
