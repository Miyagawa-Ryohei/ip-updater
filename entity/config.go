package entity


type IPUpdaterConfig struct {
	Auth *AuthOption
}

type AuthOption struct {
	Username string
	Password string
}
