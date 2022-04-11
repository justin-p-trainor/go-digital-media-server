# go-digital-media-server

A digital media server hosted on `localhost:4041` with one endpoint `/trackName` which will return an array of track data whose name includes the query string.

For example, the query `http://localhost:4041/dawn` should return:
```json
[
    {
        "id": 5,
        "trackName": "Princess of the Dawn",
        "artistName": "Accept",
        "albumName": "Restless and Wild",
        "milliseconds": 375418,
        "bytes": 6290521
    },
    {
        "id": 2475,
        "trackName": "Slow Dawn",
        "artistName": "Smashing Pumpkins",
        "albumName": "Judas 0: B-Sides and Rarities",
        "milliseconds": 192339,
        "bytes": 6269057
    }
]
```

# Building and Running
The implementation uses [this](https://github.com/mattn/go-sqlite3) sqlite3 driver which requires you to have the environment variable `CGO_ENABLED=1` as well as a `gcc` compiler in your path.

To build/run the server on `localhost:4041`, from the root of the repository: `go run ./cmd/media-server/`

To run the tests, from the root of the repository: `go test ./internal/...`

# Implementation Details and Assumptions
* Of the bonus objectives, I included the album/artist names and added testing
* Of the bonus objectives, I didn't include improvements to the search behaviour, pagination of results, or dockerization
* I used an unofficial project structure described by [this source](https://github.com/golang-standards/project-layout)
* For test cases, I used the provided sqlite test database, but for the implementation I assumed that the database could be dynamic and may be too big to load into memory. Because of this:
    * The implementation uses an SQL query per request
    * I didn't implement partial matching (one way to do it in-memory using Go could be a [fuzzy search library](https://github.com/lithammer/fuzzysearch))
* If I were going to implement pagination, I would use query parameters - likely one of the methods described [here](https://www.moesif.com/blog/technical/api-design/REST-API-Design-Filtering-Sorting-and-Pagination/) depending on time constraints and database assumptions
* I didn't implement any benchmarking, but I was happy to see that it is built into the language and looks relatively easy to use