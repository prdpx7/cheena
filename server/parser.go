package server

import "strings"

func (iptr *Interpreter) Parser(input string) {
	inputArr := strings.Split(input, " ")
	for i := 0; i < len(inputArr); i++ {
		switch i {
		case 0:
			iptr.Cmd = strings.ToUpper(inputArr[0])
		case 1:
			iptr.Key = inputArr[1]
		case 2:
			iptr.Value = inputArr[2]
		case 3:
			iptr.OptionalArgument = strings.ToUpper(inputArr[3])
		}
	}

}
