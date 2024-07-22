package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func ChangeAddressCountryToIndonesia(address *Address) {
	address.Country = "Indonesia"
}

type Man struct{
	Name string
}

func (man *Man) Married() {
	man.Name = "Mr. " + man.Name

}
func main2() {
	//Kode program pointer di method
	eko := Man{"eko"}
	eko.Married()
	fmt.Println(eko)
	//kode program pointer di function
	address := Address{"Subang", "Jawa Barat", ""}
	ChangeAddressCountryToIndonesia(&address)

	fmt.Println(address)

	//Kode dengan operator *
	address1 := Address{"Sukoharjo", "Jawa Tengah", "Indonesia"}
	address2 := &address1

	address2.City = "Bandung"
	*address2 = Address{"Jakarta", "DKI Jakarta", "Indonesia"}

	fmt.Println(address1)
	fmt.Println(address2)

	//Kode program function new
	alamat1 := new(Address)
	alamat2 := alamat1

	alamat2.Country = "Indonesia"

	fmt.Println(alamat1)
	fmt.Println(alamat2)



}
