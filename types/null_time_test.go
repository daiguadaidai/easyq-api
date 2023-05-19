package types

import (
	"encoding/json"
	"fmt"
	"testing"
)

func Test_NullTime_UnmarshalJSON(t *testing.T) {
	str := `
{
    "t1":"2021-11-05 20:20:21"
}
`

	var obj map[string]NullTime
	if err := json.Unmarshal([]byte(str), &obj); err != nil {
		t.Fatal(err.Error())
	}

	fmt.Printf("%v\n", obj)
}
