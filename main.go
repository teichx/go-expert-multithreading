package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
	"unicode"
)

type Result struct {
	Origin string
	Data   []byte
}

func getCEP(url, origin string, ch chan<- Result) {
	res, err := http.Get(url)
	if err != nil {
		return
	}
	defer res.Body.Close()
	data, err := io.ReadAll(res.Body)
	if err != nil {
		return
	}
	if res.StatusCode != http.StatusOK {
		return
	}
	ch <- Result{Origin: origin, Data: data}
}

func extractDigits(s string) string {
	var digits []byte
	for _, r := range s {
		if unicode.IsDigit(r) {
			digits = append(digits, byte(r))
		}
	}
	return string(digits)
}

func main() {
	if len(os.Args) < 2 {
		panic("Error: No CEP provided")
	}
	cep := extractDigits(os.Args[1])
	if len(cep) != 8 {
		panic("Error: CEP must have exactly 8 digits")
	}
	ch := make(chan Result)

	go getCEP("https://brasilapi.com.br/api/cep/v1/"+cep, "BrasilAPI", ch)
	go getCEP("http://viacep.com.br/ws/"+cep+"/json/", "ViaCEP", ch)

	select {
	case res := <-ch:
		fmt.Printf("{\"api\": \"%s\", \"data\": %s}\n", res.Origin, string(res.Data))
	case <-time.After(1 * time.Second):
		fmt.Println("{\"error\": \"timeout - no API responded within 1 second\"}")
	}
}
