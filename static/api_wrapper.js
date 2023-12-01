import { displayErrors } from './error_painter.js';

export class ApiWrapper {
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

    async #unpackResponse(response) {
        const data = await response.json();
        if (!response.ok) {
            console.error('API error:', data);
            displayErrors(data);
        }
        return data;
    }
}

