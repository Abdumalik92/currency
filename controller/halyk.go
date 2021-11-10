package controller

import (
	"curr/models"
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
	"strings"
)

func HALIKBANK(che chan models.Bank) {
	client := &http.Client{}
	//proxyS, err := url.Parse("http://192.168.0.8:4480")
	//if err != nil {
	//	log.Println("HALIKBANK proxy ", err.Error())
	//	che <- models.Bank{
	//		BankName:    "Халык банк",
	//		ShortName:"halyk",
	//
	//		Icon: "",
	//		Colors: models.Colors{
	//			Color1: "#219653",
	//			Color2: "#18BD5E",
	//		},
	//		AndroidLink: "https://play.google.com/store/apps/details?id=kz.kkb.homebank&hl=ru",
	//		IOSLink:     "https://apps.apple.com/cy/app/homebank-kz/id440635615",
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

	//HALIKBANK

	respHALIKBANK, err := client.Get("http://halykbank.tj/")
	if err != nil {
		log.Println("KAZCOM res err ", err)
		che <- models.Bank{
			BankName:  "Халык банк",
			ShortName: "halyk",

			Icon: "",
			Colors: models.Colors{
				Color1: "#219653",
				Color2: "#18BD5E",
			},
			AndroidLink: "https://play.google.com/store/apps/details?id=kz.kkb.homebank&hl=ru",
			IOSLink:     "https://apps.apple.com/cy/app/homebank-kz/id440635615",
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
	doc, err := goquery.NewDocumentFromReader(respHALIKBANK.Body)
	if err != nil {
		log.Println("doc HALIKBANK err ", err)
		che <- models.Bank{
			BankName:  "Халык банк",
			ShortName: "halyk",

			Icon: "",
			Colors: models.Colors{
				Color1: "#219653",
				Color2: "#18BD5E",
			},
			AndroidLink: "https://play.google.com/store/apps/details?id=kz.kkb.homebank&hl=ru",
			IOSLink:     "https://apps.apple.com/cy/app/homebank-kz/id440635615",
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
	var KAZ []string
	// Find the review items
	doc.Find(".exchange_rates").Each(func(i int, s *goquery.Selection) {
		KAZ = append(KAZ, s.Find(".currency__group").Text())
	})

	for _, arr := range KAZ {
		KAZ = strings.Split(arr, "\n")
	}
	var halikBank []string
	for _, arr := range KAZ {
		halikBank = append(halikBank, strings.ReplaceAll(arr, " ", ""))
	}

	if len(halikBank) < 2 {
		che <- models.Bank{
			BankName:  "Халык банк",
			ShortName: "halyk",

			Icon: "",
			Colors: models.Colors{
				Color1: "#219653",
				Color2: "#18BD5E",
			},
			AndroidLink: "https://play.google.com/store/apps/details?id=kz.kkb.homebank&hl=ru",
			IOSLink:     "https://apps.apple.com/cy/app/homebank-kz/id440635615",
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
		BankName:  "Халык банк",
		ShortName: "halyk",

		Icon: "",
		Colors: models.Colors{
			Color1: "#219653",
			Color2: "#18BD5E",
		},
		AndroidLink: "https://play.google.com/store/apps/details?id=kz.kkb.homebank&hl=ru",
		IOSLink:     "https://apps.apple.com/cy/app/homebank-kz/id440635615",
		Currencies: []models.Currencies{
			models.Currencies{
				Name: "USD",
				Buy:  halikBank[1],
				Sell: halikBank[2],
			},
			models.Currencies{
				Name: "RUB",
				Buy:  halikBank[7],
				Sell: halikBank[8],
			},
			models.Currencies{
				Name: "EUR",
				Buy:  halikBank[4],
				Sell: halikBank[5],
			},
		},
	}
	return
}
