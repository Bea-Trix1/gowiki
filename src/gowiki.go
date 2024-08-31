package main

import (
	"fmt"
	"net/http"
	"os"
)

func verHandler(w http.ResponseWriter, r *http.Request) {
	titulo := r.URL.Path[len("/ver/"):]
	p, _ := carregarPagina(titulo)
	fmt.Fprintf(w, "<h1>%s</h1><div>%s</div", p.Titulo, p.Corpo)
}

type Pagina struct {
	Titulo string
	Corpo  []byte
}

func main() {
	p1 := &Pagina{
		Titulo: "Pagina de Teste",
		Corpo:  []byte("Olá, essa é uma página de teste."),
	}

	p1.salvar()

	p2, _ := carregarPagina("Pagina de Teste")
	fmt.Println(string(p2.Corpo))
}

func (p *Pagina) salvar() error {
	nomeDoArquivo := p.Titulo + ".txt"
	return os.WriteFile(nomeDoArquivo, p.Corpo, 0600)
}

func carregarPagina(titulo string) (*Pagina, error) {
	nomeDoArquivo := titulo + ".txt"
	corpo, error := os.ReadFile(nomeDoArquivo)
	if error != nil {
		return nil, error
	}

	return &Pagina{
		Titulo: titulo,
		Corpo:  corpo,
	}, nil
}
