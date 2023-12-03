import {Logger} from '../handlers/Logger.js';
import * as helper from "./TestHelper.js";

describe('Logger should', () => {
    let mockDom;
    let logger;

    beforeEach(() => {
        jest.clearAllMocks();
        mockDom = helper.mockDom();
        logger = new Logger(mockDom.window.document);
    });

    describe.each([
        ['for errors', 'error', 'error-list', 'error-box'],
        ['for warnings', 'warning', 'warn-list', 'warn-box'],
    ])('%s', (_, testedMethod, listId, boxId) => {
        const messages = ['Message 1', 'Message 2'];

        it('display messages box when messages are provided', () => {
            logger[testedMethod](messages);

            const messageBox = helper.getElementFrom(mockDom, boxId);
            expect(isHidden(messageBox)).toBe(false);
        });

        it('display messages box when a single message is provided', () => {
            logger[testedMethod]('not an array');

            const messageBox = helper.getElementFrom(mockDom, boxId);
            expect(isHidden(messageBox)).toBe(false);
        });

        it('populate the message list', () => {
            logger[testedMethod](messages);

            const messageList = helper.getElementFrom(mockDom, listId);
            messages.forEach((message, index) => {
                const listItem = messageList.childNodes[index];
                expect(listItem.textContent).toBe(message);
            });
        });

        it('not display message box when no messages are provided', () => {
            logger[testedMethod]([]);

            const messageBox = helper.getElementFrom(mockDom, boxId);
            expect(isHidden(messageBox)).toBe(true);
        });

        it('disregard old messages', () => {
            logger[testedMethod](['Message 1', 'Message 2']);
            logger[testedMethod]([]);

            const messageBox = helper.getElementFrom(mockDom, boxId);
            expect(isHidden(messageBox)).toBe(true);
        });
    });

    it('reset all log messages', () => {
        logger.resetLogMessages();

        const errorBox = helper.getElementFrom(mockDom, 'error-box');
        expect(isHidden(errorBox)).toBe(true);
        const warnBox = helper.getElementFrom(mockDom, 'warn-box');
        expect(isHidden(warnBox)).toBe(true);
    });

});

function isHidden(element) {
    return element.classList.contains('hidden');
}
