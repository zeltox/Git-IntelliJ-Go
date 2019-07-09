package main

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

func check_git_username_and_repo(args string, args_len int, num int, git_username *string, git_repo_name *string) {

	index := -1
	for i, r := range args[num:] {
		if string(r) == "/" {
			index = i
			break
		}
	}

	if index == -1 || index == 0 {
		fmt.Println(errors.New("Invalid remote Git repository"))
	} else {

		*git_username = args[num : num+index]
		fmt.Println("git_username " + *git_username)

		//if args_len > num+index+1 {
		fmt.Println(num)

		if len(args[num+index+1:]) > len(".git") && args[(args_len-len(".git")):args_len] == ".git" {

			fmt.Println("")
			*git_repo_name = args[num+index+1:]
			fmt.Println("git_repo_name " + *git_repo_name)
		} else if num == 11 {

			*git_repo_name = args[num+index+1:] + ".git"
			fmt.Println("git_repo_name " + *git_repo_name)
		} else {

			fmt.Println(errors.New("Invalid remote Git repository, not ending with .git"))
		}

		//}else{
		//fmt.Println(errors.New("Invalid remote Git repository"))
		//}
		//
		// fmt.Println("git_repo_name " + args[1][15+index+1:])

	}

}

func main() {
	//	git@github.com:gopi25/Git-IntelliJ-Go.git
	//	https://github.com/gopi25/Git-IntelliJ-Go.git
	//	github.com/gopi25/Git-IntelliJ-Go
	//tonystark
	//ironman-suite
	args := os.Args
	//
	git_username := ""
	git_repo_name := ""

	no_of_args := len(args)
	if no_of_args > 2 {
		fmt.Println(errors.New("More than 1 arguments passed"))
	} else if no_of_args == 2 {

		fmt.Println("Before Trimming White Space: \"" + args[1] + "\"")
		args[1] = strings.TrimSpace(args[1])
		fmt.Println("After Trimming White Space: \"" + args[1] + "\"")

		args_len := len(args[1])
		//fmt.Println(args[1]
		if args_len == 0 {

			fmt.Println(errors.New("Blank argument passed"))

		} else {

			if args_len > 15 && args[1][:15] == "git@github.com:" {

				check_git_username_and_repo(args[1], args_len, 15, &git_username, &git_repo_name)

			} else if args_len > 19 && args[1][:19] == "https://github.com/" {

				check_git_username_and_repo(args[1], args_len, 19, &git_username, &git_repo_name)

			} else if args_len > 11 && args[1][:11] == "github.com/" {

				check_git_username_and_repo(args[1], args_len, 11, &git_username, &git_repo_name)

			} else {
				fmt.Println(errors.New("Invalid remote Git repository"))
			}

		}

	} else {
		fmt.Println(errors.New("No arguments passed"))
	}

}
