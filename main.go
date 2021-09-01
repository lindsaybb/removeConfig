package main

import (
        "bufio"
        "fmt"
        "io"
        "os"
        "strings"
)

func main() {
        var skip bool
        var list []string
        for {
                line := readFromStdin()
                // two empty lines will signify the end
                if line == "" {
                        if skip {
                                break
                        } else {
                                skip = true
                                continue
                        }
                }
                skip = false
                list = append(list, parseLine(line))
        }

        for _, l := range list {
                fmt.Printf("%s\n", l)
        }

}

func parseLine(line string) string {
        //fmt.Println(line)
        str := strings.Fields(line)
        switch {
        case str[0] == "interface":
                return line
        case str[0] == "onu":
                return "no onu serial-number"
        case str[0] == "service-profile":
                return fmt.Sprintf("no service-profile %s", str[1])
        case str[0] == "description":
                return "no description"
        case str[0] == "exit":
                return "exit"
        default:
                return line
        }
}

func readFromStdin() string {
        reader := bufio.NewReaderSize(os.Stdin, 1024*1024)
        a, _, err := reader.ReadLine()
        if err == io.EOF {
                return ""
        } else if err != nil {
                panic(err)
        }

        return strings.TrimRight(string(a), "\r\n")
}
