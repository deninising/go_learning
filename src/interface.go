package main

import "fmt"

// 接口一组方法的集合
type People interface {
	GetName() string
}

type Role interface {
	GetRole() string
}

type Student struct {
	name string
	role string
}

// 一个类型可以实现多个接口
func (student Student) GetName() string {
	return student.name
}

func (student Student) GetRole() string {
	return student.role
}

type Teacher struct {
	name string
}

func (teacher *Teacher) GetName() string {
	return teacher.name // 无论是指针还是 值都可以通过“.”来获取属性值
}

func main() {
	student := Student{
		name: "张三",
		role: "学生",
	}
	fmt.Println(student.GetName())
	fmt.Println(student.GetRole())

	var people People
	people = student // 通过值来实现接口

	var role Role = student
	fmt.Println(people.GetName())
	fmt.Println(role.GetRole())
	// 检测接口类型
	CheckPeopleInterface(student)

	teacher := Teacher{
		name: "王老师",
	}
	people = &teacher // 通过指针来实现接口,teacher 本身没有实现People接口
	fmt.Println(people.GetName())
	// 检测接口类型
	CheckPeopleInterface(&teacher)

	// 接口的零值是nil
	var person People
	if person == nil {
		fmt.Println("接口的零值是nil")
	}

}

// 由于x.(T)只能是接口类型判断，所以传参时候，传入的是接口类型
func CheckPeopleInterface(test interface{}) {
	if _, ok := test.(People); ok {
		fmt.Println("Student implements People")
	}

	if _, ok := test.(Role); ok {
		fmt.Println("Student implements Role")
	}

}
