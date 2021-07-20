package bicycleCrowdEvaluator

import (
	"encoding/json"
	"io/ioutil"
)

func Decode(filename string) map[string]interface{} {
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	var responses map[string]interface{}

	if err := json.Unmarshal(body, &responses); err != nil {
		panic(err)
	}

	//fmt.Println(responses)

	return responses
}
