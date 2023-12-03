import {Result} from "./Result.js";

export class ApiWrapper {
    constructor() {
    }

    async postRandomGame() {
        const response = await fetch('/api/randomRover', {
            method: 'POST',
        });
        return await this.#unpackResponse(response);
    }

    async postMoveRover(id, commands) {
        const response = await fetch('/api/moveSequence', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify({id, commands}),
        });
        return await this.#unpackResponse(response);
    }

    async #unpackResponse(response) {
        if (!response.ok) {
            return Result.failure(response.statusText);
        }
        return Result.success(await response.json());
    }
}
