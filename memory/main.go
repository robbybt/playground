package main

import "fmt"

type test struct {
	success *success
}

type success struct {
	hit int64
	asd int64
}

func main(){
	var a *test
	a = &test{
		success:&success{
			hit:1,
			asd:10,
		},
	}
	var b *test
	b = func(s test) *test{
		s.success.hit = 12
		return &s
	}(*a)

	fmt.Println(b.success.hit)
	fmt.Println(a.success.hit)
}
