package db

import "testing"

func checkTrackSize(t *testing.T, tracks []Track, size int) bool {
	result := len(tracks) == size

	if !result {
		t.Errorf("Got %d total tracks, expected %d", len(tracks), size)
	}

	return result
}

func checkRecord(t *testing.T, expected Track, got Track) bool {
	result := expected == got

	if !result {
		t.Errorf("Got %v, expected %v", expected, got)
	}

	return result
}

func getTracks(name string) []Track {
	db, err := NewConnection("../../Chinook_Sqlite.sqlite")
	if err != nil {
		return []Track{}
	}

	return db.GetTracks(name)
}

func TestQueryAll(t *testing.T) {
	tracks := getTracks("")
	expectedSize := 3503

	checkTrackSize(t, tracks, expectedSize)
}

func TestQueryWallCaseInsensitivity(t *testing.T) {
	tracks := getTracks("wall")
	expectedSize := 6

	checkTrackSize(t, tracks, expectedSize)
}

func TestQueryEscapeChars(t *testing.T) {
	tracks := getTracks("%a")
	expectedSize := 0

	checkTrackSize(t, tracks, expectedSize)
}

func TestQuerySingleRecord(t *testing.T) {
	tracks := getTracks("Wall Of Denial")
	expectedSize := 1
	expectedTrack := Track{
		"Wall Of Denial",
		"Stevie Ray Vaughan & Double Trouble",
		"In Step",
		336927,
		11085915,
	}

	if checkTrackSize(t, tracks, expectedSize) {
		checkRecord(t, tracks[0], expectedTrack)
	}
}