package model
import (
	"fmt"
)

//声明结构体 表示用户信息
type Customer struct{
	Id int
	Name string
	Gender string
	Age int
	Phone string
	Email string
}

//工厂模式 返回Customer实例
func NewCustomer(id int, name string, gender string,
	age int, phone string, email string) Customer{
	return Customer{
		Id: id,
		Name: name,
		Gender: gender,
		Age: age,
		Phone: phone,
		Email: email,
	}
}

//第二种创建Cutomer实例的方法 不带id
func NewCustomer2(name string, gender string,
	age int, phone string, email string) Customer{
	return Customer{
		Name: name,
		Gender: gender,
		Age: age,
		Phone: phone,
		Email: email,
	}
}

//返回用户信息 格式化字符串
func (this Customer) GetInfo() string {
	info := fmt.Sprintf("%v\t%v\t%v\t%v\t%v\t%v", this.Id, this.Name,
	this.Gender, this.Age, this.Phone, this.Email)
	return info
}