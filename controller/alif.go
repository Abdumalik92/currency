package controller

import (
	"curr/models"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func Alif(che chan models.Bank) {
	client := &http.Client{}
	//proxyS, err := url.Parse("http://192.168.0.8:4480")
	//if err != nil {
	//	log.Println("Alif proxy ", err.Error())
	//	che <- models.Bank{
	//		BankName:    "Алифбанк",
	//		ShortName:"alif",
	//
	//		Icon: "",
	//		Colors: models.Colors{
	//			Color1: "#32A371",
	//			Color2: "#39B980",
	//		},
	//		AndroidLink: "https://play.google.com/store/apps/details?id=tj.alif.mobi",
	//		IOSLink:     "https://apps.apple.com/us/app/alif-mobi/id1331374853?l=ru&ls=1",
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

	ti := time.Now().Format("2006-01-02")

	//Alif
	respAlif, err := client.Get("https://alif.tj/api/currency/index.php?currency=rub&date=" + ti)
	if err != nil {
		log.Println("Alif resp err ", err.Error())
		che <- models.Bank{
			BankName:  "Алифбанк",
			ShortName: "alif",

			Icon: "",
			Colors: models.Colors{
				Color1: "#32A371",
				Color2: "#39B980",
			},
			AndroidLink: "https://play.google.com/store/apps/details?id=tj.alif.mobi",
			IOSLink:     "https://apps.apple.com/us/app/alif-mobi/id1331374853?l=ru&ls=1",
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
	AlifRub := models.Currencies{}

	bodyAlif, err := ioutil.ReadAll(respAlif.Body)
	if err != nil {
		log.Println("Alif resp body ", err.Error())
		che <- models.Bank{
			BankName:  "Алифбанк",
			ShortName: "alif",

			Icon: "",
			Colors: models.Colors{
				Color1: "#32A371",
				Color2: "#39B980",
			},
			AndroidLink: "https://play.google.com/store/apps/details?id=tj.alif.mobi",
			IOSLink:     "https://apps.apple.com/us/app/alif-mobi/id1331374853?l=ru&ls=1",
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
	err = json.Unmarshal(bodyAlif, &AlifRub)
	if err != nil {
		log.Println("func AlifRUB can't unmarshal json ", err.Error())
		che <- models.Bank{
			BankName:  "Алифбанк",
			ShortName: "alif",

			Icon: "",
			Colors: models.Colors{
				Color1: "#32A371",
				Color2: "#39B980",
			},
			AndroidLink: "https://play.google.com/store/apps/details?id=tj.alif.mobi",
			IOSLink:     "https://apps.apple.com/us/app/alif-mobi/id1331374853?l=ru&ls=1",
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
	respAlif, err = client.Get("https://alif.tj/api/currency/index.php?currency=usd&date=" + ti)
	if err != nil {
		log.Println("Alif resp err ", err.Error())
		che <- models.Bank{
			BankName: "Алифбанк",

			Icon: "",
			Colors: models.Colors{
				Color1: "",
				Color2: "",
			},
			AndroidLink: "",
			IOSLink:     "",
			Currencies: []models.Currencies{
				models.Currencies{Name: "USD",
					Buy:  "0.0000",
					Sell: "0.0000"},
				models.Currencies{Name: "RUB",
					Buy:  "0.0000",
					Sell: "0.0000"},
				models.Currencies{Name: "EUR",
					Buy:  "0.0000",
					Sell: "0.0000"},
			},
		}
		return
	}
	AlifUsd := models.Currencies{}
	bodyAlif, err = ioutil.ReadAll(respAlif.Body)
	if err != nil {
		log.Println("AlifUSD resp body ", err.Error())
		che <- models.Bank{
			BankName:  "Алифбанк",
			ShortName: "alif",

			Icon: "",
			Colors: models.Colors{
				Color1: "#32A371",
				Color2: "#39B980",
			},
			AndroidLink: "https://play.google.com/store/apps/details?id=tj.alif.mobi",
			IOSLink:     "https://apps.apple.com/us/app/alif-mobi/id1331374853?l=ru&ls=1",
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
	err = json.Unmarshal(bodyAlif, &AlifUsd)
	if err != nil {
		log.Println("func AlifUSD can't unmarshal json ", err.Error())
		che <- models.Bank{}
		return
	}
	respAlif, err = client.Get("https://alif.tj/api/currency/index.php?currency=eur&date=" + ti)
	if err != nil {
		log.Println("Alif resp err ", err.Error())
		che <- models.Bank{
			BankName:  "Алифбанк",
			ShortName: "alif",

			Icon: "",
			Colors: models.Colors{
				Color1: "#32A371",
				Color2: "#39B980",
			},
			AndroidLink: "https://play.google.com/store/apps/details?id=tj.alif.mobi",
			IOSLink:     "https://apps.apple.com/us/app/alif-mobi/id1331374853?l=ru&ls=1",
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
	AlifEur := models.Currencies{}
	bodyAlif, err = ioutil.ReadAll(respAlif.Body)
	if err != nil {
		log.Println("Alif resp body ", err.Error())
		che <- models.Bank{
			BankName:  "Алифбанк",
			ShortName: "alif",

			Icon: "",
			Colors: models.Colors{
				Color1: "#32A371",
				Color2: "#39B980",
			},
			AndroidLink: "https://play.google.com/store/apps/details?id=tj.alif.mobi",
			IOSLink:     "https://apps.apple.com/us/app/alif-mobi/id1331374853?l=ru&ls=1",
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
	err = json.Unmarshal(bodyAlif, &AlifEur)
	if err != nil {
		log.Println("func AlifEUR can't unmarshal json ", err.Error())
		che <- models.Bank{
			BankName:  "Алифбанк",
			ShortName: "alif",

			Icon: "",
			Colors: models.Colors{
				Color1: "#32A371",
				Color2: "#39B980",
			},
			AndroidLink: "https://play.google.com/store/apps/details?id=tj.alif.mobi",
			IOSLink:     "https://apps.apple.com/us/app/alif-mobi/id1331374853?l=ru&ls=1",
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
		BankName:  "Алифбанк",
		ShortName: "alif",

		Icon: "",
		Colors: models.Colors{
			Color1: "#32A371",
			Color2: "#39B980",
		},
		AndroidLink: "https://play.google.com/store/apps/details?id=tj.alif.mobi",
		IOSLink:     "https://apps.apple.com/us/app/alif-mobi/id1331374853?l=ru&ls=1",
		Currencies: []models.Currencies{
			models.Currencies{
				Name: AlifUsd.Name,
				Buy:  AlifUsd.Buy,
				Sell: AlifUsd.Sell,
			},
			models.Currencies{
				Name: AlifRub.Name,
				Buy:  AlifRub.Buy,
				Sell: AlifRub.Sell,
			},
			models.Currencies{
				Name: AlifEur.Name,
				Buy:  AlifEur.Buy,
				Sell: AlifEur.Sell,
			},
		},
	}
	return
}

func AlifTransfer() models.Bank {
	client := &http.Client{}
	//proxyS, err := url.Parse("http://192.168.0.8:4480")
	//if err != nil {
	//	log.Println("AlifTransfer proxy ", err.Error())
	//	return models.Bank{
	//		BankName:    "Алифбанк",
	//		ShortName:"alif",
	//
	//		Icon: "",
	//		Colors: models.Colors{
	//			Color1: "#32A371",
	//			Color2: "#39B980",
	//		},
	//		AndroidLink: "https://intiqol.tj/",
	//		IOSLink:     "https://intiqol.tj/",
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
	//client.Transport = &http.Transport{Proxy: http.ProxyURL(proxyS)}

	//Alif
	respAlif, err := client.Get("https://intiqol.tj/intiqol/api/v1/get/condition")
	if err != nil {
		log.Println("AlifTransfer resp err ", err.Error())
		return models.Bank{
			BankName:  "Алифбанк",
			ShortName: "alif",

			Icon: "",
			Colors: models.Colors{
				Color1: "#32A371",
				Color2: "#39B980",
			},
			AndroidLink: "https://intiqol.tj/",
			IOSLink:     "https://intiqol.tj/",
			Currencies: []models.Currencies{
				models.Currencies{
					Name: "RUB",
					Buy:  "0.0000",
					Sell: "0.0000",
				},
			},
		}
	}
	AlifRub := models.Rate{}

	bodyAlif, err := ioutil.ReadAll(respAlif.Body)
	if err != nil {
		log.Println("AlifTransfer resp body ", err.Error())
		return models.Bank{
			BankName:  "Алифбанк",
			ShortName: "alif",

			Icon: "",
			Colors: models.Colors{
				Color1: "#32A371",
				Color2: "#39B980",
			},
			AndroidLink: "https://intiqol.tj/",
			IOSLink:     "https://intiqol.tj/",
			Currencies: []models.Currencies{
				models.Currencies{
					Name: "RUB",
					Buy:  "0.0000",
					Sell: "0.0000",
				},
			},
		}
	}

	err = json.Unmarshal(bodyAlif, &AlifRub)
	if err != nil {
		log.Println("func AlifTransfer can't unmarshal json ", err.Error())
		return models.Bank{
			BankName:  "Алифбанк",
			ShortName: "alif",

			Icon: "",
			Colors: models.Colors{
				Color1: "#32A371",
				Color2: "#39B980",
			},
			AndroidLink: "https://intiqol.tj/",
			IOSLink:     "https://intiqol.tj/",
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
		BankName:  "Алифбанк",
		ShortName: "alif",

		Icon: "",
		Colors: models.Colors{
			Color1: "#32A371",
			Color2: "#39B980",
		},
		AndroidLink: "https://intiqol.tj/",
		IOSLink:     "https://intiqol.tj/",
		Currencies: []models.Currencies{
			models.Currencies{
				Name: "RUB",
				Buy:  fmt.Sprint(AlifRub.Alif),
				Sell: "0.0000",
			},
		},
	}

}
