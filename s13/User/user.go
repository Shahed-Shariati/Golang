package user

import "errors"

func Sum(num1, num2 int) (int, error) {
	if num1 < 1 {

		//panic("number not valid")
		return 0, errors.New("number not valid")
	} else {
		return num1 + num2, nil
	}
}
