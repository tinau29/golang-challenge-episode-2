package app

type BaseInterface interface {
	SetIDAndName(int, string)
}

func (b *Base) SetIDAndName(id int, name string) {
	b.ID = id
	b.Name = name
}
