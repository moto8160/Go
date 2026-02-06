// ビルドと実行
// go run sample.go

package main //何らかに必ず属する。mainは必須。
import (
	"fmt" //fmt(フォーマット)パッケージ
	"reflect"
)

func main() {
	fmt.Println("Hello, World!")

	//
	// 変数
	//
	var num int
	num = 1
	fmt.Println(num)

	num2 := 2
	fmt.Println(num2)
	fmt.Println(reflect.TypeOf(num2)) //int

	//
	// 配列
	//
	list_1 := [2]string{"A", "B"}
	list_2 := [...]string{"A", "B"}
	fmt.Println(list_1) //[A B]
	fmt.Println(list_2) //[A B]

	list_3 := [2][2]string{{"A1", "A2"}, {"B1", "B2"}}
	fmt.Println(list_3)       //[[A1 A2] [B1 B2]]
	fmt.Println(list_3[0][1]) //A2

	//
	// 条件
	//
	string_1 := "a"
	string_2 := "b"

	if string_1 == string_2 {
		fmt.Println("true")
	} else {
		fmt.Println("else") //else
	}

	if string_3 := "a"; string_1 == string_3 { //if 簡易文;
		fmt.Println("true") //true
	} else {
		fmt.Println("else")
	}

	//
	// 繰り返し
	//
	for i := 0; i < 3; i++ {
		fmt.Println(i) //0 1 2
	}

	for i := 0; i < 3; {
		fmt.Println(i) //0 1 2
		i++
	}

	//
	// 関数
	//
	result := cal(10, 20)
	fmt.Println(result) //30

	//
	// 構造体
	//
	var user1 User
	user1.name = "Taro"
	user1.age = 20
	fmt.Println(user1) //{Taro 20}

	user2 := User{name: "Suzuki", age: 25}
	fmt.Println(user2) //{Suzuki 25}

	//
	// メソッド
	//
	user1.join() //Taroさんは20歳です。

	res := user2.join2("Hello") //Suzukiさんは25歳です。引数Helloを受け取りました。
	fmt.Println(res)
}

type User struct {
	name string
	age  int
}

func (user User) join() { // func (レシーバー)
	fmt.Println(user.name + "さんは" + fmt.Sprint(user.age) + "歳です。")
}

func (user User) join2(character string) (joinString string) {
	joinString = user.name + "さんは" + fmt.Sprint(user.age) + "歳です。引数" + character + "を受け取りました。"
	return joinString
}

// func cal(num1, num2 int) (r int) {
// 	r = num1 + num2
// 	return num1 + num2
// }

func cal(num1 int, num2 int) int {
	return num1 + num2
}
