package routes

import (
	"curr/controller"
	"curr/jobs"
	"curr/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/natefinch/lumberjack"
	"io"
	"log"
	"os"
	"time"
)

func RunRoutes() {
	r := gin.Default()
	r.Use(CORSMiddleware())
	f, err := os.OpenFile(utils.AppSettings.AppParams.LogFile, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0666)

	if err != nil {
		fmt.Println("file create error", err.Error())
		return
	}

	logger := &lumberjack.Logger{
		Filename:   f.Name(),
		MaxSize:    10, // megabytes
		MaxBackups: 100,
		MaxAge:     28,   // days
		Compress:   true, // disabled by default
	}

	log.SetOutput(logger)

	gin.DefaultWriter = io.MultiWriter(logger, os.Stdout)

	r.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {

		// your custom format
		return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
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
	}))

	r.Use(gin.Recovery())

	jobs.NPCRBank()
	jobs.C2CBank()

	r.GET("/nbt_rates", controller.NbtRates)
	r.GET("/npcr_bank_rates", controller.NPCRBankRates)
	r.GET("/c2c_bank_rates", controller.C2CBankRates)
	r.GET("/get_image/:image_name", utils.GetImages)
	_ = r.Run(utils.AppSettings.AppParams.PortRun)
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Content-Type", "application/json")
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET,PUT,POST,DELETE")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "*")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(200)
		}

		c.Next()
	}
}
