import {StorageWrapper} from "./StorageWrapper.js";

export class EventHandler {
    #dom;
    #window;
    #gameHandler;

    constructor(dom, window, gameHandler) {
        this.#dom = dom;
        this.#window = window;
        this.#gameHandler = gameHandler;
    }

    listenOnReload() {
        this.#window.onload = () => {
            const storedRoverId = StorageWrapper.getRoverId();
            if (storedRoverId === null) {
                this.#gameHandler.randomGame();
            } else {
                this.#gameHandler.randomGame();
                // TODO: uncomment once loadGame endpoint is implemented
                // this.#gameHandler.loadGame(storedRoverId);
            }
        }
    }

    listenOnKeyPress() {
        this.#dom.addEventListener('keydown', async (event) => { // NOSONAR
            switch (event.key) {
                case 'ArrowUp':
                case 'k':
                    await this.#gameHandler.moveRover('f');
                    break;
                case 'ArrowDown':
                case 'j':
                    await this.#gameHandler.moveRover('b');
                    break;
                case 'ArrowLeft':
                case 'h':
                    await this.#gameHandler.moveRover('l');
                    break;
                case 'ArrowRight':
                case 'l':
                    await this.#gameHandler.moveRover('r');
                    break;
                default:
                    break;
            }
        });
    }
}
