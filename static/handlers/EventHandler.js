export class EventHandler {
    #dom;
    #roverHandler;
    constructor(dom, roverHandler) {
        this.#dom = dom;
        this.#roverHandler = roverHandler;
    }

    listenOnReload() {
        this.#dom.addEventListener('DOMContentLoaded', () => this.#roverHandler.randomGame()); // NOSONAR
    }

    listenOnKeyPress() {
        this.#dom.addEventListener('keydown', async (event) => { // NOSONAR
            switch (event.key) {
                case 'ArrowUp':
                case 'k':
                    await this.#roverHandler.moveRover('f');
                    break;
                case 'ArrowDown':
                case 'j':
                    await this.#roverHandler.moveRover('b');
                    break;
                case 'ArrowLeft':
                case 'h':
                    await this.#roverHandler.moveRover('l');
                    break;
                case 'ArrowRight':
                case 'l':
                    await this.#roverHandler.moveRover('r');
                    break;
                default:
                    break;
            }
        });
    }
}
