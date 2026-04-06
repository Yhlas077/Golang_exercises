// package main

// import "fmt"

// func main() {

// // Konsola cykaryar
//     fmt.Println("Hello World!")

// // berlenlerin gornusleri
//     var count int64 = 0
//     var uint64 uint64 = 100
//     if int64(uint64) > int64(count) {
//         fmt.Println("uint64 is greater than int64")
//     }
// }

////////////////////////////////////////////////////////

// package main

// import "fmt"

// func main() {
//     var a int
//     fmt.Scan(&a)
//     fmt.Println("You entered: ", a)
//     var b int
//     fmt.Scan(&b)
//     fmt.Println("You entered: ", b)
//     if a > b {
//         fmt.Println("a is greater than b")
//     } else if a < b {
//         fmt.Println("a is less than b")
//     } else {
//         fmt.Println("a is equal to b")
//     }
// }

////////////////////////////////////////////////////////

// package main

// import "fmt"

// func main() {
//     var s int
//     fmt.Print("Enter a number: ")
//     fmt.Scan(&s)
//     switch s {
//     case 1:
//         fmt.Println("You entered one")
//         fallthrough
//     case 2:
//         fmt.Println("You entered two")
// 		fallthrough
//     case 3:
//         fmt.Println("You entered three")
//     default:
//         fmt.Println("You entered a number other than 1, 2, or 3")}
//     }

////////////////////////////////////////////////////////

// package main

// import "fmt"

// func main() {
//     var a int
//     fmt.Print("San giriz: ")
//     fmt.Scan(&a)
//     for i := 1; i <= a; i += 2 {
//         fmt.Println(i)
//     }
// }

////////////////////////////////////////////////////////

// While yok dine for

// package main

// import "fmt"

// func main() {
//     var a int
//     fmt.Print("San giriz: ")
//     fmt.Scan(&a)
//     for 0 < a {
//         fmt.Println(a)
//         a--
//     }
//  }

////////////////////////////////////////////////////////

// package main

// import "fmt"

// func main() {
// 	var n int
// 	var a []int
// 	fmt.Print("San giriz: ")
// 	fmt.Scan(&n)
// 	a = make([]int, n)
// 	for i := 0; i < n; i++ {
// 		fmt.Scan(&a[i])
// 	}
// 	b := a[0:2]
// 	fmt.Println(a)
// 	fmt.Println("b:", b)
// 	c := a[2:]
// 	fmt.Println("c:", c)
// 	d := a[:3]
// 	fmt.Println("d:", d)
// 	a = append(a, b...)
// 	a = append(a, 5)
// 	fmt.Println("a:", a)
// }

// /////////////////////////////////////////////

// package main

// import (
// 	"fmt"
// 	"strings"
// 	"strconv"
// )

// func main() {
// 	s := ""
// 	fmt.Print("setir giriz: ")
// 	fmt.Scan(&s)
// 	fmt.Println("You entered:", s)
// 	str := strings.Split(s, " ")
// 	fmt.Println("Split string:", str)

// 	a := 0
// 	// b := 0
// 	var err error
// 	a, err = strconv.Atoi(s)
// 	if err != nil {
// 		fmt.Println("Error converting string to int:", err)
// 	} else {
// 		fmt.Println("String to int:", a)
// 	}
// }

// ////////////////////////////////////////////////////////

// package main

// import (
// 	"fmt"
// 	"time"
// )

// func main() {
// 	fmt.Println("Welcome to the playground!")

// 	fmt.Println("The time is", time.Weekday(time.Now().Weekday()))
// }

///////////////////////////////////////////////////////////

// package main

// import (
// 	"fmt"
// 	"math/rand"
// )

// func main() {
// 	fmt.Println("My favorite number is", rand.Intn(10))
// }

///////////////////////////////////////////////////////////

// package main

// import "fmt"

// func add(x int, y int) int {
// 	return x + y
// }

// func main() {
// 	fmt.Println(add(42, 13))
// }

///////////////////////////////////////////////////////////

// package main

// import "fmt"

// func swap(x, y string) (string, string) {
// 	return y, x
// }

// func main() {
// 	a, b:= swap("hello", "world")
// 	fmt.Println(a, b)
// }

///////////////////////////////////////////////////////////

// package main

// import "fmt"

// func split(sum int) (int, int) {
// 	x, y :=0, 0
// 	x = sum * 4 / 9
// 	y = sum - x
// 	return x, y
// }

// func main() {
// 	fmt.Println(split(17))
// }

///////////////////////////////////////////////////////////

// package main

// import "fmt"

// var c, python, java bool

// func main() {
// 	var i int
// 	fmt.Println(i, c, python, java)
// }

///////////////////////////////////////////////////////////

// package main

// import "fmt"

// var i, j int = 1, 2

// func main() {
// 	var c, python, java = true, false, "no!"
// 	fmt.Println(i, j, c, python, java)
// }

///////////////////////////////////////////////////////////

// package main

// import (
// 	"fmt"
// 	"math/cmplx"
// )

// var (
// 	ToBe   bool       = false
// 	MaxInt uint64     = 1<<64 - 1
// 	z      complex128 = cmplx.Sqrt(-5 + 12i)
// )

// func main() {
// 	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
// 	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
// 	fmt.Printf("Type: %T Value: %v\n", z, z)

// a := 1 + 2i
// b := 3 + 4i

// fmt.Println(a + b) // (4+6i)
// fmt.Println(a * b) // complex multiplication
// }

//////////////////////////////////////////////////////////////

// package main

// import "fmt"

// const Pi = 3.14

// func main() {
// 	const World = "世界"
// 	fmt.Println("Hello", World)
// 	fmt.Println("Happy", Pi, "Day")

// 	const Truth = true
// 	fmt.Println("Go rules?", Truth)

// }

/////////////////////////////////////////////////////////////////

// package main

// import "fmt"

// const (
// 	// Create a huge number by shifting a 1 bit left 100 places.
// 	// In other words, the binary number that is 1 followed by 100 zeroes.
// 	Big = 1 << 100
// 	// Shift it right again 99 places, so we end up with 1<<1, or 2.
// 	Small = Big >> 99
// )

// func needInt(x int) int { return x*10 + 1 }
// func needFloat(x float64) float64 {
// 	return x * 0.1
// }

// func main() {

// 	fmt.Println(needInt(Small))
// 	fmt.Println(needFloat(Small))
// 	fmt.Println(needFloat(Big))
// }

//////////////////////////////////////////////////////////////

// package main

// import "fmt"

// func main() {
// 	sum := 1
// 	for sum < 1000 {
// 		sum += sum
// 	}
// 	fmt.Println(sum)

// b := []byte("Hello")
// b = fmt.Append(b, " World")
// b = fmt.Appendf(b, " %d", 20)

// fmt.Println(string(b))
// }

/////////////////////////////////////////////////////////////////

// package main

// import (
// 	"fmt"
// 	"runtime"
// )

// func main() {
// 	fmt.Print("Go runs on ")
// 	switch os := runtime.GOOS; os {
// 	case "darwin":
// 		fmt.Println("macOS.")
// 	case "linux":
// 		fmt.Println("Linux.")
// 	default:
// 		// freebsd, openbsd,
// 		// plan9, windows...
// 		fmt.Printf("%s.\n", os)
// 	}
// }

//////////////////////////////////////////////////////////////

// package main

// import (
// 	"fmt"
// 	"time"
// )

// func main() {
// 	fmt.Println("When's Saturday?")
// 	today := time.Now().Weekday()
// 	switch time.Saturday {
// 	case today + 0:
// 		fmt.Println("Today.")
// 	case today + 1:
// 		fmt.Println("Tomorrow.")
// 	case today + 2:
// 		fmt.Println("In two days.")
// 	default:
// 		fmt.Println("Too far away.")
// 	}
// }

/////////////////////////////////////////////////////////////

// package main

// import "fmt"

// func main() {
// 	// defer fmt.Println("world")

// 	// fmt.Println("hello")
// 	// fmt.Println("hello2")

// }

/////////////////////////////////////////////////////

// package main

// import "fmt"

// func main() {
//     defer fmt.Println("First")
//     defer fmt.Println("Second")
//     defer fmt.Println("Third")
//     fmt.Println("Done")
// }

////////////////////////////////////////////////////////

// package main

// import "fmt"

// func printNumbers() {
//     for i := 1; i <= 5; i++ {
//         defer fmt.Println(i)
//     }
// }

// func main() {
//     printNumbers()
// }

/////////////////////////////////////////////////////////

// package main

// import "fmt"

// func main() {

// 	i := 42
// 	p := &i

// 	fmt.Println(p)
// }

////////////////////////////////////////////////////////////

// package main

// import "fmt"

// type vertex struct {
// 	X int
// 	Y int
// }

// func main() {
// 	fmt.Println(vertex{1, 2})
// }

////////////////////////////////////////////////////////////

// package main

// import "fmt"

// func main() {
// 	type Machine struct {
// 		name string
// 		age int
// 		time int8
// 	}
// 	s:=""
// 	age,time:=0,0
// 	fmt.Print("Machine name: ")
// 	fmt.Scan(&s)
// 	fmt.Print("Machine age: ")
// 	fmt.Scan(&age)
// 	fmt.Print("Machine time: ")
// 	fmt.Scan(&time)

// 	var m Machine
// 	m.name = s
// 	m.age = age
// 	m.time = int8(time)
// 	fmt.Println(m)
// }

////////////////////////////////////////////////////////////

// package main

// import "fmt"

// func main() {
// 	type Student struct {
// 		name string
// 		age  int
// 		grade float64
// 	}
// 	s := Student{name: "Alex", age: 18, grade: 12}
// 	fmt.Println(s)
// }

/////////////////////////////////////////////////////////////

// package main

// import "fmt"

// type Car struct {
// 	name string
// 	year int
// 	speed int
// }

// func main() {
// 	car :=Car{}
// 	fmt.Print("Car: ")
// 	fmt.Scan(&car.name)
// 	fmt.Print("Year: ")
// 	fmt.Scan(&car.year)
// 	fmt.Print("Speed: ")
// 	fmt.Scan(&car.speed)

// 	fmt.Printf("Year: %d\nCar: %s\nSpeed: %d\n",car.year,car.name,car.speed)
// }

/////////////////////////////////////////////////////////////

// package main

// import "fmt"

// type Book struct {
// 	title string
// 	author string
// 	pages int
// }

// func main() {
// 	book := Book{}
// 	fmt.Print("Book title: ")
// 	fmt.Scan(&book.title)
// 	fmt.Print("Author: ")
// 	fmt.Scan(&book.author)
// 	fmt.Print("Pages: ")
// 	fmt.Scan(&book.pages)
// 	fmt.Printf("Title: %sAuthor: %sPages: %d", book.title, book.author, book.pages)
// }

//////////////////////////////////////////////////////////////

// package main

// import (
// 	"fmt"
// )

// func Sqrt(x int) int {
// 	z := 1
// 	for i := 0; i <= 10; i++ {
// 		z -= (z*z - x) / (2 * z)
// 		fmt.Printf("Iteration %d: ",  z)
// 	}
// }

// func main() {
// 	var x int
// 	fmt.Scan(&x)
// 	Sqrt(int(x))
// }

//////////////////////////////////////////////////////////////////

// package main

// import "fmt"
// import "time"

// func add(n int)	int {
// 	if(n == 1) {
// 		return 1
// 	}
// 	n+=add(n-1)
// 	time.sleep(1000)

// 	return n
// }

// func main() {
// 	fmt.Println(add(5))
// }

///////////////////////////////////////////////////////////////////

// package main

// import "log"

// func reverse(s string) string {
// 	if len(s) == 0 {
// 		return ""
// 	}
// 	k := ""
// 	k += reverse(s[1:])
// 	k += string(s[0])

// 	return k
// }

// func main() {

// 	s1 := reverse("Hello, World!")
// 	log.Println(s1)
// }

// ///////////////////////////////////////////////////////////////

// package main

// import "fmt"

// type Employee struct {
// 	ID     int
// 	Name   string
// 	Age    int
// 	Salary float64
// }

// type Manager struct {
// 	Employees []Employee
// }

// // AddEmployee adds a new employee to the manager's list.
// func (m *Manager) AddEmployee(e Employee) {

// }

// // RemoveEmployee removes an employee by ID from the manager's list.
// func (m *Manager) RemoveEmployee(id int) {
// 	// TODO: Implement this method
// }

// // GetAverageSalary calculates the average salary of all employees.
// func (m *Manager) GetAverageSalary() float64 {
// 	// TODO: Implement this method
// 	return 0
// }

// // FindEmployeeByID finds and returns an employee by their ID.
// func (m *Manager) FindEmployeeByID(id int) *Employee {
// 	// TODO: Implement this method
// 	return nil
// }

// func main() {
// 	manager := Manager{}
// 	manager.AddEmployee(Employee{ID: 1, Name: "Alice", Age: 30, Salary: 70000})
// 	manager.AddEmployee(Employee{ID: 2, Name: "Bob", Age: 25, Salary: 65000})
// 	manager.RemoveEmployee(1)
// 	averageSalary := manager.GetAverageSalary()
// 	employee := manager.FindEmployeeByID(2)

// 	fmt.Printf("Average Salary: %f\n", averageSalary)
// 	if employee != nil {
// 		fmt.Printf("Employee found: %+v\n", *employee)
// 	}
// }

////////////////////////////////////////////////////////////////////////////////////////

// package main

// import "fmt"

// func main() {
// 	defer fmt.Println("world")

// 	fmt.Println("hello")
// 	fmt.Println("hello")
// 	fmt.Println("hello")
// 	fmt.Println("hello")
// 	fmt.Println("hello")
// 	fmt.Println("hello")
// }

// package main

// import "fmt"

// func main() {
// 	k := 42

// 	         // point to i
// 	fmt.Println(k) // read i through the pointer
// 	// *p = 21         // set i through the pointer
// 	// fmt.Println(i)  // see the new value of i

// 	// p = &j         // point to j
// 	// *p = *p / 37   // divide j through the pointer
// 	// fmt.Println(j) // see the new value of j
// }

///////////////////////////////////////////////////////////////////////////////

// package main

// import "fmt"

// type Vertex struct {
// 	Lat, Long float64
// }

// var m = map[string]Vertex{
// 	"Bell Labs": {40.68433, -74.39967},
// 	"Google":    {37.42202, -122.08408},
// }

// func main() {
// 	fmt.Println(m)
// }

/////////////////////////////////////////////////////////////////////////////////
// package main

// import "fmt"

// func add(n int) int {
// 	if(n==1) {
// 	return 1
// 	}
// 	n += add(n-1)
// 	return n
// }

// func main() {
// 	k:=0
// 	fmt.Scan(&k)
// 	s:= add(k)
// 	fmt.Printf("1-den %d cenli sanlaryn jemi: %d ", k, s)
// }

///////////////////////////////////////////////////


/////////////////////////////////////////////////////////////

/
//////////////////////////////////////////////////////////////////////////////////

