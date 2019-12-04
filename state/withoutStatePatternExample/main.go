package main

import "fmt"

var currentState = 0

//四种状态，0代表电梯门打开，1代表电梯门关闭，2代表电梯在运行，3代表电梯停止（电梯门没有打开）
var OPENED = 0
var CLOSED = 1
var RUNNING = 2
var STOPPED = 3

func stateChange(nextState int) {
	if currentState == OPENED {
		if nextState == OPENED {
			fmt.Println("电梯已经打开")
		} else if nextState == CLOSED {
			fmt.Println("关闭电梯门")
			currentState = CLOSED
		}
	}
}
