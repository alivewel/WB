package main

type Visitor interface {
	visitHome(*Home)
	visitWork(*Work)
	visitSchool(*School)
}
