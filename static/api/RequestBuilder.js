export class RequestBuilder {
    static randomGameRequest() {
        return ['/api/randomGame', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
        }]
    }

    static moveRoverRequest(id, commands) {
        return ['/api/moveSequence', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify({id, commands}),
        }]
    }

    static loadGameRequest(id) {
        return ['/api/loadGame', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify({id}),
        }]
    }
}