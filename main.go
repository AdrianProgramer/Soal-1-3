package module_quiz

import "fmt"

func Bola(){
	var kotak = [3]int{3, 1 ,2}

	fmt.Println(kotak[1])
	fmt.Println(kotak[2])
	fmt.Println(kotak[0])
}

func Lemari(p int, l int, t int){
	lemari := p * l * t
	fmt.Println("Volume lemari adalah")
	fmt.Print(lemari)
}

func Lingkaran(r float64, l float64) {
	phi := 3.14
	luas := phi * r * r
	fmt.Println("Luas lingkaran adalah ", luas)
}

func Perlengkapan(){
	var perlengkapan = [4]string{"ikat pingang", "topi", "dasi", "sepatu"}
	var orang = [2][4]string{
		{"ikat pingang", "topi", "dasi", "sepatu"}, 
		{"ikat pingang", "topi", "", "sepatu"},
	}
	// fmt.Println(orang)

	for _, c := range orang {
		fmt.Println(c)
		if c == perlengkapan {
			fmt.Println("Sesuai")
		} else {
			fmt.Println("Tidak Sesuai")
		}
	}
}

func Soal1(x, y float64){
	fmt.Scan(&x, &y)

	f := 1/(3*x*x+10) + 10*y + 7
	fmt.Println(f)
}

func Soal2(x, y float64){
	fmt.Scan(&x)
	f := (x*x*x + 3*x) / (x*x*x*x - 3*x*x + 4)
	fmt.Println(f)
}

func Soal3(x int64){
	var d1, d2, d3 int64

	fmt.Scan(&x)

	if x <= 999 {
		d1 = (x / 100 % 10)
		d2 = (x /10 % 10)
		d3 = (x % 10)
		fmt.Println(d1, d2, d3)
	}
}