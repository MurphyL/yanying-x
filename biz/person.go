package biz

import "log"

type Person struct {
}

func NewPerson() *Person {
	log.Println("动态创建角色")
	return &Person{}
}
