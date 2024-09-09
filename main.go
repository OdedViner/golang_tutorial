package main

import (
	"fmt"
	"os"
	"strings"
	"time"

	"simple_app/stemmer"
)

type Role string

const (
	Viewer    Role = "viewer"
	Developer Role = "developer"
	Admin     Role = "admin"
)

type User struct {
	Login string
	Role  Role
}

func Promote(u *User, r Role) {
	u.Role = r
}

func OpenFileCheck(file_path string) (string, error) {
	file, err := os.Open(file_path)
	var out string
	if err != nil {
		out = file_path
	} else {
		out = "Success"
	}
	defer file.Close()
	return out, err
}

type Location struct {
	X float64
	Y float64
}

func NewLocation(x float64, y float64) (Location, error) {
	if x > 10 || y > 10 {
		return Location{}, fmt.Errorf("invalid Loc")
	}
	loc := Location{
		X: x,
		Y: y,
	}
	return loc, nil

}

func (l *Location) Move(x float64, y float64) {
	l.X = x
	l.Y = y
}

type Car struct {
	ID string
	Location
}

func NewCar(id string, l Location) (Car, error) {
	loc, err := NewLocation(1, 2)
	if err != nil {
		return Car{}, fmt.Errorf("no availbel location")
	}
	return Car{id, loc}, nil

}

type Mover interface {
	Move(float64, float64)
}

func MoveAll(items []Mover, x float64, y float64) {
	for _, item := range items {
		item.Move(x, y)
	}
}

func ParallelFunc(n int) {
	time.Sleep(100 * time.Millisecond)
	fmt.Printf("goroutine %d\n", n)
}
func ChannelEx() {
	ch := make(chan int)

	// Start a goroutine to send data
	go func() {
		ch <- 1 // Send 99 to the channel
	}()

	// Receive data from the channel
	val := <-ch // Receive from the channel
	fmt.Printf("got %d\n", val)
}
func main() {
	url := "https://www.linkedin.com"
	print(stemmer.CheckURL(url))
	tok := stemmer.Stem("oded")
	fmt.Print(tok)
	ChannelEx()
	// for i := 0; i < 5; i++ {
	// 	go ParallelFunc(i)
	// }
	// time.Sleep(2 * time.Second)

	loc, err := NewLocation(7, 9)
	fmt.Printf("%#v", loc)
	if err != nil {
		fmt.Println(err)
		return
	}
	car, err := NewCar("fiat1", loc)
	if err != nil {
		fmt.Println("error")
	}
	fmt.Printf("%#v", car)
	loc.Move(6, 5)
	fmt.Printf("%#v", loc)
	arr_objs := []Mover{&loc, &car}
	MoveAll(arr_objs, 4.4, 5.5)
	fmt.Printf("%#v", arr_objs)
	for _, arr_obj := range arr_objs {
		fmt.Printf("%#v", arr_obj)
	}

	d := Location{1, 2}
	fmt.Println(d.X)
	out, err := OpenFileCheck("oded.txt")
	if err != nil {
		fmt.Printf("Err")
	} else {
		fmt.Println(out)
	}
	u := User{"elliot", Viewer}
	fmt.Printf("%#v\n", u)
	Promote(&u, Admin)
	fmt.Printf("%#v\n", u)

	s := "hellw my name is oded viner oded ccc oded ocfr oded"
	words := strings.Split(s, " ")
	map_words := make(map[string]int)
	for _, word := range words {
		map_words[word] = map_words[word] + 1
	}
	max := 0
	com_word := ""
	for k, v := range map_words {
		if v > max {
			max = v
			com_word = k
		}
	}
	fmt.Print(com_word + "\n")
	a := CollatzStep(6)
	fmt.Printf("%d\n", a)
	var b, c string
	b, c = SplitFileName("odedpy")
	fmt.Printf("%s,   %s\n", b, c)
	val, err := ExErrOdd(2)
	if err == nil {
		fmt.Printf("val is not odd %d ", val)
	} else {
		fmt.Printf("err: %s", err)
		return
	}
	fmt.Println("Succees")

}

func CollatzStep(n int) int {
	var num int
	if n%2 == 0 {
		num = n / 2
	} else {
		num = n * 2
	}
	return num
}

func SplitFileName(file_name string) (string, string) {
	index := strings.Index(file_name, ".")
	if index == -1 {
		return file_name, ""
	} else {
		return file_name[:index], file_name[index+1:]
	}
}

func ExErrOdd(num int) (int, error) {
	if num%2 == 0 {
		return num * 2, nil
	} else {
		return num, fmt.Errorf("Erorr")
	}
}
