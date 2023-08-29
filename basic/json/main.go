package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	{
		var b []byte
		var res map[string]interface{}
		fmt.Println(json.Unmarshal(b, &res)) //unexpected end of JSON input
	}

	{
		b := []byte("{}")
		var res map[string]interface{}
		fmt.Println(json.Unmarshal(b, &res)) //<nil>
	}

}
