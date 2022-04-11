package main

import(
	"log"
	"net/http"
	"os"

	"github.com/justin-p-trainor/go-digital-media-server/internal/db"

	"github.com/gin-gonic/gin"
)

func main() {
	db, err := db.NewConnection("./Chinook_Sqlite.sqlite")
	if err != nil {
		log.Fatal("Database error: ", err)
	}

	getTracksHandler := func(c *gin.Context) {
		trackName := c.Param("trackName")
		c.IndentedJSON(http.StatusOK, db.GetTracks(trackName))
	}
	
	router := gin.Default()
	router.SetTrustedProxies(nil)
	log.Println(os.Getwd())

	router.GET("/:trackName", getTracksHandler)

	router.Run("localhost:4041")
}