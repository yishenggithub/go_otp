package main

type supervisor interface {
	init()
	startLink(string)
}
