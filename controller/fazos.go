package controller

import (
	"curr/models"
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func FazoC(che chan models.Bank) {
	client := &http.Client{}
	//proxyS, err := url.Parse("http://192.168.0.8:4480")
	//if err != nil {
	//	log.Println("FazoC proxy ", err.Error())
	//	che <- models.Bank{
	//		BankName:    "Фазо С",
	//		ShortName:"fazos",
	//
	//		Icon: "",
	//		Colors: models.Colors{
	//			Color1: "#E5935D",
	//			Color2: "#C75104",
	//		},
	//		AndroidLink: "http://fazo-s.tj/index_rus.html",
	//		IOSLink:     "http://fazo-s.tj/index_rus.html",
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

	//Fazo C
	respFazoC, err := client.Get("http://fazo-s.tj/index_rus.html")
	if err != nil {
		log.Println("FazoC resp err ", err.Error())
		che <- models.Bank{
			BankName:  "Фазо С",
			ShortName: "fazos",

			Icon: "",
			Colors: models.Colors{
				Color1: "#E5935D",
				Color2: "#C75104",
			},
			AndroidLink: "http://fazo-s.tj/index_rus.html",
			IOSLink:     "http://fazo-s.tj/index_rus.html",
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
	doc, err := goquery.NewDocumentFromReader(respFazoC.Body)
	if err != nil {
		log.Println("Doc FazoC err ", err.Error())
		che <- models.Bank{
			BankName:  "Фазо С",
			ShortName: "fazos",

			Icon: "",
			Colors: models.Colors{
				Color1: "#E5935D",
				Color2: "#C75104",
			},
			AndroidLink: "http://fazo-s.tj/index_rus.html",
			IOSLink:     "http://fazo-s.tj/index_rus.html",
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
	doc.Find(".lefttd").Each(func(i int, s *goquery.Selection) {

		temp = append(temp, s.Text())
	})
	if len(temp) == 0 {
		che <- models.Bank{
			BankName:  "Фазо С",
			ShortName: "fazos",

			Icon: "",
			Colors: models.Colors{
				Color1: "#E5935D",
				Color2: "#C75104",
			},
			AndroidLink: "http://fazo-s.tj/index_rus.html",
			IOSLink:     "http://fazo-s.tj/index_rus.html",
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

	var FazoC []string
	for _, arr := range temp {
		_, err := strconv.ParseFloat(strings.TrimSpace(arr), 32)
		if err == nil {
			FazoC = append(FazoC, strings.TrimSpace(arr))
		}
	}

	che <- models.Bank{
		BankName:  "Фазо С",
		ShortName: "fazos",

		Icon: "",
		Colors: models.Colors{
			Color1: "#E5935D",
			Color2: "#C75104",
		},
		AndroidLink: "http://fazo-s.tj/index_rus.html",
		IOSLink:     "http://fazo-s.tj/index_rus.html",
		Currencies: []models.Currencies{
			models.Currencies{
				Name: "USD",
				Buy:  FazoC[1],
				Sell: FazoC[2],
			},
			models.Currencies{
				Name: "RUB",
				Buy:  FazoC[4],
				Sell: FazoC[5],
			},
			models.Currencies{
				Name: "EUR",
				Buy:  FazoC[7],
				Sell: FazoC[8]},
		},
	}
	return
}
