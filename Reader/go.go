package Reader

import (
	"fmt"
    "os"
    "io"	
	"bufio"
)

type Hero struct {
	name string
	xp int

}
	var hero = Hero{}

func ReadFile(r string){

	file,err := os.Open(r)
	if err != nil{
		fmt.Println("Ошибка в текстовом файле!")
		os.Exit(1)
	}
	defer file.Close()
	data := make([]byte, 64)
    for{
        n, err := file.Read(data)
        if err == io.EOF{   // если конец файла
            break           // выходим из цикла
        }
        fmt.Print(string(data[:n]))
    }
}

func ReadNumberFromMenu(){
	var number int16
	fmt.Print("Введите пункт меню: ")
	fmt.Scan(&number)
	switch number{
	case 1:
		FirstRaund()
	case 2: 
		fmt.Println("Settings")
	case 3: 
		fmt.Println("exit")
	default:
		fmt.Println("Empty")
	}
}


func NameHero() string{
	fmt.Println("Введите ваше имя: ")
	in := bufio.NewScanner(os.Stdin)
	in.Scan()
if err := in.Err(); err != nil {
    fmt.Fprintln(os.Stderr, "Ошибка ввода:", err)
}
	hero.name = in.Text()
return hero.name
}

func FirstRaund(){
	fmt.Println("Приветствую тебя ", NameHero())
	ReadFile("Greeting.txt")


}

