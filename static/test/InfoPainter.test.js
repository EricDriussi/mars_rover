import {InfoPainter} from '../InfoPainter.js';
import * as helper from "./TestHelper.js";

describe('InfoPainter should', () => {
    let mockDom;
    let infoPainter;

    beforeEach(() => {
        mockDom = helper.mockDom();
        infoPainter = new InfoPainter(mockDom.window.document);
    });

    describe('display errors', () => {
        const errors = ['Error 1', 'Error 2'];
        it('when errors are provided', () => {
            infoPainter.errors(errors);

            const errorBox = helper.getErrorBoxFrom(mockDom);
            expect(isHidden(errorBox)).toBe(false);
        });

        it('populating the error list with error messages', () => {
            infoPainter.errors(errors);

            const errorList = mockDom.window.document.getElementById('error-list');
            errors.forEach((error, index) => {
                const listItem = errorList.childNodes[index];
                expect(listItem.textContent).toBe(error);
            });
        });
    });

    describe('not display error box', () => {
        it('when no errors are provided', () => {
            infoPainter.errors([]);

            const errorBox = helper.getErrorBoxFrom(mockDom);
            expect(isHidden(errorBox)).toBe(true);
        });

        it('for old errors', () => {
            infoPainter.errors(['Error 1', 'Error 2']);
            infoPainter.errors([]);

            const errorBox = helper.getErrorBoxFrom(mockDom);
            expect(isHidden(errorBox)).toBe(true);
        });
    });
});

function isHidden(element) {
    return element.classList.contains('hidden');
}

