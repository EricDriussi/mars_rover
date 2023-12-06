export class RequestBuilder {
    static randomGameRequest() {
        return ['http://localhost:4242/api/randomGame', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify({}),
        }]
    }

    static moveRoverRequest(id, commands) {
        return ['http://localhost:4242/api/moveSequence', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify({id, commands}),
        }]
    }

    static loadGameRequest(id) {
        return ['http://localhost:4242/api/loadGame', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify({id}),
        }]
    }
}