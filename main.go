package main

import (
	"bufio"
	"fmt"
	"os"
	"time"

	"errors"

	"strings"

	"github.com/dixonwille/wmenu"
)

const password = "coulson"
const rootDir = "/opt/shield/"

func main() {
	fmt.Println("S.H.I.E.L.D. Terminal Client v1.2.3434b")
	fmt.Println(".................................................")
	fmt.Print("Enter password: ")

	pass := ""
	count := 0
	for {
		fmt.Scanln(&pass)
		if strings.ToLower(pass) == password {
			break
		}
		count++
		if count < 3 {
			fmt.Print("Incorrect, please try again: ")
		} else {
			fmt.Print("Hint(your favorite agent): ")
		}

	}
	printFile(rootDir+"boot.txt", 30*time.Millisecond)

	file, _ := os.Open(rootDir + "shield.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
		time.Sleep(100 * time.Millisecond)
	}
	fmt.Println("Welcome to S.H.I.E.L.D. Agent Anastasia")
	menu := wmenu.NewMenu("What would you like to do?")
	//menu.Action(func(opts []wmenu.Opt) error { fmt.Printf(opts[0].Text + " is your favorite food."); return nil })
	menu.Option("Messages", nil, true, checkMessages)
	menu.Option("Agent Status", nil, false, agentStatus)
	menu.Option("Exit", nil, false, func(wmenu.Opt) error { return errors.New("exit") })
	var err error
	for err == nil {
		err = menu.Run()
	}
}

func checkMessages(wmenu.Opt) error {
	printFile(rootDir+"messages.txt", 50*time.Millisecond)
	return nil
}
func agentStatus(wmenu.Opt) error {
	printFile(rootDir+"agents.txt", 80*time.Millisecond)
	return nil
}

func printFile(filename string, speed time.Duration) {
	file, _ := os.Open(filename)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		data := scanner.Text()
		if data == "BREAK" {
			fmt.Println("--hit any key to continue--")
			bufio.NewReader(os.Stdin).ReadByte()
		} else {

			for _, c := range data {
				fmt.Print(string(c))
				time.Sleep(speed)
			}
			fmt.Print("\n")
		}

	}

}
