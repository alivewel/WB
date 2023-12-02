package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"syscall"
	"time"

	"github.com/shirou/gopsutil/process"
)

func main() {
	args := os.Args

	ExecuteCommand(args)
}

func ExecuteCommand(args []string) {
	if len(args) > 1 {
		nameUtility := args[1]
		switch nameUtility {
		case "cd":
			cd(args)
		case "pwd":
			pwd(args)
		case "echo":
			echo(args)
		case "kill":
			kill(args)
		case "ps":
			ps()
		case "fork":
			fork(args)
		}
	}
}

// директория меняется внутри программы, но в терминале остается прежней
func cd(args []string) {
	if len(args) == 3 {
		newDir := args[2]
		if newDir == ".." { // команда cd ..
			currentDir, _ := os.Getwd()
			newDir = filepath.Dir(currentDir)
		}
		err := os.Chdir(newDir)
		if err != nil {
			fmt.Println("Ошибка при переходе в директорию:", err)
			os.Exit(1)
		}
	}
}

func pwd(args []string) {
	if len(args) == 2 {
		currentDir, err := os.Getwd()
		if err != nil {
			fmt.Println("Ошибка при переходе в директорию:", err)
			os.Exit(1)
		}
		fmt.Println(currentDir)
	}
}

func echo(args []string) {
	if len(args) > 2 {
		for i, arg := range args[2:] {
			fmt.Printf(arg)
			if i != len(args[2:])-1 {
				fmt.Printf(" ")
			}
		}
		fmt.Printf("\n")
	}
}

func kill(args []string) {
	numArgs := len(args)
	if numArgs == 3 || numArgs == 4 {
		pid, err := strconv.Atoi(args[numArgs-1])
		if err != nil {
			fmt.Println("kill: illegal pid:", args[2])
			os.Exit(1)
		}

		_, err = os.FindProcess(pid)
		if err != nil {
			fmt.Printf("kill: kill %d failed: no such process", pid)
			os.Exit(1)
		}

		signal := syscall.SIGTERM
		if numArgs == 4 && args[numArgs-2] == "-9" {
			signal = syscall.SIGKILL
		}
		err = syscall.Kill(pid, signal)
		if err != nil {
			fmt.Println("Error sending signal:", err)
			os.Exit(1)
		}
	}
}

func ps() {
	processList, err := process.Processes()
	if err != nil {
		fmt.Println("Ошибка при получении списка процессов:", err)
		os.Exit(1)
	}

	fmt.Printf("%-5s %-10s %-10s %s\n", "PID", "TTY", "TIME", "CMD")
	for _, p := range processList {
		cmd, err := p.CmdlineSlice()
		if err != nil {
			cmd = []string{"unknown"}
		}

		startTime, err := p.CreateTime()
		if err != nil {
			fmt.Println("Ошибка при получении времени создания процесса:", err)
			continue
		}
		elapsed := time.Since(time.Unix(0, startTime*int64(time.Millisecond))).Truncate(time.Second)

		fmt.Printf("%-5d %-10s %s %-20s\n", p.Pid, "?", elapsed, cmd)
	}
}

func fork(command []string) {
	com := exec.Command(command[0], command[1:]...)

	output, err := com.Output()
	if err != nil {
		fmt.Println("zsh: command not found:", command)
	}
	fmt.Println(string(output))
}
