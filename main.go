package main

import (	
"./Reader"

)



func main(){

//StartGame()
Reader.ParsingTask()
}


func StartGame(){

Reader.ReadFile("start.txt")
Reader.ReadFile("Settings.txt")
Reader.ReadNumberFromMenu()
}