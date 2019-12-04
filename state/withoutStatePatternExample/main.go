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
		if nextState == OPEN {
			fmt.Println("电梯门已经打开")
		} else if nextState == CLOSE {
			fmt.Println("关闭电梯门")
			currentState = CLOSE
		} else if nextState == RUN {
			fmt.Println("电梯门尚未关闭，不能运行")
		} else if nextState == STOP {
			fmt.Print("状态切换错误，只能从运行状态切换到停止状态")
		}
	} else if currentState == CLOSE {
		if nextState == OPEN {
			fmt.Println("打开电梯门")
		} else if nextState == CLOSE {
			fmt.Println("电梯门已经关闭")
		} else if nextState == RUN {
			fmt.Println("电梯开始运行")
			currentState = RUN
		} else if nextState == STOP {
			fmt.Print("状态切换错误，只能从运行状态切换到停止状态")
		}
	} else if currentState == RUN {
		if nextState == OPEN {
			fmt.Println("不能打开运行中的电梯门")
		} else if nextState == CLOSE {
			fmt.Println("状态切换错误，只能从电梯门打开状态切换到关闭状态")
		} else if nextState == RUN {
			fmt.Println("电梯正在运行")
		} else if nextState == STOP {
			fmt.Print("停止电梯")
			currentState = STOP
		}
	} else if currentState == STOP {
		if nextState == OPEN {
			fmt.Println("打开电梯门")
			currentState = OPEN
		} else if nextState == CLOSE {
			fmt.Println("状态切换错误，只能从电梯门打开状态切换到关闭状态")
		} else if nextState == RUN {
			fmt.Println("电梯开始运行")
			currentState = RUN
		} else if nextState == STOP {
			fmt.Print("停止电梯")
			currentState = STOP
		}
	}

}
