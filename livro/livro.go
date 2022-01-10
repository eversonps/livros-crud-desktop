package livro

type Livro struct {
	codigo  string
	titulo  string
	autor   string
	editora string
	ano     int
}

func (l *Livro) SetCodigo(codigo string) {
	l.codigo = codigo
}

func (l Livro) GetCodigo() string {
	return l.codigo
}

func (l *Livro) SetTitulo(titulo string) {
	l.titulo = titulo
}

func (l Livro) GetTitulo() string {
	return l.titulo
}

func (l *Livro) SetAutor(autor string) {
	l.autor = autor
}

func (l Livro) GetAutor() string {
	return l.autor
}

func (l *Livro) SetEditora(editora string) {
	l.editora = editora
}

func (l Livro) GetEditora() string {
	return l.editora
}

func (l *Livro) SetAno(ano int) {
	l.ano = ano
}

func (l Livro) GetAno() int {
	return l.ano
}
