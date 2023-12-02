export class ApiWrapper {
    #errorPainter;

    constructor(errorPainter) {
        this.#errorPainter = errorPainter;
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
            console.error('API error:', response.statusText);
            this.#errorPainter.displayErrors(response.statusText);
        }
        return await response.json();
    }
}

