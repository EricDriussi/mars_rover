import * as helper from "./TestHelper.js";
import {GameHandler} from "../GameHandler.js";
import {ApiWrapper} from "../ApiWrapper";

global.fetch = jest.fn();

describe('ApiWrapper should', () => {
    let mockInfoPainter = helper.mockInfoPainter();
    let apiWrapper;

    beforeEach(() => {
        jest.clearAllMocks();
        apiWrapper = new ApiWrapper(mockInfoPainter);
    });

    describe('when calling the /api/randomRover endpoint', () => {
        const expectedFetchParams = ['/api/randomRover', { method: 'POST' }];
        it('fetch and unpack the response', async () => {
            const errorFreeResponse = {
                ok: true,
                json: jest.fn()
            };
            global.fetch.mockResolvedValue(errorFreeResponse);

            await apiWrapper.callGetEndpoint();

            expect(global.fetch).toHaveBeenCalledWith(...expectedFetchParams);
            expect(mockInfoPainter.error).not.toHaveBeenCalled();
            expect(errorFreeResponse.json).toHaveBeenCalled();
        });

        it('handle error if present in response', async () => {
            const mockErrorResponse = {
                ok: false,
                statusText: 'Not Found',
                json: jest.fn()
            };
            global.fetch.mockResolvedValue(mockErrorResponse);

            await apiWrapper.callGetEndpoint();

            expect(global.fetch).toHaveBeenCalledWith(...expectedFetchParams);
            expect(mockInfoPainter.error).toHaveBeenCalledWith('Not Found');
        });
    });

    describe('when calling the /api/moveSequence endpoint', () => {
        const expectedMoveFetchParams = [
            '/api/moveSequence',
            {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                },
                body: JSON.stringify({ id: 'roverId', commands: 'f' }),
            },
        ];

        it('fetch and unpack the response', async () => {
            const mockResponse = {
                ok: true,
                json: jest.fn()
            };
            global.fetch.mockResolvedValue(mockResponse);

            await apiWrapper.callMoveEndpoint('roverId', 'f');

            expect(global.fetch).toHaveBeenCalledWith(...expectedMoveFetchParams);
            expect(mockInfoPainter.error).not.toHaveBeenCalled();
            expect(mockResponse.json).toHaveBeenCalled();
        });

        it('handle error if present in response', async () => {
            const mockErrorResponse = {
                ok: false,
                statusText: 'Bad Request',
                json: jest.fn()
            };
            global.fetch.mockResolvedValue(mockErrorResponse);

            await apiWrapper.callMoveEndpoint('roverId', 'f');

            expect(global.fetch).toHaveBeenCalledWith(...expectedMoveFetchParams);
            expect(mockInfoPainter.error).toHaveBeenCalledWith('Bad Request');
        });
    });
});