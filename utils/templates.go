package utils
import  (
	"net/http"
	"html/template"
)

var templates *template.Template

func LoadTemplates(path string) {
	templates = template.Must(template.ParseGlob(path))
}

func ExcuteTemplate(w http.ResponseWriter, tmpl string, data interface{}) {
	templates.ExecuteTemplate(w, tmpl, data)
}