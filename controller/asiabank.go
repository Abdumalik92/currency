package controller

import (
	"curr/models"
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func BankAsia(che chan models.Bank) {
	client := &http.Client{}
	//proxyS, err := url.Parse("http://192.168.0.8:4480")
	//if err != nil {
	//	log.Println("BankAsia proxy ", err.Error())
	//	che <- models.Bank{
	//		BankName:    "Банк Азии",
	//		ShortName:"asiabank",
	//
	//		Icon: "",
	//		Colors: models.Colors{
	//			Color1: "#D8A000",
	//			Color2: "#F1B200",
	//		},
	//		AndroidLink: "http://www.bankasia.tj/ru/",
	//		IOSLink:     "http://www.bankasia.tj/ru/",
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

	//BANK ASIA
	respBankAsia, err := client.Get("http://www.bankasia.tj/ru/")
	if err != nil {
		log.Println("BankAsia resp err ", err.Error())
		che <- models.Bank{
			BankName:  "Банк Азии",
			ShortName: "asiabank",

			Icon: "",
			Colors: models.Colors{
				Color1: "#D8A000",
				Color2: "#F1B200",
			},
			AndroidLink: "http://www.bankasia.tj/ru/",
			IOSLink:     "http://www.bankasia.tj/ru/",
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
	doc, err := goquery.NewDocumentFromReader(respBankAsia.Body)
	if err != nil {
		log.Println("Doc respBankAsia err ", err.Error())
		che <- models.Bank{
			BankName:  "Банк Азии",
			ShortName: "asiabank",

			Icon: "",
			Colors: models.Colors{
				Color1: "#D8A000",
				Color2: "#F1B200",
			},
			AndroidLink: "http://www.bankasia.tj/ru/",
			IOSLink:     "http://www.bankasia.tj/ru/",
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
	doc.Find(".rate-block-body").Each(func(i int, s *goquery.Selection) {
		temp = append(temp, s.Find(".rate-body").Text())
	})
	if len(temp) == 0 {
		che <- models.Bank{
			BankName:  "Банк Азии",
			ShortName: "asiabank",

			Icon: "",
			Colors: models.Colors{
				Color1: "#D8A000",
				Color2: "#F1B200",
			},
			AndroidLink: "http://www.bankasia.tj/ru/",
			IOSLink:     "http://www.bankasia.tj/ru/",
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

	var bankAsia []string
	for _, arr := range temp {
		_, err := strconv.ParseFloat(strings.TrimSpace(arr), 32)
		if err == nil {
			bankAsia = append(bankAsia, strings.TrimSpace(arr))
		}
	}
	che <- models.Bank{
		BankName:  "Банк Азии",
		ShortName: "asiabank",

		Icon: "",
		Colors: models.Colors{
			Color1: "#D8A000",
			Color2: "#F1B200",
		},
		AndroidLink: "http://www.bankasia.tj/ru/",
		IOSLink:     "http://www.bankasia.tj/ru/",
		Currencies: []models.Currencies{
			models.Currencies{
				Name: "USD",
				Buy:  bankAsia[0],
				Sell: bankAsia[1],
			},
			models.Currencies{
				Name: "RUB",
				Buy:  bankAsia[6],
				Sell: bankAsia[7],
			},
			models.Currencies{
				Name: "EUR",
				Buy:  bankAsia[3],
				Sell: bankAsia[4],
			},
		},
	}
	return
}
