package main

import "fmt"

func main() {
	oddEven()
	fmt.Println("")
	gradeScore()
	fmt.Println("")
	showNumbers()
	fmt.Println("")
	fizzbuzz()
	fmt.Println("")
	removeDuplicatesmain()
}

func oddEven() {
	fmt.Println("Angka Genap Ganjil")
	fmt.Print("Masukkan angka : ")
	var n int
	fmt.Scanln(&n)

	if n%2 == 0 {
		fmt.Println(n, "adalah angka genap")
	} else {
		fmt.Println(n, "adalah angka ganjil")
	}
}

func gradeScore() {
	fmt.Println("Konversi Nilai dari Skor")
	fmt.Print("Masukkan nilai (0 - 4) : ")
	var score float32
	fmt.Scanln(&score)

	if score > 3.66 && score <= 4 {
		fmt.Println("Nilai A")
	} else if score > 3.33 && score <= 3.66 {
		fmt.Println("Nilai A-")
	} else if score > 3 && score <= 3.33 {
		fmt.Println("Nilai B+")
	} else if score > 2.66 && score <= 3 {
		fmt.Println("Nilai B")
	} else if score > 2.33 && score <= 2.66 {
		fmt.Println("Nilai B-")
	} else if score > 2 && score <= 2.33 {
		fmt.Println("Nilai C+")
	} else if score > 1.66 && score <= 2 {
		fmt.Println("Nilai C")
	} else if score > 2.33 && score <= 1.66 {
		fmt.Println("Nilai C-")
	} else if score > 1 && score <= 1.33 {
		fmt.Println("Nilai D+")
	} else if score > 0 && score <= 1 {
		fmt.Println("Nilai D")
	} else if score == 0 {
		fmt.Println("Nilai E")
	} else {
		fmt.Println("Masukkan nilai yang tepat")
	}
}

func showNumbers() {
	fmt.Println("Tampilkan angka antara 10 - 50")
	for i := 11; i <= 50; i++ {
		if i == 50 {
			fmt.Print(i)
		} else {
			fmt.Print(i, ", ")
		}
	}
}

func fizzbuzz() {
	fmt.Println("")
	fmt.Println("FizzBuzz")
	for i := 1; i <= 20; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("Fizzbuzz")
		} else if i%3 == 0 {
			fmt.Println("Fizz")
		} else if i%5 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(i)
		}
	}
}

func removeDuplicates(arr []string) []string {
	words_string := map[string]bool{}
	for i := range arr {
		words_string[arr[i]] = true
	}
	desired_output := []string{}
	for j := range words_string {
		desired_output = append(desired_output, j)
	}
	return desired_output
}
func removeDuplicatesmain() {
	arr := []string{"ha", "ha", "ha", "lucu", "lucu", "sekali"}
	fmt.Println(arr)
	desired_output := removeDuplicates(arr)
	fmt.Println(desired_output)
}
