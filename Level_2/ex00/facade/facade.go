package main

type Facade struct {
	subSystem1 SubSystem1
	subSystem2 SubSystem2
}

func (f Facade) OperationWrapper() {
	f.subSystem1.Suboperation()
	f.subSystem2.Suboperation()
}
