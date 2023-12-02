import {EventHandler} from '../EventHandler.js';
import * as helper from "./TestHelper.js";

describe('EventListener should', () => {
    let mockDom;
    let eventHandler;
    let mockRoverHandler;

    beforeEach(() => {
        mockDom = helper.mockDom();
        mockRoverHandler = helper.mockRoverHandler();
        eventHandler = new EventHandler(mockDom, mockRoverHandler);
    });

    it('listen on reload', () => {
        eventHandler.listenOnReload();

        expect(mockDom.addEventListener).toHaveBeenCalledWith('DOMContentLoaded', expect.any(Function));
        // Hack to get the function from within the callback passed to addEventListener
        const eventCallback = mockDom.addEventListener.mock.calls[0][1];
        eventCallback();
        expect(mockRoverHandler.getNewRoverAndPlanet).toHaveBeenCalled();
    });

    it.each([
        ['ArrowUp', 'f'],
        ['k', 'f'],
        ['ArrowDown', 'b'],
        ['j', 'b'],
        ['ArrowLeft', 'l'],
        ['h', 'l'],
        ['ArrowRight', 'r'],
        ['l', 'r'],
    ])(
        'listen on keydown event with key %s',
        async (givenKey, expectedDirection) => {
            eventHandler.listenOnKeyPress();

            expect(mockDom.addEventListener).toHaveBeenCalledWith('keydown', expect.any(Function));
            // Hack to get the function from within the callback passed to addEventListener
            const eventCallback = mockDom.addEventListener.mock.calls[0][1];
            eventCallback({key: givenKey});
            expect(mockRoverHandler.moveRover).toHaveBeenCalledWith(expectedDirection);
        }
    );

    it('listen on keydown event and do nothing with unknown key', () => {
        eventHandler.listenOnKeyPress();

        expect(mockDom.addEventListener).toHaveBeenCalledWith('keydown', expect.any(Function));
        // Hack to get the function from within the callback passed to addEventListener
        const eventCallback = mockDom.addEventListener.mock.calls[0][1];
        eventCallback({ key: 'wrong!' });
        expect(mockRoverHandler.moveRover).not.toHaveBeenCalled();

    });
});
