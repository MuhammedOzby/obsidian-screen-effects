package handler

import (
	"html/template"
	"net/http"
	"path/filepath"
)

// ViewHandler view endpointleri için bağımlılıkları tutar
type ViewHandler struct {
	// İleride buraya da YouTube Client gibi bağımlılıklar eklenebilir.
}

// IndexPageHandler OBS'te görünecek ana sayfayı render eder
func (h *ViewHandler) IndexPageHandler(w http.ResponseWriter, r *http.Request) {
	// Şablon dosyasının yolunu belirliyoruz
	tmplPath := filepath.Join("templates", "pages", "index.tmpl")

	// Şablonu parse ediyoruz
	tmpl, err := template.ParseFiles(tmplPath)
	if err != nil {
		http.Error(w, "Could not parse template", http.StatusInternalServerError)
		return
	}

	// Şablonu veri olmadan çalıştırıp response'a yazıyoruz
	tmpl.Execute(w, nil)
}
