package main
import (
    "fmt"
    "slices"
    )
    
func main() {
    var done = true
    var userChoice int
    var group string
    var groups []string
    var toDoItems= make([]map[string]string,0)
    //var toDoItem = make(map[string]string,0)
    var userInputItem string
    fmt.Println("Welcome to Your to-do List")
    for done {
        fmt.Println("Please Choose one of the below\n1)Add Group\n2)Add item to existing Group\n3)You are done\n4)Print Current To Do List")
        fmt.Scan(&userChoice)
        if userChoice == 1 {
            fmt.Println("Please Enter Group Name: ")
            fmt.Scan(&group)
            groups = append(groups, group)
            fmt.Println(groups)
        }else if userChoice == 2 {
            fmt.Println(groups)
            fmt.Println("Please Choose Existing Group Name from The Above: ")
            fmt.Scan(&group)
            if ! slices.Contains(groups, group){
                fmt.Println("You Choose non-existing Group, Please Try again")
            } else {
                fmt.Println("Please Enter Your Item:")
                fmt.Scan(&userInputItem)
                var toDoItem = make(map[string]string,0)
                toDoItem[group]=userInputItem
                //fmt.Println(toDoItem)
                toDoItems=append(toDoItems,toDoItem)
                //fmt.Println(toDoItems)
            }
        } else if userChoice == 3 {
            fmt.Println("You are Done\nExiting...\n ")
            fmt.Println(toDoItems)
            break
        } else if userChoice == 4{
            fmt.Println(toDoItems)
        }else{
            fmt.Println("Wrong Choice, Please Try again\n ")
        }
        
    }
} 
