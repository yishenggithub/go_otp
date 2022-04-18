package main

type spervisor interface {
	init ()
	start_link (string)
}