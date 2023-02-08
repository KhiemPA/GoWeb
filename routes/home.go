package routes

import (
	"fmt"
	"net/http"
	"prj/models"
	"prj/utils"

	
)
func homeGetHandler(w http.ResponseWriter, r *http.Request) {
	categories, err := models.GetCategories()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal server error"))
	}

	products, err := models.GetProducts()

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("internal server error"))
	}

	utils.ExcuteTemplate(w, "home.html", struct {
		Categories []models.Category
		Products []models.Product
	}{
		Categories: categories,
		Products: products,
	})
}

func homePostHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "text-html; charset-UTF-8")
	r.ParseForm()
	search := r.PostForm.Get("search")
	products, err := models.SearchProducts(search)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("internal server error"))
	}
	
	var html string = ""
	count := len(products)
	if count > 0 {
		html += "<table class='table table-bordered'>"
		html += fmt.Sprintf("<th> Id </th> <th>Category</th> <th>Name</th> <th>Price</th> <th>Quantity</th> <th>Amount</th>")
		for _, p := range products {
			html += "<tr>"
			html += fmt.Sprintf("<td>%d</td> <td>%s</td> <td>%s</td> <td>%.2f</td> <td>%d</td> <td>%.2f</td>", p.Id, p.Category.Description, p.Name, p.Price, p.Quantity, p.Amount)
			html += "</tr>"
		}
		html += "</table>"
	} else {
		html += fmt.Sprintf(`<p class'alert alert-info'>No result with <code>"<strong>%s</strong>"</code></p>`, search)

	}

	w.Write([]byte(html))
}