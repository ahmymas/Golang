/*GO dasturlash tilida yozilgan har bir programma paketlar (packages) bilan ishlaydi
Dastur main paketi bilan boshlanadi*/

package main 

import (
	"fmt"
)

func main() {
	fmt.Println("Salom, GO!")
}


/*import qilish
quyidagi import so'zidan keyin kelgan qavs ichidagi so'zlar kodlar jamlanmasini sizga foydalanish uchun taqdim etadi
keling matematik formulalarni "math" so'zi bilan import qilib ko'ramiz*/

package main

import (
	"fmt"
	"math"
)

//import "fmt" alohida import qilsa ham bo'ladi
//import "math" faqat vaqt ko'p ketganligi uchun jamlab import qilsangiz yaxshiroq

func main() {
	Javob := math.Sqrt(25)
	fmt.Println("25 dan ildiz olsak, javobi : ", Javob, "bo'ladi.")
}


/*Siz biror narsani export qilmoqchi bo'lsangiz o'sha so'zni bosh harf bilan boshlashingiz kerak
masalan Pi ni export qilib ko'ramiz. Pi taqriban 3.14 ga teng.*/

package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(math.Pi)
}


/*Funksiya o'zida hech qanday argumentni jamlamasligi yoki bir nechta argumentlarni jamlashi mumkin
Ushbu misolda "qushish" o'ziga ikkita parametrni qabul qilib oladi*/

package main

import "fmt"

func qushish(x int, y int) int {
	return x + y
}//(x int, y int) o'rniga (x, y int) ni qo'llasaningiz ham bo'ladi

func main() {
	fmt.Println(qushish(42, 13))//42+13 ning javobini ekranda chiqaradi
}


/*Funksiyalar bir nechta natijalarni ham bir vaqtda qaytara oladi.
quyidagi funksiya berilgan so'zlarni joyini alishtirib beradi*/

package main

import "fmt"

func almashtir(x, y string) (string, string) {
	return y, x
}
func main() {
	a, b := almashtir("salom", "o'rganuvchilarga")
	fmt.Println(a, b)
}



/*"var" operatori o'zgaruvchilar ro'yhatini tanishtirishda ishlatiladi.
Tur nomi ro'yhatdan keyin keladi.
*/

package main

import "fmt"

var c, uy, moshina bool

func main() {
	var i int
	fmt.Println(i, c, uy, moshina)
}


/*Turlarni o'zgartirish
*/

package main

import (
	"fmt"
	"math"
)

func main() {
	var x, y int = 3, 4 //int - butun sonlar
	var f float64 = math.Sqrt(float64(x*x + y*y))//int dan float64 turiga o'tdi
	var z int = int(f)//float64 dan yana int ga o'tdi
	fmt.Println(x, y, z)//bunda x, y int; z float64 dan yana int ga o'tdi
}


/*O'zgarmaslar (const)*/

package main

import "fmt"

const Pi = 3.14 //bu yerda Pi o'zgarmas son bo'lib kelmoqda

func main() {
	const Dunyo = "Dunyo"
	fmt.Println("Salom", Dunyo)
	fmt.Println("Pi ning taqribiy qiymati", Pi, "ga teng")

	const Tugri = true
	fmt.Println("To'g'rimi?", Tugri)
}


