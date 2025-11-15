package utils

import (
	"encoding/json"
	"io"
	"os"
	"strings"
)

// Load account names from the "app/assets/account_names.json" file
func LoadAccountNames(country string, lang string) (accountNames map[string]string, err error) {
	jsonFile, err := os.Open("./app/assets/account_names.json")
	if err != nil {
		return
	}
	defer jsonFile.Close()

	var _accountNames map[string]map[string]map[string]string
	byteValue, _ := io.ReadAll(jsonFile)

	if json.Unmarshal(byteValue, &_accountNames) != nil {
		return
	}

	countryCode := strings.ToUpper(country)
	language := strings.ToUpper(lang)
	accountNames = _accountNames[countryCode][language]
	return accountNames, nil
}
