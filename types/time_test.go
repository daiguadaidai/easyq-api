package types

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestTime_UnmarshalJSON(t *testing.T) {
	str := `
{
    "t1":"2021-11-05 20:20:21"
}
`

	var obj map[string]Time
	if err := json.Unmarshal([]byte(str), &obj); err != nil {
		t.Fatal(err.Error())
	}

	fmt.Printf("%v\n", obj)
}
