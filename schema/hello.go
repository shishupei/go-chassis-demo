package schema

import (
	"net/http"

	"math/rand"

	rf "github.com/go-chassis/go-chassis/v2/server/restful"
)

var num = rand.Intn(100)

// RestFulHello is a struct used for implementation of restfull hello program
type RestFulHello struct {
}

// Sayhello is a method used to reply user with hello
func (r *RestFulHello) Sayhello(b *rf.Context) {
	b.Write([]byte("hello"))
}

// URLPatterns helps to respond for corresponding API calls
func (r *RestFulHello) URLPatterns() []rf.Route {
	return []rf.Route{
		{Method: http.MethodGet, Path: "/hello", ResourceFunc: r.Sayhello,
			Returns: []*rf.Returns{{Code: 200}}},
	}
}
