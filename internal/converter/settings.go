package converter

import (
	desc "github.com/nastya-zz/fisher-protocols/gen/user_v1"
	"user/internal/model"
)

func GetDescAvailability(str string) desc.Availability {
	switch str {
	case model.Private:
		return desc.Availability_PRIVATE
	case model.Public:
		return desc.Availability_PUBLIC
	default:
		return desc.Availability_PUBLIC
	}
}

func GetDescLanguage(str string) desc.Language {
	switch str {
	case model.LangRu:
		return desc.Language_RU
	case model.LangEn:
		return desc.Language_ENG
	default:
		return desc.Language_RU
	}
}

func GetModelAvailability(str desc.Availability) string {
	switch str {
	case desc.Availability_PRIVATE:
		return model.Private
	case desc.Availability_PUBLIC:
		return model.Public
	default:
		return model.Public
	}
}

func GetModelLanguage(str desc.Language) string {
	switch str {
	case desc.Language_RU:
		return model.LangRu
	case desc.Language_ENG:
		return model.LangEn
	default:
		return model.LangRu
	}
}
