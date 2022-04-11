package main

import(
	"log"
	"net/http"

	"github.com/justin-p-trainor/go-digital-media-server/internal/db"

	"github.com/gin-gonic/gin"
)

func main() {
	db, err := db.NewConnection("./Chinook_Sqlite.sqlite")
	if err != nil {
		log.Fatal("Database error: ", err)
	}

	getEmptyHandler := func (c *gin.Context) {
		c.IndentedJSON(http.StatusOK, db.GetTracks(""))
	}

	getTracksHandler := func(c *gin.Context) {
		trackName := c.Param("trackName")
		c.IndentedJSON(http.StatusOK, db.GetTracks(trackName))
	}
	
	router := gin.Default()
	router.SetTrustedProxies(nil)

	router.GET("/", getEmptyHandler)
	router.GET("/:trackName", getTracksHandler)

	router.Run("localhost:4041")
}