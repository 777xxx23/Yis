package ynet

import "Yis/yiface"

// BaseRouter 要实现router先继承BaseRouter，然后再重写三个方法（因为有些router可能不需要全部实现）,模版模式
type BaseRouter struct {
}

func (br *BaseRouter) PreHandle(request yiface.IRequest) {

}
func (br *BaseRouter) Handle(request yiface.IRequest) {

}
func (br *BaseRouter) PostHandle(request yiface.IRequest) {

}
