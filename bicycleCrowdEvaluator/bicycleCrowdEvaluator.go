package bicycleCrowdEvaluator

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type AnnotatorInfo struct {
	NumAnnotators            int
	AverageAnnotationTimes   int
	MinAnnotationTimes       int
	MaxAnnotationTimes       int
	AnnotatorResults         map[string]int
	HighlyDisagreedQuestions []struct {
		id  string
		url string
	}
}

type Pair struct {
	TaskOutput map[string]interface{}
	ImageUrl   string
}

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

func Annotators(responses map[string]interface{}) AnnotatorInfo {

	var m map[string]int = make(map[string]int)
	var s []int
	var m2 map[string][]Pair = make(map[string][]Pair)
	//Todo: Maybe rename all instances of v to child or children or better yet firstChild like
	// *html.Node.FirstChild or just v. Update: just leave k and v the way they are named here,
	// in order to avoid confusion and ambiquity.
	for _, v1 := range responses {
		for _, v2 := range v1.(map[string]interface{}) {
			for k3, v3 := range v2.(map[string]interface{}) {
				if k3 == "results" {
					for k4, v4 := range v3.(map[string]interface{}) {
						for k5, v5 := range v4.(map[string]interface{}) {
							if k5 == "results" {
								for _, v6 := range v5.([]interface{}) {
									var tmp string
									for k7, v7 := range v6.(map[string]interface{}) {
										if k7 == "task_input" {
											//Todo: rename kNew and VNew
											for kNew, VNew := range v7.(map[string]interface{}) {

												if kNew == "image_url" {
													tmp = VNew.(string)
												}
											}
										}

										if k7 == "user" {
											for k8, v8 := range v7.(map[string]interface{}) {
												if k8 == "vendor_user_id" {
													m[v8.(string)] += 1
												}
											}
										}

										if k7 == "task_output" {
											for k8, v8 := range v7.(map[string]interface{}) {
												if k8 == "duration_ms" {
													s = append(s, int(v8.(float64)))
												}
											}

											m2[k4] = append(m2[k4], Pair{v7.(map[string]interface{}), tmp})
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

	var highlyDisagreedQs []struct {
		id  string
		url string
	}
	for k, v := range m2 {
		// Note/Todo: to compare each annotators output I could easily use a nested for loop that loops
		// over m2

		// Note that there are 10 task_output, even though for now I am only using the 0th and 1st
		// task_output
		var url string

		for _, v := range v {
			if v.ImageUrl != "" {
				url = v.ImageUrl
			}
		}

		var one Pair = v[0]
		var two Pair = v[1]

		if one.TaskOutput["answer"].(string) != two.TaskOutput["answer"].(string) {
			highlyDisagreedQs = append(highlyDisagreedQs, struct {
				id  string
				url string
			}{k, url})
		} else if one.TaskOutput["cant_solve"].(bool) != two.TaskOutput["cant_solve"].(bool) {
			highlyDisagreedQs = append(highlyDisagreedQs, struct {
				id  string
				url string
			}{k, url})
		} else {
			t1, ok1 := one.TaskOutput["corrupt_data"].(bool)
			t2, ok2 := two.TaskOutput["corrupt_data"].(bool)

			if ok1 && ok2 {
				if t1 != t2 {
					highlyDisagreedQs = append(highlyDisagreedQs, struct {
						id  string
						url string
					}{k, url})
				}
			}

		}
	}

	fmt.Println(highlyDisagreedQs)

	var annotatorInfo AnnotatorInfo = AnnotatorInfo{NumAnnotators: len(m),
		AverageAnnotationTimes: Average(s, len(m)), MinAnnotationTimes: Min(s),
		MaxAnnotationTimes: Max(s), AnnotatorResults: m, HighlyDisagreedQuestions: highlyDisagreedQs}
	return annotatorInfo
}

func Max(s []int) int {
	var max int = s[0]

	for _, v := range s {
		if v > max {
			max = v
		}
	}

	return max
}

func Min(s []int) int {
	var min = s[0]

	for _, v := range s {
		if v < min {
			min = v
		}
	}

	return min
}

func Average(s []int, n int) int {
	var sum int

	for _, v := range s {
		sum += v
	}

	var average = sum / n

	return average
}
