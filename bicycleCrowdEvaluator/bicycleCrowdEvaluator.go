package bicycleCrowdEvaluator

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math"
	"sort"
	"strings"
)

type AnnotatorInfo struct {
	NumAnnotators          int
	AverageAnnotationTimes int
	MinAnnotationTimes     int
	MaxAnnotationTimes     int
	AnnotatorResults       map[string]int
}

type Pair struct {
	TaskOutput map[string]interface{}
	ImageUrl   string
	User       map[string]interface{}
}

type Response struct {
	CantSolve   int
	CorruptData int
}

type ReferenceDataSetResult struct {
	IsBicycle    int
	IsNotBicycle int
}

type Votes struct {
	Yes int
	No  int
}

type DisagreedQuestion struct {
	ImageUrl string
	Votes    Votes
}

func Decode(filename string) map[string]interface{} {
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	var responses map[string]interface{}

	if err := json.Unmarshal(b, &responses); err != nil {
		panic(err)
	}

	return responses
}

func Annotators(responses map[string]interface{}) (AnnotatorInfo, Response, []string, map[string][]Pair) {

	var m map[string]int = make(map[string]int)
	var s []int
	var m2 map[string][]Pair = make(map[string][]Pair)

	for _, v1 := range responses {
		for _, v2 := range v1.(map[string]interface{}) {
			for k3, v3 := range v2.(map[string]interface{}) {
				if k3 == "results" {
					for k4, v4 := range v3.(map[string]interface{}) {
						for k5, v5 := range v4.(map[string]interface{}) {
							if k5 == "results" {
								var tmp string
								var user map[string]interface{} = make(map[string]interface{})
								var taskOutput map[string]interface{} = make(map[string]interface{})
								for _, v6 := range v5.([]interface{}) {
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
											user = v7.(map[string]interface{})
											for k8, v8 := range v7.(map[string]interface{}) {
												if k8 == "vendor_user_id" {
													m[v8.(string)] += 1
												}
											}
										}

										if k7 == "task_output" {
											taskOutput = v7.(map[string]interface{})
											for k8, v8 := range v7.(map[string]interface{}) {
												if k8 == "duration_ms" {
													s = append(s, int(v8.(float64)))
												}
											}

										}

									}
									if user["vendor_user_id"].(string) == "annotator_11" {
										if taskOutput["cant_solve"].(bool) {
											fmt.Println("result:", taskOutput["cant_solve"].(bool))
											fmt.Println("imageUrl:", tmp)
										}
									}

									if len(taskOutput) != 0 && tmp != "" && len(user) != 0 {
										m2[k4] = append(m2[k4], Pair{taskOutput, tmp, user})
									}
								}
							}
						}
					}
				}
			}
		}
	}

	var response Response = Response{}
	var cantSolveOrCorruptDataAnnotators []string

	for _, v := range m2 {
		for _, v := range v {
			cantSolve, ok1 := v.TaskOutput["cant_solve"].(bool)
			corruptData, ok := v.TaskOutput["corrupt_data"].(bool)

			if ok1 && cantSolve {
				response.CantSolve += 1
				val, ok := v.User["vendor_user_id"].(string)
				if ok {
					cantSolveOrCorruptDataAnnotators = append(cantSolveOrCorruptDataAnnotators, val)
				}
			} else if ok && corruptData {
				if val, ok := v.User["vendor_user_id"].(string); ok {
					cantSolveOrCorruptDataAnnotators = append(cantSolveOrCorruptDataAnnotators, val)
				}
				response.CorruptData += 1
			}
		}
	}
	fmt.Println()
	var distinct []string = Deduplicate(cantSolveOrCorruptDataAnnotators)

	var annotatorInfo AnnotatorInfo = AnnotatorInfo{NumAnnotators: len(m),
		AverageAnnotationTimes: Average(s, len(s)), MinAnnotationTimes: Min(s),
		MaxAnnotationTimes: Max(s), AnnotatorResults: m}
	return annotatorInfo, response, distinct, m2
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

func ReferenceDataSetDistribution(file string) ReferenceDataSetResult {
	var m map[string]interface{} = Decode(file)
	var referenceDataSetResult ReferenceDataSetResult = ReferenceDataSetResult{}

	for _, v := range m {
		m := v.(map[string]interface{})
		v, ok := m["is_bicycle"].(bool)
		if ok {
			if v {
				referenceDataSetResult.IsBicycle += 1
			} else {
				referenceDataSetResult.IsNotBicycle += 1
			}
		}
	}

	return referenceDataSetResult
}

func GetReferenceSet(file string) map[string]map[string]interface{} {
	var m map[string]interface{} = Decode(file)

	var m2 map[string]map[string]interface{} = make(map[string]map[string]interface{})
	for k, v := range m {
		m2[k] = v.(map[string]interface{})
	}

	return m2
}

func GetAnnotators(file string) map[string]map[string]string {
	var m map[string]interface{} = Decode(file)
	var imageToAnnotatorResponses map[string]map[string]string = make(map[string]map[string]string)

	for _, v := range m {
		var v = v.(map[string]interface{})
		for _, v := range v {
			var v = v.(map[string]interface{})
			for k, v := range v {
				if k == "results" {
					var v = v.(map[string]interface{})
					for _, v := range v {
						var v = v.(map[string]interface{})
						for _, v := range v {
							v, ok := v.([]interface{})

							if ok {
								for _, v := range v {
									var imageUrl, annotator, answer string

									if v, ok := v.(map[string]interface{})["task_input"].(map[string]interface{})["image_url"].(string); ok {
										imageUrl = v
									}

									if v, ok := v.(map[string]interface{})["user"].(map[string]interface{})["vendor_user_id"]; ok {
										annotator = v.(string)
									}

									if v, ok := v.(map[string]interface{})["task_output"].(map[string]interface{})["answer"]; ok {
										answer = v.(string)
									}

									imageToAnnotatorResponses[imageUrl] = map[string]string{annotator: answer}
								}

							}
						}
					}
				}
			}
		}
	}

	return imageToAnnotatorResponses
}

func TheGoodTheBadAnnotators(reference map[string]map[string]interface{},
	annotators map[string]map[string]string) (map[string]int, map[string]int) {
	var correct []string
	var wrong []string
	var timesCorrect map[string]int = make(map[string]int)
	var timesWrong map[string]int = make(map[string]int)

	for k1, v1 := range reference {
		for k2, v2 := range annotators {
			if strings.Contains(k2, k1) {
				for k, v := range v2 {
					if v == "yes" && v1["is_bicycle"].(bool) {
						correct = append(correct, k)
					} else if v == "no" && v1["is_bicycle"] == false {
						correct = append(correct, k)
					} else {
						wrong = append(wrong, k)
					}
				}
			}
		}
	}

	for _, v := range correct {
		timesCorrect[v] += 1
	}

	for _, v := range wrong {
		timesWrong[v] += 1
	}

	return timesCorrect, timesWrong
}

func Deduplicate(in []string) []string {
	sort.Strings(in)
	j := 0
	for i := 1; i < len(in); i++ {
		if in[j] == in[i] {
			continue
		}
		j++
		in[j] = in[i]
	}
	result := in[:j+1]
	return result
}

func Scores(timesCorrect map[string]int, timesWrong map[string]int) map[string]float32 {
	var scores map[string]float32 = make(map[string]float32)

	for k, v := range timesCorrect {
		scores[k] = (float32(v) / float32((v + timesWrong[k]))) * 100
	}

	return scores
}

func Rank(scores map[string]float32) (map[string]float32, map[string]float32) {
	var goodAnnotators map[string]float32 = make(map[string]float32)
	var badAnnotators map[string]float32 = make(map[string]float32)

	for k, v := range scores {
		// Arbitrary pass mark
		if v > 93 {
			goodAnnotators[k] = v
		} else {
			badAnnotators[k] = v
		}
	}

	return goodAnnotators, badAnnotators
}

func Contains(s []Pair, pair Pair) {

}

func QuestionYesNoAnswers(file string) map[string]*Votes {
	var annotatorResponses map[string]interface{} = Decode(file)
	var m map[string]*Votes = make(map[string]*Votes)

	var rootNode map[string]interface{} = annotatorResponses["results"].(map[string]interface{})
	results1, ok := rootNode["root_node"].(map[string]interface{})

	if ok {
		for _, v := range results1["results"].(map[string]interface{}) {
			for _, v := range v.(map[string]interface{}) {
				temp, ok := v.([]interface{})

				if ok {
					var votes *Votes = &Votes{}
					for i := 0; i < len(temp); i++ {
						var imageUrl = v.([]interface{})[i].(map[string]interface{})["task_input"].(map[string]interface{})["image_url"].(string)
						m[imageUrl] = votes

						if v.([]interface{})[i].(map[string]interface{})["task_output"].(map[string]interface{})["answer"] == "yes" {
							m[imageUrl].Yes += 1
						} else {
							m[imageUrl].No += 1
						}
					}
				}
			}
		}
	}

	return m
}

func HighDisagreedQuestions(yesNoAnswers map[string]*Votes) []DisagreedQuestion {
	var s []DisagreedQuestion
	for k, v := range yesNoAnswers {
		// Absolute difference less than average of yes and no answers for specific question means
		// annotators highly disagree
		if int((math.Abs(float64((*v).Yes) - float64((*v).No)))) < ((*v).Yes+(*v).No)/2 {
			s = append(s, DisagreedQuestion{k, *v})
		}
	}
	return s
}

func AverageTimeOfAnnotators(m map[string][]Pair, annotatorInfo AnnotatorInfo) map[string]float64 {
	var m2 map[string]float64 = map[string]float64{}

	for _, v := range m {
		for _, v := range v {
			if v.User["vendor_user_id"].(string) == "annotator_19" {
				fmt.Println(v.TaskOutput["duration_ms"].(float64))
			}
			//if v.TaskOutput["duration_ms"].(float64) > 0 {
			m2[v.User["vendor_user_id"].(string)] += v.TaskOutput["duration_ms"].(float64 /*int*/)
			//}
		}
	}

	for k, v := range m2 {
		m2[k] = v / float64(annotatorInfo.AnnotatorResults[k])
	}

	return m2
}
