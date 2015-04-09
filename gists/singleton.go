package singleton

import (
"fmt"
)

type Singleton interface {
    SaySomething()
}

type singleton struct {
    text string
}

var oneSingleton Singleton

func NewSingleton(text string) Singleton {
    if oneSingleton == nil {
        oneSingleton = &singleton{
        text: text,
        }
    }
    return oneSingleton
}

func (this *singleton) SaySomething() {
    fmt.Println(this.text)
}