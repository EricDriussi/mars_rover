import { Logger } from '../handlers/Logger.js';
import * as helper from './TestHelper.js';

describe('Logger should', () => {
    const messages = ['Message 1', 'Message 2'];
    let mockDom;
    let logger;

    beforeEach(() => {
        jest.clearAllMocks();
        mockDom = helper.mockDom();
        logger = new Logger(mockDom.window.document);
    });

    describe.each([
        ['error', 'error-box', 'error-list'],
        ['warning', 'warn-box', 'warn-list'],
    ])('for %s messages', (messageType, boxId, listId) => {
        it('display message box when messages are provided', () => {
            logger[messageType](messages);

            const messageBox = helper.getElementFrom(mockDom, boxId);
            expect(isHidden(messageBox)).toBe(false);
        });

        it('display message box when a single message is provided', () => {
            logger[messageType]('not an array');

            const messageBox = helper.getElementFrom(mockDom, boxId);
            expect(isHidden(messageBox)).toBe(false);
        });

        it('populate the message list', () => {
            logger[messageType](messages);

            const messageList = helper.getElementFrom(mockDom, listId);
            messages.forEach((message, index) => {
                const listItem = messageList.childNodes[index];
                expect(listItem.textContent).toBe(message);
            });
        });

        it('not display box when no messages are provided', () => {
            logger[messageType]([]);

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
