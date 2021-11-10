package models

type Request struct {
	Phone string `json:"phone"`
	Sms   string `json:"sms"`
	Login string `json:"login"`
	Hash  string `json:"hash"`
}
type Banks struct {
	id          int    `json:"id"`
	BankName    string `json:"bank_name"`
	ShortName   string `json:"-"`
	Icon        string `json:"icon"`
	IsChange    bool   `json:"is_change"`
	AndroidLink string `json:"android_link"`
	IOSLink     string `json:"ios_link"`
	Transfer    string `json:"-"`
}
type Bank struct {
	BankName    string `json:"bank_name"`
	ShortName   string
	Colors      Colors       `json:"colors"`
	Icon        string       `json:"icon"`
	IsChange    bool         `json:"is_change"`
	AndroidLink string       `json:"android_link"`
	IOSLink     string       `json:"ios_link"`
	Currencies  []Currencies `json:"Currency"`
}

type Colors struct {
	id       int    `json:"id"`
	Color1   string `json:"color_1"`
	Color2   string `json:"color_2"`
	BankName string `json:"bank_name,omitempty"`
	Transfer string `json:"transfer,omitempty"`
}

type Currencies struct {
	id       int    `json:"id"`
	Name     string `json:"name,omitempty"`
	FullName string `json:"full_name,omitempty"`
	Flag     string `json:"flag,omitempty"`
	Sell     string `json:"sell_value,omitempty"`
	Buy      string `json:"buy_value,omitempty"`
	BankName string `json:"-"`
	Transfer string `json:"-"`
}
type Val struct {
	Valute []Valute `xml:"Valute"`
}

type Valute struct {
	ID       string  `xml:"ID,attr" json:"id"`
	CharCode string  `xml:"CharCode" json:"name"`
	Nominal  int     `xml:"Nominal" json:"nominal"`
	Name     string  `xml:"Name" json:"full_name"`
	Flag     string  `json:"flag"`
	Value    float64 `xml:"Value" json:"value"`
}

type Imon struct {
	Buy struct {
		USD struct {
			Rate string `json:"rate"`
		} `json:"1"`
		RUB struct {
			Rate string `json:"rate"`
		} `json:"2"`
		EUR struct {
			Rate string `json:"rate"`
		} `json:"3"`
	} `json:"IMON"`
	Sell struct {
		USD struct {
			Rate string `json:"rate"`
		} `json:"4"`
		RUB struct {
			Rate string `json:"rate"`
		} `json:"5"`
		EUR struct {
			Rate string `json:"rate"`
		} `json:"6"`
	} `json:"IMON1"`
}

type Rate struct {
	Alif      float64 `json:"rateRub"`
	Bonus     float64 `json:"bonus"`
	Spitamen  string  `json:"spitamen"`
	DC        string  `json:"dc"`
	Tamvil    string  `json:"tamvil"`
	Orienbank string  `json:"CurrencyRate"`
}

type IBT struct {
	Extra struct {
		Params []struct {
			Value string `json:"value"`
		} `json:"params"`
	} `json:"extra"`
}
