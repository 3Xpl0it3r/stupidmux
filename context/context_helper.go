package context


type middleWare struct {

}

func NewMiddleWare()*middleWare {
	return nil
}



func middleWareChain(handler StupidHandler)StupidHandler{
	return func(ctx *Context) {
		handler(ctx)
	}
}