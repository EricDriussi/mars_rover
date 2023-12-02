import {ErrorPainter} from '../ErrorPainter.js';
import * as helper from "./TestHelper.js";

describe('ErrorPainter should', () => {
    let mockDom;
    let errorPainter;

    beforeEach(() => {
        mockDom = helper.mockDom();
        errorPainter = new ErrorPainter(mockDom.window.document);
    });

    describe('display errors', () => {
        const errors = ['Error 1', 'Error 2'];
        it('when errors are provided', () => {
            errorPainter.displayErrors(errors);

            const errorBox = helper.getErrorBoxFrom(mockDom);
            expect(isHidden(errorBox)).toBe(false);
        });

        it('populating the error list with error messages', () => {
            errorPainter.displayErrors(errors);

            const errorList = mockDom.window.document.getElementById('error-list');
            errors.forEach((error, index) => {
                const listItem = errorList.childNodes[index];
                expect(listItem.textContent).toBe(error);
            });
        });
    });

    describe('not display error box', () => {
        it('when no errors are provided', () => {
            errorPainter.displayErrors([]);

            const errorBox = helper.getErrorBoxFrom(mockDom);
            expect(isHidden(errorBox)).toBe(true);
        });

        it('for old errors', () => {
            errorPainter.displayErrors(['Error 1', 'Error 2']);
            errorPainter.displayErrors([]);

            const errorBox = helper.getErrorBoxFrom(mockDom);
            expect(isHidden(errorBox)).toBe(true);
        });
    });
});

function isHidden(element) {
    return element.classList.contains('hidden');
}

