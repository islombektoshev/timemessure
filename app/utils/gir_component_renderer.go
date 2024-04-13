package utils

import (
	"context"
	"github.com/a-h/templ"
	"github.com/gin-gonic/gin/render"
	"net/http"
)

type GirComponentRenderer struct {
	templ.Component
}

func (g *GirComponentRenderer) Render(writer http.ResponseWriter) error {
	return g.Component.Render(context.Background(), writer)
}

func (g *GirComponentRenderer) WriteContentType(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "text/html")
}

func ToRender(c templ.Component) render.Render {
	return &GirComponentRenderer{c}
}
