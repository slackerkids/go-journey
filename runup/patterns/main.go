package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"sync"
)

// 1. Dependency injection
type User struct{}

type UserRepo interface {
	Get(id int) (User, error)
}

type UserService struct {
	repo UserRepo
}

func NewUserService(r UserRepo) *UserService {
	return &UserService{repo: r}
}

// 2. Factory
type Logger struct{}

func NewJsonLogger() Logger {
	return Logger{}
}

func NewConsoleLogger() Logger {
	return Logger{}
}

func NewLogger(env string) Logger {
	if env == "prof" {
		return NewJsonLogger()
	}
	return NewConsoleLogger()
}

// 3. Strategy
type SortStrategy interface {
	Sort([]int)
}

type Validator func(string) error

// 4. Adapter
type PaymentProvider interface {
	Pay(amount int) error
}

// Implementation of interface.
func Pay(amount int) error {
	return nil
}

// 5. Singleton
var (
	logger *Logger
	once   sync.Once
)

func GetLogger() *Logger {
	once.Do(func() {
		logger = &Logger{}
	})
	return logger
}

// Testing functions
func Sum(x, y int) int {
	return x + y
}

func IsEven(d int) bool {
	if d%2 == 0 {
		return true
	}
	return false
}

// Api endpoints
func hello(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Hello")
}

type Req struct {
	Name string `json:"name"`
}

type Resp struct {
	Message string `json:"message"`
}

func handler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	var req Req
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid json", http.StatusBadRequest)
		return
	}

	resp := Resp{Message: "Hello " + req.Name}
	json.NewEncoder(w).Encode(resp)
}

func main() {
	// mux := http.NewServeMux()
	// mux.HandleFunc("POST /hello", handler)

	// log.Fatal(http.ListenAndServe(":8080", mux))

	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("opening file: %v", err)
	}
	defer file.Close()

	sc := bufio.NewScanner(file)
	for sc.Scan() {
		line := sc.Text()
		line = strings.ToLower(line)
		fmt.Println(line)
	}

	if err := sc.Err(); err != nil {
		log.Fatalf("scanner error: %v", err)
	}
}
