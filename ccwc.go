package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"sync"
	"unicode/utf8"
)

func main() {

	// need go routines because file Opening should be concurrent. 
	// Else it may interfere in other function's results when no flag is given and 
	// all the functions access the file at the same time
	var wg sync.WaitGroup;

	wg.Add(4);


	reader := bufio.NewReader(os.Stdin);
	input, err := reader.ReadString('\n');

	if err != nil {
		fmt.Println("Command not found");
		return;
	}

	inputArray := strings.Split(input, " ");

	command := inputArray[0];
	flag := ""
	if len(inputArray) == 3 {
		flag = inputArray[1];
	}
	fileName := ""
	if len(inputArray) == 3 {
		fileName = strings.TrimSpace(inputArray[2]);
	} else {
		fileName = strings.TrimSpace(inputArray[1]);
	}

	if command == "ccwc" {
		if flag == "-c" || flag == "" {
			go flagC(fileName, &wg);
		}

		if flag == "-l" || flag == "" {
			go flagL(fileName, &wg);
		}

		if flag == "-w" || flag == ""{
			go flagW(fileName, &wg);
		}

		if flag == "-m" || flag == ""{
			go flagM(fileName, &wg);
		}
	}
	wg.Wait();
	fmt.Println("All go routines completed");

}

func flagL(fileName string, wg *sync.WaitGroup) {
	defer wg.Done();

			file, err := os.Open(fileName);
			if err != nil {
				fmt.Println("Error opening the file");
			}
			defer file.Close();
			scanner := bufio.NewScanner(file)
			lineCount := 0
			for scanner.Scan() {
				lineCount++
			}

			if err := scanner.Err(); err != nil {
				fmt.Println(err)
			}
			fmt.Println(lineCount);
}

func flagW(fileName string, wg *sync.WaitGroup) {
defer wg.Done();
			file, err := os.Open(fileName);
			if err != nil {
				fmt.Println("Error opening the file");
			}
			defer file.Close();
			scanner := bufio.NewScanner(file);
			wordCount := 0;

			for scanner.Scan() {
				line := scanner.Text();
				words := strings.Fields(line);
				wordCount += len(words);
			}

			if err := scanner.Err(); err != nil {
				fmt.Println(err);
			}
			fmt.Println(wordCount);
}

func flagM(fileName string, wg *sync.WaitGroup) {
		defer wg.Done();
			file, err := os.Open(fileName);
			if err != nil {
				fmt.Println("Error opening the file");
			}
			defer file.Close();
			fileContents, err := os.ReadFile(fileName);

			if err != nil {
				fmt.Println(err)
			}

			fileString := string(fileContents);

			numOfUnicodeCharacters := utf8.RuneCountInString(fileString);

			fmt.Println(numOfUnicodeCharacters);
}

func flagC(fileName string, wg *sync.WaitGroup) {
	  defer wg.Done();
			file, err := os.Open(fileName);
			if err != nil {
				fmt.Println("Error opening the file");
			}
			defer file.Close();
			fileContents, err := os.ReadFile(fileName);

			if err != nil {
				fmt.Println(err);
				return;
			}

			numBytes := len(fileContents);

			fmt.Println(numBytes);

}