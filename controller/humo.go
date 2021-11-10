package controller

import (
	"curr/models"
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

func MDOHumo(che chan models.Bank) {
	client := &http.Client{}
	//proxyS, err := url.Parse("http://192.168.0.8:4480")
	//if err != nil {
	//	log.Println("MDOHumo proxy ", err.Error())
	//	che <- models.Bank{
	//		BankName:    "Хумо",
	//		ShortName:"humo",
	//
	//		Icon: "",
	//		Colors: models.Colors{
	//			Color1: "#DE5000",
	//			Color2: "#FC8D26",
	//		},
	//		AndroidLink: "https://play.google.com/store/apps/details?id=tj.humo.online",
	//		IOSLink:     "https://apps.apple.com/us/app/humo-online/id1242252363",
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

	//Humo
	respHumo, err := client.Get("https://humo.tj/ru/")
	if err != nil {
		log.Println("Humo resp err ", err.Error())
		che <- models.Bank{
			BankName:  "Хумо",
			ShortName: "humo",

			Icon: "",
			Colors: models.Colors{
				Color1: "#DE5000",
				Color2: "#FC8D26",
			},
			AndroidLink: "https://play.google.com/store/apps/details?id=tj.humo.online",
			IOSLink:     "https://apps.apple.com/us/app/humo-online/id1242252363",
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
	doc, err := goquery.NewDocumentFromReader(respHumo.Body)
	if err != nil {
		log.Println("Doc Humo err ", err.Error())
		che <- models.Bank{
			BankName:  "Хумо",
			ShortName: "humo",

			Icon: "",
			Colors: models.Colors{
				Color1: "#DE5000",
				Color2: "#FC8D26",
			},
			AndroidLink: "https://play.google.com/store/apps/details?id=tj.humo.online",
			IOSLink:     "https://apps.apple.com/us/app/humo-online/id1242252363",
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
				models.Currencies{Name: "EUR",
					Buy:  "0.0000",
					Sell: "0.0000",
				},
			},
		}
		return
	}
	temp := []string{}
	// Find the review items
	doc.Find(".kursHUMO").Each(func(i int, s *goquery.Selection) {
		temp = append(temp, s.Find(".kursBody").Text())
	})
	if len(temp) == 0 {
		che <- models.Bank{
			BankName:  "Хумо",
			ShortName: "humo",

			Icon: "",
			Colors: models.Colors{
				Color1: "#DE5000",
				Color2: "#FC8D26",
			},
			AndroidLink: "https://play.google.com/store/apps/details?id=tj.humo.online",
			IOSLink:     "https://apps.apple.com/us/app/humo-online/id1242252363",
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
				models.Currencies{Name: "EUR",
					Buy:  "0.0000",
					Sell: "0.0000",
				},
			},
		}
		return
	}
	temp = strings.Split(strings.TrimSpace(temp[0]), "\n")

	var Humo []string
	for _, arr := range temp {
		_, err := strconv.ParseFloat(strings.TrimSpace(arr), 32)
		if err == nil {
			Humo = append(Humo, strings.TrimSpace(arr))
		}
	}

	che <- models.Bank{
		BankName:  "Хумо",
		ShortName: "humo",

		Icon: "",
		Colors: models.Colors{
			Color1: "#DE5000",
			Color2: "#FC8D26",
		},
		AndroidLink: "https://play.google.com/store/apps/details?id=tj.humo.online",
		IOSLink:     "https://apps.apple.com/us/app/humo-online/id1242252363",
		Currencies: []models.Currencies{
			models.Currencies{
				Name: "USD",
				Buy:  Humo[0],
				Sell: Humo[1],
			},
			models.Currencies{
				Name: "RUB",
				Buy:  Humo[4],
				Sell: Humo[5],
			},
			models.Currencies{
				Name: "EUR",
				Buy:  Humo[2],
				Sell: Humo[3],
			},
		},
	}
	return
}

func MDOHumoTransfer(che chan models.Bank) {
	client := &http.Client{}
	proxyS, err := url.Parse("http://192.168.0.8:4480")
	if err != nil {
		log.Println("MDOHumo proxy ", err.Error())
		che <- models.Bank{
			BankName:  "Хумо",
			ShortName: "humo",

			Icon: "",
			Colors: models.Colors{
				Color1: "#DE5000",
				Color2: "#FC8D26",
			},
			AndroidLink: "https://play.google.com/store/apps/details?id=tj.humo.online",
			IOSLink:     "https://apps.apple.com/us/app/humo-online/id1242252363",
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
	client.Transport = &http.Transport{Proxy: http.ProxyURL(proxyS)}

	//Humo
	respHumo, err := client.Get("https://humo.tj/ru/")
	if err != nil {
		log.Println("Humo resp err ", err.Error())
		che <- models.Bank{
			BankName:  "Хумо",
			ShortName: "humo",

			Icon: "",
			Colors: models.Colors{
				Color1: "#DE5000",
				Color2: "#FC8D26",
			},
			AndroidLink: "https://play.google.com/store/apps/details?id=tj.humo.online",
			IOSLink:     "https://apps.apple.com/us/app/humo-online/id1242252363",
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
	doc, err := goquery.NewDocumentFromReader(respHumo.Body)
	if err != nil {
		log.Println("Doc Humo err ", err.Error())
		che <- models.Bank{
			BankName:  "Хумо",
			ShortName: "humo",

			Icon: "",
			Colors: models.Colors{
				Color1: "#DE5000",
				Color2: "#FC8D26",
			},
			AndroidLink: "https://play.google.com/store/apps/details?id=tj.humo.online",
			IOSLink:     "https://apps.apple.com/us/app/humo-online/id1242252363",
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
				models.Currencies{Name: "EUR",
					Buy:  "0.0000",
					Sell: "0.0000",
				},
			},
		}
		return
	}
	temp := []string{}
	// Find the review items
	doc.Find(".kursHUMO").Each(func(i int, s *goquery.Selection) {
		temp = append(temp, s.Find(".kursBody").Text())
	})
	if len(temp) == 0 {
		che <- models.Bank{
			BankName:  "Хумо",
			ShortName: "humo",

			Icon: "",
			Colors: models.Colors{
				Color1: "#DE5000",
				Color2: "#FC8D26",
			},
			AndroidLink: "https://play.google.com/store/apps/details?id=tj.humo.online",
			IOSLink:     "https://apps.apple.com/us/app/humo-online/id1242252363",
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
				models.Currencies{Name: "EUR",
					Buy:  "0.0000",
					Sell: "0.0000",
				},
			},
		}
		return
	}
	temp = strings.Split(strings.TrimSpace(temp[0]), "\n")

	var Humo []string
	for _, arr := range temp {
		_, err := strconv.ParseFloat(strings.TrimSpace(arr), 32)
		if err == nil {
			Humo = append(Humo, strings.TrimSpace(arr))
		}
	}

	che <- models.Bank{
		BankName:  "Хумо",
		ShortName: "humo",

		Icon: "",
		Colors: models.Colors{
			Color1: "#DE5000",
			Color2: "#FC8D26",
		},
		AndroidLink: "https://play.google.com/store/apps/details?id=tj.humo.online",
		IOSLink:     "https://apps.apple.com/us/app/humo-online/id1242252363",
		Currencies: []models.Currencies{
			models.Currencies{
				Name: "USD",
				Buy:  Humo[0],
				Sell: Humo[1],
			},
			models.Currencies{
				Name: "RUB",
				Buy:  Humo[4],
				Sell: Humo[5],
			},
			models.Currencies{
				Name: "EUR",
				Buy:  Humo[2],
				Sell: Humo[3],
			},
		},
	}
	return
}
