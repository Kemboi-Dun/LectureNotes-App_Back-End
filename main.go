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

	// FILE END POINTS
	// r.POST("/upload", server.FileCreate)
	// r.GET("/download", server.FileGet)

	// USER END POINTS
	r.POST("/user", controllers.UserCreate)
	r.GET("/users", controllers.GetUsers)
	r.GET("/user/:id", controllers.GetUserId)
	r.PUT("/user/:id", controllers.UpdateUser)
	r.DELETE("/user/:id", controllers.DeleteUser)

	// FOLDER END POINTS
	r.POST("/folder", controllers.FolderCreate)
	r.GET("/folders", controllers.GetFolders)
	r.GET("/folder/:id", controllers.GetFolderId)
	r.PUT("/folder/:id", controllers.UpdateFolder)
	r.DELETE("/folder/:id", controllers.DeleteFolder)

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

	// YEAR END POINTS
	r.POST("/year", controllers.YearCreate)
	r.GET("/years", controllers.GetYears)
	r.GET("/year/:id", controllers.GetYearId)
	r.PUT("/year/:id", controllers.UpdateYear)
	r.DELETE("/year/:id", controllers.DeleteYear)

	// COURSE END POINTS
	r.POST("/course", controllers.CourseCreate)
	r.GET("/courses", controllers.GetCourses)
	r.GET("/course/:id", controllers.GetCourseId)
	r.PUT("/course/:id", controllers.UpdateCourse)
	r.DELETE("/course/:id", controllers.DeleteCourse)

	// COURSETYPE END POINTS
	r.POST("/courseType", controllers.CourseTypeCreate)
	r.GET("/courseTypes", controllers.GetCourseTypes)
	r.GET("/courseType/:id", controllers.GetCourseTypeId)
	r.PUT("/courseType/:id", controllers.UpdateCourseType)
	r.DELETE("/courseType/:id", controllers.DeleteCourseType)

	// SEMESTERNAME END POINTS
	r.POST("/semesterName", controllers.SemesterNameCreate)
	r.GET("/semesterNames", controllers.GetSemesterNames)
	r.GET("/semesterName/:id", controllers.GetSemesterNameId)
	r.PUT("/semesterName/:id", controllers.UpdateSemesterName)
	r.DELETE("/semesterName/:id", controllers.DeleteSemesterName)

	// SEMESTER END POINTS
	r.POST("/semester", controllers.SemesterCreate)
	r.GET("/semesters", controllers.GetSemesters)
	r.GET("/semester/:id", controllers.GetSemesterId)
	r.PUT("/semester/:id", controllers.UpdateSemester)
	r.DELETE("/semester/:id", controllers.DeleteSemester)

	// COURSECLASS END POINTS
	r.POST("/courseClass", controllers.CourseClassCreate)
	r.GET("/courseClasses", controllers.GetCourseClasses)
	r.GET("/courseClass/:id", controllers.GetCourseClassId)
	r.PUT("/courseClass/:id", controllers.UpdateCourseClass)
	r.DELETE("/courseClass/:id", controllers.DeleteCourseClass)

	// UNIT END POINTS
	r.POST("/unit", controllers.UnitCreate)
	r.GET("/units", controllers.GetUnits)
	r.GET("/unit/:id", controllers.GetUnitId)
	r.PUT("/unit/:id", controllers.UpdateUnit)
	r.DELETE("/unit/:id", controllers.DeleteUnit)

	// COMMENT END POINTS
	r.POST("/comment", controllers.CommentCreate)
	r.GET("/comments", controllers.GetComments)
	r.GET("/comment/:id", controllers.GetCommentId)
	r.PUT("/comment/:id", controllers.UpdateComment)
	r.DELETE("/comment/:id", controllers.DeleteComment)

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
