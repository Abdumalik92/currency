package jobs

import (
	"curr/controller"
	"curr/db"
	"curr/models"
	"fmt"
	"log"
	"strconv"
)

func C2CBank() {
	var (
		che   = []models.Bank{}
		str   models.Bank
		color models.Colors
		bank  models.Banks
		c1    models.Banks
		c2    models.Colors
		c3    models.Currencies
	)
	//
	che = append(che, controller.IBTTranser())
	che = append(che, controller.TamvilTransfer())
	che = append(che, controller.AlifTransfer())
	che = append(che, controller.SpitamenTransfer())
	che = append(che, controller.CBTTransfer())
	che = append(che, controller.DCTransfer())
	che = append(che, controller.OrienTransfer())

	for i := 0; i < len(che); i++ {
		str = che[i]
		b1, _ := strconv.ParseFloat(str.Currencies[0].Buy, 32)
		s1, _ := strconv.ParseFloat(str.Currencies[0].Sell, 32)

		str.Currencies[0].Buy = fmt.Sprintf("%.4f", b1)
		str.Currencies[0].Sell = fmt.Sprintf("%.4f", s1)

		str.Currencies[0].Flag = "https://transfer.humo.tj/currency-app/v1/get_image/flags-" + str.Currencies[0].Name + ".png"
		str.Currencies[0].FullName = "Российский рубль"
		str.Icon = "https://transfer.humo.tj/currency-app/v1/get_image/logo-" + str.ShortName + ".png"

		if str.ShortName == "imon" || str.ShortName == "agroinvest" || str.ShortName == "pakbankintj" {
			str.IsChange = true
		} else {
			str.IsChange = false
		}
		str.Currencies = append(str.Currencies, models.Currencies{
			Name:     "TJS",
			FullName: "Таджикский сомони",
			Flag:     "https://transfer.humo.tj/currency-app/v1/get_image/flags-TJS.png",
			Sell:     "1.0000",
			Buy:      "1.0000",
		})
		str.Colors.BankName = str.ShortName

		str.Colors.Transfer = "trans"
		color = str.Colors
		bank = models.Banks{
			BankName:    str.BankName,
			ShortName:   str.ShortName,
			Icon:        str.Icon,
			IsChange:    str.IsChange,
			IOSLink:     str.IOSLink,
			AndroidLink: str.AndroidLink,
			Transfer:    "trans",
		}
		if err := db.GetDBConn().Where("short_name=? and transfer='trans'", bank.ShortName).Find(&c1).Error; err != nil {
			log.Println("select bank err", err)
		}
		if c1 == (models.Banks{}) {
			if err := db.GetDBConn().Create(&bank).Error; err != nil {
				log.Println("Create err", err)
			}
		}
		////////////////////////////////////////////////////////////////////////////////////////
		//Currency
		////////////////////////////////////////////////////////////////////////////////////////
		for _, arr := range str.Currencies {
			arr.BankName = str.ShortName
			arr.Transfer = "trans"

			if err := db.GetDBConn().Where("bank_name=? and transfer='trans' and name = ?", str.ShortName, arr.Name).Find(&c3).Error; err != nil {
				log.Println("select curr err", err)
			}
			if c3 == (models.Currencies{}) {
				if err := db.GetDBConn().Create(&arr).Error; err != nil {
					log.Println("Create  curr err", err)
				}
			} else {
				if err := db.GetDBConn().Model(&c3).Where("bank_name=? and transfer='trans' and name = ?", str.ShortName, arr.Name).Updates(models.Currencies{Buy: arr.Buy, Sell: arr.Sell}).Error; err != nil {
					log.Println("select curr err", err)
				}
			}

		}
		/////////////////////////////////////////////////////////////////////////////////////
		//Color
		////////////////////////////////////////////////////////////////////////////////////

		if err := db.GetDBConn().Where("bank_name=? and transfer='trans'", str.ShortName).Find(&c2).Error; err != nil {
			log.Println("select bank err", err)
		}
		if c2 == (models.Colors{}) {
			if err := db.GetDBConn().Create(&color).Error; err != nil {
				log.Println("Create err", err)
			}
		}
	}
	return
}

func NPCRBank() {
	var (
		che   = make(chan models.Bank)
		c1    models.Banks
		c2    models.Colors
		c3    models.Currencies
		str   models.Bank
		color models.Colors
		bank  models.Banks
	)

	go controller.Alif(che)
	//go controller.Agroinvest(che)
	go controller.Amonatbank(che)
	go controller.Arvand(che)
	go controller.BankAsia(che)
	go controller.KAFOLAT(che)
	go controller.Imon(che)
	go controller.HALIKBANK(che)
	go controller.FazoC(che)
	go controller.NBP(che)
	go controller.Eskhata(che)
	go controller.Finca(che)
	go controller.FMBT(che)
	go controller.MDOHumo(che)
	go controller.CBT(che)
	go controller.IBT(che)
	go controller.Spitamen(che)

	for i := 0; i < 16; i++ {
		str = <-che
		if str.ShortName == "imon" || str.ShortName == "agroinvest" || str.ShortName == "pakbankintj" {
			str.IsChange = true
		} else {
			str.IsChange = false
		}

		b1, _ := strconv.ParseFloat(str.Currencies[0].Buy, 32)
		b2, _ := strconv.ParseFloat(str.Currencies[1].Buy, 32)
		b3, _ := strconv.ParseFloat(str.Currencies[2].Buy, 32)
		s1, _ := strconv.ParseFloat(str.Currencies[0].Sell, 32)
		s2, _ := strconv.ParseFloat(str.Currencies[1].Sell, 32)
		s3, _ := strconv.ParseFloat(str.Currencies[2].Sell, 32)

		str.Currencies[0].Flag = "https://transfer.humo.tj/currency-app/v1/get_image/flags-" + str.Currencies[0].Name + ".png"
		str.Currencies[1].Flag = "https://transfer.humo.tj/currency-app/v1/get_image/flags-" + str.Currencies[1].Name + ".png"
		str.Currencies[2].Flag = "https://transfer.humo.tj/currency-app/v1/get_image/flags-" + str.Currencies[2].Name + ".png"
		str.Currencies[0].FullName = "Доллар США"
		str.Currencies[1].FullName = "Российский рубль"
		str.Currencies[2].FullName = "ЕВРО"
		str.Currencies[0].Buy = fmt.Sprintf("%.4f", b1)
		str.Currencies[1].Buy = fmt.Sprintf("%.4f", b2)
		str.Currencies[2].Buy = fmt.Sprintf("%.4f", b3)
		str.Currencies[0].Sell = fmt.Sprintf("%.4f", s1)
		str.Currencies[1].Sell = fmt.Sprintf("%.4f", s2)
		str.Currencies[2].Sell = fmt.Sprintf("%.4f", s3)
		str.Currencies[1].FullName = "Российский рубль"
		str.Currencies[2].FullName = "ЕВРО"
		str.Icon = "https://transfer.humo.tj/currency-app/v1/get_image/logo-" + str.ShortName + ".png"

		str.Currencies = append(str.Currencies, models.Currencies{
			Name:     "TJS",
			FullName: "Таджикский сомони",
			Flag:     "https://transfer.humo.tj/currency-app/v1/get_image/flags-TJS.png",
			Sell:     "1.0000",
			Buy:      "1.0000",
		})
		str.Colors.BankName = str.ShortName

		color = str.Colors
		bank = models.Banks{
			BankName:    str.BankName,
			ShortName:   str.ShortName,
			Icon:        str.Icon,
			IsChange:    str.IsChange,
			IOSLink:     str.IOSLink,
			AndroidLink: str.AndroidLink,
			Transfer:    "",
		}
		if err := db.GetDBConn().Where("short_name=? and transfer=''", bank.ShortName).Find(&c1).Error; err != nil {
			log.Println("select bank err", err)
		}
		if c1 == (models.Banks{}) {
			if err := db.GetDBConn().Create(&bank).Error; err != nil {
				log.Println("Create err", err)
			}
		}

		for _, arr := range str.Currencies {
			arr.BankName = str.ShortName
			arr.Transfer = ""

			if err := db.GetDBConn().Where("bank_name=? and transfer='' and name = ?", str.ShortName, arr.Name).Find(&c3).Error; err != nil {
				log.Println("select curr err", err)
			}
			if c3 == (models.Currencies{}) {
				if err := db.GetDBConn().Create(&arr).Error; err != nil {
					log.Println("Create  curr err", err)
				}
			} else {
				if err := db.GetDBConn().Model(&c3).Where("bank_name=? and transfer='' and name = ?", str.ShortName, arr.Name).Updates(models.Currencies{Buy: arr.Buy, Sell: arr.Sell}).Error; err != nil {
					log.Println("select curr err", err)
				}
			}

		}
		if err := db.GetDBConn().Where("bank_name=? and transfer=''", str.ShortName).Find(&c2).Error; err != nil {
			log.Println("select bank err", err)
		}
		if c2 == (models.Colors{}) {
			if err := db.GetDBConn().Create(&color).Error; err != nil {
				log.Println("Create err", err)
			}
		}

	}

	return
}
