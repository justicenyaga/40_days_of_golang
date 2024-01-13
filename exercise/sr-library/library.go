//--Summary:
//  Create a program to manage lending of library books.
//
//--Requirements:
//* The library must have books and members, and must include:
//  - Which books have been checked out
//  - What time the books were checked out
//  - What time the books were returned
//* Perform the following:
//  - Add at least 4 books and at least 3 members to the library
//  - Check out a book
//  - Check in a book
//  - Print out initial library information, and after each change
//* There must only ever be one copy of the library in memory at any time
//
//--Notes:
//* Use the `time` package from the standard library for check in/out times
//* Liberal use of type aliases, structs, and maps will help organize this project

package main

import (
	"fmt"
	"time"
)

type Name string
type Title string

type LendAudit struct {
	checkOut time.Time
	checkIn  time.Time
}

type Member struct {
	name  Name
	books map[Title]LendAudit
}

type BookEntry struct {
	lent  int
	total int
}

type Library struct {
	members map[Name]Member
	books   map[Title]BookEntry
}

func printMemberAudits(member *Member) {
	for title, audit := range member.books {
		var returnTime string
		if audit.checkIn.IsZero() {
			returnTime = "[not returned yet]"
		} else {
			returnTime = audit.checkIn.String()
		}
		fmt.Println(member.name, ":", title, ":", audit.checkOut.String(), "through", returnTime)
	}
}

func printMembersAudits(library *Library) {
	for _, member := range library.members {
		printMemberAudits(&member)
	}
}

func printLibraryBooks(library *Library) {
	fmt.Println()
	for title, book := range library.books {
		fmt.Println(title, "/ lent", book.lent, "/ total", book.total)
	}
	fmt.Println()
}

func checkoutBook(library *Library, title Title, member *Member) bool {
	book, found := library.books[title]
	if !found {
		fmt.Println("Book not part of the library")
		return false
	}

	if book.lent == book.total {
		fmt.Println("No more books to lend")
		return false
	}

	book.lent += 1
	library.books[title] = book

	member.books[title] = LendAudit{checkOut: time.Now()}
	return true
}

func returnBook(library *Library, title Title, member *Member) bool {
	book, found := library.books[title]
	if !found {
		fmt.Println("Book not part of the library")
		return false
	}

	audit, found := member.books[title]
	if !found {
		fmt.Println("Member did not checkout this book")
		return false
	}

	book.lent -= 1
	library.books[title] = book

	audit.checkIn = time.Now()
	member.books[title] = audit
	return true
}

func main() {
	library := Library{members: make(map[Name]Member), books: make(map[Title]BookEntry)}

	library.books["Webapps in Go"] = BookEntry{total: 4, lent: 0}
	library.books["Go for Networks"] = BookEntry{total: 7, lent: 0}
	library.books["Go for Desktop Apps"] = BookEntry{total: 2, lent: 0}

	library.members["Jason"] = Member{"Jason", make(map[Title]LendAudit)}
	library.members["Caleb"] = Member{"Caleb", make(map[Title]LendAudit)}

	fmt.Println("\nInitial:")
	printMembersAudits(&library)
	printLibraryBooks(&library)

	member := library.members["Caleb"]
	checkedOut := checkoutBook(&library, "Go for Networks", &member)
	fmt.Println("\nCheckout a book")
	if checkedOut {
		printMembersAudits(&library)
		printLibraryBooks(&library)
	}

	returned := returnBook(&library, "Go for Networks", &member)
	fmt.Println("\nReturn a book")
	if returned {
		printMembersAudits(&library)
		printLibraryBooks(&library)
	}
}
