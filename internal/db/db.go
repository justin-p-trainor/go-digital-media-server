package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

const queryTracks = `
SELECT Track.Name as TrackName, Track.Milliseconds, Track.Bytes, Artist.Name as ArtistName, Album.Title
FROM (SELECT * FROM Track WHERE instr(lower(Track.Name), lower(?))) as Track
LEFT JOIN Album ON Track.AlbumId = Album.AlbumId
LEFT JOIN Artist ON Album.ArtistId = Artist.ArtistId
`

// Holds a connection to an SQL database as well as the prepared statements for interacting with the database
type Connection struct {
	database *sql.DB
	queryTrackNameStatement *sql.Stmt
}

// Holds the data that is queryable from the database
type Track struct {
	TrackName    string `json:"trackName"`
	ArtistName   string `json:"artistName"`
	AlbumName    string `json:"albumName"`
	Milliseconds uint   `json:"milliseconds"`
	Bytes        uint   `json:"bytes"`
}

// Creates a new connection to the given database file that can be queried for tracks
func NewConnection(databaseFile string) (*Connection, error) {
	database, err := sql.Open("sqlite3", databaseFile)
	if err != nil {
		return nil, err
	}

	queryTrackNameStatement, err := database.Prepare(queryTracks)
	if err != nil {
		return nil, err
	}

	return &Connection{ database: database, queryTrackNameStatement: queryTrackNameStatement }, nil
}

// Queries the database for tracks containing the given trackName. This is case-insensitive
func (c *Connection) GetTracks(trackName string) []Track {
	result := []Track{}

	rows, err := c.queryTrackNameStatement.Query(trackName)
	if err != nil {
		return result
	}
	defer rows.Close()

	for rows.Next() {
		var track Track

		if err := rows.Scan(&track.TrackName, &track.Milliseconds, &track.Bytes, &track.ArtistName, &track.AlbumName); err != nil {
			return result
		}

		result = append(result, track)
	}

	return result
}