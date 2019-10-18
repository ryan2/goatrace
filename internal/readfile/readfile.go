package readfile

import (
	"fmt"
	"os"
	"bufio"
)

func Readfile(){
	if len(os.Args)<2{
		fmt.Println("Need to pass a filename!")
		return
	}

	file, err := os.Open(os.Args[1])
	if err!=nil{
		fmt.Println("Can't open file")
		return
	}
	defer file.Close()

	fmt.Println("Rendering ", os.Args[1])
	scanner := bufio.NewScanner(file)
	var line string
	for scanner.Scan(){
		line = scanner.Text()
		if len(line)<1 || line[0]=='#'{
			continue
		}
		fmt.Println(line)
	}


	if err := scanner.Err();err!=nil{
		fmt.Println("Reading Error")
		return
	}
}


