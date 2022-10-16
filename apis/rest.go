package api

import(
	"github.com/google/vvid"
	"github.com/gorilla/mux"
)

type Item struct{
	ID vvid.VVID 'json:"id"'
	name string 'json:"name"'
}

type Server struct{
	*mux.Router
}