package bot_libs

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func Output(s string) {
	fmt.Println(s)
}

// takes any .<cmd> <input> string and removes the .cmd
func stripCmd(s string) (string, error) {
	split := strings.Fields(s)
	if len(split) < 2 {
		return "", errors.New("Not enough arguments")
	}
	joined := strings.Join(split[1:], " ")
	return joined, nil
}

func Stock(s string) (string, error) {
	ticker, err := stripCmd(s)
	if err != nil {
		return "", err
	}
	url := fmt.Sprintf("http://download.finance.yahoo.com/d/quotes.csv?s=%s&f=sl1c1p2", ticker)
	response, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", err
	}
	output := strings.Replace(string(body), ",", " ", -1)
	output = "Symbol Last Change Percent\n" + output
	return output, nil
}
