package controller

import (
	"curr/models"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func OrienTransfer() models.Bank {
	client := &http.Client{}
	//proxyS, err := url.Parse("http://192.168.0.8:4480")
	//if err != nil {
	//	log.Println("OrienTransfer proxy ", err.Error())
	//	return models.Bank{
	//		BankName:    "Ориён банк",
	//		ShortName:"orien",
	//		Icon: "",
	//		Colors: models.Colors{
	//			Color1: "#393185",
	//			Color2: "#D81425",
	//		},
	//		AndroidLink: "https://www.orienpay.tj/",
	//		IOSLink:     "https://www.orienpay.tj/",
	//		Currencies: []models.Currencies{
	//			models.Currencies{
	//				Name: "RUB",
	//				Buy:  "0.0000",
	//				Sell: "0.0000",
	//			},
	//		},
	//	}
	//}
	//client.Transport = &http.Transport{Proxy: http.ProxyURL(proxyS)}
	payload := strings.NewReader("{\"CardNo\": \"992915482424\"}")
	respOrien, err := client.Post("https://www.orienpay.tj/api/CardInfo", "application/json", payload)
	if err != nil {
		log.Println("OrienTransfer resp err ", err.Error())
		return models.Bank{
			BankName:  "Ориён банк",
			ShortName: "orien",
			Icon:      "",
			Colors: models.Colors{
				Color1: "#393185",
				Color2: "#D81425",
			},
			AndroidLink: "https://www.orienpay.tj/",
			IOSLink:     "https://www.orienpay.tj/",
			Currencies: []models.Currencies{
				models.Currencies{
					Name: "RUB",
					Buy:  "0.0000",
					Sell: "0.0000",
				},
			},
		}
	}

	bodyOrien, err := ioutil.ReadAll(respOrien.Body)
	if err != nil {
		log.Println("Body OrienTransfer err ", err.Error())
		return models.Bank{
			BankName:  "Ориён банк",
			ShortName: "orien",
			Icon:      "",
			Colors: models.Colors{
				Color1: "#393185",
				Color2: "#D81425",
			},
			AndroidLink: "https://www.orienpay.tj/",
			IOSLink:     "https://www.orienpay.tj/",
			Currencies: []models.Currencies{
				models.Currencies{
					Name: "RUB",
					Buy:  "0.0000",
					Sell: "0.0000",
				},
			},
		}
	}
	IBTRub := models.Rate{}
	err = json.Unmarshal(bodyOrien, &IBTRub)
	if err != nil {
		log.Println("Unmarshal OrienTransfer err ", err.Error())
		return models.Bank{
			BankName:  "Ориён банк",
			ShortName: "orien",

			Icon: "",
			Colors: models.Colors{
				Color1: "#393185",
				Color2: "#D81425",
			},
			AndroidLink: "https://www.orienpay.tj/",
			IOSLink:     "https://www.orienpay.tj/",
			Currencies: []models.Currencies{
				models.Currencies{
					Name: "RUB",
					Buy:  "0.0000",
					Sell: "0.0000",
				},
			},
		}
	}
	return models.Bank{
		BankName:  "Ориён банк",
		ShortName: "orien",
		Icon:      "",
		Colors: models.Colors{
			Color1: "#393185",
			Color2: "#D81425",
		},
		AndroidLink: "https://www.orienpay.tj/",
		IOSLink:     "https://www.orienpay.tj/",
		Currencies: []models.Currencies{
			models.Currencies{
				Name: "RUB",
				Buy:  "0" + IBTRub.Orienbank,
				Sell: "0.0000",
			},
		},
	}

}
