package src

/*
接口定义一组方法
根据类型可以执行的操作而不是其所能容纳的数据类型来设计抽象
所有定义了该方法的类型我们称之为实现了Animal接口
*/

type Animal interface {
	Speak() string
}

// 狗
type Dog struct{

}

func (d Dog) Speak() string {
	return "Woof!"
}

// 猫
type Cat struct {

}

func (c Cat) Speak() string {
	return "Meow!"
}

// 骆驼
type Llama struct {

}

func (l Llama) Speak() string {
	return "mie!"
}

// JavaProgrammer
type JavaProgrammer struct {

}

func (j JavaProgrammer) Speak() string {
	return "Design pattern!"
}


//func main() {
//	animals := []Animal{Dog{}, Cat{}, Llama{}, JavaProgrammer{}}
//	for _, animal := range animals {
//		fmt.Println(animal.Speak())
//	}
//}
