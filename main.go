package main

import (
	"fmt"

	"strconv"

	"os"

	"github.com/ccaneke/bicycleProjectCrowdEvaluation/bicycleCrowdEvaluator"
	"github.com/wcharczuk/go-chart"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
)

func PlotBarChartAnnotationTimes(v1 int, v2 int, v3 int) {
	groupA := plotter.Values{float64(v1), float64(v2), float64(v3)}

	p := plot.New()

	//p.Title.Text = "Bar chart"
	p.Y.Label.Text = "Annotation times (milliseconds)"

	w := vg.Points(20)

	barsA, err := plotter.NewBarChart(groupA, w*2)
	if err != nil {
		panic(err)
	}
	barsA.LineStyle.Width = vg.Length(0)
	barsA.Color = plotutil.Color(0)
	barsA.Offset = 0

	p.Add(barsA)
	p.Legend.Top = true
	p.NominalX("Average", "Min", "Max")

	if err := p.Save(5*vg.Inch, 3*vg.Inch, "barchart.png"); err != nil {
		panic(err)
	}
}

func PlotBarChartAnnotatorResults(m map[string]int) {
	var m2 map[string]float64 = make(map[string]float64)

	for k, v := range m {
		m2[k] = float64(v)
	}

	groupA := plotter.Values{m2["annotator_01"], m2["annotator_02"], m2["annotator_03"], m2["annotator_04"],
		m2["annotator_05"], m2["annotator_06"], m2["annotator_07"], m2["annotator_08"], m2["annotator_09"],
		m2["annotator_10"], m2["annotator_11"], m2["annotator_12"], m2["annotator_13"], m2["annotator_14"],
		m2["annotator_15"], m2["annotator_16"], m2["annotator_17"], m2["annotator_18"], m2["annotator_19"],
		m2["annotator_20"], m2["annotator_21"], m2["annotator_22"],
	}

	p := plot.New()

	//p.Title.Text = "Bar chart"
	p.Y.Label.Text = "Annotator results"
	p.X.Label.Text = "Annotators"

	w := vg.Points(20)

	barsA, err := plotter.NewBarChart(groupA, w)
	if err != nil {
		panic(err)
	}
	barsA.LineStyle.Width = vg.Length(0)
	barsA.Color = plotutil.Color(0)
	barsA.Offset = 0

	p.Add(barsA)
	p.Legend.Top = true
	p.NominalX("1", "2", "3", "4", "5", "6",
		"7", "8", "9", "10", "11", "12",
		"13", "14", "15", "16", "17", "18",
		"19", "20", "21", "22")

	if err := p.Save(10*vg.Inch, 10*vg.Inch, "barchart2.png"); err != nil {
		panic(err)
	}
}

func PlotBarChartHighlyDisagreedQuestions(s []bicycleCrowdEvaluator.DisagreedQuestion) {
	var first plotter.Values
	var second plotter.Values
	var xAxisNames []string

	for _, v := range s {
		first = append(first, float64(v.Votes.Yes))
		second = append(second, float64(v.Votes.No))
		xAxisNames = append(xAxisNames, v.ImageUrl[65:])
	}

	j := 16
	k := 0
	for i := 0; i < len(first); i += 16 {
		firstTemp := first[i:j]
		secondTemp := second[i:j]
		xAxisNamesTemp := xAxisNames[i:j]

		p := plot.New()

		p.Title.Text = "Highly disagreed questions"
		p.Y.Label.Text = "Votes"
		p.X.Label.Text = "Images"

		w := vg.Points(20)

		barsA, err := plotter.NewBarChart(firstTemp, w)
		if err != nil {
			panic(err)
		}
		barsA.LineStyle.Width = vg.Length(0)
		barsA.Color = plotutil.Color(0)
		barsA.Offset = -w

		barsB, err := plotter.NewBarChart(secondTemp, w)
		if err != nil {
			panic(err)
		}
		barsB.LineStyle.Width = vg.Length(0)
		barsB.Color = plotutil.Color(1)

		p.Add(barsA, barsB)
		p.Legend.Add("Yes", barsA)
		p.Legend.Add("No", barsB)
		p.Legend.Top = true
		p.Legend.XOffs = -10
		p.NominalX(xAxisNamesTemp...)

		// Aspect ratio of 1.6667 is used for best results
		if err := p.Save(17*vg.Inch, 10*vg.Inch, "/tmp/testdata/barchart"+strconv.Itoa(k)+".png"); err != nil {
			panic(err)
		}
		k++
		j += 16
	}
}

func PlotAverageAnnotationTimesForEachAnnotator(m map[string]float64) {
	groupA := plotter.Values{m["annotator_01"], m["annotator_02"], m["annotator_03"], m["annotator_04"],
		m["annotator_05"], m["annotator_06"], m["annotator_07"], m["annotator_08"], m["annotator_09"],
		m["annotator_10"], m["annotator_11"], m["annotator_12"], m["annotator_13"], m["annotator_14"],
		m["annotator_15"], m["annotator_16"], m["annotator_17"], m["annotator_18"], m["annotator_19"],
		m["annotator_20"], m["annotator_21"], m["annotator_22"],
	}

	p := plot.New()

	p.Title.Text = "Average annotation times for annotators"
	p.Y.Label.Text = "Annotation times (milliseconds)"
	p.X.Label.Text = "Annotators"

	w := vg.Points(20)

	barsA, err := plotter.NewBarChart(groupA, w)
	if err != nil {
		panic(err)
	}
	barsA.LineStyle.Width = vg.Length(0)
	barsA.Color = plotutil.Color(0)
	barsA.Offset = 0

	p.Add(barsA)
	p.Legend.Top = true
	p.NominalX("1", "2", "3", "4", "5", "6",
		"7", "8", "9", "10", "11", "12",
		"13", "14", "15", "16", "17", "18",
		"19", "20", "21", "22")

	if err := p.Save(10*vg.Inch, 10*vg.Inch, "barchart5.png"); err != nil {
		panic(err)
	}
}

func PlotPieChart(d bicycleCrowdEvaluator.ReferenceDataSetResult) {
	pie := chart.PieChart{
		Width:  512,
		Height: 512,
		Values: []chart.Value{
			chart.Value{Value: float64(d.IsBicycle), Label: "Bicycle Images: " + strconv.Itoa(d.IsBicycle)},
			chart.Value{Value: float64(d.IsNotBicycle), Label: "Not bicycle images: " + strconv.Itoa(d.IsNotBicycle)},
		},
	}

	f, _ := os.Create("output.png")
	defer f.Close()
	pie.Render(chart.PNG, f)
}

func PlotGoodOrBadAnnotators(annotators map[string]float32, title string) {
	var group plotter.Values
	var keys []string
	for k, v := range annotators {
		group = append(group, float64(v))
		keys = append(keys, k)
	}

	p := plot.New()

	p.Title.Text = title
	p.Y.Label.Text = "Annotation scores (Percentage)"
	p.X.Label.Text = "Annotators"

	w := vg.Points(20)

	barsA, err := plotter.NewBarChart(group, w*2)
	if err != nil {
		panic(err)
	}
	barsA.LineStyle.Width = vg.Length(0)
	barsA.Color = plotutil.Color(2)
	barsA.Offset = 0

	p.Add(barsA)
	p.Legend.Top = true
	p.NominalX(keys...)

	if err := p.Save(17*vg.Inch, 10*vg.Inch, "barchart"+title+".png"); err != nil {
		panic(err)
	}
}

func main() {
	var responses = bicycleCrowdEvaluator.Decode("/tmp/anonymized_project.json")

	annotatorsInfo, response, cantSolveOrCorruptDataAnnotators, m2 := bicycleCrowdEvaluator.Annotators(responses)

	fmt.Println("Task 1:")
	fmt.Printf("%+v", annotatorsInfo)

	var yesNoAnswers map[string]*bicycleCrowdEvaluator.Votes = bicycleCrowdEvaluator.QuestionYesNoAnswers("/tmp/anonymized_project.json")

	fmt.Println()
	fmt.Println()
	fmt.Println("HighlyDisagreedQuestions...")
	disagreedQuestions := bicycleCrowdEvaluator.HighDisagreedQuestions(yesNoAnswers)
	fmt.Println(disagreedQuestions)

	fmt.Println()
	fmt.Println("Task 2:")

	var averageTimeOfAnnotators map[string]float64 = bicycleCrowdEvaluator.AverageTimeOfAnnotators(m2, annotatorsInfo)
	fmt.Println("CantSolve:", response.CantSolve, "CorruptData:", response.CorruptData)
	fmt.Println("Annotators who selected can't solve or corrupt data:", cantSolveOrCorruptDataAnnotators)
	fmt.Println("Trend shows that annotators that selected true for cant_solve or corrupt_data had" +
		" lower annotation times (duration) (when they chose either of these two options for a" +
		"question), and of course didn't not answer yes or no")
	fmt.Println("AverageTimeOfAllAnnotators:", averageTimeOfAnnotators)
	PlotAverageAnnotationTimesForEachAnnotator(averageTimeOfAnnotators)

	fmt.Println("Task 3:")
	var referenceDataSetDistribution bicycleCrowdEvaluator.ReferenceDataSetResult = bicycleCrowdEvaluator.ReferenceDataSetDistribution("/tmp/references.json")

	PlotPieChart(referenceDataSetDistribution)
	fmt.Println("IsBicycle", referenceDataSetDistribution.IsBicycle, "IsNotBicycle",
		referenceDataSetDistribution.IsNotBicycle, "\nReference Dataset is balanced because the number"+
			" of images of bicycles and images that are not bicycles are approximately the same")

	fmt.Println("Task 4:")
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
	PlotGoodOrBadAnnotators(good, "Good annotators")
	PlotGoodOrBadAnnotators(bad, "Bad annotators")
}
