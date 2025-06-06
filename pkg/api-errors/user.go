package api_errors

const (
	UserIdRequired = "Не указан id пользователя"
	UserIdNotFound = "Пользователь с таким Id не найден"

	UserEmailEmpty           = "Не указан email пользователя"
	UserEmailNotMatchPattern = "Некорректный email"

	UserUpdateFailed = "Не удалось обновить пользователя"
)
