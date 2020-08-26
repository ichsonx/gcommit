/**
 * @Author: sonic
 * @File:  main.go
 * @Date: 2020/8/12 16:32
 * @Description:
 */
package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

var (
	h bool
	p bool
	t string
	tm string
)

func main() {

	flag.BoolVar(&h, "h", false, "this help")
	flag.BoolVar(&p, "p", false, "push after commiting code")
	flag.StringVar(&t, "t", "", "tag the version code")
	flag.StringVar(&tm, "tm", "", "tag annotated, ignore if without -t")

	flag.Parse()

	if h {
		flag.Usage = usage
		flag.Usage()
		os.Exit(0)
	}

	template := "<type>(<scope>): <subject>\n\n<body>\n\n<footer>"
	template = readType(template)
	template = readScope(template)
	template = readSubject(template)
	template = readBody(template)
	template = readFooter(template)

	cmd := exec.Command("git", "commit", "-m", template)
	output, _ := cmd.CombinedOutput()
	fmt.Printf("%s\n", output)

	if t != "" {
		var cmd *exec.Cmd
		if tm != ""{
			cmd = exec.Command("git", "tag", "-a", t, "-m", tm)
		}else {
			cmd = exec.Command("git", "tag", t)
		}
		output, _ := cmd.CombinedOutput()

		fmt.Printf("%s\n", output)
	}

	if p {
		cmd := exec.Command("git", "push")
		output, _ := cmd.CombinedOutput()

		fmt.Printf("%s\n", output)
	}
}

func usage()  {
	fmt.Fprintf(os.Stderr, `gcommit version: v0.3.0
Usage: gcommit [-hp]

Options:
`)
	flag.PrintDefaults()
}

func readType(t string) (msg string) {
	fmt.Println("【TYPE】(Required)：")
	fmt.Println("【0】feat：new feature")
	fmt.Println("【1】fix：fix bugs")
	fmt.Println("【2】docs：documentation")
	fmt.Println("【3】style：Changes that do not affect code execution")
	fmt.Println("【4】refactor：It is neither a new feature nor a code change to modify a bug")
	fmt.Println("【5】test：Add test")
	fmt.Println("【6】build：（named chore）builds the system(scripts,configuration,code changes etc...)")
	fmt.Println("【7】perf：Performance improvement related changes")
	fmt.Println("【8】vendor：Update the version of dependency and package")
	fmt.Print("【input the code】: ")

	input := ""
	_, err := fmt.Scanln(&input)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}

	switch input {
	case "0":
		input = "feat"
	case "1":
		input = "fix"
	case "2":
		input = "docs"
	case "3":
		input = "style"
	case "4":
		input = "refactor"
	case "5":
		input = "test"
	case "6":
		input = "build"
	case "7":
		input = "perf"
	case "8":
		input = "vendor"
	default:
		fmt.Printf("Incorrect input，quit now %T\n", input)
		os.Exit(1)
	}

	msg = strings.Replace(t, "<type>", input, 1)
	return msg
}

func readScope(t string) (msg string) {
	input := ""
	fmt.Println("\n【SCOPE】(Optional) press【Enter】to skip:")
	//reader := bufio.NewReader(os.Stdin)
	//input, _ = reader.ReadString('\n')

	// 以下方式读取控制台输入，只获取文本内容。对于没输入直接【回车】，不会获取【回车】的换行符内容
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input = scanner.Text()

	if strings.TrimSpace(input) == "" {
		msg = strings.Replace(t, "(<scope>)", "", 1)
	} else {
		msg = strings.Replace(t, "<scope>", input, 1)
	}

	return msg
}

func readSubject(t string) (msg string) {
	input := ""
	fmt.Println("\n【SUBJECT】(Required): ")
	// 以下方式读取控制台输入，只获取文本内容。对于没输入直接【回车】，不会获取【回车】的换行符内容
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input = scanner.Text()

	if strings.TrimSpace(input) == "" {
		fmt.Printf("Incorrect input，quit now %T\n", input)
		os.Exit(1)
	}

	msg = strings.Replace(t, "<subject>", input, 1)

	return msg
}

func readBody(t string) string {
	input := ""
	fmt.Println("\n【BODY】(Optional) press【Enter】to skip:")
	fmt.Println("input【eof】to finish")

	var lines []string

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("-> ")
		scanner.Scan()
		line := scanner.Text()
		lines = append(lines, line)

		if len(lines) == 1 {
			if strings.TrimSpace(lines[0]) == "" {
				break
			}
		}

		if strings.TrimSpace(line) == "eof" {
			lines = lines[:len(lines)-1]
			break
		}
	}

	for _, v := range lines {
		input = input + "\n" + v
	}

	if strings.TrimSpace(input) == "" {
		return t
	}

	r := strings.Replace(t, "<body>", strings.TrimSpace(input), 1)

	return r
}

func readFooter(t string) string {
	input := ""
	fmt.Println("\n【FOOTER】(Optional) press【Enter】to skip:")
	fmt.Println(" input【eof】to finish:")
	fmt.Println(" notice:If there is no upward compatible, input【BREAKING CHANGE】:")

	var lines []string

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("-> ")
		scanner.Scan()
		line := scanner.Text()
		lines = append(lines, line)

		if len(lines) == 1 {
			if strings.TrimSpace(lines[0]) == "" {
				break
			}
		}

		if strings.TrimSpace(line) == "eof" {
			lines = lines[:len(lines)-1]
			break
		}
	}

	for _, v := range lines {
		input = input + "\n" + v
	}

	if strings.TrimSpace(input) == "" {
		return t
	}

	r := strings.Replace(t, "<footer>", strings.TrimSpace(input), 1)

	return r
}
