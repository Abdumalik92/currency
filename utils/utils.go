package utils

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"github.com/gin-gonic/gin"
	"log"
	"strings"
)

// SHA256 HMAC makes signed hash of string
func GetSha256(text string, secret []byte) string {

	// Create a new HMAC by defining the hash type and the key (as byte array)
	h := hmac.New(sha256.New, secret)

	// Write Data to it
	h.Write([]byte(text))

	// Get result and encode as hexadecimal string
	hash := hex.EncodeToString(h.Sum(nil))

	return hash
}

func GetImages(c *gin.Context) {
	imageName := c.Param("image_name")

	// used this for taking the path of file, made this for not making updates in application and sort logos and images in folders
	str := strings.ReplaceAll(imageName, "-", "/")

	log.Println("image_name:", str)

	// returning requested image
	c.File("files/" + str)
}
