import {displayErrors} from "./ErrorPainter.js";

export class RoverHandler {
    #roverId;
    #lastRoverPosition;
    #canvas;
    #apiWrapper;
    #canvasPainter;

    constructor(apiWrapper, canvasPainter) {
        this.#apiWrapper = apiWrapper;
        this.#canvasPainter = canvasPainter;
        this.#canvas = canvasPainter.getCanvas();
    }

    async getRandomRover() {
        const gameData = await this.#apiWrapper.callGetEndpoint();

        this.#roverId = gameData.Rover.Id;
        this.#lastRoverPosition = {
            x: gameData.Rover.Coordinate.X, y: gameData.Rover.Coordinate.Y
        };
        const ctx = this.#canvas.getContext('2d');
        ctx.clearRect(0, 0, this.#canvas.width, this.#canvas.height);
        this.#canvasPainter.drawPlanet(gameData.Planet);
        this.#canvasPainter.drawObstacles(gameData.Planet.Obstacles);
        this.#canvasPainter.drawRover(gameData.Rover);
    }

    async moveRover(commands) {
        if (!this.#roverId) {
            displayErrors('Rover ID not available. Call getRandomRover first.');
            return;
        }

        const moveData = await this.#apiWrapper.callMoveEndpoint(this.#roverId, commands)
        this.#clearPreviousCell();
        let x = moveData.Rover.Coordinate.X;
        let y = moveData.Rover.Coordinate.Y;
        this.#canvasPainter.clearCell(x, y);
        this.#canvasPainter.drawRover(moveData.Rover);
        this.#lastRoverPosition = {x, y};
        displayErrors(moveData.Errors);
    }

    #clearPreviousCell() {
        this.#canvasPainter.clearCell(this.#lastRoverPosition.x, this.#lastRoverPosition.y);
    }
}

