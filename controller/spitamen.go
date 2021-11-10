package controller

import (
	"curr/models"
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func Spitamen(che chan models.Bank) {
	client := &http.Client{}
	//proxyS, err := url.Parse("http://192.168.0.8:4480")
	//if err != nil {
	//	log.Println("Spitamen proxy ", err.Error())
	//	che <- models.Bank{
	//		BankName:    "Спитамен банк",
	//		ShortName:"spitamen",
	//
	//		Icon: "",
	//		Colors: models.Colors{
	//			Color1: "#292E36",
	//			Color2: "#535F75",
	//		},
	//		AndroidLink: "https://play.google.com/store/apps/details?id=com.bank.spitamen.pay.spitamenpay",
	//		IOSLink:     "https://apps.apple.com/ru/app/spitamen-pay/id1454198081",
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

	//Spitamen
	respSpitamen, err := client.Get("https://www.spitamenbank.tj/")
	if err != nil {
		log.Println("Spitamen resp err ", err.Error())
		che <- models.Bank{
			BankName:  "Спитамен банк",
			ShortName: "spitamen",

			Icon: "",
			Colors: models.Colors{
				Color1: "#292E36",
				Color2: "#535F75",
			},
			AndroidLink: "https://play.google.com/store/apps/details?id=com.bank.spitamen.pay.spitamenpay",
			IOSLink:     "https://apps.apple.com/ru/app/spitamen-pay/id1454198081",
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
	doc, err := goquery.NewDocumentFromReader(respSpitamen.Body)
	if err != nil {
		log.Println("Doc Spitamen err ", err.Error())
		che <- models.Bank{
			BankName:  "Спитамен банк",
			ShortName: "spitamen",

			Icon: "",
			Colors: models.Colors{
				Color1: "#292E36",
				Color2: "#535F75",
			},
			AndroidLink: "https://play.google.com/store/apps/details?id=com.bank.spitamen.pay.spitamenpay",
			IOSLink:     "https://apps.apple.com/ru/app/spitamen-pay/id1454198081",
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
	doc.Find(".conversation__table").Each(func(i int, s *goquery.Selection) {
		temp = append(temp, s.Find(".conversation__row").Text())
	})
	if len(temp) == 0 {
		che <- models.Bank{
			BankName:  "Спитамен банк",
			ShortName: "spitamen",

			Icon: "",
			Colors: models.Colors{
				Color1: "#292E36",
				Color2: "#535F75",
			},
			AndroidLink: "https://play.google.com/store/apps/details?id=com.bank.spitamen.pay.spitamenpay",
			IOSLink:     "https://apps.apple.com/ru/app/spitamen-pay/id1454198081",
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
	temp = strings.Split(strings.TrimSpace(temp[3]), "\n")

	var Spitamen []string
	for _, arr := range temp {
		_, err := strconv.ParseFloat(strings.TrimSpace(arr), 32)
		if err == nil {
			Spitamen = append(Spitamen, strings.TrimSpace(arr))
		}
	}

	che <- models.Bank{
		BankName:  "Спитамен банк",
		ShortName: "spitamen",

		Icon: "",
		Colors: models.Colors{
			Color1: "#292E36",
			Color2: "#535F75",
		},
		AndroidLink: "https://play.google.com/store/apps/details?id=com.bank.spitamen.pay.spitamenpay",
		IOSLink:     "https://apps.apple.com/ru/app/spitamen-pay/id1454198081",
		Currencies: []models.Currencies{
			models.Currencies{
				Name: "USD",
				Buy:  Spitamen[0],
				Sell: Spitamen[1],
			},
			models.Currencies{
				Name: "RUB",
				Buy:  Spitamen[4],
				Sell: Spitamen[5],
			},
			models.Currencies{
				Name: "EUR",
				Buy:  Spitamen[2],
				Sell: Spitamen[3],
			},
		},
	}
	return
}

func SpitamenTransfer() models.Bank {
	client := &http.Client{}
	//proxyS, err := url.Parse("http://192.168.0.8:4480")
	//if err != nil {
	//	log.Println("Spitamen proxy ", err.Error())
	//	return models.Bank{
	//		BankName:    "Спитамен банк",
	//		ShortName:"spitamen",
	//
	//		Icon: "",
	//		Colors: models.Colors{
	//			Color1: "#292E36",
	//			Color2: "#535F75",
	//		},
	//		AndroidLink: "https://c2c.spitamen.com",
	//		IOSLink:     "https://c2c.spitamen.com",
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

	//Spitamen
	respSpitamen, err := client.Get("https://c2c.spitamen.com")
	if err != nil {
		log.Println("Spitamen resp err ", err.Error())
		return models.Bank{
			BankName:  "Спитамен банк",
			ShortName: "spitamen",

			Icon: "",
			Colors: models.Colors{
				Color1: "#292E36",
				Color2: "#535F75",
			},
			AndroidLink: "https://c2c.spitamen.com",
			IOSLink:     "https://c2c.spitamen.com",
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
	doc, err := goquery.NewDocumentFromReader(respSpitamen.Body)
	if err != nil {
		log.Println("Doc Spitamen err ", err.Error())
		return models.Bank{
			BankName:  "Спитамен банк",
			ShortName: "spitamen",

			Icon: "",
			Colors: models.Colors{
				Color1: "#292E36",
				Color2: "#535F75",
			},
			AndroidLink: "https://c2c.spitamen.com",
			IOSLink:     "https://c2c.spitamen.com",
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
	doc.Find(".c-form-sec-item").Each(func(i int, s *goquery.Selection) {
		title := s.Find("#kurs").Text()
		if title != "" {
			temp = strings.TrimSpace(title)
		}
	})

	return models.Bank{
		BankName:  "Спитамен банк",
		ShortName: "spitamen",

		Icon: "",
		Colors: models.Colors{
			Color1: "#292E36",
			Color2: "#535F75",
		},
		AndroidLink: "https://c2c.spitamen.com",
		IOSLink:     "https://c2c.spitamen.com",
		Currencies: []models.Currencies{
			models.Currencies{
				Name: "RUB",
				Buy:  temp,
				Sell: "0.0000",
			},
		},
	}
}
