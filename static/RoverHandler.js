export class RoverHandler {
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

    async getNewRoverAndPlanet() {
        const gameData = await this.#apiWrapper.callGetEndpoint();
        this.#roverId = gameData.Rover.Id;
        this.#lastRoverPosition = gameData.Rover.Coordinate;

        this.#canvasPainter.drawPlanet(gameData.Planet);
        this.#canvasPainter.drawObstacles(gameData.Planet.Obstacles);
        this.#canvasPainter.drawRover(gameData.Rover);
    }

    async moveRover(commands) {
        if (!this.#roverId) {
            this.#infoPainter.errors('Rover ID not available. Call getRandomRover first.');
            return;
        }

        const movementData = await this.#apiWrapper.callMoveEndpoint(this.#roverId, commands)
        this.#clearCell(movementData.Rover.Coordinate);
        this.#canvasPainter.drawRover(movementData.Rover);
        this.#lastRoverPosition = movementData.Rover.Coordinate;
        this.#infoPainter.errors(movementData.Errors);
    }

    #clearCell(coordinate) {
        this.#canvasPainter.clearCell({
            x: this.#lastRoverPosition.X,
            y: this.#lastRoverPosition.Y
        });
        this.#canvasPainter.clearCell(coordinate);
    }
}

