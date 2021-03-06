// Package cli implements a colored text handler suitable for command-line interfaces.
package cli

import (
	"fmt"
	"io"
	"os"
	"sort"
	"sync"
	"time"

	"github.com/yewno/log"
)

// Default handler outputting to stderr.
var Default = New(os.Stderr)

// start time.
var start = time.Now()

// colors.
const (
	none   = 0
	red    = 31
	green  = 32
	yellow = 33
	blue   = 34
	gray   = 37
)

// Colors mapping.
var Colors = [...]int{
	log.DebugLevel: blue,
	log.InfoLevel:  green,
	log.WarnLevel:  yellow,
	log.ErrorLevel: red,
	log.FatalLevel: red,
}

// Strings mapping.
var Strings = [...]string{
	log.DebugLevel: "•",
	log.InfoLevel:  "[INFO]",
	log.WarnLevel:  "[WARNING]",
	log.ErrorLevel: "[ERROR]",
	log.FatalLevel: "[FATAL]",
}

// field used for sorting.
type field struct {
	Name  string
	Value interface{}
}

// by sorts projects by call count.
type byName []field

func (a byName) Len() int           { return len(a) }
func (a byName) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a byName) Less(i, j int) bool { return a[i].Name < a[j].Name }

// Handler implementation.
type Handler struct {
	mu      sync.Mutex
	Writer  io.Writer
	Padding int
}

// New handler.
func New(w io.Writer) *Handler {
	return &Handler{
		Writer:  w,
		Padding: 3,
	}
}

// HandleLog implements log.Handler.
func (h *Handler) HandleLog(e *log.Entry) error {
	color := Colors[e.Level]
	level := Strings[e.Level]

	var fields []field

	for k, v := range e.Fields {
		fields = append(fields, field{k, v})
	}

	sort.Sort(byName(fields))

	h.mu.Lock()
	defer h.mu.Unlock()
	tm := time.Now().Format("2006-01-02 15:04:05")

	fmt.Fprintf(h.Writer, "%v \033[%dm%*s\033[0m %-25s", tm, color, h.Padding+1, level, e.Message)

	for _, f := range fields {
		fmt.Fprintf(h.Writer, " \033[%dm%s\033[0m=%v", color, f.Name, f.Value)
	}

	fmt.Fprintln(h.Writer)

	return nil
}
