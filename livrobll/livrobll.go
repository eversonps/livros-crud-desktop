package livrobll

import (
	"POO/erro"
	"POO/livro"
	"POO/livrodal"
)

/* Quando chamada pelo livroihm.go irá chamar a função pertencente
ao pacote livrodal que abrirá uma conexão ao banco de dados */

func Conecta() {
	livrodal.Conecta()
}

/* Quando chamada irá executar a função do pacote livrodal.go
que irá desfazer a conexão com o banco de dados */
func Desconecta() {
	livrodal.Desconecta()
}

func ValidaDados(e *erro.Erro, l livro.Livro, op string) {
	e.SetErro(false)

	if l.GetCodigo() == "" {
		e.SetMens("O campo código é de preenchimento obrigatório")
		e.SetErro(true)
		return
	}

	if l.GetTitulo() == "" {
		e.SetMens("O campo titulo é de preenchimento obrigatório")
		e.SetErro(true)
		return
	}

	if l.GetAutor() == "" {
		e.SetMens("O campo autor é de preenchimento obrigatório")
		e.SetErro(true)
		return
	}

	if l.GetEditora() == "" {
		e.SetMens("O campo editora é de preenchimento obrigatório")
		e.SetErro(true)
		return
	}

	if l.GetAno() <= 0 {
		e.SetMens("O campo ano tem de ser maior do que 0")
		e.SetErro(true)
		return
	}

	/* Caso a opção digitada seja "i" é porque o usuario deseja inserir um livro,
	caso contrário ele deseja atualizar um livro */
	if op == "i" {
		/* Com as informações retornadas pelo livrodal.go ele irá inserir na struct Erro se houve erro ou não
		e a sua mensagem */
		valorErro, valorMens := livrodal.InserirLivro(l)
		e.SetErro(valorErro)
		e.SetMens(valorMens)
	} else {
		/* Com as informações retornadas pelo livrodal.go ele irá inserir na struct Erro se houve erro ou não
		e a sua mensagem */
		valorErro, valorMens := livrodal.AtualizaUmLivro(l)
		e.SetErro(valorErro)
		e.SetMens(valorMens)
	}
}

func ValidaCodigo(e *erro.Erro, l *livro.Livro, op string) {
	e.SetErro(false)
	if l.GetCodigo() == "" {
		/* Com as informações retornadas pelo livrodal.go ele irá inserir na struct Erro se houve erro ou não
		e a sua mensagem */
		e.SetMens("O campo código é de preenchimento obrigatório")
		e.SetErro(true)
		return
	}

	/* Caso a opção digitada seja "c" é porque o usuario deseja consultar um livro,
	caso contrário ele deseja deletar um livro */
	if op == "c" {
		/* Com as informações retornadas pelo livrodal.go ele irá inserir na struct Erro se houve erro ou não
		e a sua mensagem */
		livro, valorErro, valorMens := livrodal.ConsultaUmLivro(*l)
		e.SetErro(valorErro)
		e.SetMens(valorMens)

		/* Se a consulta ao livro foi um sucesso ele setará os valores no livro que faz referencia
		ao livro do arquivo livroihm.go */
		if valorErro == false {
			l.SetCodigo(livro.GetCodigo())
			l.SetTitulo(livro.GetTitulo())
			l.SetAutor(livro.GetAutor())
			l.SetEditora(livro.GetEditora())
			l.SetAno(livro.GetAno())
		}
	} else {
		/* Com as informações retornadas pelo livrodal.go ele irá inserir na struct Erro se houve erro ou não
		e a sua mensagem */
		valorErro, valorMens := livrodal.ExcluiUmLivro(*l)
		e.SetErro(valorErro)
		e.SetMens(valorMens)
	}
}
