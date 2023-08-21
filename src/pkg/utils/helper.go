package utils

import (
	"encoding/json"
	"fmt"
)

func Debug[T any](obj T) {
	raw, _ := json.MarshalIndent(&obj, "", "\t")
	fmt.Println(string(raw))
}
