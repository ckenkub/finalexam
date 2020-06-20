package main

import "github.com/ckenkub/finalexam/customer"

func main() {
	r := customer.SetupRouter()
	r.Run(":2019")
}
