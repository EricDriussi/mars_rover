export class GameHandler {
    #roverId;
    #lastRoverPosition;
    #apiWrapper;
    #canvasPainter;
    #infoPainter;

    constructor(apiWrapper, canvasPainter, infoPainter) {
        this.#apiWrapper = apiWrapper;
        this.#canvasPainter = canvasPainter;
        this.#infoPainter = infoPainter;
    }

    async randomGame() {
        const apiResult = await this.#apiWrapper.postRandomGame();
        if (apiResult.isFailure()) {
            this.#infoPainter.error(apiResult.value());
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
            this.#infoPainter.error('Rover ID not available. Call getRandomRover first.');
            return;
        }
        const apiResult = await this.#apiWrapper.postMoveRover(this.#roverId, commands)
        if (apiResult.isFailure()) {
            this.#infoPainter.error(apiResult.value());
            return;
        }

        const movementData = apiResult.value();
        this.#clearCell(movementData.Rover.Coordinate);
        this.#canvasPainter.drawRover(movementData.Rover);
        this.#lastRoverPosition = movementData.Rover.Coordinate;
        this.#infoPainter.warning(movementData.Errors);
    }

    #clearCell(coordinate) {
        this.#canvasPainter.clearCell({
            x: this.#lastRoverPosition.X,
            y: this.#lastRoverPosition.Y
        });
        this.#canvasPainter.clearCell(coordinate);
    }
}

