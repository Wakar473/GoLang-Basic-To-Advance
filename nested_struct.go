package main

import "fmt"
//nested struct is a struct type that contains another struct type as one of its fields.
//This allows you to create more complex data structures by composing multiple structs together. 
// Define an inner struct
type Address struct {
    Street  string
    City    string
    ZipCode string
}

// Define an outer struct containing the inner struct
type Person struct {
    Name    string
    Age     int
    Address Address // Nested struct as a field
}

func main() {
    // Create a new Person instance with nested Address struct
    person := Person{
        Name: "John Doe",
        Age:  30,
        Address: Address{
            Street:  "123 Main Street",
            City:    "New York",
            ZipCode: "10001",
        },
    }

    // Accessing fields of the nested structs
    fmt.Println("Name:", person.Name)
    fmt.Println("Age:", person.Age)
    fmt.Println("Street:", person.Address.Street)
    fmt.Println("City:", person.Address.City)
    fmt.Println("ZipCode:", person.Address.ZipCode)
}
