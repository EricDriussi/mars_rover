import {EventHandler} from '../handlers/EventHandler.js';
import * as helper from "./TestHelper.js";
import {StorageWrapper} from "../handlers/StorageWrapper.js";

jest.mock("../handlers/StorageWrapper.js");

describe('EventListener should', () => {
    let mockDom;
    let eventHandler;
    let mockGameHandler;
    const mockWindow = {};

    beforeEach(() => {
        jest.clearAllMocks();
        mockDom = helper.mockDom();
        mockGameHandler = helper.mockGameHandler();
        eventHandler = new EventHandler(mockDom, mockWindow, mockGameHandler);
    });

    it('call randomGame on reload if no roverId is found', () => {
        StorageWrapper.getRoverId.mockReturnValueOnce(null);

        eventHandler.listenOnReload();
        mockWindow.onload();

        expect(mockGameHandler.randomGame).toHaveBeenCalled();
    });

    // TODO: unskip once loadGame endpoint is implemented
    it.skip('call loadGame on reload if a roverId is found', () => {
        const storedRoverId = 'mockedRoverId';
        StorageWrapper.getRoverId.mockReturnValueOnce(storedRoverId);

        eventHandler.listenOnReload();
        mockWindow.onload();

        expect(mockGameHandler.loadGame).toHaveBeenCalledWith(storedRoverId);
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
