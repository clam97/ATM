package main

import (
	"bufio"
	"fmt"
	"os"

)
var a int32
var money  int32

func Save()  {//存款
	fmt.Println("存   款")
	fmt.Println("请输入存款金额：")
	fmt.Scan(&a)
	money=money+a
	fmt.Println("已存入金额：",a)

}

func GetMoney()  {
	fmt.Println("取   款")
	fmt.Println("请输入要取款的金额")
	fmt.Scan(&a)
	if a>money {
		fmt.Println("余额不足请充值！")

	}else if a<=money {
		money=money-a
		fmt.Println("已取款：",a)
	}


}

func Search()  {
	fmt.Println("查询余额")

	fmt.Println("当前余额：",money)
}
func Atm()  string{
    var i string
	fmt.Println("请选择要操作的业务:"+"(1.存   款"+"2.取   款"+"3.查询余额"+"4.退   出)")
	input := bufio.NewScanner(os.Stdin)
	input.Scan()
	i=input.Text()
	fmt.Println("您选择的业务为：", i)
	return i


}
func exit()  {
	fmt.Println("已退出ATM！")

}



func main() {
  var b int
	fmt.Println("Holle  Go!")
	fmt.Println("请输入您的帐号")
	fmt.Scan(&a)
	fmt.Println("请输入您的密码")
	fmt.Scan(&b)
	if a==123456 && b==123123{
		fmt.Println("欢迎登录ATM取款机！")

		var i string
		for true{
			i=Atm()
			switch i {
			case "1":
				Save()


			case "2":
				GetMoney()

			case "3":
				Search()


			}
			if i=="4"{
				exit()
				break
			}
	}

	}
	//fmt.Println(reflect.TypeOf(input.Text()))
	//var a string
	//fmt.Scan(&a)
	//fmt.Println(a)
	//for true{
	//	Atm(xx)
	//}




}