import {CanvasPainter} from './CanvasPainter.js';
import {Logger} from "./handlers/Logger.js";
import {EventHandler} from "./handlers/EventHandler.js";
import {GameHandler} from "./handlers/GameHandler.js";

const canvasPainter = new CanvasPainter(document.getElementById('canvas'), 20);
const gameHandler = new GameHandler(canvasPainter, new Logger(document));
const eventHandler = new EventHandler(document, window, gameHandler);
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
window.loadGame = function () {
    const roverId = document.getElementById('roverId').value
    gameHandler.loadGame(roverId)
        .then()
        .catch(e => console.log(e));
};
