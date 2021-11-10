package controller

import (
	"curr/models"
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func Amonatbank(che chan models.Bank) {
	client := &http.Client{}
	//proxyS, err := url.Parse("http://192.168.0.8:4480")
	//if err != nil {
	//	log.Println("Amonatbank proxy ", err.Error())
	//	che <- models.Bank{}
	//}
	//client.Transport = &http.Transport{Proxy: http.ProxyURL(proxyS)}

	//Amonatbank
	respAmonatbank, err := client.Get("http://amonatbonk.tj/tj/")
	if err != nil {
		log.Println("Amonatbank resp err ", err.Error())
		che <- models.Bank{
			BankName:  "Амонатбанк",
			ShortName: "amonat",

			Icon: "",
			Colors: models.Colors{
				Color1: "#1F6608",
				Color2: "#2C9F07",
			},
			AndroidLink: "https://play.google.com/store/apps/details?id=tj.itservice.amonatbonk",
			IOSLink:     "https://apps.apple.com/us/app/%D0%B0%D0%BC%D0%BE%D0%BD%D0%B0%D1%82-%D0%BC%D0%BE%D0%B1%D0%B0%D0%B9%D0%BB/id1513528648",
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
	doc, err := goquery.NewDocumentFromReader(respAmonatbank.Body)
	if err != nil {
		log.Println("Doc Amonatbank err ", err)
		che <- models.Bank{
			BankName:  "Амонатбанк",
			ShortName: "amonat",

			Icon: "",
			Colors: models.Colors{
				Color1: "#1F6608",
				Color2: "#2C9F07",
			},
			AndroidLink: "https://play.google.com/store/apps/details?id=tj.itservice.amonatbonk",
			IOSLink:     "https://apps.apple.com/us/app/%D0%B0%D0%BC%D0%BE%D0%BD%D0%B0%D1%82-%D0%BC%D0%BE%D0%B1%D0%B0%D0%B9%D0%BB/id1513528648",
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
	temp := []string{}
	// Find the review items
	doc.Find("#tab2").Each(func(i int, s *goquery.Selection) {
		s.Find(".coll3").Each(func(i int, t *goquery.Selection) {
			for _, arr := range strings.Split(t.Text(), "|") {
				temp = append(temp, strings.TrimSpace(arr))
			}
		})
	})
	if len(temp) == 0 {
		che <- models.Bank{
			BankName:  "Амонатбанк",
			ShortName: "amonat",

			Icon: "",
			Colors: models.Colors{
				Color1: "#1F6608",
				Color2: "#2C9F07",
			},
			AndroidLink: "https://play.google.com/store/apps/details?id=tj.itservice.amonatbonk",
			IOSLink:     "https://apps.apple.com/us/app/%D0%B0%D0%BC%D0%BE%D0%BD%D0%B0%D1%82-%D0%BC%D0%BE%D0%B1%D0%B0%D0%B9%D0%BB/id1513528648",
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
	var Amonatbank []string
	for _, arr := range temp {
		_, err := strconv.ParseFloat(strings.ReplaceAll(arr, ",", "."), 32)
		if err == nil {
			Amonatbank = append(Amonatbank, strings.ReplaceAll(arr, ",", "."))
		}
	}

	che <- models.Bank{
		BankName:  "Амонатбанк",
		ShortName: "amonat",

		Icon: "",
		Colors: models.Colors{
			Color1: "#1F6608",
			Color2: "#2C9F07",
		},
		AndroidLink: "https://play.google.com/store/apps/details?id=tj.itservice.amonatbonk",
		IOSLink:     "https://apps.apple.com/us/app/%D0%B0%D0%BC%D0%BE%D0%BD%D0%B0%D1%82-%D0%BC%D0%BE%D0%B1%D0%B0%D0%B9%D0%BB/id1513528648",
		Currencies: []models.Currencies{
			models.Currencies{
				Name: "USD",
				Buy:  Amonatbank[0],
				Sell: Amonatbank[1],
			},
			models.Currencies{
				Name: "RUB",
				Buy:  Amonatbank[4],
				Sell: Amonatbank[5],
			},
			models.Currencies{
				Name: "EUR",
				Buy:  Amonatbank[2],
				Sell: Amonatbank[3],
			},
		},
	}
	return
}
