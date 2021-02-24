package lambda

type IList interface {
	Array
	Insert(index int, ele interface{}) error
	IndexOf(ele interface{}) int
	RemoveAt(index int)
}

func List(source interface{}) IList {
	arr := LambdaArray(source)
	return arr.ToList()
}

func (p *_obj) Insert(index int, ele interface{}) error {
	panic("implement me")
}

func (p *_obj) IndexOf(ele interface{}) int {
	panic("implement me")
}

func (p *_obj) RemoveAt(index int) {
	panic("implement me")
}
