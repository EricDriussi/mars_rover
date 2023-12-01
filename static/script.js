import {CanvasPainter} from './canvas_painter.js';
import {displayErrors} from './error_painter.js';
import {ApiWrapper} from "./api_wrapper.js";
import {EventHandler} from "./event_listener_handler.js";

const eventHandler = new EventHandler(document);
eventHandler.listenOnReload();
eventHandler.listenOnKeyPress();

let roverId;
let lastRoverPosition;

export async function getRandomRover(canvas) {
    const apiWrapper = new ApiWrapper();
    const gameData = await apiWrapper.callGetEndpoint();

    roverId = gameData.Rover.Id;
    lastRoverPosition = {
        x: gameData.Rover.Coordinate.X, y: gameData.Rover.Coordinate.Y
    };
    const ctx = canvas.getContext('2d');
    ctx.clearRect(0, 0, canvas.width, canvas.height);
    const roverDrawer = new CanvasPainter(canvas, 20);
    roverDrawer.drawPlanet(gameData.Planet);
    roverDrawer.drawObstacles(gameData.Planet.Obstacles);
    roverDrawer.drawRover(gameData.Rover);
}

export async function moveRover(commands, canvas) {
    if (!roverId) {
        displayErrors('Rover ID not available. Call getRandomRover first.');
        return;
    }

    const apiWrapper = new ApiWrapper();
    const moveData = await apiWrapper.callMoveEndpoint(roverId, commands)
    const roverDrawer = new CanvasPainter(canvas, 20);
    clearPreviousCell();
    let x = moveData.Rover.Coordinate.X;
    let y = moveData.Rover.Coordinate.Y;
    roverDrawer.clearCell(x, y);
    roverDrawer.drawRover(moveData.Rover);
    lastRoverPosition = {x, y};
    displayErrors(moveData.Errors);
}

function clearPreviousCell() {
    const canvas = document.getElementById('canvas');
    const roverDrawer = new CanvasPainter(canvas, 20);
    roverDrawer.clearCell(lastRoverPosition.x, lastRoverPosition.y);
}


window.random = function (dom) {
    getRandomRover(dom.getElementById('canvas'))
        .then()
        .catch(e => console.log(e));
};
window.move = function (dom) {
    const commands = dom.getElementById('commands').value
    const canvas = dom.getElementById('canvas');
    moveRover(commands, canvas)
        .then()
        .catch(e => console.log(e));
};
