import {InfoPainter} from '../InfoPainter.js';
import * as helper from "./TestHelper.js";

describe('InfoPainter should', () => {
    let mockDom;
    let infoPainter;

    beforeEach(() => {
        mockDom = helper.mockDom();
        infoPainter = new InfoPainter(mockDom.window.document);
    });

    describe.each([
        ['for errors', 'error', 'error-list', 'error-box'],
        ['for warnings', 'warning', 'warn-list', 'warn-box'],
    ])('%s', (_, testedMethod, listId, boxId) => {
        const messages = ['Message 1', 'Message 2'];

        it('display messages box when messages are provided', () => {
            infoPainter[testedMethod](messages);

            const messageBox = helper.getElementFrom(mockDom, boxId);
            expect(isHidden(messageBox)).toBe(false);
        });

        it('display messages box when a single message is provided', () => {
            infoPainter[testedMethod]('not an array');

            const messageBox = helper.getElementFrom(mockDom, boxId);
            expect(isHidden(messageBox)).toBe(false);
        });

        it('populate the message list', () => {
            infoPainter[testedMethod](messages);

            const messageList = helper.getElementFrom(mockDom, listId);
            messages.forEach((message, index) => {
                const listItem = messageList.childNodes[index];
                expect(listItem.textContent).toBe(message);
            });
        });

        it('not display message box when no messages are provided', () => {
            infoPainter[testedMethod]([]);

            const messageBox = helper.getElementFrom(mockDom, boxId);
            expect(isHidden(messageBox)).toBe(true);
        });

        it('disregard old messages', () => {
            infoPainter[testedMethod](['Message 1', 'Message 2']);
            infoPainter[testedMethod]([]);

            const messageBox = helper.getElementFrom(mockDom, boxId);
            expect(isHidden(messageBox)).toBe(true);
        });
    });
});

function isHidden(element) {
    return element.classList.contains('hidden');
}
