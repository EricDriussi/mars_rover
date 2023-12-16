import * as helper from "./TestHelper.js";
import {GameHandler} from "../handlers/GameHandler.js";
import {ApiWrapper} from "../api/ApiWrapper.js";
import {StorageWrapper} from "../handlers/StorageWrapper";

jest.mock("../api/ApiWrapper.js");
jest.mock("../handlers/StorageWrapper.js");

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
            const mockApiResponse = helper.creationApiResponse();
            ApiWrapper.postRandomGame.mockResolvedValue(mockApiResponse);
            StorageWrapper.setRoverId.mockResolvedValue();

            await gameHandler.randomGame();

            expect(mockCanvasPainter.drawPlanet).toHaveBeenCalledWith(mockApiResponse.value().Planet);
            expect(mockCanvasPainter.drawObstacles).toHaveBeenCalledWith(mockApiResponse.value().Planet.Obstacles);
            expect(mockCanvasPainter.drawRover).toHaveBeenCalledWith(mockApiResponse.value().Rover.Direction, mockApiResponse.value().Rover.Coordinate);
        });

        it('use the info painter to draw the error obtained from the api wrapper', async () => {
            const mockApiResponse = helper.failedApiResponse();
            ApiWrapper.postRandomGame.mockResolvedValue(mockApiResponse);

            await gameHandler.randomGame();

            expect(mockLogger.error).toHaveBeenCalledWith(mockApiResponse.value());
        });
    });

    describe('when moving the rover', () => {
        it('use the canvas painter to draw the data obtained from the api wrapper', async () => {
            ApiWrapper.postRandomGame.mockResolvedValue(helper.creationApiResponse());
            const mockApiResponse = helper.movementApiResponse();
            ApiWrapper.postMoveRover.mockResolvedValue(mockApiResponse);

            await gameHandler.randomGame();
            await gameHandler.moveRover("x");

            const movementResult = mockApiResponse.value().Results[0];
            expect(mockCanvasPainter.drawRover).toHaveBeenCalledWith(movementResult.Direction, movementResult.Coordinate);
            expect(mockLogger.warning).toHaveBeenCalledWith(movementResult.Issue);
        });

        it('error when trying to move a rover before it is created', async () => {
            StorageWrapper.getRoverId.mockReturnValueOnce(null);
            await gameHandler.moveRover();

            expect(mockCanvasPainter.drawRover).not.toHaveBeenCalled();
        });

        it('use the info painter to draw the error obtained from the api wrapper', async () => {
            const mockApiResponse = helper.failedApiResponse();
            StorageWrapper.getRoverId.mockReturnValueOnce('aRoverId');
            ApiWrapper.postMoveRover.mockResolvedValue(mockApiResponse);

            await gameHandler.moveRover();

            expect(mockLogger.error).toHaveBeenCalledWith(mockApiResponse.value());
        });
    });

    describe('when loading a game', () => {
        it('use the canvas painter to draw the data obtained from the api wrapper', async () => {
            const mockApiResponse = helper.creationApiResponse();
            ApiWrapper.getLoadGame.mockResolvedValue(mockApiResponse);

            await gameHandler.loadGame();

            expect(mockCanvasPainter.drawPlanet).toHaveBeenCalledWith(mockApiResponse.value().Planet);
            expect(mockCanvasPainter.drawObstacles).toHaveBeenCalledWith(mockApiResponse.value().Planet.Obstacles);
            expect(mockCanvasPainter.drawRover).toHaveBeenCalledWith(mockApiResponse.value().Rover.Direction, mockApiResponse.value().Rover.Coordinate);
        });

        it('use the info painter to draw the error obtained from the api wrapper', async () => {
            const mockApiResponse = helper.failedApiResponse();
            ApiWrapper.getLoadGame.mockResolvedValue(mockApiResponse);

            await gameHandler.loadGame();

            expect(mockLogger.error).toHaveBeenCalledWith(mockApiResponse.value());
        });
    });

    it('reset logs on every operation', async () => {
        await gameHandler.loadGame('123');
        await gameHandler.randomGame();
        await gameHandler.moveRover('321');

        expect(mockLogger.resetLogMessages).toHaveBeenCalledTimes(3);
    });
});
