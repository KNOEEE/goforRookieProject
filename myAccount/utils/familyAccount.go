package utils
import (
	"fmt"
)

type FamilyAccount struct {
	//接收用户输入的选项
	key string
	//控制是否退出循环
	loop bool
	//账户余额
	balance float64
	//每次收支的金额
	money float64
	//每次收支的说明
	note string
	//定义变量 记录是否有收支行为
	flag bool
	//收支详情使用string记录
	//收支后更新details
	details string
}

//给结构绑定方法

func (this *FamilyAccount) ShowDetails(){
	fmt.Println("--------------当前收支明细记录--------------")
	if this.flag {
		fmt.Println(this.details)
	}else{
		fmt.Println("当前没有收支明细，来一笔吧！")
	}
}

func (this *FamilyAccount) income(){
	fmt.Println("本次收入金额：")
	fmt.Scanln(&this.money)
	this.balance += this.money //modify balance
	fmt.Println("本次收入说明：")
	fmt.Scanln(&this.note)
	//将收入情况 拼到details
	this.details += fmt.Sprintf("\n收入\t%v\t\t%v\t\t%v", 
	this.balance, this.money, this.note)
	this.flag = true
}

func (this *FamilyAccount) pay(){
	fmt.Println("本次支出金额：")
	fmt.Scanln(&this.money)
	//此处需要判断
	if this.money > this.balance {
		fmt.Println("余额金额不足")
		return
	}
	this.balance -= this.money
	fmt.Println("本次支出说明")
	fmt.Scanln(&this.note)
	//将收入情况 拼到details
	this.details += fmt.Sprintf("\n收入\t%v\t\t%v\t\t%v", 
	this.balance, this.money, this.note)
	this.flag = true
}

//将退出封装到方法
func (this *FamilyAccount) exit(){
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
		this.loop = false 
	}
}

//编写工程模式的构造方法
func NewFamilyAccount() *FamilyAccount{
	return &FamilyAccount{
		key: "",
		loop: true,
		balance: 10000.0,
		money: 0.0,
		note: "",
		flag: false,
		details: "收支\t账户金额\t收支金额\t说明",
	}
}

//显示主菜单
func (this *FamilyAccount) MainMenu() {
	for {
		fmt.Println("\n--------------家庭收支记账软件--------------")
		fmt.Println("               1 收支明细")
		fmt.Println("               2 登记收入")
		fmt.Println("               3 登记支出")
		fmt.Println("               4   退出")
		fmt.Print("请选择(1-4):")
		fmt.Scanln(&this.key)
		switch this.key {
		case "1":
			this.ShowDetails()
		case "2":
			this.income()
		case "3":
			this.pay()
		case "4":
			this.exit()
			// 此处break只能退出switch
		default:
			fmt.Println("请输入正确选项")
		}
		if !this.loop {
			break
		}
	}
}