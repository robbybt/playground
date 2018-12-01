package main

import (
	"fmt"
	"sync"
)

type asd struct{
	text string
}
type cBreakersObj struct{
	list map[string]asd
}

var cBreakers cBreakersObj
func (c *cBreakersObj) setCircuitBreaker(list []string, v *asd) {
	if c.list == nil {
		c.list = make(map[string]asd)
	}
	for _, cbList := range list {
		c.list[cbList] = *v
	}
}
func main(){
	cBreakers.SetUpCircuitBreaker([]string{"a","b"},&asd{text:"zz"})
	fmt.Println(cBreakers.list)
}
