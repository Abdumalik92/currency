package controller

import (
	"curr/models"
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func CBT(che chan models.Bank) {
	client := &http.Client{}
	//proxyS, err := url.Parse("http://192.168.0.8:4480")
	//if err != nil {
	//	log.Println("CBT proxy ", err.Error())
	//	che <- models.Bank{
	//		BankName:    "Коммерсбанк Таджикистана",
	//		ShortName:"commerce",
	//
	//		Icon: "",
	//		Colors: models.Colors{
	//			Color1: "#126351",
	//			Color2: "#16987B",
	//		},
	//		AndroidLink: "https://play.google.com/store/apps/details?id=com.doocat.banking.favri",
	//		IOSLink:     "https://apps.apple.com/ru/app/favri/id1397302137",
	//		Currencies: []models.Currencies{
	//			models.Currencies{
	//				Name: "USD",
	//				Buy:  "0.0000",
	//				Sell: "0.0000"},
	//			models.Currencies{
	//				Name: "RUB",
	//				Buy:  "0.0000",
	//				Sell: "0.0000"},
	//			models.Currencies{
	//				Name: "EUR",
	//				Buy:  "0.0000",
	//				Sell: "0.0000"},
	//		},
	//	}
	//	return
	//}
	//client.Transport = &http.Transport{Proxy: http.ProxyURL(proxyS)}

	//CBT
	respCBT, err := client.Get("https://cbt.tj/")
	if err != nil {
		log.Println("CBT resp err ", err.Error())
		che <- models.Bank{
			BankName:  "Коммерсбанк Таджикистана",
			ShortName: "commerce",

			Icon: "",
			Colors: models.Colors{
				Color1: "#126351",
				Color2: "#16987B",
			},
			AndroidLink: "https://play.google.com/store/apps/details?id=com.doocat.banking.favri",
			IOSLink:     "https://apps.apple.com/ru/app/favri/id1397302137",
			Currencies: []models.Currencies{
				models.Currencies{
					Name: "USD",
					Buy:  "0.0000",
					Sell: "0.0000"},
				models.Currencies{
					Name: "RUB",
					Buy:  "0.0000",
					Sell: "0.0000"},
				models.Currencies{
					Name: "EUR",
					Buy:  "0.0000",
					Sell: "0.0000"},
			},
		}
		return
	}
	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(respCBT.Body)
	if err != nil {
		log.Println("Doc CBT err ", err.Error())
		che <- models.Bank{
			BankName:  "Коммерсбанк Таджикистана",
			ShortName: "commerce",

			Icon: "",
			Colors: models.Colors{
				Color1: "#126351",
				Color2: "#16987B",
			},
			AndroidLink: "https://play.google.com/store/apps/details?id=com.doocat.banking.favri",
			IOSLink:     "https://apps.apple.com/ru/app/favri/id1397302137",
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
	doc.Find(".cbt-currency").Each(func(i int, s *goquery.Selection) {
		temp = append(temp, strings.TrimSpace(strings.ReplaceAll(s.Text(), "\n", " ")))
	})

	if len(temp) == 0 {
		che <- models.Bank{
			BankName:  "Коммерсбанк Таджикистана",
			ShortName: "commerce",

			Icon: "",
			Colors: models.Colors{
				Color1: "#126351",
				Color2: "#16987B",
			},
			AndroidLink: "https://play.google.com/store/apps/details?id=com.doocat.banking.favri",
			IOSLink:     "https://apps.apple.com/ru/app/favri/id1397302137",
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

	temp = strings.Split(strings.TrimSpace(temp[1]), " ")

	var CBT []string
	for _, arr := range temp {
		_, err := strconv.ParseFloat(strings.TrimSpace(arr), 32)
		if err == nil {
			CBT = append(CBT, strings.TrimSpace(arr))
		}
	}

	che <- models.Bank{
		BankName:  "Коммерсбанк Таджикистана",
		ShortName: "commerce",

		Icon: "",
		Colors: models.Colors{
			Color1: "#126351",
			Color2: "#16987B",
		},
		AndroidLink: "https://play.google.com/store/apps/details?id=com.doocat.banking.favri",
		IOSLink:     "https://apps.apple.com/ru/app/favri/id1397302137",
		Currencies: []models.Currencies{
			models.Currencies{
				Name: "USD",
				Buy:  CBT[0],
				Sell: CBT[1],
			},
			models.Currencies{
				Name: "RUB",
				Buy:  CBT[4],
				Sell: CBT[5],
			},
			models.Currencies{
				Name: "EUR",
				Buy:  CBT[2],
				Sell: CBT[3],
			},
		},
	}
	return
}

func CBTTransfer() models.Bank {
	client := &http.Client{}
	//proxyS, err := url.Parse("http://192.168.0.8:4480")
	//if err != nil {
	//	log.Println("CBTTransfer proxy ", err.Error())
	//	return models.Bank{
	//		BankName:    "Коммерсбанк Таджикистана",
	//		ShortName:"commerce",
	//
	//		Icon: "",
	//		Colors: models.Colors{
	//			Color1: "#126351",
	//			Color2: "#16987B",
	//		},
	//		AndroidLink: "https://play.google.com/store/apps/details?id=tj.cbt.favri.intiqol&hl=en_US&gl=US",
	//		IOSLink:     "https://apps.apple.com/ru/app/favri-intiqol/id1537354404",
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
	//CBT
	respCBT, err := client.Get("https://online.cbt.tj/ex-rate-view-main.action")
	if err != nil {
		log.Println("CBT resp err ", err.Error())
		return models.Bank{
			BankName:  "Коммерсбанк Таджикистана",
			ShortName: "commerce",

			Icon: "",
			Colors: models.Colors{
				Color1: "#126351",
				Color2: "#16987B",
			},
			AndroidLink: "https://play.google.com/store/apps/details?id=tj.cbt.favri.intiqol&hl=en_US&gl=US",
			IOSLink:     "https://apps.apple.com/ru/app/favri-intiqol/id1537354404",
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
	doc, err := goquery.NewDocumentFromReader(respCBT.Body)
	if err != nil {
		log.Println("Doc CBT err ", err.Error())
		return models.Bank{
			BankName:  "Коммерсбанк Таджикистана",
			ShortName: "commerce",

			Icon: "",
			Colors: models.Colors{
				Color1: "#126351",
				Color2: "#16987B",
			},
			AndroidLink: "https://play.google.com/store/apps/details?id=tj.cbt.favri.intiqol&hl=en_US&gl=US",
			IOSLink:     "https://apps.apple.com/ru/app/favri-intiqol/id1537354404",
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
	doc.Find("#rates_favri .currency_today").Each(func(i int, s *goquery.Selection) {
		title := s.Text()
		if title != "" {
			temp = strings.TrimSpace(title)
		}
	})

	if len(temp) == 0 {
		return models.Bank{
			BankName:  "Коммерсбанк Таджикистана",
			ShortName: "commerce",

			Icon: "",
			Colors: models.Colors{
				Color1: "#126351",
				Color2: "#16987B",
			},
			AndroidLink: "https://play.google.com/store/apps/details?id=tj.cbt.favri.intiqol&hl=en_US&gl=US",
			IOSLink:     "https://apps.apple.com/ru/app/favri-intiqol/id1537354404",
			Currencies: []models.Currencies{
				models.Currencies{
					Name: "RUB",
					Buy:  "0.0000",
					Sell: "0.0000",
				},
			},
		}
	}

	return models.Bank{
		BankName:  "Коммерсбанк Таджикистана",
		ShortName: "commerce",

		Icon: "",
		Colors: models.Colors{
			Color1: "#126351",
			Color2: "#16987B",
		},
		AndroidLink: "https://play.google.com/store/apps/details?id=com.doocat.banking.favri",
		IOSLink:     "https://apps.apple.com/ru/app/favri/id1397302137",
		Currencies: []models.Currencies{
			models.Currencies{
				Name: "RUB",
				Buy:  temp,
				Sell: "0.0000",
			},
		},
	}
}
