package main

import (
	"fmt"
	"xp-study/DesignPatternsForGo/SingletonPattern/singleton"
)

func main() {
	instance := singleton.GetInstance()
	fmt.Println(&instance)
	instance = singleton.GetInstance()
	fmt.Println(&instance)
}
