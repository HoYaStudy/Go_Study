package explorer

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/HoYaStudy/Go_Study/hcoin/blockchain"
)

type homeData struct {
	PageTitle string
	Blocks    []*blockchain.Block
}

const (
	port        string = ":4000"
	templateDir string = "explorer/templates/"
)

var htmlTemplates *template.Template

func Start(port int) {
	handler := http.NewServeMux()
	htmlTemplates = template.Must(template.ParseGlob(templateDir + "pages/*.gohtml"))
	htmlTemplates = template.Must(htmlTemplates.ParseGlob(templateDir + "partials/*.gohtml"))
	handler.HandleFunc("/", home)
	handler.HandleFunc("/add_block", addBlock)
	fmt.Printf("Listening on http://localhost:%d\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), handler))
}

func home(rw http.ResponseWriter, r *http.Request) {
	data := homeData{"Home", nil}
	htmlTemplates.ExecuteTemplate(rw, "home", data)
}

func addBlock(rw http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		htmlTemplates.ExecuteTemplate(rw, "add_block", nil)
	case "POST":
		blockchain.Blockchain().AddBlock()
		http.Redirect(rw, r, "/", http.StatusPermanentRedirect)
	}
}
