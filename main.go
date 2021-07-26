package main

import (
	"fmt"

	"strconv"

	"github.com/ccaneke/bicycleProjectCrowdEvaluation/bicycleCrowdEvaluator"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
)

func PlotBarChartAnnotationTimes(v1 int, v2 int, v3 int) {
	// values := plotter.Values{float64( /*v1*/ 100), float64( /*v2*/ 200), float64( /*v3*/ -300)}
	// verticalLabels := []string{"Average, Min, Max"}
	// //horizontalLabels := []string{"Average", "Min", "Max"}

	// p1 := plot.New()
	// verticalBarChart, err := plotter.NewBarChart(values, 0.5*vg.Centimeter)
	// if err != nil {
	// 	log.Panic(err)
	// }

	// p1.Add(verticalBarChart)
	// p1.NominalX(verticalLabels...)
	// err = p1.Save(100, 100, "/tmp/verticalBarChart.png")
	// if err != nil {
	// 	log.Panic(err)
	// }
	groupA := plotter.Values{float64(v1), float64(v2), float64(v3) /*-20, 35, 30, 35, 27*/}
	//groupB := plotter.Values{25 /*, 32, 34, 20, 25*/}
	//groupC := plotter.Values{12 /*, 28, 15, 21, 8*/}

	p := plot.New()

	p.Title.Text = "Bar chart"
	p.Y.Label.Text = "Annotation times (milliseconds)"

	w := vg.Points(20)

	barsA, err := plotter.NewBarChart(groupA, w*2)
	if err != nil {
		panic(err)
	}
	barsA.LineStyle.Width = vg.Length(0)
	barsA.Color = plotutil.Color(0)
	barsA.Offset = /*-w*/ 0

	// barsB, err := plotter.NewBarChart(groupB, w)
	// if err != nil {
	// 	panic(err)
	// }
	// barsB.LineStyle.Width = vg.Length(0)
	// barsB.Color = plotutil.Color(1)

	// barsC, err := plotter.NewBarChart(groupC, w)
	// if err != nil {
	// 	panic(err)
	// }
	// barsC.LineStyle.Width = vg.Length(0)
	// barsC.Color = plotutil.Color(2)
	// barsC.Offset = w

	p.Add(barsA /*, barsB, barsC*/)
	//p.Legend.Add("Group A", barsA)
	// p.Legend.Add("Group B", barsB)
	// p.Legend.Add("Group C", barsC)
	p.Legend.Top = true
	p.NominalX("Average", "Min", "Max")

	if err := p.Save(5*vg.Inch, 3*vg.Inch, "barchart.png"); err != nil {
		panic(err)
	}
}

func PlotBarChartAnnotatorResults(m map[string]int) {
	var s []float64
	for _, v := range m {
		s = append(s, float64(v))
	}

	groupA := plotter.Values{s[0], s[1], s[2], s[3], s[4], s[5], s[6], s[7], s[8], s[9], s[10], s[11],
		s[12], s[13], s[14], s[15], s[16], s[17], s[18], s[19], s[20], s[21] /*float64(v1), float64(v2), float64(v3)*/ /*-20, 35, 30, 35, 27*/}
	//groupB := plotter.Values{25 /*, 32, 34, 20, 25*/}
	//groupC := plotter.Values{12 /*, 28, 15, 21, 8*/}

	p := plot.New()

	p.Title.Text = "Bar chart"
	p.Y.Label.Text = "Annotator results"
	p.X.Label.Text = "Annotators"

	w := vg.Points(20)

	barsA, err := plotter.NewBarChart(groupA, w /**2*/)
	if err != nil {
		panic(err)
	}
	barsA.LineStyle.Width = vg.Length(0)
	barsA.Color = plotutil.Color(0)
	barsA.Offset = /*-w*/ 0

	// barsB, err := plotter.NewBarChart(groupB, w)
	// if err != nil {
	// 	panic(err)
	// }
	// barsB.LineStyle.Width = vg.Length(0)
	// barsB.Color = plotutil.Color(1)

	// barsC, err := plotter.NewBarChart(groupC, w)
	// if err != nil {
	// 	panic(err)
	// }
	// barsC.LineStyle.Width = vg.Length(0)
	// barsC.Color = plotutil.Color(2)
	// barsC.Offset = w

	p.Add(barsA /*, barsB, barsC*/)
	//p.Legend.Add("Group A", barsA)
	// p.Legend.Add("Group B", barsB)
	// p.Legend.Add("Group C", barsC)
	p.Legend.Top = true
	p.NominalX("1", "2", "3", "4", "5", "6",
		"7", "8", "9", "10", "11", "12",
		"13", "14", "15", "16", "17", "18",
		"19", "20", "21", "22")

	if err := p.Save( /*5*/ 10*vg.Inch /*3*/, 10*vg.Inch, "barchart2.png"); err != nil {
		panic(err)
	}
}

func PlotBarChartHighlyDisagreedQuestions(s []bicycleCrowdEvaluator.DisagreedQuestion) {
	// groupA := plotter.Values{20, 35, 30, 35, 27}
	// groupB := plotter.Values{25, 32, 34, 20, 25}
	// groupC := plotter.Values{12, 28, 15, 21, 8}
	// groupD := plotter.Values{20, 35, 30, 35, 27}
	// groupE := plotter.Values{25, 32, 34, 20, 25}
	// groupF := plotter.Values{12, 28, 15, 21, 8}

	// var s2 []plotter.Values = make([]plotter.Values, len(s)* 2)

	// var j int
	// for i, v := range s {
	// 	s2[]
	// }
	var first plotter.Values
	var second plotter.Values
	var xAxisNames []string
	// for i := 0; i < len(s /*s2*/); i++ /*i += 2*/ {
	// 	//s2[i] = plotter.Values{}
	// 	first = append(first, float64(s[i].Votes.Yes))
	// 	second = append(second, float64(s[i].Votes.No))
	// }
	for _, v := range s {
		first = append(first, float64(v.Votes.Yes))
		second = append(second, float64(v.Votes.No))
		xAxisNames = append(xAxisNames, v.ImageUrl[65:])
	}

	j := 16
	k := 0
	for i := 0; i < len(first); i += 16 {
		firstTemp := first[i: /*16*/ j]
		secondTemp := second[i:j]
		xAxisNamesTemp := xAxisNames[i: /*16*/ j]

		p := plot.New()

		p.Title.Text = "Highly disagreed questions"
		p.Y.Label.Text = "Votes"
		p.X.Label.Text = "Images"

		w := vg.Points(20)

		barsA, err := plotter.NewBarChart(firstTemp /*groupA*/, w)
		if err != nil {
			panic(err)
		}
		barsA.LineStyle.Width = vg.Length(0)
		barsA.Color = plotutil.Color(0)
		barsA.Offset = -w

		barsB, err := plotter.NewBarChart(secondTemp /*groupB*/, w)
		if err != nil {
			panic(err)
		}
		barsB.LineStyle.Width = vg.Length(0)
		barsB.Color = plotutil.Color(1)

		// barsC, err := plotter.NewBarChart(groupC, w)
		// if err != nil {
		// 	panic(err)
		// }
		// barsC.LineStyle.Width = vg.Length(0)
		// barsC.Color = plotutil.Color(2)
		// barsC.Offset = w

		p.Add(barsA, barsB /*, barsC*/)
		p.Legend.Add( /*"Group A"*/ "Yes", barsA)
		p.Legend.Add( /*"Group B"*/ "No", barsB)
		//p.Legend.Add("Group C", barsC)
		p.Legend.Top = true
		p.Legend.XOffs = -10
		p.NominalX( /*xAxisNames[0]*/ xAxisNamesTemp... /*"img"*/ /*"One"*/ /*, "Two", "Three", "Four", "Five"*/)

		// Aspect ratio of 1.6667 is used for best results
		if err := p.Save( /*5**/ 17* /*20**/ /*23**/ /*30**/ vg.Inch /*3*/ /*3**/, 10* /*12**/ /*14**/ /*10**/ vg.Inch, "/tmp/testdata/barchart"+strconv.Itoa(k)+".png"); err != nil {
			panic(err)
		}
		k++
		j += 16
	}
}

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

	PlotBarChartAnnotationTimes(annotatorsInfo.AverageAnnotationTimes, annotatorsInfo.MinAnnotationTimes,
		annotatorsInfo.MaxAnnotationTimes)
	PlotBarChartAnnotatorResults(annotatorsInfo.AnnotatorResults)
	PlotBarChartHighlyDisagreedQuestions(disagreedQuestions)
}
