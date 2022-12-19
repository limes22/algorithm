package main

import "fmt"

type PhoneBook struct {
	Name   string
	Number int
	PhBook *PhoneBook
}

func newFriedPhoneInfo(name string, number int, phBook *PhoneBook) PhoneBook {
	return PhoneBook{
		Name:   name,
		Number: number,
		PhBook: phBook,
	}
}

func main() {
	initPhoneBook := newFriedPhoneInfo("osj", 1234, nil)
	onePhoneBook := newFriedPhoneInfo("osj1", 1, &initPhoneBook)
	twoPhoneBook := newFriedPhoneInfo("osj2", 2, &onePhoneBook)
	threePhoneBook := newFriedPhoneInfo("osj3", 3, &twoPhoneBook)

	phoneBook := &threePhoneBook
	var frontAddrBook *PhoneBook
	for {
		if phoneBook.Name == "pjh2" {
			frontAddrBook.PhBook = phoneBook.PhBook
			break
		}
		if phoneBook.PhBook == nil {
			break
		}
		frontAddrBook = phoneBook
		phoneBook = phoneBook.PhBook
	}

	curPhoneBook := &threePhoneBook
	for {
		fmt.Println(curPhoneBook.Name)
		fmt.Println(curPhoneBook.Number)
		if curPhoneBook.PhBook == nil {
			break
		}
		curPhoneBook = curPhoneBook.PhBook
	}
}
