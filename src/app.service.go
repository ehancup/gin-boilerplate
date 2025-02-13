package src

import "gin-boilerplate/src/app/example"

func GetExampleServie() *example.Service{
	return &example.Service{}
}