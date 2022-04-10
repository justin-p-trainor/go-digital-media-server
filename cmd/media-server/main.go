package main

import(
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type track struct {
	TrackName    string `json:"trackName"`
	ArtistName   string `json:"artistName"`
	AlbumName    string `json:"albumName"`
	Milliseconds uint   `json:"milliseconds"`
	Bytes        uint   `json:"bytes"`
}

var tracks = []track{
	{"Alpha", "One", "Trees", 50000, 120000},
	{"Omega", "Seven", "Shrubs", 25121, 224095},
	{"Beta", "Two", "Flowers", 100, 1500},
}

func getTracks(c *gin.Context) {
	trackName := c.Param("trackName")
	log.Println(trackName)
	c.IndentedJSON(http.StatusOK, tracks)
}

func main() {
	router := gin.Default()
	router.SetTrustedProxies(nil)
	log.Println(os.Getwd())

	router.GET("/:trackName", getTracks)

	router.Run("localhost:4041")
}