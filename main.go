package main

import (
	"flag"
	"fmt"

	"github.com/c-sto/gocryptopals/cmd/sets/Set1"
	"github.com/c-sto/gocryptopals/cmd/sets/Set2"
	"github.com/c-sto/gocryptopals/cmd/sets/Set3"
)

func main() {

	//check for args -s ##
	//if no args:
	var challNumber int
	flag.IntVar(&challNumber, "c", -1, "Challenge number to run")
	flag.Parse()
	if challNumber < 0 {
		fmt.Println("Enter challenge number to run")
	}
	if challNumber < 0 {
		panic("Bad challenge number")
	}

	//can probably tidy this up with some sort of refle\ct? idk
	switch challNumber {

	case 1:
		Set1.Challenge1()
	case 2:
		Set1.Challenge2()
	case 3:
		Set1.Challenge3()
	case 4:
		Set1.Challenge4()
	case 5:
		Set1.Challenge5()
	case 6:
		Set1.Challenge6()
	case 7:
		Set1.Challenge7()
	case 8:
		Set1.Challenge8()
	case 9:
		Set2.Challenge9()
	case 10:
		Set2.Challenge10()
	case 11:
		Set2.Challenge11()
	case 12:
		Set2.Challenge12()
	case 13:
		Set2.Challenge13()
	case 14:
		Set2.Challenge14()
	case 15:
		Set2.Challenge15()
	case 16:
		Set2.Challenge16()
	case 17:
		Set3.Challenge17()
	default:
		fmt.Println("Can't find specified challenge")
	}
}
