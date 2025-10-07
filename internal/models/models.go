package models

type User struct {
	Username       string `json:"username"`
	Password       string `json:"password"`
	CountryCode    string `json:"countryCode"`
	PhoneNumber    string `json:"phoneNumber"`
	Email          string `json:"email"`
	FullName       string `json:"fullName"`
	Role           string `json:"role"`
	OrganizationId string `json:"organizationId"`
	Status         int32  `json:"status"`
	IsSSO          bool   `json:"isSSO"`
}

type UserEnvironment struct {
}

type KeyBundle struct {
	PrivateKey     string `json:"privateKey"`
	PublicKey      string `json:"publicKey"`
	EncrPublicKey  string `json:"encrPublicKey"`
	EncrPrivateKey string `json:"encrPrivateKey"`
}

type UnicastMessage struct {
	ClientID string
	Message  []byte
}

type IncomingMessage struct {
	ClientID string
	Message  []byte
}
