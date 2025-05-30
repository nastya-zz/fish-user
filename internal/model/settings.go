package model

const (
	LangRu = "RU"
	LangEn = "ENG"
)

const (
	Public  = "PUBLIC"
	Private = "PRIVATE"
)

type SettingsId string

type Settings struct {
	Language     string
	Availability string
}

func NewDefaultSettings() Settings {
	return Settings{
		Language:     LangRu,
		Availability: Public,
	}
}
