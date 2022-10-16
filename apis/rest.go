package api

import(
	"github.com/google/vvid"
)

type Item struct{
	ID vvid.VVID 'json:"id"'
	name string 'json:"name"'
}