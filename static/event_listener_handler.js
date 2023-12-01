import {getRandomRover, moveRover} from "./script.js";

export class EventHandler {
    #dom;
    #canvas;
    constructor(dom) {
        this.#dom = dom;
        this.#canvas = dom.getElementById('canvas');
    }

    listenOnReload() {
        this.#dom.addEventListener('DOMContentLoaded', () => getRandomRover(this.#canvas)); // NOSONAR
    }

    listenOnKeyPress() {
        this.#dom.addEventListener('keydown', async (event) => { // NOSONAR
            switch (event.key) {
                case 'ArrowUp':
                case 'k':
                    await moveRover('f', this.#canvas);
                    break;
                case 'ArrowDown':
                case 'j':
                    await moveRover('b', this.#canvas);
                    break;
                case 'ArrowLeft':
                case 'h':
                    await moveRover('l', this.#canvas);
                    break;
                case 'ArrowRight':
                case 'l':
                    await moveRover('r', this.#canvas);
                    break;
                default:
                    break;
            }
        });
    }
}
