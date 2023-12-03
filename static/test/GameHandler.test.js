import * as helper from "./TestHelper.js";
import {GameHandler} from "../GameHandler.js";

describe('GameHandler should', () => {
    let mockApiWrapper = helper.mockApiWrapper();
    let mockCanvasPainter = helper.mockCanvasPainter();
    let mockInfoPainter = helper.mockInfoPainter();
    let gameHandler;

    beforeEach(() => {
        jest.clearAllMocks();
        gameHandler = new GameHandler(mockApiWrapper, mockCanvasPainter, mockInfoPainter);
    });

    describe('when creating a new game', () => {
        it('use the canvas painter to draw the data obtained from the api wrapper', async () => {
            const mockApiResponse = successfulApiResponse();
            mockApiWrapper.callGetEndpoint.mockResolvedValue(mockApiResponse);

            await gameHandler.randomGame();

            expect(mockApiWrapper.callGetEndpoint).toHaveBeenCalled();
            expect(mockCanvasPainter.drawPlanet).toHaveBeenCalledWith(mockApiResponse.value().Planet);
            expect(mockCanvasPainter.drawObstacles).toHaveBeenCalledWith(mockApiResponse.value().Planet.Obstacles);
            expect(mockCanvasPainter.drawRover).toHaveBeenCalledWith(mockApiResponse.value().Rover);
        });

        it('use the info painter to draw the error obtained from the api wrapper', async () => {
            const mockApiResponse = failedApiResponse();
            mockApiWrapper.callGetEndpoint.mockResolvedValue(mockApiResponse);

            await gameHandler.randomGame();

            expect(mockApiWrapper.callGetEndpoint).toHaveBeenCalled();
            expect(mockInfoPainter.error).toHaveBeenCalledWith(mockApiResponse.value());
        });
    });

    describe('when moving the rover', () => {
        it('use the canvas painter to draw the data obtained from the api wrapper', async () => {
            const mockApiResponse = successfulApiResponse();
            mockApiWrapper.callGetEndpoint.mockResolvedValue(mockApiResponse);
            mockApiWrapper.callMoveEndpoint.mockResolvedValue(mockApiResponse);

            await gameHandler.randomGame();
            await gameHandler.moveRover();

            expect(mockApiWrapper.callMoveEndpoint).toHaveBeenCalled();
            expect(mockCanvasPainter.drawRover).toHaveBeenCalledWith(mockApiResponse.value().Rover);
            expect(mockInfoPainter.warning).toHaveBeenCalledWith(mockApiResponse.value().Errors);
        });

        it('error when trying to move a rover before it is created', async () => {
            await gameHandler.moveRover();

            expect(mockApiWrapper.callMoveEndpoint).not.toHaveBeenCalled();
            expect(mockCanvasPainter.drawRover).not.toHaveBeenCalled();
        });

        it('use the info painter to draw the error obtained from the api wrapper', async () => {
            const mockApiResponse = failedApiResponse();
            mockApiWrapper.callGetEndpoint.mockResolvedValue(successfulApiResponse());
            mockApiWrapper.callMoveEndpoint.mockResolvedValue(mockApiResponse);

            await gameHandler.randomGame();
            await gameHandler.moveRover();

            expect(mockApiWrapper.callGetEndpoint).toHaveBeenCalled();
            expect(mockApiWrapper.callMoveEndpoint).toHaveBeenCalled();
            expect(mockInfoPainter.error).toHaveBeenCalledWith(mockApiResponse.value());
        });
    });
});

function successfulApiResponse() {
    return {
        value: () => ({
            Rover: {Id: 'aRoverId', Coordinate: {X: 1, Y: 2}},
            Planet: "aPlanet",
            Errors: "anError"
        }),
        isFailure: () => false,
    };
}

function failedApiResponse() {
    return {
        value: () => "sadface :(",
        isFailure: () => true,
    };
}

