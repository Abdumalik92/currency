package controller

import (
	"curr/models"
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func Eskhata(che chan models.Bank) {
	client := &http.Client{}
	//proxyS, err := url.Parse("http://192.168.0.8:4480")
	//if err != nil {
	//	log.Println("Eskhata proxy ", err.Error())
	//	che <- models.Bank{
	//		BankName:    "Банк Эсхата",
	//		ShortName:   "eskhata",
	//
	//		Icon: "",
	//		Colors: models.Colors{
	//			Color1: "#002269",
	//			Color2: "#0042CD",
	//		},
	//		AndroidLink: "https://play.google.com/store/apps/details?id=com.eskhata.online",
	//		IOSLink:     "https://apps.apple.com/tj/app/%D1%8D%D1%81%D1%85%D0%B0%D1%82%D0%B0-%D0%BE%D0%BD%D0%BB%D0%B0%D0%B9%D0%BD/id1438481790",
	//		Currencies: []models.Currencies{
	//			models.Currencies{Name: "USD",
	//				Buy:  "0.0000",
	//				Sell: "0.0000"},
	//			models.Currencies{Name: "RUB",
	//				Buy:  "0.0000",
	//				Sell: "0.0000"},
	//			models.Currencies{Name: "EUR",
	//				Buy:  "0.0000",
	//				Sell: "0.0000"},
	//		},
	//	}
	//	return
	//}
	//client.Transport = &http.Transport{Proxy: http.ProxyURL(proxyS)}

	//Eskhata
	respEskhata, err := client.Get("http://eskhata.com/")
	if err != nil {
		log.Println("Eskhata resp err ", err.Error())
		che <- models.Bank{
			BankName:  "Банк Эсхата",
			ShortName: "eskhata",
			Icon:      "",
			Colors: models.Colors{
				Color1: "#002269",
				Color2: "#0042CD",
			},
			AndroidLink: "https://play.google.com/store/apps/details?id=com.eskhata.online",
			IOSLink:     "https://apps.apple.com/tj/app/%D1%8D%D1%81%D1%85%D0%B0%D1%82%D0%B0-%D0%BE%D0%BD%D0%BB%D0%B0%D0%B9%D0%BD/id1438481790",
			Currencies: []models.Currencies{
				models.Currencies{Name: "USD",
					Buy:  "0.0000",
					Sell: "0.0000"},
				models.Currencies{Name: "RUB",
					Buy:  "0.0000",
					Sell: "0.0000"},
				models.Currencies{Name: "EUR",
					Buy:  "0.0000",
					Sell: "0.0000"},
			},
		}
		return
	}
	defer respEskhata.Body.Close()
	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(respEskhata.Body)
	if err != nil {
		log.Println("Doc Eskhata err ", err.Error())
		che <- models.Bank{
			BankName:  "Банк Эсхата",
			ShortName: "eskhata",
			Icon:      "",
			Colors: models.Colors{
				Color1: "#002269",
				Color2: "#0042CD",
			},
			AndroidLink: "https://play.google.com/store/apps/details?id=com.eskhata.online",
			IOSLink:     "https://apps.apple.com/tj/app/%D1%8D%D1%81%D1%85%D0%B0%D1%82%D0%B0-%D0%BE%D0%BD%D0%BB%D0%B0%D0%B9%D0%BD/id1438481790",
			Currencies: []models.Currencies{
				models.Currencies{Name: "USD",
					Buy:  "0.0000",
					Sell: "0.0000"},
				models.Currencies{Name: "RUB",
					Buy:  "0.0000",
					Sell: "0.0000"},
				models.Currencies{Name: "EUR",
					Buy:  "0.0000",
					Sell: "0.0000"},
			},
		}
		return
	}
	temp := []string{}
	// Find the review items
	doc.Find(".eb-second-page").Each(func(i int, s *goquery.Selection) {
		s.Find(".curs_value").Each(func(i int, t *goquery.Selection) {
			temp = append(temp, t.Text())
		})
	})

	if len(temp) == 0 {
		che <- models.Bank{
			BankName:  "Банк Эсхата",
			ShortName: "eskhata",
			Icon:      "",
			Colors: models.Colors{
				Color1: "#002269",
				Color2: "#0042CD",
			},
			AndroidLink: "https://play.google.com/store/apps/details?id=com.eskhata.online",
			IOSLink:     "https://apps.apple.com/tj/app/%D1%8D%D1%81%D1%85%D0%B0%D1%82%D0%B0-%D0%BE%D0%BD%D0%BB%D0%B0%D0%B9%D0%BD/id1438481790",
			Currencies: []models.Currencies{
				models.Currencies{Name: "USD",
					Buy:  "0.0000",
					Sell: "0.0000"},
				models.Currencies{Name: "RUB",
					Buy:  "0.0000",
					Sell: "0.0000"},
				models.Currencies{Name: "EUR",
					Buy:  "0.0000",
					Sell: "0.0000"},
			},
		}
		return
	}

	var Eskhata []string
	for _, arr := range temp {

		_, err := strconv.ParseFloat(strings.ReplaceAll(arr, ",", "."), 32)
		if err == nil {
			Eskhata = append(Eskhata, strings.ReplaceAll(arr, ",", "."))
		}
	}

	che <- models.Bank{
		BankName:  "Банк Эсхата",
		ShortName: "eskhata",

		Icon: "",
		Colors: models.Colors{
			Color1: "#002269",
			Color2: "#0042CD",
		},
		AndroidLink: "https://play.google.com/store/apps/details?id=com.eskhata.online",
		IOSLink:     "https://apps.apple.com/tj/app/%D1%8D%D1%81%D1%85%D0%B0%D1%82%D0%B0-%D0%BE%D0%BD%D0%BB%D0%B0%D0%B9%D0%BD/id1438481790",
		Currencies: []models.Currencies{
			models.Currencies{
				Name: "USD",
				Buy:  Eskhata[10],
				Sell: Eskhata[11],
			},
			models.Currencies{
				Name: "RUB",
				Buy:  Eskhata[8],
				Sell: Eskhata[9],
			},
			models.Currencies{
				Name: "EUR",
				Buy:  Eskhata[6],
				Sell: Eskhata[7],
			},
		},
	}
}
