package controller

import (
	"curr/models"
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
	"strings"
)

func KAFOLAT(che chan models.Bank) {
	client := &http.Client{}
	//proxyS, err := url.Parse("http://192.168.0.8:4480")
	//if err != nil {
	//	log.Println("KAFOLAT proxy ", err.Error())
	//	che <- models.Bank{
	//		BankName:    "Кафолат банк",
	//		ShortName:"kafolat",
	//
	//		Icon: "",
	//		Colors: models.Colors{
	//			Color1: "#1F2B5D",
	//			Color2: "#475387",
	//		},
	//		AndroidLink: "https://kafolatbank.tj/",
	//		IOSLink:     "https://kafolatbank.tj/",
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

	respKAFOLAT, err := client.Get("http://kafolatbank.tj/")
	if err != nil {
		log.Println("KAFOLAT res ", err.Error())
		che <- models.Bank{
			BankName:  "Кафолат банк",
			ShortName: "kafolat",

			Icon: "",
			Colors: models.Colors{
				Color1: "#1F2B5D",
				Color2: "#475387",
			},
			AndroidLink: "https://kafolatbank.tj/",
			IOSLink:     "https://kafolatbank.tj/",
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
	doc, err := goquery.NewDocumentFromReader(respKAFOLAT.Body)
	if err != nil {
		log.Println("KAFOLAT doc ", err.Error())
		che <- models.Bank{
			BankName:  "Кафолат банк",
			ShortName: "kafolat",

			Icon: "",
			Colors: models.Colors{
				Color1: "#1F2B5D",
				Color2: "#475387",
			},
			AndroidLink: "https://kafolatbank.tj/",
			IOSLink:     "https://kafolatbank.tj/",
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
	var temp []string
	// Find the review items
	doc.Find(".container .currency").Each(func(i int, s *goquery.Selection) {
		temp = append(temp, s.Find(".currency").Text())
	})
	if len(temp) == 0 {
		che <- models.Bank{
			BankName:  "Кафолат банк",
			ShortName: "kafolat",

			Icon: "",
			Colors: models.Colors{
				Color1: "#1F2B5D",
				Color2: "#475387",
			},
			AndroidLink: "https://kafolatbank.tj/",
			IOSLink:     "https://kafolatbank.tj/",
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
	temp = strings.Split(strings.ReplaceAll(temp[0], " ", ""), "\n")

	var KAFOLAT []string
	for _, arr := range temp {
		if arr != "" {
			KAFOLAT = append(KAFOLAT, strings.ReplaceAll(arr, " ", ""))
		}
	}
	che <- models.Bank{
		BankName:  "Кафолат банк",
		ShortName: "kafolat",

		Icon: "",
		Colors: models.Colors{
			Color1: "#1F2B5D",
			Color2: "#475387",
		},
		AndroidLink: "https://kafolatbank.tj/",
		IOSLink:     "https://kafolatbank.tj/",
		Currencies: []models.Currencies{
			models.Currencies{
				Name: "USD",
				Buy:  KAFOLAT[1],
				Sell: KAFOLAT[3],
			},
			models.Currencies{
				Name: "RUB",
				Buy:  KAFOLAT[5],
				Sell: KAFOLAT[7],
			},
			models.Currencies{
				Name: "EUR",
				Buy:  KAFOLAT[9],
				Sell: KAFOLAT[11],
			},
		},
	}
	return
}
