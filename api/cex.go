package api

import (
	"fmt"
	"io"
	"net/http"
	"strings"

	"example.com/go/crypto/datatypes"
)

const url = "https://cex.io/api/ticker/%s/USD"

func main() {
	GetRate("usd")
}
func GetRate(currency string) (*datatypes.Rate, error) {
	cur := strings.ToUpper(currency)
	res, err := http.Get(fmt.Sprintf(url, cur))
	if err != nil {
		return nil, err
	}
	if res.StatusCode == http.StatusOK {
		bodyBytes, err := io.ReadAll(res.Body)
		if err != nil {
			return nil, err
		}
		json := string(bodyBytes)
		fmt.Println(json)
	} else {
		return nil, fmt.Errorf("status code received: %v", res.StatusCode)
	}
		rate := datatypes.Rate{Currency: currency, Price: 20}
	return &rate, err
}
