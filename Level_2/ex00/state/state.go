package main

type State interface {
	charge() error
	shoot() error
	switchFuse() error
}
