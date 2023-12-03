import {EventHandler} from '../EventHandler.js';
import * as helper from "./TestHelper.js";

describe('EventListener should', () => {
    let mockDom;
    let eventHandler;
    let mockGameHandler;

    beforeEach(() => {
        jest.clearAllMocks();
        mockDom = helper.mockDom();
        mockGameHandler = helper.mockGameHandler();
        eventHandler = new EventHandler(mockDom, mockGameHandler);
    });

    it('listen on reload', () => {
        eventHandler.listenOnReload();

        expect(mockDom.addEventListener).toHaveBeenCalledWith('DOMContentLoaded', expect.any(Function));
        // Hack to get the function from within the callback passed to addEventListener
        const eventCallback = mockDom.addEventListener.mock.calls[0][1];
        eventCallback();
        expect(mockGameHandler.randomGame).toHaveBeenCalled();
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
            expect(mockGameHandler.moveRover).toHaveBeenCalledWith(expectedDirection);
        }
    );

    it('listen on keydown event and do nothing with unknown key', () => {
        eventHandler.listenOnKeyPress();

        expect(mockDom.addEventListener).toHaveBeenCalledWith('keydown', expect.any(Function));
        // Hack to get the function from within the callback passed to addEventListener
        const eventCallback = mockDom.addEventListener.mock.calls[0][1];
        eventCallback({ key: 'wrong!' });
        expect(mockGameHandler.moveRover).not.toHaveBeenCalled();

    });
});
