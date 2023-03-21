package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kemboi-dun/controllers"
	"github.com/kemboi-dun/initializers"
	"github.com/rs/cors"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectDB()

}

func main() {
	r := gin.Default()

	// USER END POINTS
	r.POST("/user", controllers.UserCreate)
	r.GET("/users", controllers.GetUsers)
	r.GET("/user/:id", controllers.GetUserId)
	r.PUT("/user/:id", controllers.UpdateUser)
	r.DELETE("/user/:id", controllers.DeleteUser)

	// FILE END POINTS
	r.POST("/file", controllers.FileCreate)
	r.GET("/files", controllers.GetFiles)
	r.GET("/file/:id", controllers.GetFileId)
	r.PUT("/file/:id", controllers.UpdateFile)
	r.DELETE("/file/:id", controllers.DeleteFile)

	// FILEACCESS END POINTS
	r.POST("/fileAccess", controllers.FileAccessCreate)
	r.GET("/fileAccesses", controllers.GetFileAccess)
	r.GET("/fileAccess/:id", controllers.GetFileAccessId)
	r.PUT("/fileAccess/:id", controllers.UpdateFileAccess)
	r.DELETE("/fileAccess/:id", controllers.DeleteFileAccess)

	// COMMENT END POINTS
	r.POST("/comment", controllers.CommentCreate)
	r.GET("/comments", controllers.GetComments)
	r.GET("/comment/:id", controllers.GetCommentId)
	r.PUT("/comment/:id", controllers.UpdateComment)
	r.DELETE("/comment/:id", controllers.DeleteComment)

	// ARTICLE END POINTS
	r.POST("/article", controllers.ArticleCreate)
	r.GET("/articles", controllers.GetArticles)
	r.GET("/article/:id", controllers.GetArticleId)
	r.PUT("/article/:id", controllers.UpdateArticle)
	r.DELETE("/article/:id", controllers.DeleteArticle)

	// DOCUMENT END POINTS
	r.POST("/DOC", controllers.UploadDoc)
	r.GET("/docs", controllers.GetDoc)
	r.GET("/doc/:id", controllers.GetDocId)
	r.PUT("/doc/:id", controllers.UpdateDoc)
	r.DELETE("/doc/:id", controllers.DeleteDoc)

	// r.GET("/view_file", func(c *gin.Context) {
	// 	filePath := "/assets/uploads/"  // Replace with the actual file path on your server
	// 	fileName := "KEMBOI DUNCAN.pdf" // Replace with the actual file name
	// 	c.Header("Content-Disposition", "inline; filename="+fileName)
	// 	c.Header("Content-Type", "application/octet-stream")
	// 	c.File(filePath)
	// })

	r.GET("/DOC_DOWNLOAD/:fileId", controllers.DownloadFile)

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
