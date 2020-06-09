package main
import "fmt"
import "time"

func main(){
  i :=2
  fmt.Print("Write", i , "as")
  switch i {
  case 1:
    fmt.println("one")
  case 2:
    fmt.println("two")
  case 3:
      fmt.println("three")
  default:
      fmt.println("its a weekday")

  }

}
