package util

import (
	"fmt"
	"math/rand"
	"strings"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

func RandomFloat(min, max float64) float64 {
	return min + (rand.Float64() * (max - min))
}

func RadomString(n int) string {
	var sb strings.Builder

	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}
	return sb.String()

}

func RandomOwner() string {
	return RadomString(6)
}

func RandomMoney() float64 {
	return RandomFloat(0, 1000)
}

func RandomCurrency() string {
	currencies := []string{EUR, CAD, USD}
	n := len(currencies)
	return currencies[rand.Intn(n)]
}

func RandomEmail() string {
	return fmt.Sprintf("%s@email.com", RadomString(6))
}
