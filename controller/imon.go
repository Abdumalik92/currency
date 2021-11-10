package controller

import (
	"curr/models"
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
)

func Imon(che chan models.Bank) {
	client := &http.Client{}
	//proxyS, err := url.Parse("http://192.168.0.8:4480")
	//if err != nil {
	//	log.Println("Imon proxy ", err.Error())
	//	che <- models.Bank{
	//		BankName:    "Имон Интернейшнал",
	//		ShortName:"imon",
	//
	//		Icon: "",
	//		Colors: models.Colors{
	//			Color1: "#007CC2",
	//			Color2: "#4698D5",
	//		},
	//		AndroidLink: "https://play.google.com/store/apps/details?id=tj.imon.ikfl",
	//		IOSLink:     "https://apps.apple.com/us/app/imon-online/id1479158764",
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

	//Imon International
	respImon, err := client.Get("https://imon.tj/")
	if err != nil {
		log.Println("Imon resp err ", err.Error())
		che <- models.Bank{
			BankName:  "Имон Интернейшнал",
			ShortName: "imon",

			Icon: "",
			Colors: models.Colors{
				Color1: "#007CC2",
				Color2: "#4698D5",
			},
			AndroidLink: "https://play.google.com/store/apps/details?id=tj.imon.ikfl",
			IOSLink:     "https://apps.apple.com/us/app/imon-online/id1479158764",
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
	defer respImon.Body.Close()
	Imon := models.Imon{}
	doc, err := goquery.NewDocumentFromReader(respImon.Body)
	if err != nil {
		log.Println("Imon resp body err ", err.Error())
		che <- models.Bank{
			BankName:  "Имон Интернейшнал",
			ShortName: "imon",

			Icon: "",
			Colors: models.Colors{
				Color1: "#007CC2",
				Color2: "#4698D5",
			},
			AndroidLink: "https://play.google.com/store/apps/details?id=tj.imon.ikfl",
			IOSLink:     "https://apps.apple.com/us/app/imon-online/id1479158764",
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
	doc.Find(".row .mt-2 .text-center .text-md-left .text-lg-left .currancy").Each(func(i int, s *goquery.Selection) {
		s.Find(".col-6 .text-right .text-md-left").Each(func(i int, t *goquery.Selection) {
			temp = append(temp, t.Text())
		})
	})

	if len(temp) == 0 {
		che <- models.Bank{
			BankName:  "Имон Интернейшнал",
			ShortName: "imon",

			Icon: "",
			Colors: models.Colors{
				Color1: "#007CC2",
				Color2: "#4698D5",
			},
			AndroidLink: "https://play.google.com/store/apps/details?id=tj.imon.ikfl",
			IOSLink:     "https://apps.apple.com/us/app/imon-online/id1479158764",
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

	che <- models.Bank{
		BankName:  "Имон Интернейшнал",
		ShortName: "imon",

		Icon: "",
		Colors: models.Colors{
			Color1: "#007CC2",
			Color2: "#4698D5",
		},
		AndroidLink: "https://play.google.com/store/apps/details?id=tj.imon.ikfl",
		IOSLink:     "https://apps.apple.com/us/app/imon-online/id1479158764",
		Currencies: []models.Currencies{
			models.Currencies{
				Name: "USD",
				Buy:  Imon.Buy.USD.Rate,
				Sell: Imon.Sell.USD.Rate,
			},
			models.Currencies{
				Name: "RUB",
				Buy:  Imon.Buy.RUB.Rate,
				Sell: Imon.Sell.RUB.Rate,
			},
			models.Currencies{
				Name: "EUR",
				Buy:  Imon.Buy.EUR.Rate,
				Sell: Imon.Sell.EUR.Rate,
			},
		},
	}
	return
}
