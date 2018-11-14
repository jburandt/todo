package cmd

import (
	"bufio"
	"bytes"
	"fmt"
	"github.com/spf13/cobra"
	"io"
	"log"
	"os"
	"os/user"
	"strings"
)

func init() {
	RootCmd.AddCommand(addCmd)
}

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "add item to list",
	Long:  "Use to add new item to list",
	Run: func(cmd *cobra.Command, args []string) {
		usr, err := user.Current()
		if err != nil {
			log.Fatal(err)
		}

		homeDir := usr.HomeDir
		dirCheck := strings.Join([]string{homeDir, "/.todo"}, "")
		configFile := strings.Join([]string{dirCheck, "/config"}, "")

		if CheckExists(configFile) == false {
			CreateConfig(dirCheck)
		}

		output, err := ReadConfig(configFile)
		if err != nil {
			fmt.Println("Could not read config file")
			return
		}

		if len(os.Args) < 3 {
			fmt.Println("Please specify something to add.")
			fmt.Println("todo add order a pizza")
			os.Exit(1)
		}
		i := os.Args[2:]
		s := strings.Join(i, " ")
		newoutput := append(output, s)
		WriteConfig(newoutput, configFile)


		// debug
		//f, err := os.Open(configFile)
		//if err != nil {
		//	fmt.Printf("Error opening file: %v\n", err)
		//	os.Exit(1)
		//}
		//counter, err := lineCounter(f)
		//fmt.Printf("DEBUG: %v\n", newoutput)
		//fmt.Printf("DEBUG: Line count is: %v\n", counter)

	},
}

func CheckExists(name string) bool {
	if _, err := os.Stat(name); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

func CreateConfig(file string) {
	config := strings.Join([]string{file, "/config"}, "")

	fmt.Printf("The file %s does not exist\n", config)
	fmt.Print("Would you like to create it? (yes/no)? ")

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		input := scanner.Text()
		input = strings.ToLower(strings.TrimSpace(input))
		if input == "yes" || input == "y" {
			os.MkdirAll(file, 0755)
			os.Create(config)
			return
		} else if input == "no" || input == "n" {
			return
		}
		fmt.Printf("Please type 'yes' or 'no': ")
	}
}

func ReadConfig(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func WriteConfig(lines []string, path string) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	w := bufio.NewWriter(file)
	for _, line := range lines {
		fmt.Fprintln(w, line)
	}
	return w.Flush()
}

func lineCounter(r io.Reader) (int, error) {

	var readSize int
	var err error
	var count int
	// create slice of bytes to hold read
	buf := make([]byte, 1024)

	for {
		// use read method on new slice
		readSize, err = r.Read(buf)
		if err != nil {
			break
		}

		var buffPosition int
		for {
			// search line for the break \n
			i := bytes.IndexByte(buf[buffPosition:], '\n')
			// break if -1 err, or the length of bytes read from Read method (readSize) matches the bytes from buffPosition
			if i == -1 || readSize == buffPosition {
				break
			}
			buffPosition += i + 1
			count++
		}
	}
	if readSize > 0 && count == 0 || count > 0 {
		count++
	}
	if err == io.EOF {
		return count, nil
	}

	return count, err
}
