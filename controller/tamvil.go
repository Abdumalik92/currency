package controller

import (
	"curr/models"
	"io/ioutil"
	"log"
	"net/http"
)

func TamvilTransfer() models.Bank {
	client := &http.Client{}
	//proxyS, err := url.Parse("http://192.168.0.8:4480")
	//if err != nil {
	//	log.Println("Tamvil proxy ", err.Error())
	//	return models.Bank{
	//		BankName:    "Тамвил",
	//		ShortName:   "tamvil",
	//
	//		Icon: "",
	//		Colors: models.Colors{
	//			Color1: "#292E36",
	//			Color2: "#535F75",
	//		},
	//		AndroidLink: "https://oddi.tamvil.tj",
	//		IOSLink:     "https://oddi.tamvil.tj",
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
	resTamvil, err := client.Get("https://oddi.tamvil.tj/ajax.asp?fn=get_exch_rate&ccy=RUB")
	if err != nil {
		log.Println("Tamvil resp err ", err.Error())
		return models.Bank{
			BankName:  "Тамвил",
			ShortName: "tamvil",

			Icon: "",
			Colors: models.Colors{
				Color1: "#292E36",
				Color2: "#535F75",
			},
			AndroidLink: "https://oddi.tamvil.tj",
			IOSLink:     "https://oddi.tamvil.tj",
			Currencies: []models.Currencies{
				models.Currencies{
					Name: "RUB",
					Buy:  "0.0000",
					Sell: "0.0000",
				},
			},
		}
	}

	bodyTamvil, err := ioutil.ReadAll(resTamvil.Body)
	if err != nil {
		log.Println("Body Tamvil err ", err.Error())
		return models.Bank{
			BankName:  "Тамвил",
			ShortName: "tamvil",

			Icon: "",
			Colors: models.Colors{
				Color1: "#292E36",
				Color2: "#535F75",
			},
			AndroidLink: "https://oddi.tamvil.tj",
			IOSLink:     "https://oddi.tamvil.tj",
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
		BankName:  "Тамвил",
		ShortName: "tamvil",

		Icon: "",
		Colors: models.Colors{
			Color1: "#292E36",
			Color2: "#535F75",
		},
		AndroidLink: "https://oddi.tamvil.tj",
		IOSLink:     "https://oddi.tamvil.tj",
		Currencies: []models.Currencies{
			models.Currencies{
				Name: "RUB",
				Buy:  string(bodyTamvil),
				Sell: "0.0000",
			},
		},
	}
}
