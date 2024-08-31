package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
)

type Pagina struct {
	Titulo string
	Corpo  []byte
}

func main() {
	http.HandleFunc("/ver/", verHandler)
	http.HandleFunc("/editar/", editarHandler)
	http.HandleFunc("/salvar/", verHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func verHandler(w http.ResponseWriter, r *http.Request) {
	titulo := r.URL.Path[len("/ver/"):]
	p, error := carregarPagina(titulo)
	if error != nil {
		http.Redirect(w, r, "/editar/"+titulo, http.StatusFound)
		return
	}
	renderTemplate(w, "ver", p)
}

func editarHandler(w http.ResponseWriter, r *http.Request) {
	titulo := r.URL.Path[len("/editar/"):]
	p, error := carregarPagina(titulo)
	if error != nil {
		p = &Pagina{Titulo: titulo}
	}

	renderTemplate(w, "editar", p)
}

var templates = template.Must(template.ParseFiles("editar.html", "ver.html"))

func renderTemplate(w http.ResponseWriter, tmpl string, p *Pagina) {
	err := templates.ExecuteTemplate(w, tmpl+".html", p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func salvarHandler(w http.ResponseWriter, r *http.Request) {
	titulo := r.URL.Path[len("/salvar/"):]
	corpo := r.FormValue("corpo")

	p := &Pagina{Titulo: titulo, Corpo: []byte(corpo)}
	err := p.salvar()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/ver/"+titulo, http.StatusFound)
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
