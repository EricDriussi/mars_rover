export class ApiWrapper {
    #infoPainter;

    constructor(infoPainter) {
        this.#infoPainter = infoPainter;
    }

    async callGetEndpoint() {
        const response = await fetch('/api/randomRover', {
            method: 'POST',
        });
        return await this.#unpackResponse(response);
    }

    async callMoveEndpoint(id, commands) {
        const response = await fetch('/api/moveSequence', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify({id, commands}),
        });
        return await this.#unpackResponse(response);
    }

    // TODO: should return Error() if response is not ok
    async #unpackResponse(response) {
        if (!response.ok) {
            this.#infoPainter.error(response.statusText);
        }
        return await response.json();
    }
}

