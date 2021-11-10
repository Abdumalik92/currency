package controller

import (
	"curr/models"
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func Agroinvest(che chan models.Bank) {
	client := &http.Client{}
	//proxyS, err := url.Parse("http://192.168.0.8:4480")
	//if err != nil {
	//	log.Println("Agroinvest proxy ", err.Error())
	//	che <- models.Bank{
	//		BankName:    "Агроинвестбанк",
	//		ShortName:"agroinvest",
	//
	//		Icon: "",
	//		Colors: models.Colors{
	//			Color1: "#006CB5",
	//			Color2: "#057ED0",
	//		},
	//		AndroidLink: "http://www.agroinvestbank.tj/ru/index.php",
	//		IOSLink:     "http://www.agroinvestbank.tj/ru/index.php",
	//		Currencies: []models.Currencies{
	//			models.Currencies{
	//				Name: "USD",
	//				Buy:  "0.0000",
	//				Sell: "0.0000",
	//			},
	//			models.Currencies{
	//				Name: "RUB",
	//				Buy:  "0.0000",
	//				Sell: "0.0000",
	//			},
	//			models.Currencies{
	//				Name: "EUR",
	//				Buy:  "0.0000",
	//				Sell: "0.0000",
	//			},
	//		},
	//	}
	//	return
	//}
	//client.Transport = &http.Transport{Proxy: http.ProxyURL(proxyS)}

	//Agroinvest bank
	respAgroinvest, err := client.Get("http://www.agroinvestbank.tj/ru/index.php")
	if err != nil {
		log.Println("Agroinvest resp err ", err.Error())
		che <- models.Bank{
			BankName:  "Агроинвестбанк",
			ShortName: "agroinvest",

			Icon: "",
			Colors: models.Colors{
				Color1: "#006CB5",
				Color2: "#057ED0",
			},
			AndroidLink: "http://www.agroinvestbank.tj/ru/index.php",
			IOSLink:     "http://www.agroinvestbank.tj/ru/index.php",
			Currencies: []models.Currencies{
				models.Currencies{
					Name: "USD",
					Buy:  "0.0000",
					Sell: "0.0000",
				},
				models.Currencies{
					Name: "RUB",
					Buy:  "0.0000",
					Sell: "0.0000",
				},
				models.Currencies{
					Name: "EUR",
					Buy:  "0.0000",
					Sell: "0.0000",
				},
			},
		}
		return
	}
	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(respAgroinvest.Body)
	if err != nil {
		log.Println("Doc Agroinvest ", err.Error())
		che <- models.Bank{
			BankName:  "Агроинвестбанк",
			ShortName: "agroinvest",

			Icon: "",
			Colors: models.Colors{
				Color1: "#006CB5",
				Color2: "#057ED0",
			},
			AndroidLink: "http://www.agroinvestbank.tj/ru/index.php",
			IOSLink:     "http://www.agroinvestbank.tj/ru/index.php",
			Currencies: []models.Currencies{
				models.Currencies{
					Name: "USD",
					Buy:  "0.0000",
					Sell: "0.0000",
				},
				models.Currencies{
					Name: "RUB",
					Buy:  "0.0000",
					Sell: "0.0000",
				},
				models.Currencies{
					Name: "EUR",
					Buy:  "0.0000",
					Sell: "0.0000",
				},
			},
		}
		return
	}
	temp := []string{}
	// Find the review items
	doc.Find(".currency_lvl-2").Each(func(i int, s *goquery.Selection) {

		temp = append(temp, s.Text())
	})
	if len(temp) == 0 {
		che <- models.Bank{
			BankName:  "Агроинвестбанк",
			ShortName: "agroinvest",

			Icon: "",
			Colors: models.Colors{
				Color1: "#006CB5",
				Color2: "#057ED0",
			},
			AndroidLink: "http://www.agroinvestbank.tj/ru/index.php",
			IOSLink:     "http://www.agroinvestbank.tj/ru/index.php",
			Currencies: []models.Currencies{
				models.Currencies{
					Name: "USD",
					Buy:  "0.0000",
					Sell: "0.0000",
				},
				models.Currencies{
					Name: "RUB",
					Buy:  "0.0000",
					Sell: "0.0000",
				},
				models.Currencies{
					Name: "EUR",
					Buy:  "0.0000",
					Sell: "0.0000",
				},
			},
		}
		return
	}

	var Agroinvest []string
	for _, arr := range temp {
		_, err := strconv.ParseFloat(strings.TrimSpace(arr), 32)
		if err == nil {
			Agroinvest = append(Agroinvest, strings.TrimSpace(arr))
		}
	}

	che <- models.Bank{
		BankName:  "Агроинвестбанк",
		ShortName: "agroinvest",

		Icon: "",
		Colors: models.Colors{
			Color1: "#006CB5",
			Color2: "#057ED0",
		},
		AndroidLink: "http://www.agroinvestbank.tj/ru/index.php",
		IOSLink:     "http://www.agroinvestbank.tj/ru/index.php",
		Currencies: []models.Currencies{
			models.Currencies{
				Name: "USD",
				Buy:  Agroinvest[0],
				Sell: Agroinvest[3],
			},
			models.Currencies{
				Name: "RUB",
				Buy:  Agroinvest[2],
				Sell: Agroinvest[5],
			},
			models.Currencies{
				Name: "EUR",
				Buy:  Agroinvest[1],
				Sell: Agroinvest[4],
			},
		},
	}
	return

}
