package relationship

import (
	i "github.com/luskaner/ageLANServer/server/internal"
	"net/http"
)

func ClearRelationship(w http.ResponseWriter, _ *http.Request) {
	i.JSON(&w, i.A{0})
}
