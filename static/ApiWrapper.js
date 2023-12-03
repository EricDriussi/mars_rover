import {Result} from "./Result.js";
import {RequestBuilder} from "./RequestBuilder";

export class ApiWrapper {
    static async postRandomGame() {
        const request = RequestBuilder.randomGameRequest();
        const response = await fetch(...request);
        return await ApiWrapper.#unpackResponse(response);
    }

    static async postMoveRover(id, commands) {
        const request = RequestBuilder.moveRoverRequest(id, commands);
        const response = await fetch(...request);
        return await ApiWrapper.#unpackResponse(response);
    }

    static async #unpackResponse(response) {
        if (!response.ok) {
            return Result.failure(response.statusText);
        }
        return Result.success(await response.json());
    }
}
