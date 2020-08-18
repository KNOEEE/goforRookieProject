package service
import (
	"strconv"
	"go_code/projectLearn/customerManage/model"
)

//该结构完成对Customer的操作
//包括增删改查
type CustomerService struct{
	customers []model.Customer
	//声明字段 表示切片含有多少个客户
	//还可以作为新客户的id+1
	customerNum int
}

//
func NewCustomerService() *CustomerService{
	//为了能看到客户在切片中 初始化一个客户
	customerService := &CustomerService{}
	customerService.customerNum = 1
	customer := model.NewCustomer(1, "张三", "男", 20, "232", 
	"zs@mt.com")
	customerService.customers = append(customerService.customers,
	customer)
	return customerService
}

//返回用户切片
func (this *CustomerService) List() []model.Customer {
	return this.customers
}

//添加客户到customers切片
//此处必须绑定引用类型 否则无法改变实例
//每次使用的都是第一次创建的customerService实例
func (this *CustomerService) Add(customer model.Customer) bool {
	//定义分配id的规则 就是添加用户的顺序
	this.customerNum++
	customer.Id = this.customerNum
	this.customers = append(this.customers, customer)
	return true
}

//根据id删除客户（从切片中）
//添加删除功能 并没有改变model
//并没有改变customerNum 实际上就是表示id
func (this *CustomerService) Delete(id int)bool {
	index := this.FindById(id)
	if index == -1 {
		//说明没有这个客户
		return false
	}
	//从切片中删除一个元素
	this.customers = append(this.customers[:index],
	this.customers[index + 1:]...)
	return true
}

func (this *CustomerService) Update(id int, name string, gender string,
	age string, phone string, email string) bool {
	if name != "" {
		this.customers[id].Name = name
	}
	if gender != "" {
		this.customers[id].Gender = gender
	}
	if age != "" {
		this.customers[id].Age, _ = strconv.Atoi(age)
	}
	if phone != "" {
		this.customers[id].Phone = phone
	}
	if email != "" {
		this.customers[id].Email = email
	}
	return true
}

//根据id查找客户在切片中对应的下标
//若没有该客户 返回-1
func (this *CustomerService) FindById(id int) int {
	index := -1
	//遍历customers切片
	for i := 0; i < len(this.customers); i++ {
		if this.customers[i].Id == id {
			//found
			index = i
		}
	}
	return index
}