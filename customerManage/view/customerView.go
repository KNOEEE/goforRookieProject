package main
import (
	"fmt"
	"go_code/projectLearn/customerManage/service"
	"go_code/projectLearn/customerManage/model"
)

type customerView struct{
	key string //接收用户输入
	loop bool 
	customerService *service.CustomerService
}

// 显示所有客户信息
func (this *customerView) list() {
	//获取当前所有的客户信息 在切片中
	customers := this.customerService.List()
	//show
	fmt.Println("----------------客户列表----------------")
	fmt.Println("编号\t姓名\t性别\t年龄\t电话\t邮箱")
	for i := 0; i < len(customers); i++ {
		//此处没有import model但是可以使用GetInfo方法
		fmt.Println(customers[i].GetInfo())
	}
	fmt.Println("--------------客户列表完成--------------")
}

//得到用户输入的信息 构建新的customer实例
//调用service的Add方法
func (this *customerView) add() {
	fmt.Println("----------------添加客户----------------")
	fmt.Print("姓名：")
	name := ""
	fmt.Scanln(&name)
	fmt.Print("性别：")
	gender := ""
	fmt.Scanln(&gender)
	fmt.Print("年龄：")
	age := 0
	fmt.Scanln(&age)
	fmt.Print("电话：")
	phone := ""
	fmt.Scanln(&phone)
	fmt.Print("邮箱：")
	email := ""
	fmt.Scanln(&email)
	//构建customer实例
	//id不由用户输入 系统分配
	customer := model.NewCustomer2(name, gender, age, phone, email)
	if this.customerService.Add(customer) {
		fmt.Println("----------------添加完成----------------")
	}else{
		fmt.Println("----------------添加失败----------------")
	}
}

//得到用户输入id 删除该客户
func (this *customerView) delete() {
	fmt.Println("----------------删除用户----------------")
	fmt.Println("请输入待删除客户编号（-1退出）：")
	id := -1
	fmt.Scanln(&id)
	if id == -1 {
		return
	}
	fmt.Println("确认是否删除（Y/N）：")
	choice := ""
	fmt.Scanln(&choice)
	//以上全部是界面的控制
	if choice == "y" || "Y" == choice {
		if this.customerService.Delete(id) {
			fmt.Println("----------------删除完成----------------")
		}else{
			fmt.Println("--------删除失败，输入的id不存在---------")
		}
	}
}

func (this *customerView) update() {
	fmt.Println("----------------修改客户----------------")
	fmt.Println("请输入待修改客户编号（-1退出）：")
	fmt.Println("直接回车表示不修改")
	id := -1
	fmt.Scanln(&id)
	if id == -1 {
		return
	}
	id = this.customerService.FindById(id)
	if id == -1 {
		fmt.Println("--------修改失败，输入的id不存在---------")
		return
	}
	customers := this.customerService.List()
	fmt.Printf("姓名(%v): ", customers[id].Name)
	name := ""
	fmt.Scanln(&name)
	fmt.Printf("性别(%v): ", customers[id].Gender)
	gender := ""
	fmt.Scanln(&gender)
	fmt.Printf("年龄(%v): ", customers[id].Age)
	age := "" // 为达到不输入不变 先用string表示
	fmt.Scanln(&age)
	fmt.Printf("电话(%v): ", customers[id].Phone)
	phone := ""
	fmt.Scanln(&phone)
	fmt.Printf("邮箱(%v): ", customers[id].Email)
	email := ""
	fmt.Scanln(&email)
	if this.customerService.Update(id, name, gender, age, 
		phone, email) {
		fmt.Println("----------------修改完成----------------")
	}else{
		fmt.Println("----------------修改失败----------------")
	}
}

//退出软件
func (this *customerView) exit(){
	fmt.Println("确认是否退出（y/n）：")
	for {
		fmt.Scanln(&this.key)
		if this.key == "Y" || "y" == this.key || "N" == this.key || "n" == this.key {
			break
		}
		fmt.Println("输入有误（y/n）")
	}
	if this.key == "y" || "Y" == this.key {
		this.loop = false
	}
	
}
//显示主菜单
func (this *customerView) mainMenu(){
	for {
		fmt.Println("\n\n--------------客户信息管理软件--------------")
		fmt.Println("              1 添加客户")
		fmt.Println("              2 修改客户")
		fmt.Println("              3 删除客户")
		fmt.Println("              4 客户列表")
		fmt.Println("              5 退    出")
		//不会自动换行
		fmt.Print("请选择（1-5）：")
		fmt.Scanln(&this.key)
		switch this.key {
		case "1":
			this.add()
		case "2":
			this.update()
		case "3":
			this.delete()
		case "4":
			this.list()
		case "5":
			this.exit()
		default:
			fmt.Println("输入有误，请重新输入（1-5）：")
		}
		if !this.loop {
			break
		}
	}
	fmt.Println("已退出客户管理系统")
}

func main(){
	//创建一个customerView 运行并显示主菜单
	customerView := customerView{
		key : "",
		loop : true,
	}
	//完成对结构体的customerService字段的初始化
	customerView.customerService = service.NewCustomerService()
	customerView.mainMenu()
}