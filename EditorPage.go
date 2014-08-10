package main

import (
	"html/template"
	"net/http"
	"path/filepath"
)

var (
	tmpl = template.Must(template.ParseGlob(filepath.Join("./views", "*.html")))
)

type EditorPage struct {
	filePath string
}

func NewEditor(file string) (*EditorPage, error) {
	return &EditorPage{filePath: file}, nil
}

func (editor *EditorPage) FilePath() string { return editor.filePath }

func (editor *EditorPage) Render(res http.ResponseWriter) error {
	return tmpl.ExecuteTemplate(res, "index.html", editor)
}
