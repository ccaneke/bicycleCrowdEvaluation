# BicycleCrowdEvaluator
An application for evaluating the performance of annotators in annotating datasets of bicycle images

# Instructions to run the application
Open a terminal and run the following commands:

```
//clone repository to your local computer
$ git clone https://github.com/ccaneke/bicycleCrowdEvaluation

// cd into the bicycleProjectCrowdEvaluation directory
$ cd bicycleProjectCrowdEvaluation

// compile the main package to generate an executable
$ go build main.go

// run the executable (on Windows run aboutpage without the "./" instead)
$ ./main
```

Alternatively:

```
// compile and install the main package
$ go install https://github.com/ccaneke/bicycleCrowdEvaluation

// run the following command in the bicycleProjectCrowdEvaluation directory to add the install directory path for the main package to your PATH
$ export PATH=$PATH:$(dirname $(go list -f '{{.Target}} .)

// you can then run the executable from anywhere on your system
$ main

```
Another option is to run the program without building an executable first:

```
// run the following in the bicycleProjectCrowdEvaluation directory where the main.go file is
$ go run main.go
```