package hello

import "fmt"


func hello(){
  var a =1
  for a < 15 {
    if a ==5 {
      a + = a
       continue
    }
    a++
    if a > 10 {
      break
    }
  }
  if a ==11 {
    goto END
  }

  END:
    println("End")

}
