package main

import (
	"encoding/json"
	"fmt"

	"github.com/pkg/profile"
)

func main() {
	defer profile.Start(profile.ProfilePath(".")).Stop()

	var jsonBlob = []byte(`[
			{"Name": "Platypus", "Order": "Monotremata"},
			{"Name": "Quoll",    "Order": "Dasyuromorphia"}
		]`)
	type Animal struct {
		Name  string
		Order string
	}

	for i := 0; i < 1000000; i++ {
		var animals []Animal
		err := json.Unmarshal(jsonBlob, &animals)
		if err != nil {
			fmt.Println("error:", err)
		}
		fmt.Printf("%+v", animals)
		fmt.Printf("%d", i)
	}
}
