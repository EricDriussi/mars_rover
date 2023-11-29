import {CanvasPainter} from '../painter.js';

describe('CanvasPainter', () => {

    const mockContext = {
        fillStyle: null,
        strokeStyle: null,
        fillRect: jest.fn(),
        strokeRect: jest.fn(),
        beginPath: jest.fn(),
        moveTo: jest.fn(),
        lineTo: jest.fn(),
        fill: jest.fn(),
    };

    const mockCanvas = {
        getContext: jest.fn(() => mockContext),
        width: 0,
        height: 0,
    };

    const mockPlanetSize = 5;
    const mockPlanet = {
        Width: mockPlanetSize,
        Height: mockPlanetSize,
        Obstacles: [],
    };

    const mockRover = {
        Coordinate: {X: 2, Y: 3},
        Direction: 'N',
    };

    beforeEach(() => {
        mockCanvas.getContext.mockClear();
    });

    const cellSize = 20;
    const canvasPainter = new CanvasPainter(mockCanvas, cellSize);

    it('should draw a planet with expected size', () => {
        canvasPainter.drawPlanet(mockPlanet);

        const expectedSize = mockPlanetSize * cellSize;
        expect(mockCanvas.width).toBe(expectedSize);
        expect(mockCanvas.height).toBe(expectedSize);
        expect(mockCanvas.getContext().fillRect).toHaveBeenCalledWith(0, 0, expectedSize, expectedSize);
        const expectedCells = mockPlanetSize * mockPlanetSize;
        expect(mockCanvas.getContext().strokeRect).toHaveBeenCalledTimes(expectedCells);
    });

    it('should draw a white planet with lightgrey cell borders', () => {
        canvasPainter.drawPlanet(mockPlanet);

        expect(mockCanvas.getContext().fillStyle).toEqual('white');
        expect(mockCanvas.getContext().strokeStyle).toEqual('lightgrey');
    });

    it('should draw a red rover', () => {
        canvasPainter.drawRover(mockRover);

        expect(mockCanvas.getContext().fillStyle).toEqual('red');
        expect(mockCanvas.getContext().fill).toHaveBeenCalled();
    });

    it('should draw a rover in the expected position', () => {
        canvasPainter.drawRover(mockRover);

        const expectedX = mockRover.Coordinate.X * cellSize;
        const expectedY = mockCanvas.height - (mockRover.Coordinate.Y + 1) * cellSize;
        expect(mockCanvas.getContext().moveTo).toHaveBeenCalledWith(expectedX, expectedY);
        expect(mockCanvas.getContext().fill).toHaveBeenCalled();
    });

    it('should draw obstacles correctly', () => {
        const pos1 = {X: 1, Y: 1};
        const pos2 = {X: 2, Y: 2};
        const planetWithObstacles = {
            ...mockPlanet, Obstacles: [
                {Coordinate: [pos1]},
                {Coordinate: [pos2]}
            ]
        };
        canvasPainter.drawObstacles(planetWithObstacles);
        expect(mockCanvas.getContext().fillStyle).toEqual('black');
        expect(mockCanvas.getContext().fillRect).toHaveBeenCalledWith(pos1.X * cellSize, mockCanvas.height - (pos1.Y + 1) * cellSize, cellSize, cellSize);
    });

    it('should draw planet and rover correctly', () => {
        canvasPainter.drawPlanetAndRover(mockPlanet, mockRover);
        // Add assertions for the expected calls on the mocked canvas context
        // Add more assertions as needed
    });
});
