package controller

import (
	"curr/models"
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func Arvand(che chan models.Bank) {
	client := &http.Client{}
	//proxyS, err := url.Parse("http://192.168.0.8:4480")
	//if err != nil {
	//	log.Println("Arvand proxy ", err.Error())
	//	che <- models.Bank{
	//		BankName:    "Банк Авранд",
	//		ShortName:"arvand",
	//
	//		Icon: "",
	//		Colors: models.Colors{
	//			Color1: "#1B98A0",
	//			Color2: "#52DDE5",
	//		},
	//		AndroidLink: "https://play.google.com/store/apps/details?id=tj.itservice.arvand&hl=ru-RU",
	//		IOSLink:     "https://apps.apple.com/app/id1462273028",
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

	//Arvand
	respArvand, err := client.Get("https://www.arvand.tj/")
	if err != nil {
		log.Println("Arvand resp err ", err.Error())
		che <- models.Bank{
			BankName:  "Банк Авранд",
			ShortName: "arvand",

			Icon: "",
			Colors: models.Colors{
				Color1: "#1B98A0",
				Color2: "#52DDE5",
			},
			AndroidLink: "https://play.google.com/store/apps/details?id=tj.itservice.arvand&hl=ru-RU",
			IOSLink:     "https://apps.apple.com/app/id1462273028",
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
	doc, err := goquery.NewDocumentFromReader(respArvand.Body)
	if err != nil {
		log.Println("Doc Arvand err ", err.Error())
		che <- models.Bank{
			BankName:  "Банк Авранд",
			ShortName: "arvand",

			Icon: "",
			Colors: models.Colors{
				Color1: "#1B98A0",
				Color2: "#52DDE5",
			},
			AndroidLink: "https://play.google.com/store/apps/details?id=tj.itservice.arvand&hl=ru-RU",
			IOSLink:     "https://apps.apple.com/app/id1462273028",
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
	doc.Find(".currencyValue").Each(func(i int, s *goquery.Selection) {
		temp = append(temp, s.Text())
	})
	if len(temp) == 0 {
		che <- models.Bank{
			BankName:  "Банк Авранд",
			ShortName: "arvand",

			Icon: "",
			Colors: models.Colors{
				Color1: "#1B98A0",
				Color2: "#52DDE5",
			},
			AndroidLink: "https://play.google.com/store/apps/details?id=tj.itservice.arvand&hl=ru-RU",
			IOSLink:     "https://apps.apple.com/app/id1462273028",
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
	var Arvand []string
	for _, arr := range temp {
		_, err := strconv.ParseFloat(strings.TrimSpace(arr), 32)
		if err == nil {
			Arvand = append(Arvand, strings.TrimSpace(arr))
		}
	}

	che <- models.Bank{
		BankName:  "Банк Авранд",
		ShortName: "arvand",

		Icon: "",
		Colors: models.Colors{
			Color1: "#1B98A0",
			Color2: "#52DDE5",
		},
		AndroidLink: "https://play.google.com/store/apps/details?id=tj.itservice.arvand&hl=ru-RU",
		IOSLink:     "https://apps.apple.com/app/id1462273028",
		Currencies: []models.Currencies{
			models.Currencies{
				Name: "USD",
				Buy:  Arvand[3],
				Sell: Arvand[4],
			},
			models.Currencies{
				Name: "RUB",
				Buy:  Arvand[7],
				Sell: Arvand[8],
			},
			models.Currencies{
				Name: "EUR",
				Buy:  Arvand[5],
				Sell: Arvand[6],
			},
		},
	}
	return
}
