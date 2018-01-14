package ltropnr

import (
	"html/template"
	"io"
	"io/ioutil"
	"net/mail"

	"github.com/pkg/errors"
)

type templateData struct {
	headers mail.Header
	Body    string
}

// newTemplateData checks required keys in mailData map
// and return initialized templateData.
func newTemplateData(r io.Reader) (*templateData, error) {
	msg, err := mail.ReadMessage(r)
	if err != nil {
		return nil, errors.Wrap(err, "read message")
	}

	body, err := ioutil.ReadAll(msg.Body)
	if err != nil {
		return nil, errors.Wrap(err, "read message body")
	}

	td := templateData{
		headers: msg.Header,
		Body:    string(body),
	}

	return &td, nil
}

func (t *templateData) GetHeader(key string) string {
	return t.headers.Get(key)
}

func (t *templateData) HTML(value string) template.HTML {
	return template.HTML(value)
}
