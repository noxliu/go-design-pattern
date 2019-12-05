package main

import "fmt"

var currentState = 0

//四种状态，0代表电梯门打开，1代表电梯门关闭，2代表电梯在运行，3代表电梯停止（电梯门没有打开）
var OPEN = 0
var CLOSE = 1
var RUN = 2
var STOP = 3

func stateChange(nextState int) {
	if currentState == OPEN {
		//打开状态只能切换到关闭
		fmt.Println("关闭电梯门")
		currentState = CLOSE
	} else if currentState == CLOSE {
		//关闭状态可以切换到打开或者运行
		if nextState == OPEN {
			fmt.Println("打开电梯门")
			currentState = OPEN
		} else if nextState == RUN {
			fmt.Println("电梯开始运行")
			currentState = RUN
		}
	} else if currentState == RUN {
		//运行状态只能选择停止
		fmt.Print("停止电梯")
		currentState = STOP
	} else if currentState == STOP {
		//停止状态可以选择再次运行或者打开电梯门
		if nextState == RUN {
			fmt.Println("电梯开始运行")
			currentState = RUN
		} else if nextState == OPEN {
			fmt.Println("打开电梯门")
			currentState = OPEN
		}
	}
}
