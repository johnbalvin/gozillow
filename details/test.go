package details

import (
	"encoding/json"
	"log"
	"os"
)

func Test() {
	test1()
}
func test0() {
	body, err := os.ReadFile("./me.html")
	if err != nil {
		log.Println("err: ", err)
		return
	}
	data, err := ParseBodyDetails(body)
	if err != nil {
		log.Println("err: ", err)
		return
	}
	f, _ := os.Create("./house.json")
	json.NewEncoder(f).Encode(data)
}
func test1() {
	houseURL := "https://www.zillow.com/homedetails/858-Shady-Grove-Ln-Harrah-OK-73045/339897685_zpid/"
	data, err := FromPropertyURL(houseURL, nil)
	if err != nil {
		log.Println("err: ", err)
		return
	}
	f, _ := os.Create("./house.json")
	json.NewEncoder(f).Encode(data)
}
