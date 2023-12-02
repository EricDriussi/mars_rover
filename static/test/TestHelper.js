import {JSDOM} from "jsdom";

export function mockContext() {
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
    };
}

export function mockCanvas(context) {
    return {
        getContext: jest.fn(() => context),
        width: 0,
        height: 0,
    };
}

export function mockPlanet(size) {
    return {
        Width: size,
        Height: size,
        Obstacles: [],
    };
}

export function mockRover() {
    return {
        Coordinate: {X: 2, Y: 3},
        Direction: 'N',
    };
}

export function mockRoverFacing(direction) {
    return {
        Coordinate: {X: 2, Y: 3},
        Direction: direction,
    };
}

export function mockRoverHandler() {
    return {
        getNewRoverAndPlanet: jest.fn(),
        moveRover: jest.fn(),
    };
}

export function mockDom() {
    const dom = new JSDOM('' +
        '<html><body>' +
        '<div id="error-box" class="hidden">' +
        '<ul id="error-list"></ul>' +
        '</div>' +
        '<div id="warn-box" class="hidden">' +
        '<ul id="warn-list"></ul>' +
        '</div>' +
        '</body></html>'
    );
    Object.assign(dom, {
        addEventListener: jest.fn()
    });
    return dom;
}

export function getElementFrom(mockDom, elementId) {
    return mockDom.window.document.getElementById(elementId);
}

