import {ApiWrapper} from "./ApiWrapper.js";
import {CanvasPainter} from './CanvasPainter.js';
import {InfoPainter} from "./InfoPainter.js";
import {EventHandler} from "./EventListenerHandler.js";
import {RoverHandler} from "./RoverHandler.js";

const canvasPainter = new CanvasPainter(document.getElementById('canvas'), 20);
const infoPainter = new InfoPainter(document);
const roverHandler = new RoverHandler(new ApiWrapper(infoPainter), canvasPainter, infoPainter);
const eventHandler = new EventHandler(document, roverHandler);
eventHandler.listenOnReload();
eventHandler.listenOnKeyPress();

window.newGame = function () {
    roverHandler.getNewRoverAndPlanet()
        .then()
        .catch(e => console.log(e));
};
window.move = function () {
    const commands = document.getElementById('commands').value
    roverHandler.moveRover(commands)
        .then()
        .catch(e => console.log(e));
};
