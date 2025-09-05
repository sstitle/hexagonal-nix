package driven

import "net/http"

// HTTPPresenter writes the greeting to an http.ResponseWriter.
type HTTPPresenter struct {
	w http.ResponseWriter
}

func NewHTTPPresenter(w http.ResponseWriter) *HTTPPresenter {
	return &HTTPPresenter{w: w}
}

func (p *HTTPPresenter) PresentGreeting(message string) {
	p.w.Header().Set("Content-Type", "text/html; charset=utf-8")
	_, _ = p.w.Write([]byte("<!doctype html><html><body>" +
		"<h1>" + message + "</h1>" +
		"<p><a href=\"/\">Back</a></p>" +
		"</body></html>"))
}
