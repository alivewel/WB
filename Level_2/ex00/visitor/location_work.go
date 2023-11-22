package main

type Work struct {
	action string
	want   string
}

func (w *Work) accept(v Visitor) {
	v.visitWork(w)
}

func newWork() *Work {
	return &Work{
		action: "Do tasks",
		want:   "Want to receive money"}
}
