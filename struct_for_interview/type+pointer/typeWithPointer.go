package type_pointer

type personal struct {
	Name string
}

func (p *personal) UpdateNameP(newName string) {
	p.Name = newName
}

func (p personal) UpdateName(newName string) {
	p.Name = newName
}
