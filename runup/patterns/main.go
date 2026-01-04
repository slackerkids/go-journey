package main

import (
	"fmt"
	"sync"
)

// Singleton design pattern - thread safe
type Config struct {
	AppName string
}

var (
	instance *Config
	once     sync.Once
)

func GetConfigInstance() *Config {
	once.Do(func() {
		instance = &Config{
			AppName: "Example App",
		}
	})
	return instance
}

// One more example
type Logger struct{}

var (
	logger *Logger
	onceL  sync.Once
)

func GetLogger() *Logger {
	once.Do(func() {
		logger = &Logger{}
	})
	return logger
}

// Factory design pattern
type Shape interface {
	Draw()
}

type Circle struct{}

func (c Circle) Draw() { fmt.Println("Circle") }

type Square struct{}

func (s Square) Draw() { fmt.Println("Square") }

func NewShape(t string) Shape {
	switch t {
	case "circle":
		return Circle{}
	case "square":
		return Square{}
	default:
		return nil
	}
}

func main() {
	c1 := GetConfigInstance()
	c2 := GetConfigInstance()

	fmt.Println(c1 == c2)
	fmt.Println(c1.AppName)

	square := NewShape("square")

	square.Draw()
}
