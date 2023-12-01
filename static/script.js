import {CanvasPainter} from './CanvasPainter.js';
import {ApiWrapper} from "./ApiWrapper.js";
import {EventHandler} from "./EventListenerHandler.js";
import {RoverHandler} from "./RoverHandler.js";

const canvasPainter = new CanvasPainter(document.getElementById('canvas'), 20);
const roverHandler = new RoverHandler(new ApiWrapper(), canvasPainter);
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
