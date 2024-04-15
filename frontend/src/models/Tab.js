export class Tab {
    constructor({ name, query, id, error }) {
        this.id = id ?? Date.now()
        this.name = name
        this.query = query
        this.error = error
        this.table = {}
    }
}