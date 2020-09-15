package utils

import (
	"encoding/json"
	"fmt"
	"log"
)

func PrettyPrint(result interface{}) {
	bytes, err := json.MarshalIndent(result, "", " ")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("\n%s\n", string(bytes))
}
