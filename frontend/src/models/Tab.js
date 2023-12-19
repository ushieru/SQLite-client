export class Tab {
    constructor(name, query, id) {
        this.id = id ?? Date.now()
        this.name = name
        this.query = query
        this.table = []
    }
}