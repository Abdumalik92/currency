package controller

import (
	"curr/db"
	"curr/models"
	"encoding/xml"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
)

func NbtRates(c *gin.Context) {
	var val models.Val
	var ar models.Val
	client := &http.Client{}
	t := time.Now().Format("2006-01-02")
	////proxyS, err := url.Parse("http://192.168.0.8:4480")
	//
	//if err != nil {
	//	log.Println("NbtRates proxy ", err.Error())
	//	return
	//}

	//client.Transport = &http.Transport{Proxy: http.ProxyURL(proxyS)}
	// Request the HTML page.
	resp, err := client.Get("http://nbt.tj/ru/kurs/export_xml.php?date=" + t + "&export=xmlout")
	if err != nil {
		log.Println("NbtRates res ", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"reason": "Сервис времмено не доступен"})
		return
	}

	bodyText, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("NbtRates body ", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"reason": "Что-то пошло не так"})
		return
	}

	s := strings.Replace(string(bodyText), "encoding=\"windows-1251\"", "", 1)
	s = strings.Replace(s, "<?xml version=\"1.0\"  ?>", "", 1)

	err = xml.Unmarshal([]byte(s), &val)

	if err != nil {
		log.Println("NbtRates unmarshal xml ", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"reason": "Что-то пошло не так"})
		return
	}

	val.Valute = append(val.Valute, models.Valute{
		Value:    1,
		Name:     "Таджикский сомони",
		CharCode: "TJS",
		Nominal:  1,
	})
	for _, arr := range val.Valute {
		arr.Flag = "https://transfer.humo.tj/currency-app/v1/get_image/flags-" + arr.CharCode + ".png"
		ar.Valute = append(ar.Valute, arr)
	}
	c.JSON(http.StatusOK, ar.Valute)
	return
}

func NPCRBankRates(c *gin.Context) {
	var (
		resp []models.Banks
		curr []models.Currencies
		str  []models.Bank
		col  []models.Colors
	)

	db.GetDBConn().Where("transfer=''").Find(&resp)
	db.GetDBConn().Where("transfer=''").Find(&curr)
	db.GetDBConn().Where("transfer=''").Find(&col)
	for _, arr := range resp {
		str = append(str, models.Bank{
			BankName:    arr.BankName,
			Icon:        arr.Icon,
			ShortName:   arr.ShortName,
			AndroidLink: arr.AndroidLink,
			IOSLink:     arr.IOSLink,
			IsChange:    arr.IsChange,
		})
	}
	for i := 0; i < len(str); i++ {
		for k := 0; k < len(col); k++ {
			if str[i].ShortName == col[k].BankName {
				str[i].Colors.Color1 = col[k].Color1
				str[i].Colors.Color2 = col[k].Color2
			}
		}
	}

	for i := 0; i < len(str); i++ {
		for k := 0; k < len(curr); k++ {
			if str[i].ShortName == curr[k].BankName {
				str[i].Currencies = append(str[i].Currencies, curr[k])
			}
		}
	}

	c.JSON(http.StatusOK, str)
	return
}

func C2CBankRates(c *gin.Context) {
	var (
		resp []models.Banks
		curr []models.Currencies
		str  []models.Bank
		col  []models.Colors
	)

	db.GetDBConn().Where("transfer='trans'").Find(&resp)
	db.GetDBConn().Where("transfer='trans'").Find(&curr)
	db.GetDBConn().Where("transfer='trans'").Find(&col)
	for _, arr := range resp {
		str = append(str, models.Bank{
			BankName:    arr.BankName,
			Icon:        arr.Icon,
			ShortName:   arr.ShortName,
			AndroidLink: arr.AndroidLink,
			IOSLink:     arr.IOSLink,
			IsChange:    arr.IsChange,
		})
	}
	for i := 0; i < len(str); i++ {
		for k := 0; k < len(col); k++ {
			if str[i].ShortName == col[k].BankName {
				str[i].Colors.Color1 = col[k].Color1
				str[i].Colors.Color2 = col[k].Color2
			}
		}
	}

	for i := 0; i < len(str); i++ {
		for k := 0; k < len(curr); k++ {
			if str[i].ShortName == curr[k].BankName {
				str[i].Currencies = append(str[i].Currencies, curr[k])
			}
		}
	}

	c.JSON(http.StatusOK, str)
	return
}
