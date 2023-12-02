import {ApiWrapper} from "./ApiWrapper.js";
import {CanvasPainter} from './CanvasPainter.js';
import {InfoPainter} from "./InfoPainter.js";
import {EventHandler} from "./EventHandler.js";
import {GameHandler} from "./GameHandler.js";

const canvasPainter = new CanvasPainter(document.getElementById('canvas'), 20);
const infoPainter = new InfoPainter(document);
const roverHandler = new GameHandler(new ApiWrapper(infoPainter), canvasPainter, infoPainter);
const eventHandler = new EventHandler(document, roverHandler);
eventHandler.listenOnReload();
eventHandler.listenOnKeyPress();

window.newGame = function () {
    roverHandler.randomGame()
        .then()
        .catch(e => console.log(e));
};
window.move = function () {
    const commands = document.getElementById('commands').value
    roverHandler.moveRover(commands)
        .then()
        .catch(e => console.log(e));
};
