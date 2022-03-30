package main

import "fmt"

func main() {
	//Insert Code here
	var data [7]string
	// activity 1
	data[0] = "Operating System List"
	data[1] = "Windows XP"
	data[2] = "Linux 1.0"
	data[3] = "Raspbian Teensy"
	data[4] = "MacOS"
	data[5] = "IOS"
	data[6] = "Google Android"

	fmt.Println(len(data))

	// activity 2
	fmt.Println(data)
	data[1] = "Windows 10"
	data[2] = "Linux 16.04"
	data[3] = "Raspbian Buster"

	fmt.Println(data)

	// activity 3
	var new_data [10]string
	for i, value := range data {
		new_data[i] = value
	}

	new_data[7] = "Ubuntu"
	new_data[8] = "MS-Dos"
	new_data[9] = "Solaris"

	fmt.Println(new_data)

	// activity 4
	fmt.Println(new_data[:3])
	fmt.Println(new_data[3:6])
	fmt.Println(new_data[len(new_data)-3:])

}
