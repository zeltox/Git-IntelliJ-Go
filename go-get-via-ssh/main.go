package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
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
		log.Fatalf("Invalid remote Git repository\n")
	} else {

		*git_username = args[num : num+index]
		// ("git_username " + *git_username)

		//fmt.Println(num)

		if len(args[num+index+1:]) > len(".git") && args[(args_len-len(".git")):args_len] == ".git" {

			*git_repo_name = args[num+index+1:]
			//fmt.Println("git_repo_name " + *git_repo_name)
		} else if num == 11 {

			*git_repo_name = args[num+index+1:] + ".git"
			fmt.Println("git_repo_name " + *git_repo_name)
		} else {
			log.Fatalf("Invalid remote Git repository, not ending with .git\n")
		}

	}

}

func main() {
	//	git@github.com:user_name/repo_name.git
	//	https://github.com/user_name/repo_name.git
	//	github.com/user_name/repo_name

	args := os.Args

	git_username := ""
	git_repo_name := ""

	no_of_args := len(args)
	if no_of_args > 2 {
		log.Fatalf("More than 1 arguments passed\n")
	} else if no_of_args == 2 {

		//fmt.Println("Before Trimming White Space: \"" + args[1] + "\"")
		args[1] = strings.TrimSpace(args[1])
		//fmt.Println("After Trimming White Space: \"" + args[1] + "\"")

		args_len := len(args[1])

		if args_len == 0 {

			log.Fatalf("Blank argument passed\n")

		} else {

			if args_len > 15 && args[1][:15] == "git@github.com:" {

				check_git_username_and_repo(args[1], args_len, 15, &git_username, &git_repo_name)

			} else if args_len > 19 && args[1][:19] == "https://github.com/" {

				check_git_username_and_repo(args[1], args_len, 19, &git_username, &git_repo_name)

			} else if args_len > 11 && args[1][:11] == "github.com/" {

				check_git_username_and_repo(args[1], args_len, 11, &git_username, &git_repo_name)

			} else {
				log.Fatalf("Invalid remote Git repository\n")
			}

			if git_username != "" && git_repo_name != "" {

				git_ssh_url := "git@github.com:" + git_username + "/" + git_repo_name

				cmd_name := "git" //
				cmd_args := []string{"ls-remote",
					git_ssh_url}
				cmd1 := exec.Command(cmd_name, cmd_args...)
				output1, err := cmd1.CombinedOutput()

				Reader := bytes.NewReader(output1) // output is slice of bytes
				scanner := bufio.NewScanner(Reader)

				if err != nil {
					for scanner.Scan() {
						line := strings.TrimSpace(scanner.Text())
						fmt.Println(line)
					}
					log.Fatalf("cmd.Run() failed with %s\n", err)
				} else {

					path := strings.TrimSpace(os.Getenv("GOPATH"))

					if path == "" {
						log.Fatalf("GOPATH environmental variable not set or blank %s\n", err)
					} else {

						_, err := os.Stat(path)
						if err != nil {
							os.IsNotExist(err)
							{
								log.Fatalf("GOPATH environmental variable does not point to a valid directory %s\n", err)
							}
						} else {

							complete_path := path + "\\" + "src" + "\\" + "github.com" + "\\" + git_username + "\\" + git_repo_name[:len(git_repo_name)-4]

							err := os.MkdirAll(complete_path, 0755)
							if err != nil {
								panic(err)
							} else {

								cmd_name := "git"
								cmd_args := []string{"clone",
									"--progress",
									git_ssh_url,
									complete_path,
								}
								cmd1 := exec.Command(cmd_name, cmd_args...)

								output1, err := cmd1.CombinedOutput()

								Reader := bytes.NewReader(output1) // output is slice of bytes
								scanner := bufio.NewScanner(Reader)

								if err != nil {
									for scanner.Scan() {
										line := strings.TrimSpace(scanner.Text())
										fmt.Println(line)
									}
									log.Fatalf("cmd.Run() failed with %s\n", err)
								} else {
									for scanner.Scan() {
										line := strings.TrimSpace(scanner.Text())
										fmt.Println(line)
									}
								}

							}

						}

					}

				}

			}

		}

	} else {
		log.Fatalf("No arguments passed\n")
	}

}
