/**
 * @Author: sonic
 * @File:  main.go
 * @Date: 2020/8/12 16:32
 * @Description:
 */
package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	template := "<type>(<scope>): <subject>\n\n<body>\n\n<footer>"

	template = readType(template)
	template = readScope(template)
	template = readSubject(template)
	template = readBody(template)
	template = readFooter(template)

	cmd := exec.Command("git", "commit", "-m", template)
	output, _ := cmd.CombinedOutput()

	fmt.Printf("%s\n",output)
}

func readType(t string) (msg string) {
	//fmt.Println("【TYPE-必选】：")
	//fmt.Println("【0】feat：新功能（feature）")
	//fmt.Println("【1】fix：修补bug")
	//fmt.Println("【2】docs：文档（documentation）")
	//fmt.Println("【3】style：格式（不影响代码运行的变动）")
	//fmt.Println("【4】refactor：重构（即不是新增功能，也不是修改bug的代码变动）")
	//fmt.Println("【5】test：增加测试")
	//fmt.Println("【6】build：（以前称chore）构建系统（涉及脚本、配置或工具）和包依赖项相关的开发更改")
	//fmt.Println("【7】perf：性能提升相关的更改")
	//fmt.Println("【8】vendor：更新依赖项、包的版本")
	//fmt.Print("【填入选择的代码】: ")

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
	for  {
		fmt.Print("-> ")
		scanner.Scan()
		line := scanner.Text()
		lines = append(lines, line)

		if len(lines) == 1{
			if strings.TrimSpace(lines[0]) == ""{
				break
			}
		}

		if strings.TrimSpace(line) == "eof"{
			lines = lines[:len(lines)-1]
			break
		}
	}

	for _, v := range lines{
		input = input + "\n" +v
	}

	if strings.TrimSpace(input) == ""{
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
	for  {
		fmt.Print("-> ")
		scanner.Scan()
		line := scanner.Text()
		lines = append(lines, line)

		if len(lines) == 1{
			if strings.TrimSpace(lines[0]) == ""{
				break
			}
		}

		if strings.TrimSpace(line) == "eof"{
			lines = lines[:len(lines)-1]
			break
		}
	}

	for _, v := range lines{
		input = input + "\n" +v
	}

	if strings.TrimSpace(input) == ""{
		return t
	}

	r := strings.Replace(t, "<footer>", strings.TrimSpace(input), 1)

	return r
}