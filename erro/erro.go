package erro

type Erro struct {
	mens string
	erro bool
}

func (e *Erro) SetErro(erro bool) {
	e.erro = erro
}

func (e Erro) GetErro() bool {
	return e.erro
}

func (e *Erro) SetMens(mens string) {
	e.mens = mens
}

func (e Erro) GetMens() string {
	return e.mens
}
