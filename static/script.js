import {CanvasPainter} from './CanvasPainter.js';
import {Logger} from "./Logger.js";
import {EventHandler} from "./EventHandler.js";
import {GameHandler} from "./GameHandler.js";

const canvasPainter = new CanvasPainter(document.getElementById('canvas'), 20);
const gameHandler = new GameHandler(canvasPainter, new Logger(document));
const eventHandler = new EventHandler(document, gameHandler);
eventHandler.listenOnReload();
eventHandler.listenOnKeyPress();

window.newGame = function () {
    gameHandler.randomGame()
        .then()
        .catch(e => console.log(e));
};
window.move = function () {
    const commands = document.getElementById('commands').value
    gameHandler.moveRover(commands)
        .then()
        .catch(e => console.log(e));
};
