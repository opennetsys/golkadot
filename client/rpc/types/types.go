package rpctypes

// NilArgs is used by methods that don't require any arguments.
//
// This is required because the gorpc Register method requires methods of two arguments. The first being arguments, the second being a response.
// https://godoc.org/github.com/libp2p/go-libp2p-gorpc#Server.Register
type NilArgs interface{}

// NilResponse is used by rpc methods that don't require a response.
//
// This is required because the gorpc Register method requires methods of two arguments. The first being arguments, the second being a response.
// https://godoc.org/github.com/libp2p/go-libp2p-gorpc#Server.Register
type NilResponse interface{}
