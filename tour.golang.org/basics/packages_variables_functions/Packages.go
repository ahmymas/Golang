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
