import * as helper from "./TestHelper.js";
import {GameHandler} from "../GameHandler.js";

describe('GameHandler should', () => {
    let mockApiWrapper = helper.mockApiWrapper();
    let mockCanvasPainter = helper.mockCanvasPainter();
    let mockInfoPainter = helper.mockInfoPainter();
    let gameHandler;

    beforeEach(() => {
        gameHandler = new GameHandler(mockApiWrapper, mockCanvasPainter, mockInfoPainter);
        jest.clearAllMocks();
    });

    describe('when creating a new game', () => {
        it('use the canvas painter to draw the data obtained from the api wrapper', async () => {
            const mockApiResponse = {
                Rover: "aRover",
                Planet: "aPlanet",
            };
            mockApiWrapper.callGetEndpoint.mockResolvedValue(mockApiResponse);

            await gameHandler.randomGame();

            expect(mockApiWrapper.callGetEndpoint).toHaveBeenCalled();
            expect(mockCanvasPainter.drawPlanet).toHaveBeenCalledWith(mockApiResponse.Planet);
            expect(mockCanvasPainter.drawObstacles).toHaveBeenCalledWith(mockApiResponse.Planet.Obstacles);
            expect(mockCanvasPainter.drawRover).toHaveBeenCalledWith(mockApiResponse.Rover);
        });
    });

    describe('when moving the rover', () => {
        it('use the canvas painter to draw the data obtained from the api wrapper', async () => {
            const mockApiResponse = {
                Rover: {Id: 'aRoverId', Coordinate: {X: 1, Y: 2}},
                Planet: "aPlanet",
                Errors: "anError"
            };
            mockApiWrapper.callGetEndpoint.mockResolvedValue(mockApiResponse);
            mockApiWrapper.callMoveEndpoint.mockResolvedValue(mockApiResponse);

            await gameHandler.randomGame();
            await gameHandler.moveRover();

            expect(mockApiWrapper.callMoveEndpoint).toHaveBeenCalled();
            expect(mockCanvasPainter.drawRover).toHaveBeenCalledWith(mockApiResponse.Rover);
            expect(mockInfoPainter.warning).toHaveBeenCalledWith(mockApiResponse.Errors);
        });

        it('error when trying to move a rover before it is created', async () => {
            await gameHandler.moveRover();

            expect(mockApiWrapper.callMoveEndpoint).not.toHaveBeenCalled();
            expect(mockCanvasPainter.drawRover).not.toHaveBeenCalled();
        });
    });

});

