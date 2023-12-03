export class Logger {
    #dom;

    constructor(dom) {
        this.#dom = dom;
    }

    error(errors) {
        this.#render('error', errors);
    }

    warning(warnings) {
        this.#render('warn', warnings);
    }

    #render(type, messages) {
        const listId = `${type}-list`;
        const boxId = `${type}-box`;

        const messageList = this.#dom.getElementById(listId);
        this.#clear(messageList);

        const messageBox = this.#dom.getElementById(boxId);
        if (!messages || messages.length === 0) {
            this.#hide(messageBox);
            return;
        }

        this.#paintMessages(messages, messageList);
        this.#reveal(messageBox);
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
