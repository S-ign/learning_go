package main

import (
	"log"
	"os"
	"text/template"
)

type Part struct {
	Name string
	Count int
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func executeTemplate(text string, data interface{}) {
	tmpl, err := template.New("test").Parse(text)
	check(err)
	err = tmpl.Execute(os.Stdout, data)
	check(err)
}

func main() {
	//{{.}} places text in the template
	executeTemplate("Dot is: {{.}}!\n", "ABC")
	executeTemplate("Dot is: {{.}}!\n", "123.5")

	//{{if .}} places text only if true {{end}}
	executeTemplate("start {{if .}}Dot is true!{{end}} finish\n", true)
	executeTemplate("start {{if .}}Dot is false!{{end}} finish\n", false)

	//{{range .}} loops therough slice indexes when in between range action. {{end}}
	//otherwise prints out the entire slice.
	templateText := "Before loop: {{.}}\n{{range .}}In Loop: {{.}}\n{{end}}After loop: {{.}}\n"
	executeTemplate(templateText, []string{"do","re","mi"})

	templateText = "Prices:\n{{range .}}${{.}}\n{{end}}"
	executeTemplate(templateText, []float64{1.25, 0.99, 27})

	//Access struct fields by placing the name of the field after the dot {{.Name}}
	templateText = "Name: {{.Name}}\nCount: {{.Count}}\n"
	//Pass in the newly created struct.
	executeTemplate(templateText, Part{Name: "Fuses", Count: 5})
	executeTemplate(templateText, Part{Name: "Cables", Count: 2})
}
