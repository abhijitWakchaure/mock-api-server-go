package main

import (
	"flag"
	"fmt"
	"net"
	"net/http"
	"os"
	"os/exec"
	"os/signal"
	"runtime"
	"syscall"
	"time"

	"github.com/abhijitWakchaure/mock-api-server-go/db"
	"github.com/abhijitWakchaure/mock-api-server-go/mylogger"
	"github.com/abhijitWakchaure/mock-api-server-go/user"
	"github.com/gorilla/mux"
)

var port *int
var c = &user.Controller{
	Users: db.MockUserData,
}

func main() {
	quit := make(chan bool)
	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, syscall.SIGKILL, syscall.SIGHUP, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGTERM)
	port = flag.Int("p", 8080, "Change the default port for the server")
	help := flag.Bool("h", false, "Print the help section")
	flag.Parse()
	if *help {
		printHelpText()
		return
	}
	checkPortUsed(*port)
	userRouter := mux.NewRouter()
	userRouter.HandleFunc("/users", loggingHandler(c.ListUsers)).Methods("GET")
	userRouter.HandleFunc("/users", loggingHandler(c.CreateUser)).Methods("POST")
	userRouter.HandleFunc("/users/{id}", loggingHandler(c.ReadUser)).Methods("GET")
	userRouter.HandleFunc("/users/{id}", loggingHandler(c.UpdateUser)).Methods("PUT")
	userRouter.HandleFunc("/users/{id}", loggingHandler(c.DeleteUser)).Methods("DELETE")
	go func() {
		if err := http.ListenAndServe(fmt.Sprintf(":%d", *port), userRouter); err != nil {
			mylogger.ErrorLog("Unexpected server error occurred: ", err.Error())
			quit <- true
		}
	}()
	// OS Signal Handler
	go func() {
		fmt.Printf("\nReceived signal: %v\n", <-sigChan)
		quit <- true
	}()
	mylogger.InfoLog("Mock API Server [v1.0.0] started on port %d", *port)
	for {
		select {
		case <-quit:
			mylogger.InfoLog("Shutting down the server...")
			return
		default:
			// …
		}
	}
}

func executeCommand(cmd *exec.Cmd) {
	var waitStatus syscall.WaitStatus
	if err := cmd.Run(); err != nil {
		if err != nil {
			os.Stderr.WriteString(fmt.Sprintf("Error: %s\n", err.Error()))
		}
		if exitError, ok := err.(*exec.ExitError); ok {
			waitStatus = exitError.Sys().(syscall.WaitStatus)
			mylogger.InfoLog("Error during killing process using port %d(exit code: %s)", *port, []byte(fmt.Sprintf("%d", waitStatus.ExitStatus())))
		}
	} else {
		waitStatus = cmd.ProcessState.Sys().(syscall.WaitStatus)
		mylogger.InfoLog("Successfully killed the process using port %d (exit code: %s)", *port, []byte(fmt.Sprintf("%d", waitStatus.ExitStatus())))
	}
}

func checkPortUsed(port int) {
	timeout := time.Second
	conn, _ := net.DialTimeout("tcp", fmt.Sprintf(":%d", port), timeout)

	if conn != nil {
		defer conn.Close()
		mylogger.InfoLog("Port %d is already in use, trying to kill the process...", port)
		if runtime.GOOS == "windows" {
			command := fmt.Sprintf("(Get-NetTCPConnection -LocalPort %d).OwningProcess -Force", port)
			executeCommand(exec.Command("Stop-Process", "-Id", command))
		} else {
			command := fmt.Sprintf("lsof -i tcp:%d | grep LISTEN | awk '{print $2}' | xargs kill -9", port)
			executeCommand(exec.Command("bash", "-c", command))
		}
	}
}

func printHelpText() {
	helpText := `
---------------------------------------------------------------------------
                             Mock API Server
---------------------------------------------------------------------------

Exposed APIs:

Method: [   GET]	Path: [/users]      
Method: [  POST]	Path: [/users]       
Method: [   GET]	Path: [/users/{id}]  
Method: [   PUT]	Path: [/users/{id}]  
Method: [DELETE]	Path: [/users/{id}]

Schema for User:

{
	"id": "60624180893d170927d32e27",
	"username": "john@example.com",
	"password": "EQWMJYq40spmT#g",
	"fullname": "John Doe",
	"mobile": "+91 9999999999",
	"createdAt": 1538919475135,
	"modifiedAt": 1599340945571,
	"blocked": false,
	"roles": [
		"ROLE_USER"
	]
}
`
	fmt.Println(helpText)
}

func loggingHandler(handler http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		handler.ServeHTTP(w, r)
		mylogger.InfoLog(
			"[%s] [%s] %s",
			fmt.Sprintf("%6s", r.Method),
			r.RequestURI,
			time.Since(start))
	})
}
