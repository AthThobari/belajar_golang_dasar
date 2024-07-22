package main

import (
	"fmt"
)

func test(){
	fmt.Println("===========Konversi type data (1)=========")
	var nilai32 int32 = 32768
	var nilai64 int64 = int64(nilai32)
	
	var nilai16 int16 = int16(nilai32)
	
	fmt.Println(nilai32)
	fmt.Println(nilai64)
	fmt.Println(nilai16)

	fmt.Println("===========Konversi type data (2)=========")
	var name = "Agaton"
	var e = name[0]
	var eString = string(e)
	
	fmt.Println(name)
	fmt.Println(eString)

	fmt.Println("===========Kode program type declaration=========")
	type NoKTP string
	
	var KTPAgaton NoKTP = "123456"
	fmt.Println(KTPAgaton)
	fmt.Println(NoKTP("11111111"))
	
	fmt.Println("===========Kode program operasi matematika=========")
	var a = 10
	var b = 10
	var c = a + b
	fmt.Println(c)

	fmt.Println("===========Kode program augmented assignment=========")
	var i = 0
	i += 10
	fmt.Println(i)

	fmt.Println("===========Kode program unary operator=========")
	var j = 1
	j++
	j++
	fmt.Println(j)

	fmt.Println("===========Kode program perbandingan=========")
	var name1  = "Agaton"
	var name2 = "Agaton"
	var result bool = name1 == name2
	fmt.Println(result)

	fmt.Println("===========Kode program operasi boolean=========")
	var nilaiAkhir = 90
	var absensi = 80

	var lulusNilaiAkhir bool = nilaiAkhir > 80
	var lulusAbsensi bool = absensi > 80

	var lulus bool = lulusNilaiAkhir && lulusAbsensi

	fmt.Println(lulus)

	fmt.Println("===========Kode program array=========")
	var nama [3]string
	nama[0] = "Agaton"
	nama[1] = "kakon"
	nama[2] = "Anopheles"

	fmt.Println(nama[0])
	fmt.Println(nama[1])
	fmt.Println(nama[2])

	var values = [3]int{
		12,
		15,
		6,
	}
	fmt.Println(values)
	fmt.Println(len(values))
	values[0] = 100
	fmt.Println(values)

	fmt.Println("===========Kode program slice=========")
	names := [...]string{"Agaton", "Eko", "Kuriniawan", "Khanedy", "Kakon", "Anopheles"}
	slice := names[4:6]
	
	fmt.Println(slice[0])
	fmt.Println(slice[1])
	
	fmt.Println("===========Kode program append slice=========")
	var days = [...]string{
		"senin",
		"selasa",
		"rabu",
		"kamis",
		"jum'at",
		"sabtu",
		"minggu",
	}

	daySlice1 := days[5:]
	daySlice1[0] = "sabtu baru"
	daySlice1[1] = "minggu baru"

	fmt.Println(days)

	daySlice2 := append(daySlice1, "buka bersama")
	daySlice2[0] = "ups"
	fmt.Println(daySlice2)
	fmt.Println(days)

	fmt.Println("===========Kode program make slice=========")
	newSlice := make([]string, 3, 5)
	newSlice[0] = "Agaton"
	newSlice[1] = "Agaton"
	newSlice[2] = "Agaton"

	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))

	fmt.Println("===========Kode program copy slice=========")
	fromSlice := days[:]
	toSlice := make([]string, len(fromSlice), cap(fromSlice))
	
	copy(toSlice, fromSlice)
	
	fmt.Print(toSlice)

	fmt.Println("\n===========Kode program array vs slice=========")
	iniArray := [...]int{1,2,3,4,5}
	iniSlice := []int{1,2,3,4,5}

	fmt.Println(iniArray)
	fmt.Println(iniSlice)

	fmt.Println("===========Kode program map=========")
	person := map[string]string{
		"name" : "Agaton",
		"address" : "Sukoharjo",
	}
	
	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["address"])

	book := make(map[string]string)
	book["title"] = "Belajar Golang"
	book["author"] = "Eko"
	book["wrong"] = "ups"

	delete(book, "wrong")
	fmt.Println(book)

	fmt.Println("===========Kode program if statement=========")
	namaString := "joko"

	if namaString == "eko" {
		fmt.Println("hello eko")
	} else if namaString == "joko" {
		fmt.Println("Hallo, jok")
	} else {
		fmt.Println("Anda siapa?")
	}

	fmt.Println("===========Kode program if short statement=========")
	namaString = "Giyan"
	
	if length := len(namaString); length > 5{
		fmt.Println("nama terlalu panjang")
	} else {
		fmt.Println("nama anda sudah benar")
	}

	fmt.Println("===========Kode program switch statement=========")
	namaString = "Agaton"

	switch namaString {
	case "Eko":
		fmt.Println("Halo, Eko")
	case "Agaton":
		fmt.Println("Halo, Agaton")
	default:
		fmt.Println("Halo, Siapa ya?")
	}

	fmt.Println("===========Kode program switch short statement=========")
	namaString = "Agaton"

	switch length := len(namaString); length > 5{
	case true:
		fmt.Println("Nama anda sudah benar")

	case false:
		fmt.Println("Nama anda kurang panjang")
	}

	fmt.Println("===========Kode program switch tanda kondisi=========")

	namaString = "Agaton"
	length := len(namaString)
	switch {
	    case length > 5:
			fmt.Println("nama lumayan panjang")
			
	    case length > 10:
			fmt.Println("nama terlalu panjang")

		default:
			fmt.Println("nama sudah benar")
	}

	fmt.Println("===========Kode program for=========")

	counter := 1

	for counter <= 10{
		fmt.Println("perulangan ke ", counter)
		counter++
	}
	// Kode Program for Dengan Statement
	for counter = 1; counter <= 10; counter++{
		fmt.Println("Perulangan dengan statement ke", counter)
	}

	// Kode Program for Range

	namesRange := []string{"Eko", "Kurniawan", "Khanedy"}
	for index, name := range namesRange{
		fmt.Println("index", index, "=", name )
	}

	// break Dan continue

	for i := 0; i <= 10; i++{
		fmt.Println("Perulangan ke", i)
		if i == 5 {
			break
		}
	}
	for i := 0; i <= 10; i++{

		if i%2 == 0 {
			continue
		}
		fmt.Println("Perulangan ke", i)
	}

}