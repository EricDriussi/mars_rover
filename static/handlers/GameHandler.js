import {ApiWrapper} from "../api/ApiWrapper.js";

export class GameHandler {
    #roverId;
    #lastRoverPosition;
    #canvasPainter;
    #logger;

    constructor(canvasPainter, logger) {
        this.#canvasPainter = canvasPainter;
        this.#logger = logger;
    }

    async randomGame() {
        const apiResult = await ApiWrapper.postRandomGame()
        if (apiResult.isFailure()) {
            this.#logger.error(apiResult.value());
            return;
        }

        const gameData = apiResult.value();
        this.#roverId = gameData.Rover.Id;
        this.#lastRoverPosition = gameData.Rover.Coordinate;
        this.#paintGame(gameData);
    }

    #paintGame(gameData) {
        this.#canvasPainter.drawPlanet(gameData.Planet);
        this.#canvasPainter.drawObstacles(gameData.Planet.Obstacles);
        this.#canvasPainter.drawRover(gameData.Rover);
    }

    async moveRover(commands) {
        if (!this.#roverId) {
            this.#logger.error('Rover ID not available. Call getRandomRover first.');
            return;
        }
        const apiResult = await ApiWrapper.postMoveRover(this.#roverId, commands)
        if (apiResult.isFailure()) {
            this.#logger.error(apiResult.value());
            return;
        }

        const movementData = apiResult.value();
        this.#clearCell(movementData.Rover.Coordinate);
        this.#canvasPainter.drawRover(movementData.Rover);
        this.#lastRoverPosition = movementData.Rover.Coordinate;
        this.#logger.warning(movementData.Errors);
    }

    #clearCell(coordinate) {
        this.#canvasPainter.clearCell({
            x: this.#lastRoverPosition.X,
            y: this.#lastRoverPosition.Y
        });
        this.#canvasPainter.clearCell(coordinate);
    }
}

