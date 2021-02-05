package main

import "fmt"

type Nomer struct {
	num int
}

type Shaxs struct{
ism, familiya string
Nomer
}

func main() {
	i := Shaxs{
		ism: "Alisher",
		familiya: "Valisherov",
		Nomer:   Nomer{998901234567},
	}

	fmt.Println("Alisher Valisherovning nomeri : ", i.num)
}
