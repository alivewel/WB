package main

type School struct {
	action string
	want   string
}

func (s *School) accept(v Visitor) {
	v.visitSchool(s)
}

func newSchool() *School {
	return &School{
		action: "Starting to study",
		want:   "Want to know more"}
}
