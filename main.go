package main

import (
	"errors"
	_ "expvar"
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"

	"github.com/bakins/net-http-recover"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/gorilla/rpc"
	"github.com/gorilla/rpc/json"
	"github.com/justinas/alice"
)

// modified versions of examples found at http://golang.org/pkg/net/rpc/
type Args struct {
	A, B int
}

type Quotient struct {
	Quo, Rem int
}

type Arith int

func (t *Arith) Multiply(r *http.Request, args *Args, reply *int) error {
	*reply = args.A * args.B
	return nil
}

func (t *Arith) Divide(r *http.Request, args *Args, quo *Quotient) error {
	if args.B == 0 {
		return errors.New("divide by zero")
	}
	quo.Quo = args.A / args.B
	quo.Rem = args.A % args.B
	return nil
}

func main() {

	r := mux.NewRouter()

	s := rpc.NewServer()
	s.RegisterCodec(json.NewCodec(), "application/json")

	arith := new(Arith)
	s.RegisterService(arith, "")

	chain := alice.New(
		func(h http.Handler) http.Handler {
			return handlers.CombinedLoggingHandler(os.Stdout, h)
		},
		handlers.CompressHandler,
		func(h http.Handler) http.Handler {
			return recovery.Handler(os.Stderr, h, true)
		})

	// httprof and expvar endpoints
	r.PathPrefix("/debug/").Handler(chain.Then(http.DefaultServeMux))

	r.Handle("/rpc", chain.Then(s))
	log.Fatal(http.ListenAndServe(":8080", r))
}
