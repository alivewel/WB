package main

type Home struct {
	action string
	want   string
}

func (h *Home) accept(v Visitor) {
	v.visitHome(h)
}

func newHome() *Home {
	return &Home{
		action: "Go to sleep",
		want:   "Want to go home"}
}
