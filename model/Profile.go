package model

type BaseProfile struct {
	DNI         DNI
	PicturePath string
	Name        Name
	Gender      Gender
	Address     string
	Age         int
	Email       string
	Username    string
	Password    string
}
