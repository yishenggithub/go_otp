package main
type genServer interface {
	init ()
	handleCall (string)
	handleCast (string)
	handleInfo (string)
	terminate ()
}