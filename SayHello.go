package test_pzn_module

import "strconv"

func SayHello(name string, age int) string {
	return "Hello World! " + name + " with age " + strconv.Itoa(age)
}
