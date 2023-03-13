package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rs/cors"
)

func main() {
	r := gin.Default()
	r.Static("/assets", "./assets")
	r.LoadHTMLGlob("templates/*")
	r.MaxMultipartMemory = 8 << 20
	// 8 MiB
	r.GET("/download", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{})
	})

	r.POST("/upload", func(c *gin.Context) {

		// GET  file
		file, err := c.FormFile("image")

		if err != nil {
			c.HTML(http.StatusOK, "index.html", gin.H{
				"error": "Failed to upload image...",
			})
			return
		}

		// SAVE THE FILE
		err = c.SaveUploadedFile(file, "assets/uploads/"+file.Filename)
		if err != nil {
			c.HTML(http.StatusOK, "index.html", gin.H{
				"error": "Failed to upload image...",
			})
			return
		}

		// RENDER THE IMAGE
		c.HTML(http.StatusOK, "index.html", gin.H{
			"image": "assets/uploads/" + file.Filename,
		})
	})
	// HANDLE CORS
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
		AllowedMethods:   []string{"GET", "DELETE", "POST", "PUT"},
	})

	handler := c.Handler(r)
	log.Fatal((http.ListenAndServe(":3000", handler)))
	http.Handle("/", r)
	r.Run()
}
