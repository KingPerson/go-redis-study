package main

import (
	"fmt"
	"go-redis-study/sortedset"
)

const COMMAND_DESC = `
NAME:
   go-redis - Using golang to learn the data structure of redis

USAGE:
   go-redis [global options] command [command options] [arguments...]

VERSION:
   0.0.1

COMMANDS:
   zadd  

GLOBAL OPTIONS:
`
const VERSION = "0.0.1"

func main() {
	//fmt.Println("go-redis-study")
	//fmt.Println(COMMAND_DESC)

	level := sortedset.ZslRandomLevel()
	fmt.Println(level)
}
