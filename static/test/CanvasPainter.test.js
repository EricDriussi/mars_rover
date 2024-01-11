import {CanvasPainter} from '../CanvasPainter.js';
import * as helper from "./TestHelper.js";

describe('CanvasPainter should', () => {
    const mockPlanetSize = 5;
    const mockPlanet = helper.mockPlanet(mockPlanetSize);

    const cellSize = 20;
    const borderSize = 1;
    let canvasPainter;
    let mockCanvas;

    beforeEach(() => {
        jest.clearAllMocks();
        mockCanvas = helper.mockCanvas(helper.mockContext());
        canvasPainter = new CanvasPainter(mockCanvas, cellSize, borderSize);
    });

    describe('draw a planet', () => {
        it('with expected size based on cell size', () => {
            canvasPainter.drawPlanet(mockPlanet);

            const expectedSize = mockPlanetSize * cellSize;
            expect(mockCanvas.width).toBe(expectedSize);
            expect(mockCanvas.height).toBe(expectedSize);
            expect(mockCanvas.getContext().fillRect)
                .toHaveBeenCalledWith(0, 0, expectedSize, expectedSize);
        });

        it('with expected number of cells based on size', () => {
            canvasPainter.drawPlanet(mockPlanet);

            const expectedNumberOfCells = mockPlanetSize * mockPlanetSize;
            expect(mockCanvas.getContext().strokeRect)
                .toHaveBeenCalledTimes(expectedNumberOfCells);
        });

        it('with white background and lightgrey cell borders', () => {
            canvasPainter.drawPlanet(mockPlanet);

            expect(mockCanvas.getContext().fillStyle).toEqual('white');
            expect(mockCanvas.getContext().fillRect).toHaveBeenCalled();
            expect(mockCanvas.getContext().strokeStyle).toEqual('lightgrey');
            expect(mockCanvas.getContext().strokeRect).toHaveBeenCalled();
        });
    });

    describe('draw obstacles', () => {
        const anObstacleXPosition = 1;
        const anObstacleYPosition = 1;
        const anotherObstacleXPosition = 2;
        const anotherObstacleYPosition = 2;
        const obstacles = [
            [{X: anObstacleXPosition, Y: anObstacleYPosition}],
            [{X: anotherObstacleXPosition, Y: anotherObstacleYPosition}]
        ]
        it('at the right position', () => {
            canvasPainter.drawObstacles(obstacles);

            const anExpectedXGridPosition = anObstacleXPosition * cellSize + borderSize;
            const anExpectedYGridPosition = mockCanvas.height - (anObstacleYPosition + 1) * cellSize + borderSize;
            expect(mockCanvas.getContext().fillRect)
                .toHaveBeenCalledWith(
                    anExpectedXGridPosition,
                    anExpectedYGridPosition,
                    cellSize - borderSize * 2,
                    cellSize - borderSize * 2);
            const anotherExpectedXGridPosition = anotherObstacleXPosition * cellSize + borderSize;
            const anotherExpectedYGridPosition = mockCanvas.height - (anotherObstacleYPosition + 1) * cellSize + borderSize;
            expect(mockCanvas.getContext().fillRect)
                .toHaveBeenCalledWith(
                    anotherExpectedXGridPosition,
                    anotherExpectedYGridPosition,
                    cellSize - borderSize * 2,
                    cellSize - borderSize * 2);
        });
    })

    describe('draw a rover', () => {
        const mockRover = helper.mockRover();

        it('painted red', () => {
            canvasPainter.drawRover(mockRover.Direction, mockRover.Coordinate);

            expect(mockCanvas.getContext().fillStyle).toEqual('red');
            expect(mockCanvas.getContext().fill).toHaveBeenCalled();
        });

        it('at the right position', () => {
            canvasPainter.drawRover(mockRover.Direction, mockRover.Coordinate);

            const roverXGridPosition = mockRover.Coordinate.X * cellSize;
            const roverYGridPosition = mockCanvas.height - (mockRover.Coordinate.Y + 1) * cellSize;
            const halfCellSize = cellSize / 2;
            const centeredRoverXGridPosition = roverXGridPosition + halfCellSize;
            const centeredRoverYGridPosition = roverYGridPosition + halfCellSize;
            expect(mockCanvas.getContext().translate).toHaveBeenCalledWith(
                centeredRoverXGridPosition,
                centeredRoverYGridPosition);
            expect(mockCanvas.getContext().fill).toHaveBeenCalled();
        });

        it.each([
            ['N', Math.PI],
            ['S', 0],
            ['E', 3 * Math.PI / 2],
            ['W', Math.PI / 2],
        ])('facing %s', (direction, expectedRotation) => {
            const mockRover = helper.mockRoverFacing(direction);
            canvasPainter.drawRover(mockRover.Direction, mockRover.Coordinate);

            expect(mockCanvas.getContext().rotate).toHaveBeenCalledWith(expectedRotation);
        });
    });
});
