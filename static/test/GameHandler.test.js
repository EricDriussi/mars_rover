import * as helper from "./TestHelper.js";
import {GameHandler} from "../handlers/GameHandler.js";
import {ApiWrapper} from "../api/ApiWrapper.js";

jest.mock("../api/ApiWrapper.js");

describe('GameHandler should', () => {
    let mockCanvasPainter = helper.mockCanvasPainter();
    let mockLogger = helper.mockLogger();
    let gameHandler;

    beforeEach(() => {
        jest.clearAllMocks();
        gameHandler = new GameHandler(mockCanvasPainter, mockLogger);
    });

    describe('when creating a new game', () => {
        it('use the canvas painter to draw the data obtained from the api wrapper', async () => {
            const mockApiResponse = genericSuccessfulApiResponse();
            ApiWrapper.postRandomGame.mockResolvedValue(mockApiResponse);

            await gameHandler.randomGame();

            expect(mockCanvasPainter.drawPlanet).toHaveBeenCalledWith(mockApiResponse.value().Planet);
            expect(mockCanvasPainter.drawObstacles).toHaveBeenCalledWith(mockApiResponse.value().Planet.Obstacles);
            expect(mockCanvasPainter.drawRover).toHaveBeenCalledWith(mockApiResponse.value().Rover);
        });

        it('use the info painter to draw the error obtained from the api wrapper', async () => {
            const mockApiResponse = genericFailedApiResponse();
            ApiWrapper.postRandomGame.mockResolvedValue(mockApiResponse);

            await gameHandler.randomGame();

            expect(mockLogger.error).toHaveBeenCalledWith(mockApiResponse.value());
        });
    });

    describe('when moving the rover', () => {
        it('use the canvas painter to draw the data obtained from the api wrapper', async () => {
            const mockApiResponse = genericSuccessfulApiResponse();
            ApiWrapper.postRandomGame.mockResolvedValue(mockApiResponse);
            ApiWrapper.postMoveRover.mockResolvedValue(mockApiResponse);

            await gameHandler.randomGame();
            await gameHandler.moveRover();

            expect(mockCanvasPainter.drawRover).toHaveBeenCalledWith(mockApiResponse.value().Rover);
            expect(mockLogger.warning).toHaveBeenCalledWith(mockApiResponse.value().Errors);
        });

        it('error when trying to move a rover before it is created', async () => {
            await gameHandler.moveRover();

            expect(mockCanvasPainter.drawRover).not.toHaveBeenCalled();
        });

        it('use the info painter to draw the error obtained from the api wrapper', async () => {
            const mockApiResponse = genericFailedApiResponse();
            ApiWrapper.postRandomGame.mockResolvedValue(mockApiResponse);
            ApiWrapper.postMoveRover.mockResolvedValue(mockApiResponse);

            await gameHandler.randomGame();
            await gameHandler.moveRover();

            expect(mockLogger.error).toHaveBeenCalledWith(mockApiResponse.value());
        });
    });

    describe('when loading a game', () => {
        it('use the canvas painter to draw the data obtained from the api wrapper', async () => {
            const mockApiResponse = genericSuccessfulApiResponse();
            ApiWrapper.getLoadGame.mockResolvedValue(mockApiResponse);

            await gameHandler.loadGame();

            expect(mockCanvasPainter.drawPlanet).toHaveBeenCalledWith(mockApiResponse.value().Planet);
            expect(mockCanvasPainter.drawObstacles).toHaveBeenCalledWith(mockApiResponse.value().Planet.Obstacles);
            expect(mockCanvasPainter.drawRover).toHaveBeenCalledWith(mockApiResponse.value().Rover);
        });

        it('use the info painter to draw the error obtained from the api wrapper', async () => {
            const mockApiResponse = genericFailedApiResponse();
            ApiWrapper.getLoadGame.mockResolvedValue(mockApiResponse);

            await gameHandler.loadGame();

            expect(mockLogger.error).toHaveBeenCalledWith(mockApiResponse.value());
        });
    });

    it('reset logs on every operation', async () => {
        const mockApiResponse = genericFailedApiResponse();
        ApiWrapper.getLoadGame.mockResolvedValue(mockApiResponse);

        await gameHandler.loadGame('123');
        await gameHandler.randomGame();
        await gameHandler.moveRover('321');

        expect(mockLogger.resetLogMessages).toHaveBeenCalledTimes(3);
    });
});

function genericSuccessfulApiResponse() {
    return {
        value: () => ({
            Rover: {Id: 'aRoverId', Coordinate: {X: 1, Y: 2}},
            Planet: "aPlanet",
            Errors: "anError"
        }),
        isFailure: () => false,
    };
}

function genericFailedApiResponse() {
    return {
        value: () => "sadface :(",
        isFailure: () => true,
    };
}

