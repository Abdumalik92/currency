package controller

import (
	"curr/models"
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func NBP(che chan models.Bank) {
	client := &http.Client{}
	//proxyS, err := url.Parse("http://192.168.0.8:4480")
	//if err != nil {
	//	log.Println("NBP proxy ", err.Error())
	//	che <- models.Bank{
	//		BankName:    "Национальный банк Пакистана в Таджикистане",
	//		ShortName:"pakbankintj",
	//
	//		Icon: "",
	//		Colors: models.Colors{
	//			Color1: "#052F0D",
	//			Color2: "#077A1D",
	//		},
	//		AndroidLink: "http://www.nbp.tj/",
	//		IOSLink:     "http://www.nbp.tj/",
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

	//National bank of Pakistan in Tajikistan
	respNBP, err := client.Get("http://www.nbp.tj/")
	if err != nil {
		log.Println("NBP resp err ", err.Error())
		che <- models.Bank{
			BankName:  "Национальный банк Пакистана в Таджикистане",
			ShortName: "pakbankintj",

			Icon: "",
			Colors: models.Colors{
				Color1: "#052F0D",
				Color2: "#077A1D",
			},
			AndroidLink: "http://www.nbp.tj/",
			IOSLink:     "http://www.nbp.tj/",
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
	doc, err := goquery.NewDocumentFromReader(respNBP.Body)
	if err != nil {
		log.Println("Doc NBP err ", err.Error())
		che <- models.Bank{
			BankName:  "Национальный банк Пакистана в Таджикистане",
			ShortName: "pakbankintj",

			Icon: "",
			Colors: models.Colors{
				Color1: "#052F0D",
				Color2: "#077A1D",
			},
			AndroidLink: "http://www.nbp.tj/",
			IOSLink:     "http://www.nbp.tj/",
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
	doc.Find("#block-block-6").Each(func(i int, s *goquery.Selection) {

		temp = append(temp, s.Text())
	})
	if len(temp) == 0 {
		che <- models.Bank{
			BankName:  "Национальный банк Пакистана в Таджикистане",
			ShortName: "pakbankintj",

			Icon: "",
			Colors: models.Colors{
				Color1: "#052F0D",
				Color2: "#077A1D",
			},
			AndroidLink: "http://www.nbp.tj/",
			IOSLink:     "http://www.nbp.tj/",
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
	temp = strings.Split(temp[0], "\n")

	var NBP []string
	for _, arr := range temp {
		_, err := strconv.ParseFloat(strings.TrimSpace(arr), 32)
		if err == nil {
			NBP = append(NBP, strings.TrimSpace(arr))
		}
	}

	che <- models.Bank{
		BankName:  "Национальный банк Пакистана в Таджикистане",
		ShortName: "pakbankintj",

		Icon: "",
		Colors: models.Colors{
			Color1: "#052F0D",
			Color2: "#077A1D",
		},
		AndroidLink: "http://www.nbp.tj/",
		IOSLink:     "http://www.nbp.tj/",
		Currencies: []models.Currencies{
			models.Currencies{
				Name: "USD",
				Buy:  NBP[0],
				Sell: NBP[1],
			},
			models.Currencies{
				Name: "RUB",
				Buy:  "0.0000",
				Sell: "0.0000",
			},
			models.Currencies{
				Name: "EUR",
				Buy:  NBP[2],
				Sell: NBP[3],
			},
		},
	}
}
