/******************************************************************************

                            Online Go Lang Compiler.
                Code, Compile, Run and Debug Go Lang program online.
Write your code in this editor and press "Run" button to execute it.

*******************************************************************************/


package main
import "fmt"

func main() {
    var day int 
    fmt.Print("Enter The Day Number:")
    fmt.Scanln(&day)
    switch(day){
        case 1:fmt.Println("Monday")
         case 2:fmt.Println("Tuesday")
          case 3:fmt.Println("Wednesday")
           case 4:fmt.Println("Thursday")
            case 5:fmt.Println("Friday")
             case 6:fmt.Println("Saturday")
              case 7:fmt.Println("Sunday")
               default:fmt.Println("Syntax Error!")
    }
}
