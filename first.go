


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

