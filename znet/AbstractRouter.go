package znet

import "binary/wz/kronus/ziface"

// AbstractRouter 路由的抽象类
type AbstractRouter struct {
}

func (router *AbstractRouter) PreHandle(request ziface.IRequest) {

}

func (router *AbstractRouter) Handle(request ziface.IRequest) {

}

func (router *AbstractRouter) PostHandle(request ziface.IRequest) {

}
