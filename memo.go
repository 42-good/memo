//Copyright(c) 2025 a42-good
//本程序是自由软件：你可以再分发之和/或依照由自由软件基金会发布的 GNU 通用公共许可证修改之，无论是版本 3 许可证，还是任何以后版都可以。

//发布该程序是希望它能有用，但是并无保障;甚至连可销售和符合某个特定的目的都不保证。请参看 GNU 通用公共许可证，了解详情。

//你应该随程序获得一份 GNU 通用公共许可证的复本。如果没有，请看 <https://www.gnu.org/licenses/>。 
package main

import (
	"fmt"
)

func main () {

var matter []string

for {
	fmt.Println("1.添加待定事项")	
	fmt.Println("2.列出待定事项")	
	fmt.Println("3.删除待定事项")	
	fmt.Println("4.退出")
	fmt.Print(": ")
	var choice int
	fmt.Scan(&choice)
	if choice == 1 {
		fmt.Println("请输入")
		var userinput string
		fmt.Scanln(&userinput)
		matter = append(matter, userinput)
	}else if choice == 2 {
		if len(matter) == 0 {
			fmt.Println("没有事项")
		}
		for i, matter :=  range matter {
			fmt.Println("\n", i+1, matter)
		}
	}else if choice == 3 {
		fmt.Println("输入序号")
		var number int
		fmt.Scan(&number)
		if number <1 || number > len(matter) {
			fmt.Println("错误序号")
			return
		}
		matter = append(matter[:number-1], matter[number:]...)
		fmt.Println("OK")
	}else if choice == 4 {
		return
	}else {
		fmt.Println("错误")
	}
}

}
