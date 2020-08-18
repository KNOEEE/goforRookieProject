package main
import (
	"fmt"
)

func main() {
	//接收用户输入的选项
	key := ""
	//控制是否退出循环
	loop := true
	//账户余额
	balance := 10000.0
	//每次收支的金额
	money := 0.0
	//每次收支的说明
	note := ""
	//定义变量 记录是否有收支行为
	flag := false
	//收支详情使用string记录
	//收支后更新details
	details := "收支\t账户金额\t收支金额\t说明"
	//显示主菜单
	for {
		fmt.Println("\n--------------家庭收支记账软件--------------")
		fmt.Println("               1 收支明细")
		fmt.Println("               2 登记收入")
		fmt.Println("               3 登记支出")
		fmt.Println("               4   退出")
		fmt.Print("请选择(1-4):")
		fmt.Scanln(&key)
		switch key {
		case "1":
			fmt.Println("--------------当前收支明细记录--------------")
			if flag {
				fmt.Println(details)
			}else{
				fmt.Println("当前没有收支明细，来一笔吧！")
			}
			
		case "2":
			fmt.Println("本次收入金额：")
			fmt.Scanln(&money)
			balance += money //modify balance
			fmt.Println("本次收入说明：")
			fmt.Scanln(&note)
			//将收入情况 拼到details
			details += fmt.Sprintf("\n收入\t%v\t\t%v\t\t%v", 
			balance, money, note)
			flag = true
		case "3":
			fmt.Println("本次支出金额：")
			fmt.Scanln(&money)
			//此处需要判断
			if money > balance {
				fmt.Println("余额金额不足")
				//跳出switch
				break
			}
			balance -= money
			fmt.Println("本次支出说明")
			fmt.Scanln(&note)
			//将收入情况 拼到details
			details += fmt.Sprintf("\n收入\t%v\t\t%v\t\t%v", 
			balance, money, note)
			flag = true
		case "4":
			fmt.Println("你确定要退出吗？ y/n")
			choice := ""
			for {
				fmt.Scanln(&choice)
				if choice == "y" || "n" == choice {
					break
				}
				fmt.Println("输入有误，请重新输入 y/n")
			}
			if choice == "y" {
				loop = false 
			}
			// 此处break只能退出switch
		default:
			fmt.Println("请输入正确选项")
		}
		if !loop {
			break
		}
	}
	fmt.Println("已退出")
}