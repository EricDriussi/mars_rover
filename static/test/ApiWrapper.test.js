import {ApiWrapper} from "../api/ApiWrapper.js";
import {RequestBuilder} from "../api/RequestBuilder.js";

describe('ApiWrapper should', () => {

    global.fetch = jest.fn();
    beforeEach(() => {
        jest.clearAllMocks();
    });

    const errorResponse = {
        ok: false,
        statusText: 'Bad Request',
        json: jest.fn()
    };

    const errorFreeResponse = {
        ok: true,
        json: jest.fn()
    };

    describe('when calling the random game endpoint', () => {
        const expectedFetchParams = RequestBuilder.randomGameRequest();
        it('fetch and unpack the response', async () => {
            global.fetch.mockResolvedValue(errorFreeResponse);

            const result = await ApiWrapper.postRandomGame();

            expect(global.fetch).toHaveBeenCalledWith(...expectedFetchParams);
            expect(errorFreeResponse.json).toHaveBeenCalled();
            expect(result.isFailure()).toBe(false);
        });

        it('handle error if present in response', async () => {
            global.fetch.mockResolvedValue(errorResponse);

            const result = await ApiWrapper.postRandomGame();

            expect(global.fetch).toHaveBeenCalledWith(...expectedFetchParams);
            expect(result.isFailure()).toBe(true);
        });
    });

    describe('when calling the move rover endpoint', () => {
        const requestData = ['roverId', 'f']
        const expectedMoveFetchParams = RequestBuilder.moveRoverRequest(...requestData);

        it('fetch and unpack the response', async () => {
            global.fetch.mockResolvedValue(errorFreeResponse);

            const result = await ApiWrapper.postMoveRover(...requestData);

            expect(global.fetch).toHaveBeenCalledWith(...expectedMoveFetchParams);
            expect(errorFreeResponse.json).toHaveBeenCalled();
            expect(result.isFailure()).toBe(false);
        });

        it('handle error if present in response', async () => {
            global.fetch.mockResolvedValue(errorResponse);

            const result = await ApiWrapper.postMoveRover(...requestData);

            expect(global.fetch).toHaveBeenCalledWith(...expectedMoveFetchParams);
            expect(result.isFailure()).toBe(true);
        });
    });
});