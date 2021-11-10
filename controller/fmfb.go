package controller

import (
	"curr/models"
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
	"strings"
)

func FMBT(che chan models.Bank) {
	var fmfb []string
	client := &http.Client{}
	//proxyS, err := url.Parse("http://192.168.0.8:4480")
	//if err != nil {
	//	log.Println("FMBT proxy ", err.Error())
	//	che <- models.Bank{
	//		BankName:    "Первый микрофинансовый банк",
	//		ShortName:   "microfin",
	//
	//		Icon: "",
	//		Colors: models.Colors{
	//			Color1: "#0069BC",
	//			Color2: "#0080E5",
	//		},
	//		AndroidLink: "https://fmfb.tj/ru/",
	//		IOSLink:     "https://fmfb.tj/ru/",
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

	//FMFB

	// Request the HTML page.
	respFMFB, err := client.Get("https://fmfb.tj/ru/")
	if err != nil {
		log.Println("res FMFB err ", err.Error())
		che <- models.Bank{
			BankName:  "Первый микрофинансовый банк",
			ShortName: "microfin",

			Icon: "",
			Colors: models.Colors{
				Color1: "#0069BC",
				Color2: "#0080E5",
			},
			AndroidLink: "https://fmfb.tj/ru/",
			IOSLink:     "https://fmfb.tj/ru/",
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
	doc, err := goquery.NewDocumentFromReader(respFMFB.Body)
	if err != nil {
		log.Println("res FMFB err ", err.Error())
		che <- models.Bank{
			BankName:  "Первый микрофинансовый банк",
			ShortName: "microfin",

			Icon: "",
			Colors: models.Colors{
				Color1: "#0069BC",
				Color2: "#0080E5",
			},
			AndroidLink: "https://fmfb.tj/ru/",
			IOSLink:     "https://fmfb.tj/ru/",
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

	// Find the review items
	doc.Find(".new-currency-last").Each(func(i int, s *goquery.Selection) {
		fmfb = append(fmfb, s.Text())
	})
	if len(fmfb) == 0 {
		che <- models.Bank{
			BankName:  "Первый микрофинансовый банк",
			ShortName: "microfin",

			Icon: "",
			Colors: models.Colors{
				Color1: "#0069BC",
				Color2: "#0080E5",
			},
			AndroidLink: "https://fmfb.tj/ru/",
			IOSLink:     "https://fmfb.tj/ru/",
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
	s1 := strings.Trim(strings.ReplaceAll(fmfb[1], "\n", ""), "Покупка ")
	s2 := strings.ReplaceAll(s1, "Продажа ", " ")
	fmfb = strings.Split(s2, " ")
	che <- models.Bank{
		BankName:  "Первый микрофинансовый банк",
		ShortName: "microfin",

		Icon: "",
		Colors: models.Colors{
			Color1: "#0069BC",
			Color2: "#0080E5",
		},
		AndroidLink: "https://fmfb.tj/ru/",
		IOSLink:     "https://fmfb.tj/ru/",
		Currencies: []models.Currencies{
			models.Currencies{
				Name: "USD",
				Buy:  fmfb[0],
				Sell: fmfb[3]},
			models.Currencies{
				Name: "RUB",
				Buy:  fmfb[2],
				Sell: fmfb[5],
			},
			models.Currencies{
				Name: "EUR",
				Buy:  fmfb[1],
				Sell: fmfb[4],
			},
		},
	}
	return
}
