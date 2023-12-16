import {ApiWrapper} from "../api/ApiWrapper.js";
import {StorageWrapper} from "./StorageWrapper.js";

export class GameHandler {
    #lastRoverPosition;
    #canvasPainter;
    #logger;

    constructor(canvasPainter, logger) {
        this.#canvasPainter = canvasPainter;
        this.#logger = logger;
    }

    async randomGame() {
        this.#logger.resetLogMessages()
        const apiResult = await ApiWrapper.postRandomGame()
        this.#handleNewGameResult(apiResult);
    }

    async loadGame(roverId) {
        this.#logger.resetLogMessages()
        const apiResult = await ApiWrapper.getLoadGame(roverId)
        this.#handleNewGameResult(apiResult);
    }

    #handleNewGameResult(apiResult) {
        if (apiResult.isFailure()) {
            this.#logger.error(apiResult.value());
            return;
        }

        const gameData = apiResult.value();
        StorageWrapper.setRoverId(gameData.Rover.Id);
        this.#lastRoverPosition = gameData.Rover.Coordinate;
        this.#paintGame(gameData);
    }

    #paintGame(gameData) {
        this.#canvasPainter.drawPlanet(gameData.Planet);
        this.#canvasPainter.drawObstacles(gameData.Planet.Obstacles);
        this.#canvasPainter.drawRover(gameData.Rover.Direction, gameData.Rover.Coordinate);
    }

    async moveRover(commands) {
        this.#logger.resetLogMessages()
        const roverId = StorageWrapper.getRoverId();
        if (roverId === null) {
            this.#logger.error('Rover ID not available. Call getRandomRover first.');
            return;
        }
        const apiResult = await ApiWrapper.postMoveRover(roverId, commands)
        if (apiResult.isFailure()) {
            this.#logger.error(apiResult.value());
            return;
        }

        const movementData = apiResult.value();
        if (movementData.Results.length !== commands.length) {
            this.#logger.warning("Invalid commands skipped!");
        }
        for (const result of movementData.Results) {
            this.#clearCell(result.Coordinate);
            this.#canvasPainter.drawRover(result.Direction, result.Coordinate);
            this.#lastRoverPosition = result.Coordinate;
            this.#logger.warning(result.Issue);
            await this.#sleep(200);
        }
    }

    #sleep(milliseconds) {
        return new Promise(resolve => setTimeout(resolve, milliseconds));
    }

    #clearCell(coordinate) {
        this.#canvasPainter.clearCell({
            x: this.#lastRoverPosition.X,
            y: this.#lastRoverPosition.Y
        });
        this.#canvasPainter.clearCell(coordinate);
    }
}

