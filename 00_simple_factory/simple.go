package simplefactory

import "fmt"

//API is interface
//这个API interface是我们工厂返回的对象类型
type API interface {
	Say(name string) string
}

//NewAPI return Api instance by type
//这个就是工厂函数啦
func NewAPI(t int) API {
	if t == 1 {
		return &hiAPI{}
	} else if t == 2 {
		return &helloAPI{}
	}
	return nil
}

//下面是各种具体的类型，以及它们必须实现的Say方法
//hiAPI is one of API implement
type hiAPI struct{}

//Say hi to name
func (*hiAPI) Say(name string) string {
	return fmt.Sprintf("Hi, %s", name)
}

//HelloAPI is another API implement
type helloAPI struct{}

//Say hello to name
func (*helloAPI) Say(name string) string {
	return fmt.Sprintf("Hello, %s", name)
}
