package controller

import (
	"curr/models"
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func Finca(che chan models.Bank) {
	client := &http.Client{}
	//proxyS, err := url.Parse("http://192.168.0.8:4480")
	//if err != nil {
	//	log.Println("Finca proxy ", err.Error())
	//	che <- models.Bank{
	//		BankName:    "Финка",
	//		ShortName:"finca",
	//
	//		Icon: "",
	//		Colors: models.Colors{
	//			Color1: "#A1143A",
	//			Color2: "#B91844",
	//		},
	//		AndroidLink: "https://www.finca.tj/",
	//		IOSLink:     "https://www.finca.tj/",
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

	//Finca
	respFinca, err := client.Get("https://www.finca.tj/%d0%be%d0%b1%d0%bc%d0%b5%d0%bd-%d0%b2%d0%b0%d0%bb%d1%8e%d1%82/")
	if err != nil {
		log.Println("Finca resp err ", err.Error())
		che <- models.Bank{
			BankName:  "Финка",
			ShortName: "finca",

			Icon: "",
			Colors: models.Colors{
				Color1: "#A1143A",
				Color2: "#B91844",
			},
			AndroidLink: "https://www.finca.tj/",
			IOSLink:     "https://www.finca.tj/",
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
	doc, err := goquery.NewDocumentFromReader(respFinca.Body)
	if err != nil {
		log.Println("Doc Finca err ", err.Error())
		che <- models.Bank{
			BankName:  "Финка",
			ShortName: "finca",

			Icon: "",
			Colors: models.Colors{
				Color1: "#A1143A",
				Color2: "#B91844",
			},
			AndroidLink: "https://www.finca.tj/",
			IOSLink:     "https://www.finca.tj/",
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
	doc.Find(".exchange_rate_table").Each(func(i int, s *goquery.Selection) {
		s.Find("td").Each(func(i int, t *goquery.Selection) {
			temp = append(temp, t.Text())
		})
	})
	if len(temp) == 0 {
		che <- models.Bank{
			BankName:  "Финка",
			ShortName: "finca",

			Icon: "",
			Colors: models.Colors{
				Color1: "#A1143A",
				Color2: "#B91844",
			},
			AndroidLink: "https://www.finca.tj/",
			IOSLink:     "https://www.finca.tj/",
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
	var Finca []string
	for _, arr := range temp {
		_, err := strconv.ParseFloat(strings.TrimSpace(arr), 32)
		if err == nil {
			Finca = append(Finca, strings.TrimSpace(arr))
		}
	}

	che <- models.Bank{
		BankName:  "Финка",
		ShortName: "finca",

		Icon: "",
		Colors: models.Colors{
			Color1: "#A1143A",
			Color2: "#B91844",
		},
		AndroidLink: "https://www.finca.tj/",
		IOSLink:     "https://www.finca.tj/",
		Currencies: []models.Currencies{
			models.Currencies{
				Name: "USD",
				Buy:  Finca[2],
				Sell: Finca[3],
			},
			models.Currencies{
				Name: "RUB",
				Buy:  Finca[0],
				Sell: Finca[1],
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
