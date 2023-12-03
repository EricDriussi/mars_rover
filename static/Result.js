export class Result {
    #success;
    #value;

    constructor(success, value) {
        this.#success = success;
        this.#value = value;
    }

    static success(value) {
        return new Result(true, value);
    }

    static failure(value) {
        return new Result(false, value);
    }

    isFailure() {
        return !this.#success;
    }

    value() {
        return this.#value;
    }
}
