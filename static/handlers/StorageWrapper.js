const roverIdKey = "roverId";

export class StorageWrapper {
    static getRoverId() {
        return localStorage.getItem(roverIdKey);
    }

    static setRoverId(id) {
        localStorage.setItem(roverIdKey, id);
    }
}
