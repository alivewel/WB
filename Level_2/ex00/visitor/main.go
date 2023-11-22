package main

func main() {
	home := newHome()
	work := newWork()
	school := newSchool()

	visitorAction := VisitorAction{}

	visitorAction.visitHome(home)
	visitorAction.visitWork(work)
	visitorAction.visitSchool(school)

	visitorWant := VisitorWant{}

	visitorWant.visitHome(home)
	visitorWant.visitWork(work)
	visitorWant.visitSchool(school)
}

// Команда для запуска:
// go run .
