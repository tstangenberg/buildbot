package main

//go:generate moq -out application_moq_test.go . Application

type Application interface {
	Run(args []string) (err error)
}
