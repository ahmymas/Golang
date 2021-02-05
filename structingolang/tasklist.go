package main

import "fmt"

type vazifa struct {
	ish string
}

type vaqti struct{
haftakuni string
vazifa
}

func main() {
	i := vaqti{
		haftakuni: "Shanba",
	  vazifa:   vazifa {"Uylarni yig'ishtirish"},
	}

	fmt.Println(i.haftakuni, "kuni qilinadigan ish bu : ", i.vazifa)
}
