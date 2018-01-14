package ltropnr

import (
	"fmt"
	"html/template"
	"io"
	"io/ioutil"
	"os/exec"
	"runtime"

	"github.com/pkg/errors"
)

const (
	tmpPrefix = "ltropnr_"
)

// Sender is the interface what wraps basic Send method.
type Sender interface {
	Send(r io.Reader) error
}

// Option for initializer.
type Option func(*Opener)

// LightMode option sets layout as the lightTemplate.
func LightMode() Option {
	return func(opnr *Opener) {
		opnr.layout = template.Must(template.New("layout").Parse(lightTemplate))
	}
}

// Opener holds layout and command name
// to open default browser.
type Opener struct {
	cmdName string
	layout  *template.Template
}

// New returns initialized Opener.
func New(options ...Option) *Opener {
	opnr := Opener{
		cmdName: cmdName(),
		layout:  template.Must(template.New("layout").Parse(defaultTemplate)),
	}

	for _, option := range options {
		option(&opnr)
	}

	return &opnr
}

// Send implements the Sender interface. Send will
// create a tmp file, execute layout template with
// given data into it and then open it in browser.
func (opnr *Opener) Send(msg io.Reader) error {
	tmpfile, err := tmpFile()
	if err != nil {
		return errors.Wrap(err, "tempfile creation failed")
	}
	defer tmpfile.Close()

	if err := opnr.prepareFile(tmpfile, msg); err != nil {
		return errors.Wrap(err, "tmp file perpare")
	}

	url := fmt.Sprintf("file:///%s", tmpfile.Name())
	cmd := exec.Command(opnr.cmdName, url)
	if err := cmd.Run(); err != nil {
		return errors.Wrap(err, "open letter in the browser failed")
	}

	return nil
}

func (opnr *Opener) prepareFile(w io.Writer, msg io.Reader) error {
	templateData, err := newTemplateData(msg)
	if err != nil {
		return errors.Wrap(err, "new template data")
	}

	if err := opnr.layout.Execute(w, templateData); err != nil {
		return errors.Wrap(err, "template execution failed")
	}

	return nil
}

type writeCloseNamer interface {
	io.WriteCloser
	Name() string
}

func tmpFile() (writeCloseNamer, error) {
	tmpfile, err := ioutil.TempFile("", tmpPrefix)
	if err != nil {
		return tmpfile, err
	}

	return tmpfile, nil
}

func cmdName() string {
	switch runtime.GOOS {
	case "darwin":
		return "open"
	case "windows":
		return "cmd /c start"
	default:
		return "xdg-open"
	}
}
