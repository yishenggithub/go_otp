package main

type genServer interface {
	init(chan string)
	startLink(chan string)
	//handleCall (string)
	//handleCast (string)
	handleInfo(chan string)
	//terminate ()
}
