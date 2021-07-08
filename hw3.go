package main

import "fmt"

func main()  {
	a:=9.0
	b:=8.0
	c:=7.0
	d:=6.3
	e:=-1

	fmt.Println(a+b+c)
	fmt.Println(a*b*c)
	fmt.Println(b+c+d)
	fmt.Println(b*c*d)
	fmt.Println(a/b)
	fmt.Println(int(a)%int(b))
	fmt.Println(int(a)/e)
	fmt.Println((int(a)*int(a)+int(b)*int(b))-(int(c)^2-int(d)^2))
}

