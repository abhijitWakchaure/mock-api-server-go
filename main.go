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
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

var port *int

// VERSION ...
const VERSION = "v1.0.4"

func main() {
	flag.Usage = printHelpText
	quit := make(chan bool)
	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, syscall.SIGHUP, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGTERM)
	port = flag.Int("port", 8080, "Change the default port for the server")
	endpointFlag := flag.String("endpoint", "/api/users", "Change the default endpoint for the server")
	help := flag.Bool("help", false, "Print the help section")
	version := flag.Bool("version", false, "Print the server version")
	flag.Parse()
	if *help {
		printHelpText()
		os.Exit(0)
	}
	if *version {
		printVersion()
		os.Exit(0)
	}
	endpoint := *endpointFlag
	if endpoint[0] != '/' {
		mylogger.ErrorLog("API endpoint must start with a slash ('/')")
		os.Exit(1)
	}
	if endpoint[len(endpoint)-1] != '/' {
		endpoint = endpoint + "/"
	}
	checkPortUsed(*port)
	db.InitDB()
	c := &user.Controller{
		Users: db.MockUserData,
	}
	mylogger.InfoLog("Server port is set to %d", *port)
	mylogger.InfoLog("Server endpoint is set to [%s]", endpoint)
	userRouter := mux.NewRouter().StrictSlash(true)
	userRouter.HandleFunc("/api/users", c.ListUsers).Methods("GET")
	userRouter.HandleFunc(endpoint, c.CreateUser).Methods("POST")
	userRouter.HandleFunc(endpoint+"{id}", c.ReadUser).Methods("GET")
	userRouter.HandleFunc(endpoint+"{id}", c.UpdateUser).Methods("PUT")
	userRouter.HandleFunc(endpoint+"{id}", c.DeleteUser).Methods("DELETE")

	allowedOrigins := handlers.AllowedOrigins([]string{"*"})
	allowedMethods := handlers.AllowedMethods([]string{"GET", "POST", "OPTIONS", "DELETE", "PUT"})
	allowedHeaders := handlers.AllowedHeaders([]string{"Accept", "Accept-Encoding", "Access-Control-Allow-Headers", "Access-Control-Allow-Origin", "Access-Control-Request-Headers", "Access-Control-Request-Method", "Allow", "Authorization", "Connection", "Content-Length", "Content-Type", "Forwarded", "Keep-Alive", "Origin", "Proxy-Authenticate", "Proxy-Authorization", "Referer", "User-Agent", "X-CSRF-Token", "X-Forwarded-For", "X-Requested-With", "X-Total-Count"})

	go func() {
		handler := handlers.CombinedLoggingHandler(os.Stdout, handlers.CORS(allowedHeaders, allowedMethods, allowedOrigins)(userRouter))
		handler = myCORSHandler(handler)
		if err := http.ListenAndServe(fmt.Sprintf(":%d", *port), handler); err != nil {
			mylogger.ErrorLog("Unexpected server error occurred: ", err.Error())
			quit <- true
		}
	}()
	// OS Signal Handler
	go func() {
		mylogger.InfoLog("\nReceived signal: %v\n", <-sigChan)
		quit <- true
	}()
	t := time.Now().Format("02/Jan/2006 15:04:05 -0700")
	mylogger.InfoLog("Mock API Server [%s] is started at %s", VERSION, t)
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
			mylogger.ErrorLog("Error during killing process using port %d(exit code: %s)", *port, []byte(fmt.Sprintf("%d", waitStatus.ExitStatus())))
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

func printVersion() {
	fmt.Printf("Mock API Server:\n")
	fmt.Printf("Github: %s\n", "https://github.com/abhijitWakchaure/mock-api-server-go")
	fmt.Printf("Version: %s\n", VERSION)
}

func printHelpText() {
	helpText := `
---------------------------------------------------------------------------
                             Mock API Server
---------------------------------------------------------------------------

Usage:
    	
    -endpoint string
        Change the default endpoint for the server (default "/api/users")
    -port int
        Change the default port for the server (default 8080)
    -version
        Print the server version
    -help
        Print the help section


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

func loggingHandlerFunc(handlerFunc http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		handlerFunc.ServeHTTP(w, r)
		mylogger.InfoLog(
			"[%s] [%s] %s \t %s",
			fmt.Sprintf("%6s", r.Method),
			r.RequestURI,
			r.RemoteAddr,
			time.Since(start))
	})
}

func loggingHandler(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		handler.ServeHTTP(w, r)
		mylogger.InfoLog(
			"[%s] [%s] %s \t %s",
			fmt.Sprintf("%6s", r.Method),
			r.RequestURI,
			r.RemoteAddr,
			time.Since(start))
	})
}

func myCORSHandler(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// if r.Method == "OPTIONS" {
		// 	fmt.Println("Got OPTIONS request")
		// 	// w.WriteHeader(http.StatusOK)
		// 	// return
		// }
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Vary", "Origin")
		w.Header().Set("Vary", "Accept-Encoding")
		w.Header().Set("Vary", "Access-Control-Request-Method")
		w.Header().Set("Vary", "Access-Control-Request-Headers")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Origin, Accept, token")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS, PUT, DELETE")
		handler.ServeHTTP(w, r)
	})
}
