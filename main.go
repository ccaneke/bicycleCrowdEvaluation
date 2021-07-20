package main

import (
	"fmt"

	"github.com/ccaneke/bicycleProjectCrowdEvaluation/bicycleCrowdEvaluator"
)

func main() {
	var responses = bicycleCrowdEvaluator.Decode("/home/DeepLearning/Downloads/anonymized_project.json")

	var numAnnotators = bicycleCrowdEvaluator.NumAnnotators(responses)

	fmt.Println(numAnnotators)
}
