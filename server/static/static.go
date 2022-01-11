package static

import (
	"embed"
	"fmt"
	"log"
	"net/http"
	"strings"
)

var (
	//go:embed content
	res embed.FS

	pages = map[string]string{
		// html
		"/":              "content/html/index.html",
		"/index.html":    "content/html/index.html",
		"/404.html":      "content/html/404.html",
		"/index-ui.html": "content/html/index-ui.html",

		// css
		"/css/main.css":                "content/css/main.css",
		"/css/jquery-ui.css":           "content/css/jquery-ui.css",
		"/css/jquery-ui.structure.css": "content/css/jquery-ui.structure.css",

		// js
		"/js/jquery-3.6.0.min.js": "content/js/jquery-3.6.0.min.js",
		"/js/jquery-ui.min.js":    "content/js/jquery-ui.min.js",
		"/js/jquery-ui.js":        "content/js/jquery-ui.js",

		// image
		"/img/ui-bg_diagonals-thick_90_eeeeee_40x40.png": "content/img/ui-bg_diagonals-thick_90_eeeeee_40x40.png",
		"/img/ui-bg_glass_50_3baae3_1x400.png":           "content/img/ui-bg_glass_50_3baae3_1x400.png",
		"/img/ui-bg_glass_80_d7ebf9_1x400.png":           "content/img/ui-bg_glass_80_d7ebf9_1x400.png",
		"/img/ui-bg_glass_100_e4f1fb_1x400.png":          "content/img/ui-bg_glass_100_e4f1fb_1x400.png",
		"/img/ui-bg_highlight-hard_70_000000_1x100.png":  "content/img/ui-bg_highlight-hard_70_000000_1x100.png",
		"/img/ui-bg_highlight-hard_100_f2f5f7_1x100.png": "content/img/ui-bg_highlight-hard_100_f2f5f7_1x100.png",
		"/img/ui-bg_highlight-soft_25_ffef8f_1x100.png":  "content/img/ui-bg_highlight-soft_25_ffef8f_1x100.png",
		"/img/ui-bg_highlight-soft_100_deedf7_1x100.png": "content/img/ui-bg_highlight-soft_100_deedf7_1x100.png",
		"/img/ui-icons_2e83ff_256x240.png":               "content/img/ui-icons_2e83ff_256x240.png",
		"/img/ui-icons_3d80b3_256x240.png":               "content/img/ui-icons_3d80b3_256x240.png",
		"/img/ui-icons_72a7cf_256x240.png":               "content/img/ui-icons_72a7cf_256x240.png",
		"/img/ui-icons_2694e8_256x240.png":               "content/img/ui-icons_2694e8_256x240.png",
		"/img/ui-icons_ffffff_256x240.png":               "content/img/ui-icons_ffffff_256x240.png",
	}
)

func Serve(w http.ResponseWriter, req *http.Request) {
	log.Println(req.URL.Path)

	page, ok := pages[req.URL.Path]

	if !ok {
		w.Header().Set("Content-Type", "text/html")
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, "Page not found 404")
		return
	}

	w.Header().Set("Content-Type", contentType(page))
	w.WriteHeader(http.StatusOK)

	bytes, error := res.ReadFile(page)
	if error != nil {
		fmt.Fprintf(w, "Server error %s", error)
	}

	w.Write(bytes)
}

func contentType(page string) string {
	result := "text/html"

	if strings.HasSuffix(page, ".css") {
		result = "text/css"
	} else if strings.HasSuffix(page, ".js") {
		result = "application/javascript"
	} else if strings.HasSuffix(page, ".png") {
		result = "image/png"
	}

	return result
}
