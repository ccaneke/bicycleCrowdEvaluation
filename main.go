package main

import (
	"fmt"

	"github.com/ccaneke/bicycleProjectCrowdEvaluation/bicycleCrowdEvaluator"
)

func main() {
	var responses = bicycleCrowdEvaluator.Decode("/tmp/anonymized_project.json")

	annotatorsInfo, response := bicycleCrowdEvaluator.Annotators(responses)

	fmt.Println("Task 1:")
	//fmt.Println("AnnotatorInfo")
	fmt.Printf("%+v", annotatorsInfo)

	//fmt.Println(annotatorsInfo)

	var yesNoAnswers map[string]*bicycleCrowdEvaluator.Votes = bicycleCrowdEvaluator.QuestionYesNoAnswers("/tmp/anonymized_project.json")

	fmt.Println()
	fmt.Println()
	fmt.Println("HighlyDisagreedQuestions...")
	disagreedQuestions := bicycleCrowdEvaluator.HighDisagreedQuestions(yesNoAnswers)
	fmt.Println(disagreedQuestions)

	fmt.Println()
	fmt.Println("Task 2:")

	fmt.Println("CantSolve:", response.CantSolve, "CorruptData:", response.CorruptData)
	fmt.Println("Trend shows that annotators that selected true for cant_solve or corrupt_data had" +
		"lower annotation times (duration), and of course didn't not answer yes or no")

	fmt.Println("Task 3:")
	var referenceDataSetDistribution bicycleCrowdEvaluator.ReferenceDataSetResult = bicycleCrowdEvaluator.ReferenceDataSetDistribution("/tmp/references.json")

	fmt.Println("IsBicycle", referenceDataSetDistribution.IsBicycle, "IsNotBicycle",
		referenceDataSetDistribution.IsNotBicycle, "\nReference Dataset is balanced because the number"+
			" of images of bicycles and images that are not bicycles are approximately the same")

	fmt.Println("Getting ReferenceSet...")
	var referenceSet = bicycleCrowdEvaluator.GetReferenceSet("/tmp/references.json")

	fmt.Println()

	fmt.Println("Getting Annotators...")
	var annotators = bicycleCrowdEvaluator.GetAnnotators("/tmp/anonymized_project.json")

	timesCorrect, timesWrong := bicycleCrowdEvaluator.TheGoodTheBadAnnotators(referenceSet, annotators)

	fmt.Println("Times annotators were correct:")
	fmt.Println(timesCorrect)

	fmt.Println()
	fmt.Println("Times annotators were wrong:")
	fmt.Println(timesWrong)

	fmt.Println()
	var scores map[string]float32 = bicycleCrowdEvaluator.Scores(timesCorrect, timesWrong)
	fmt.Println("Score in percentage of annotators' correct answers:")
	fmt.Println(scores)

	fmt.Println()
	good, bad := bicycleCrowdEvaluator.Rank(scores)
	fmt.Println("Good Annotators:")
	fmt.Println(good)
	fmt.Println("")
	fmt.Println("Bad Annotators:")
	fmt.Println(bad)

}
