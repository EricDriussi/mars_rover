import {Result} from "../Result.js";
import {RequestBuilder} from "./RequestBuilder.js";

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

    static async getLoadGame(id) {
        const request = RequestBuilder.loadGameRequest(id);
        const response = await fetch(...request);
        return await ApiWrapper.#unpackResponse(response);
    }

    static async #unpackResponse(response) {
        if (!response.ok) {
            const errMsg = await response.text();
            return Result.failure(`${response.statusText}: ${errMsg}`);
        }
        return Result.success(await response.json());
    }
}
