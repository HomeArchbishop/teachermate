package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/homearchbishop/teachermate-auto/internal/client/robot"
	"github.com/homearchbishop/teachermate-auto/internal/client/winapi"
	"github.com/homearchbishop/teachermate-auto/internal/client/wsclient"
	"github.com/homearchbishop/teachermate-auto/internal/shared"
	"github.com/spf13/viper"
)

type Config struct {
	Client struct {
		Server string
	}
}

var config Config

var (
	lessonId  string
	studentId string
)

func init() {
	flag.StringVar(&lessonId, "lesson", "", "Lesson ID")
	flag.StringVar(&studentId, "student", "", "Student ID")
}

func main() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}
	if err := viper.Unmarshal(&config); err != nil {
		log.Fatalf("Unable to decode into struct, %v", err)
	}

	flag.Parse()

	reader := bufio.NewReader(os.Stdin)
	for lessonId == "" {
		fmt.Print("Enter lesson ID: ")
		lessonId, _ = reader.ReadString('\n')
		lessonId = strings.TrimSpace(lessonId)
		if lessonId == "" {
			fmt.Println("lesson ID cannot be empty. Please retry.")
		}
	}
	for studentId == "" {
		fmt.Print("Enter student ID: ")
		studentId, _ = reader.ReadString('\n')
		studentId = strings.TrimSpace(studentId)
		if studentId == "" {
			fmt.Println("studentId cannot be empty. Please retry.")
		}
	}

	wsclient.RegisterHandler("sign_signal", func(msg *shared.SignSignalMsgType) {
		doSignOnce(msg.SignUrl)
	})

	// this is a blocking call
	wsclient.StartClient(config.Client.Server, lessonId, studentId)

}

func doSignOnce(signUrl string) {
	hwnd, err := winapi.FindWindow("", "文件传输助手")
	if err != nil {
		fmt.Println("FindWindow error:", err)
		return
	}

	winapi.ForcePosition(hwnd, 0, 0, 1000, 800)

	robot.OperateOnce(
		signUrl,
		robot.NewPos(20, 750),
		robot.NewPos(931, 773),
		robot.NewPos(880, 773),
	)
}
