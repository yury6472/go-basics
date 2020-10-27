package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Student struct {
	LastName   string
	FirstName  string
	MiddleName string
	Birthday   string
	Address    string
	Phone      string
	Rating     []int
}

type User struct {
	ID       int
	Number   string
	Year     int
	Students []Student
}

type Res struct {
	Average float64
}

func main() {
	var u1 User

	var res Res

	file, _ := os.Open("data1.json")
	dt, _ := ioutil.ReadAll(file)

	//data, err := ioutil.ReadAll(os.Stdin)

	//fmt.Println(json.Valid(dt))

	//var data interface{}
	err := json.Unmarshal(dt, &u1)

	if err != nil {
		panic(err)
	}

	//Print json
	//fmt.Println(u1)

	//fmt.Println(u1.Students[1].Rating[2])

	i := len(u1.Students)

	//allaver := 0

	var avernum float64 = 0.0
	var k int = 0

	for j := 0; j < i; j++ {

		k = len(u1.Students[j].Rating)
		//aver := 0

		avernum = avernum + float64(k)

		/* for n := 0; n < k; n++ {
			//u1.Students[j].Rating[n]
			aver = aver + u1.Students[j].Rating[n]
		}
		fmt.Println("Средний бал ", j, "ого =", aver/k)
		*/

		//allaver = allaver + aver/k
	}

	//fmt.Println("Средний бал =", allaver/i)

	//s := fmt.Sprintf("%.1f", avernum/float64(i))
	//fmt.Println(avernum / float64(i))
	res.Average = avernum / float64(i)

	//fmt.Println(s)

	jsondata, _ := json.MarshalIndent(res, "", "    ")
	fmt.Println(string(jsondata))
}
