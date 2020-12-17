package model

import (
	"fmt"
	"math"
	"strconv"
)

// RUT : Body-Token
//	e.g: 12345678-9
type DNI struct {
	Body string		// minLength:7 - maxLenght:8
	Token string	// 0 to 9 or 'K'
}

func NewDNI(body string, token string) *DNI {
	dni := &DNI{
		Body: body,
		Token: token,
	}
	return dni
}

func (dni *DNI) Validate() bool {
	multiplySeries := [6]int{2, 3, 4, 5, 6, 7}
	if !dni.checkLength() {
		return false
	}

	// Let's get algorithmically the validation number (token).
	// according to the body provided
	sum := 0
	msi := 0
	for _, c := range dni.Body {
		number := int(c - '0')
		multiplier := multiplySeries[msi]
		delta := number * multiplier
		sum += delta

		msi = msi + 1
		if msi == 6 {
			msi = 0
		}

		fmt.Printf("Number (%d) * Multiplier[%d] = {%d}\n", number, multiplier, delta)
	}
	fmt.Printf("SUM OF BODY MULTIPLIED IS: %d\n", sum)

	// Apply Modulus to the number divided by (11)
	mod := math.Mod(float64(sum), float64(11))
	dvNumber := 11 - int(mod) + 1
	var dv string
	if dvNumber == 10 {
		dv = "K"
	} else {
		dv = strconv.Itoa(dvNumber)
	}
	fmt.Printf("MOD is: %.0f\n", mod)
	fmt.Printf("DV is: %s\n", dv)
	if dv == dni.Token {
		return true
	} else {
		return false
	}
}

func (dni *DNI) checkLength() bool {
	// Body must have 7 or 8 number extension
	if len(dni.Body) == 7 || len(dni.Body) == 8 {
		return true
	} else {
		return false
	}
}

