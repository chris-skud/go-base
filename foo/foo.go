package foo

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	log "github.com/Sirupsen/logrus"
	"github.com/chris-skud/go-base/correlation"
)

// Foo example service
type Foo struct{}

// New Foo
func New() *Foo { return &Foo{} }

// AddHandler demonstrates getting traceID for a log message
func (f *Foo) AddHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.WithField(correlation.KeyName,
			correlation.ID(r.Context()),
		).Infoln("AddHandler")

		fmt.Fprintf(w, f.AddFoo(r.Context(), 2, 3))
	})
}

// AddFoo example service method
func (f *Foo) AddFoo(ctx context.Context, op1, op2 int) string {
	return strconv.Itoa(op1 + op2)
}
