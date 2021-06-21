package main

import (
	"fmt"
	Apis "gin/apis"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	const PORT = ":8080"
	r := gin.Default()
	runningDir, _ := os.Getwd()
	// ========= LOGGER =========
	errLogFile, _ := os.OpenFile(fmt.Sprintf("%s/gin_error.log", runningDir), os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0600)
	accessLogFile, _ := os.OpenFile(fmt.Sprintf("%s/gin_access.log", runningDir), os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0600)
	gin.DefaultErrorWriter = errLogFile
	gin.DefaultWriter = accessLogFile
	// Standard
	// r.Use(gin.Logger())
	// Custom format logger
	r.Use(gin.LoggerWithFormatter(formatLogger))
	// Disabled path
	r.Use(gin.LoggerWithWriter(gin.DefaultWriter, "/profile"))
	// ========= LOGGER =========

	// ========= UPLOAD FILE =========
	r.MaxMultipartMemory = 8 << 20 // 8 MiB

	r.POST("/upload", func(c *gin.Context) {
		file, _ := c.FormFile("file")
		// Get extension
		extension := filepath.Ext(file.Filename)
		fmt.Printf("extension: %s\n", extension)
		username := c.PostForm("username")
		token := c.PostForm("token")
		fmt.Printf("username: %s, token: %s\n", username, token)
		c.SaveUploadedFile(file, fmt.Sprintf("%s/%s", runningDir, file.Filename))
		c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))
	})
	// ========= UPLOAD FILE =========

	// Grouping route
	Apis.Setup(r)

	// ========= STATIC FILE =========
	r.GET("/image", func(c *gin.Context) {
		c.File("static/like.png")
	})
	r.GET("/html", func(c *gin.Context) {
		c.File("static/index.html")
	})
	r.GET("/download", func(c *gin.Context) {
		c.Header("Content-Description", "Simulation File Download")
		c.Header("Content-Transfer-Encoding", "binary")
		c.Header("Content-Disposition", "attechment; filename="+"download.png")
		c.Header("Content-Type", "application/octet-stream")
		c.File("static/like.png")
	})

	// pubilc static file
	r.Static("/assets", "./assets/pubilc")

	// ========= STATIC FILE =========

	r.Run(PORT)
}

func formatLogger(param gin.LogFormatterParams) string {
	return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s\"%s\" %s\"\n",
		param.ClientIP,
		param.TimeStamp.Format(time.RFC1123),
		param.Method,
		param.Path,
		param.Request.Proto,
		param.StatusCode,
		param.Latency,
		param.Request.UserAgent(),
		param.ErrorMessage,
	)
}
