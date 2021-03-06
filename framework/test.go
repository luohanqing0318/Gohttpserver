package framework

import (

	"fmt"
)



func Test1() ControllerHandler {
	// 使用函数回调
	return func(c *Context) error {
		fmt.Println("middleware pre test1")
		c.Next()  // 调用Next往下调用，会自增contxt.index
		fmt.Println("middleware post test1")
		return nil
	}
}

func Test2() ControllerHandler {
	// 使用函数回调
	return func(c *Context) error {
		fmt.Println("middleware pre test2")
		c.Next() // 调用Next往下调用，会自增contxt.index
		fmt.Println("middleware post test2")
		return nil
	}
}

func Test3() ControllerHandler {
	// 使用函数回调
	return func(c *Context) error {
		fmt.Println("middleware pre test3")
		c.Next()
		fmt.Println("middleware post test3")
		return nil
	}
}