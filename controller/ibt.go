package controller

import (
	"curr/models"
	"encoding/json"
	"github.com/PuerkitoBio/goquery"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func IBT(che chan models.Bank) {
	client := &http.Client{}
	//proxyS, err := url.Parse("http://192.168.0.8:4480")
	//if err != nil {
	//	log.Println("IBT proxy ", err.Error())
	//	che <- models.Bank{
	//		BankName:    "Международный банк Таджикистана",
	//		ShortName:"ibt",
	//
	//		Icon: "",
	//		Colors: models.Colors{
	//			Color1: "#145F84",
	//			Color2: "#117EB4",
	//		},
	//		AndroidLink: "https://play.google.com/store/apps/details?id=tj.ibt.app",
	//		IOSLink:     "https://apps.apple.com/us/app/ibt-24/id1515094392",
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

	//IBT
	respIBT, err := client.Get("https://www.ibt.tj/")
	if err != nil {
		log.Println("IBT resp err ", err.Error())
		che <- models.Bank{
			BankName:  "Международный банк Таджикистана",
			ShortName: "ibt",

			Icon: "",
			Colors: models.Colors{
				Color1: "#145F84",
				Color2: "#117EB4",
			},
			AndroidLink: "https://play.google.com/store/apps/details?id=tj.ibt.app",
			IOSLink:     "https://apps.apple.com/us/app/ibt-24/id1515094392",
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
	doc, err := goquery.NewDocumentFromReader(respIBT.Body)
	if err != nil {
		log.Println("Doc IBT err ", err.Error())
		che <- models.Bank{
			BankName:  "Международный банк Таджикистана",
			ShortName: "ibt",

			Icon: "",
			Colors: models.Colors{
				Color1: "#145F84",
				Color2: "#117EB4",
			},
			AndroidLink: "https://play.google.com/store/apps/details?id=tj.ibt.app",
			IOSLink:     "https://apps.apple.com/us/app/ibt-24/id1515094392",
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
	doc.Find("#ibt").Each(func(i int, s *goquery.Selection) {
		temp = append(temp, s.Text())
	})
	if len(temp) == 0 {
		che <- models.Bank{
			BankName:  "Международный банк Таджикистана",
			ShortName: "ibt",

			Icon: "",
			Colors: models.Colors{
				Color1: "#145F84",
				Color2: "#117EB4",
			},
			AndroidLink: "https://play.google.com/store/apps/details?id=tj.ibt.app",
			IOSLink:     "https://apps.apple.com/us/app/ibt-24/id1515094392",
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
	temp = strings.Split(strings.TrimSpace(temp[0]), "\n")

	var IBT []string
	for _, arr := range temp {
		_, err := strconv.ParseFloat(strings.TrimSpace(arr), 32)
		if err == nil {
			IBT = append(IBT, strings.TrimSpace(arr))
		}
	}
	che <- models.Bank{
		BankName:  "Международный банк Таджикистана",
		ShortName: "ibt",

		Icon: "",
		Colors: models.Colors{
			Color1: "#145F84",
			Color2: "#117EB4",
		},
		AndroidLink: "https://play.google.com/store/apps/details?id=tj.ibt.app",
		IOSLink:     "https://apps.apple.com/us/app/ibt-24/id1515094392",
		Currencies: []models.Currencies{
			models.Currencies{
				Name: "USD",
				Buy:  IBT[0],
				Sell: IBT[1],
			},
			models.Currencies{
				Name: "RUB",
				Buy:  IBT[2],
				Sell: IBT[3],
			},
			models.Currencies{
				Name: "EUR",
				Buy:  IBT[4],
				Sell: IBT[5],
			},
		},
	}
	return
}

func IBTTranser() models.Bank {

	client := &http.Client{}
	//proxyS, err := url.Parse("http://192.168.0.8:4480")
	//if err != nil {
	//	log.Println("IBT proxy ", err.Error())
	//	return models.Bank{
	//		BankName:    "Международный банк Таджикистана",
	//		ShortName:"ibt",
	//
	//		Icon: "",
	//		Colors: models.Colors{
	//			Color1: "#145F84",
	//			Color2: "#117EB4",
	//		},
	//		AndroidLink: "https://play.google.com/store/apps/details?id=com.intervale.ibt&referrer=utm_source%3Dwidget%26utm_medium%3Dfree%26utm_campaign%3Dgeneral",
	//		IOSLink:     "https://apps.apple.com/us/app/ibt-direct-%D0%BF%D0%B5%D1%80%D0%B5%D0%B2%D0%BE%D0%B4-%D0%B4%D0%B5%D0%BD%D0%B5%D0%B3/id1473465023?ls=1",
	//		Currencies: []models.Currencies{
	//			models.Currencies{
	//				Name: "RUB",
	//				Buy:  "0.0000",
	//				Sell: "0.0000",
	//			},
	//		},
	//	}
	//
	//}
	//
	//client.Transport = &http.Transport{Proxy: http.ProxyURL(proxyS)}
	url1 := "https://openapi-entry.intervale.ru/api/v4/P2PIBT706DC183EF5DA6BE1D1A6CB9D9/payment/terms"
	payload := strings.NewReader("dst.pan=5058270301046610&dst.type=card&lang=RU&paymentId=C2C_IBT")

	//IBT
	respIBT, err := client.Post(url1, "application/x-www-form-urlencoded", payload)
	if err != nil {
		log.Println("IBT resp err ", err.Error())
		return models.Bank{
			BankName:  "Международный банк Таджикистана",
			ShortName: "ibt",

			Icon: "",
			Colors: models.Colors{
				Color1: "#145F84",
				Color2: "#117EB4",
			},
			AndroidLink: "https://play.google.com/store/apps/details?id=com.intervale.ibt&referrer=utm_source%3Dwidget%26utm_medium%3Dfree%26utm_campaign%3Dgeneral",
			IOSLink:     "https://apps.apple.com/us/app/ibt-direct-%D0%BF%D0%B5%D1%80%D0%B5%D0%B2%D0%BE%D0%B4-%D0%B4%D0%B5%D0%BD%D0%B5%D0%B3/id1473465023?ls=1",
			Currencies: []models.Currencies{
				models.Currencies{
					Name: "RUB",
					Buy:  "0.0000",
					Sell: "0.0000",
				},
			},
		}

	}
	IBTRub := models.IBT{}

	bodyIBT, err := ioutil.ReadAll(respIBT.Body)
	if err != nil {
		log.Println("IBT resp body ", err.Error())
		return models.Bank{
			BankName:  "Международный банк Таджикистана",
			ShortName: "ibt",

			Icon: "",
			Colors: models.Colors{
				Color1: "#145F84",
				Color2: "#117EB4",
			},
			AndroidLink: "https://play.google.com/store/apps/details?id=com.intervale.ibt&referrer=utm_source%3Dwidget%26utm_medium%3Dfree%26utm_campaign%3Dgeneral",
			IOSLink:     "https://apps.apple.com/us/app/ibt-direct-%D0%BF%D0%B5%D1%80%D0%B5%D0%B2%D0%BE%D0%B4-%D0%B4%D0%B5%D0%BD%D0%B5%D0%B3/id1473465023?ls=1",
			Currencies: []models.Currencies{
				models.Currencies{
					Name: "RUB",
					Buy:  "0.0000",
					Sell: "0.0000",
				},
			},
		}

	}

	err = json.Unmarshal(bodyIBT, &IBTRub)
	if err != nil {
		log.Println("IBT can't unmarshal json ", err.Error())
		return models.Bank{
			BankName:  "Международный банк Таджикистана",
			ShortName: "ibt",

			Icon: "",
			Colors: models.Colors{
				Color1: "#145F84",
				Color2: "#117EB4",
			},
			AndroidLink: "https://play.google.com/store/apps/details?id=com.intervale.ibt&referrer=utm_source%3Dwidget%26utm_medium%3Dfree%26utm_campaign%3Dgeneral",
			IOSLink:     "https://apps.apple.com/us/app/ibt-direct-%D0%BF%D0%B5%D1%80%D0%B5%D0%B2%D0%BE%D0%B4-%D0%B4%D0%B5%D0%BD%D0%B5%D0%B3/id1473465023?ls=1",
			Currencies: []models.Currencies{
				models.Currencies{
					Name: "RUB",
					Buy:  "0.0000",
					Sell: "0.0000",
				},
			},
		}

	}

	if len(IBTRub.Extra.Params) == 0 {
		return models.Bank{
			BankName:  "Международный банк Таджикистана",
			ShortName: "ibt",

			Icon: "",
			Colors: models.Colors{
				Color1: "#145F84",
				Color2: "#117EB4",
			},
			AndroidLink: "https://play.google.com/store/apps/details?id=com.intervale.ibt&referrer=utm_source%3Dwidget%26utm_medium%3Dfree%26utm_campaign%3Dgeneral",
			IOSLink:     "https://apps.apple.com/us/app/ibt-direct-%D0%BF%D0%B5%D1%80%D0%B5%D0%B2%D0%BE%D0%B4-%D0%B4%D0%B5%D0%BD%D0%B5%D0%B3/id1473465023?ls=1",
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
		BankName:  "Международный банк Таджикистана",
		ShortName: "ibt",
		Icon:      "",
		Colors: models.Colors{
			Color1: "#145F84",
			Color2: "#117EB4",
		},
		AndroidLink: "https://play.google.com/store/apps/details?id=com.intervale.ibt&referrer=utm_source%3Dwidget%26utm_medium%3Dfree%26utm_campaign%3Dgeneral",
		IOSLink:     "https://apps.apple.com/us/app/ibt-direct-%D0%BF%D0%B5%D1%80%D0%B5%D0%B2%D0%BE%D0%B4-%D0%B4%D0%B5%D0%BD%D0%B5%D0%B3/id1473465023?ls=1",
		Currencies: []models.Currencies{
			models.Currencies{
				Name: "RUB",
				Buy:  IBTRub.Extra.Params[2].Value,
				Sell: "0.0000",
			},
		},
	}
}
