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

	return responses
}

func NumAnnotators(responses map[string]interface{}) int {

	var m map[string]int = make(map[string]int)
	for _, v1 := range responses {
		for _, v2 := range v1.(map[string]interface{}) {
			for k3, v3 := range v2.(map[string]interface{}) {
				if k3 == "results" {
					for _, v4 := range v3.(map[string]interface{}) {
						for k5, v5 := range v4.(map[string]interface{}) {
							if k5 == "results" {
								for _, v6 := range v5.([]interface{}) {
									for k7, v7 := range v6.(map[string]interface{}) {
										if k7 == "user" {
											for k8, v8 := range v7.(map[string]interface{}) {
												if k8 == "vendor_user_id" {
													m[v8.(string)] += 1
												}
											}
										}
									}
								}
							}
						}
					}
				}
			}
		}
	}
	return len(m)
}
