package controller

import (
	"curr/models"
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
	"strings"
)

func DCTransfer() models.Bank {
	client := &http.Client{}
	//proxyS, err := url.Parse("http://192.168.0.8:4480")
	//if err != nil {
	//	log.Println("DCTransfer proxy ", err.Error())
	//	return models.Bank{
	//		BankName:    "Душанбе сити",
	//		ShortName:"dc",
	//
	//		Icon: "",
	//		Colors: models.Colors{
	//			Color1: "#292E36",
	//			Color2: "#535F75",
	//		},
	//		AndroidLink: "http://online.dc.tj/",
	//		IOSLink:     "http://online.dc.tj/",
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

	respDC, err := client.Get("http://online.dc.tj/")
	if err != nil {
		log.Println("DCTransfer resp err ", err.Error())
		return models.Bank{
			BankName:  "Душанбе сити",
			ShortName: "dc",

			Icon: "",
			Colors: models.Colors{
				Color1: "#292E36",
				Color2: "#535F75",
			},
			AndroidLink: "http://online.dc.tj/",
			IOSLink:     "http://online.dc.tj/",
			Currencies: []models.Currencies{
				models.Currencies{
					Name: "RUB",
					Buy:  "0.0000",
					Sell: "0.0000",
				},
			},
		}
	}

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(respDC.Body)
	if err != nil {
		log.Println("Doc DCTransfer err ", err.Error())
		return models.Bank{
			BankName:  "Душанбе сити",
			ShortName: "dc",

			Icon: "",
			Colors: models.Colors{
				Color1: "#292E36",
				Color2: "#535F75",
			},
			AndroidLink: "http://online.dc.tj/",
			IOSLink:     "http://online.dc.tj/",
			Currencies: []models.Currencies{
				models.Currencies{
					Name: "RUB",
					Buy:  "0.0000",
					Sell: "0.0000",
				},
			},
		}
	}
	temp := ""
	// Find the review items
	doc.Find(".amount-input").Each(func(i int, s *goquery.Selection) {
		title := s.Find(".curr").Text()
		if title != "" {
			temp = strings.Trim(strings.Trim(title, "Курс 1₽ = "), " TJS")
		}
	})

	return models.Bank{
		BankName:  "Душанбе сити",
		ShortName: "dc",

		Icon: "",
		Colors: models.Colors{
			Color1: "#292E36",
			Color2: "#535F75",
		},
		AndroidLink: "http://online.dc.tj/",
		IOSLink:     "http://online.dc.tj/",
		Currencies: []models.Currencies{
			models.Currencies{
				Name: "RUB",
				Buy:  temp,
				Sell: "0.0000",
			},
		},
	}
}
