package main
import (
	"fmt"
	"go_code/projectLearn/myAccount/utils"
)

func main(){
	fmt.Println("面向对象方式")
	utils.NewFamilyAccount().MainMenu()
}