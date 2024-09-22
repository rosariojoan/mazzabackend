package utils

import (
	"encoding/json"
	"fmt"
)

/* Pretty print a struct or json */
func PP(v any) {
	b, err := json.MarshalIndent(v, "", " ")
	if err == nil {
		fmt.Println(string(b))
	}
}
