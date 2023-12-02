export function newMockContext() {
    return {
        fillRect: jest.fn(),
        strokeRect: jest.fn(),
        beginPath: jest.fn(),
        moveTo: jest.fn(),
        lineTo: jest.fn(),
        fill: jest.fn(),
        save: jest.fn(),
        translate: jest.fn(),
        rotate: jest.fn(),
        closePath: jest.fn(),
        restore: jest.fn(),
        clearRect: jest.fn(),
    };
}

export function newMockCanvas(context) {
    return {
        getContext: jest.fn(() => context),
        width: 0,
        height: 0,
    };
}

export function newMockPlanet(size) {
    return {
        Width: size,
        Height: size,
        Obstacles: [],
    };
}

export function newMockRover() {
    return {
        Coordinate: {X: 2, Y: 3},
        Direction: 'N',
    };
}

export function newMockRoverFacing(direction) {
    return {
        Coordinate: {X: 2, Y: 3},
        Direction: direction,
    };
}
