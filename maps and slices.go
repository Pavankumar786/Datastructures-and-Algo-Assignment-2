package main 
import "fmt"
func main() { 
    m_a_p := map[int]string{ 
  
        90: "Dog", 
        91: "Cat", 
        92: "Cow", 
        93: "Bird", 
        94: "Rabbit", 
    } 
  
    // Iterating map using for rang loop 
    for id, pet := range m_a_p { 
  
        fmt.Println(id, pet) 
    } 
} 
$go run main.go
94 Rabbit
90 Dog
91 Cat
92 Cow
93 Bird