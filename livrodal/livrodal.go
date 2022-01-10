package livrodal

import (
	"POO/livro"
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

// Variaveis que armazenará o estado do banco de dados que será aberto na função Conecta()
var db *sql.DB
var erroBD error

/* Variaveis criadas com intuito de retornar valores para o package livro.bll
que utilizará essas informações para a validação */
var valorMens string
var valorErro bool

/* Como já visto abrirá uma conexão com o banco de dados "cadastrolivro" que criamos no workbench
e foi ativado no wampserver */
func Conecta() {
	db, erroBD = sql.Open("mysql", "root:@tcp(localhost:3306)/cadastrolivro")
	if erroBD != nil {
		panic(erroBD.Error())
	}

}

/* Quando chamada irá desfazer a conexão com o bando de dados */
func Desconecta() {
	defer db.Close()
}

/* Função responsável por inserir um livro na nossa tabela do banco de dados.
Ela receberá um livro instanciado e setará as informações da struct na tabela do nosso banco de dados */
func InserirLivro(livro livro.Livro) (bool, string) {
	// Armazenando a pesquisa responsavel por inserir um livro na tabela do mysql
	var pesquisa = "insert into livros (codigo,titulo,autor,editora,ano) values (?,?,?,?,?)"

	// Preparando a pesquisa
	stmt, erroPesquisa := db.Prepare(pesquisa)

	// Caso ocorra um erro na pesquisa será retornado um erro ao livrobll.go
	if erroPesquisa != nil {
		valorErro = true
		valorMens = "erro ao preparar a consulta ao banco de dados!"
		return valorErro, valorMens
	}

	// Substitituindo os "?" da variavel "pesquisa" pelos valores do livro instanciado e inserindo os dados na tabela
	respostaInserir, erroInserir := stmt.Exec(livro.GetCodigo(), livro.GetTitulo(), livro.GetAutor(), livro.GetEditora(), livro.GetAno())

	// Caso ocorra um erro ao inserir os dados na tabela será retornado um erro ao livrobll.go
	if erroInserir != nil {
		valorErro = true
		valorMens = "erro ao inserir o livro ao banco de dados!"
		return valorErro, valorMens
	}

	fmt.Print(respostaInserir)
	// Relatando ao livrobll.go que não houve erro ao inserir o livro
	valorErro = false
	valorMens = "livro cadastrado com sucesso!"
	return valorErro, valorMens
}

/* Essa será responsável por excluir um livro da nossa tabela do nosso banco de dados. Receberá um livro instanciado
com o código digitado pelo usuário e excluirá o livro com base nele */
func ExcluiUmLivro(livro livro.Livro) (bool, string) {
	// Armazenando a pesquisa responsavel por deletar um livro na tabela do mysql
	var pesquisa = "delete from livros where codigo = (?)"

	// Preparando a pesquisa ao banco de dados
	stmt, erroPesquisa := db.Prepare(pesquisa)

	// Caso ocorra um erro na pesquisa será retornado um erro ao livrobll.go
	if erroPesquisa != nil {
		valorErro = true
		valorMens = "erro ao preparar a consulta ao banco de dados!"
		return valorErro, valorMens
	}

	// Substitituindo o "?" da variavel "pesquisa" pelo código do livro instanciado e deletando os dados da tabela
	respostaExclui, erroExclui := stmt.Exec(livro.GetCodigo())

	// Caso ocorra um erro ao deletar os dados da tabela do banco de dados será retornado um erro ao livrobll.go
	if erroExclui != nil {
		valorErro = true
		valorMens = "Houve um erro ao excluir o livro!"
		return valorErro, valorMens
	}

	fmt.Print(respostaExclui)
	// Relatando ao livrobll.go que não houve erro ao deletar o livro
	valorErro = false
	valorMens = "o livro foi excluído com sucesso"
	return valorErro, valorMens
}

/* Essa função será responsável por atualizar um livro existente na nossa tabela do nosso banco de dados.
Receberá um livro instanciado e com base no código digitado pelo usuário atualizará o livro com as informações digitadas */
func AtualizaUmLivro(livro livro.Livro) (bool, string) {
	// Armazenando a pesquisa responsavel por atualizar um livro na tabela do mysql
	var pesquisa = "update livros set titulo=(?), autor=(?), editora=(?), ano=(?) where codigo = (?)"

	// Preparando a pesquisa ao banco de dados
	stmt, erroPesquisa := db.Prepare(pesquisa)

	// Caso ocorra um erro na pesquisa será retornado um erro ao livrobll.go
	if erroPesquisa != nil {
		valorErro = true
		valorMens = "erro ao preparar a consulta ao banco de dados!"
		return valorErro, valorMens
	}

	// Substitituindo os "?" da variavel "pesquisa" pelas informações do livro instanciado e atualizando os dados da tabela
	respostaAtualiza, erroAtualiza := stmt.Exec(livro.GetTitulo(), livro.GetAutor(), livro.GetEditora(), livro.GetAno(), livro.GetCodigo())

	// Caso ocorra um erro ao atualizar o livro será retornado um erro ao livrobll.go
	if erroAtualiza != nil {
		valorErro = true
		valorMens = "não foi possível alterar o livro!"
		return valorErro, valorMens
	}

	// Relatando ao livrobll.go que não houve erro ao deletar o livro
	fmt.Print(respostaAtualiza)
	valorErro = false
	valorMens = "o livro foi alterado com sucesso!"
	return valorErro, valorMens
}

/* Essa função será responsável por consultar um livro existente na nossa tabela do nosso banco de dados.
Receberá um livro instanciado e com base no código digitado pelo usuário retornará o livro pesquisado
na tabela do banco de dados. Note que nessa função além de retornar o valor do erro e a mensagem de erro,
estamos retornando também o livro para o livrobll.go que será utilizado no livroihm.go */
func ConsultaUmLivro(livro livro.Livro) (livro.Livro, bool, string) {
	// Armazenando a pesquisa responsavel por consultar um livro na tabela do mysql
	var pesquisa = "select * from livros where codigo=(?)"
	var entrou = false

	// Preparando a pesquisa ao banco de dados
	smtr, erroPesquisa := db.Prepare(pesquisa)

	// Caso ocorra um erro na pesquisa será retornado um erro ao livrobll.go
	if erroPesquisa != nil {
		valorErro = true
		valorMens = "erro ao preparar a consulta ao banco de dados!"
		return livro, valorErro, valorMens
	}

	/* Dessa vez não utilizaremos o Exec(), pois queremos fazer uma consulta e não alterar diretamente os valores da tabela
	do banco de dados, para isso utilizamos o a função Query() */
	respostaConsulta, erroConsulta := smtr.Query(livro.GetCodigo())

	/* Nessa parte do código definimos os valores padrão da struct e dos valores da mensagem e de erro,
	pois se a função Query() não achar o livro com o código consultado a struct livro retornará o livro
	que foi consultado anteriormente. Note que não damos "return", pois alteraremos este valor mais para frente
	caso a função Query() ache um livro na tabela do banco de dados */
	livro.SetCodigo("")
	livro.SetTitulo("")
	livro.SetAutor("")
	livro.SetEditora("")
	livro.SetAno(0)
	valorErro = true
	valorMens = "erro ao consultar o livro!"

	// Caso ocorra um erro brusco ao consultar o livro será retornado um erro ao livrobll.go
	if erroConsulta != nil {
		valorErro = true
		valorMens = "erro ao consultar o livro!"
		return livro, valorErro, valorMens
	}

	// Esse for irá percorrer o livro somente se existir na tabela um livro com o código digitado pelo usuário
	for respostaConsulta.Next() {
		entrou = true

		var codigo string
		var titulo string
		var autor string
		var editora string
		var ano int

		// A função Scan irá pegar os valores pesquisados na tabela e jogará nas váriaveis criadas
		erroLeituraDados := respostaConsulta.Scan(&codigo, &titulo, &autor, &editora, &ano)

		// Caso haja um erro ao setar os valores da tabela nas variaveis criadas será retornado um erro ao livrobll.go
		if erroLeituraDados != nil {
			valorErro = true
			valorMens = "houve um erro ao ler os dados da tabela do banco de dados!"
			return livro, valorErro, valorMens
		}

		/* E finalmente caso não haja nenhum erro setaremos os valores pesquisados na struct livro que será
		retornada ao livro bll.go */
		livro.SetCodigo(codigo)
		livro.SetTitulo(titulo)
		livro.SetAutor(autor)
		livro.SetEditora(editora)
		livro.SetAno(ano)
	}

	/* Caso ele tenha entrado no for e chegou até essa parte do código é porque deu tudo certo.
	Portanto relatamos ao livrobll.go que a consulta foi um sucesso */
	if entrou {
		valorErro = false
		valorMens = "livro consultado com sucesso"
	}

	// Retornando o livro consultado, o valor do erro e a sua mensagem ao livrobll.go
	return livro, valorErro, valorMens
}
