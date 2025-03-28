package stat

import (
	"fmt"
	"go/adv-demo/configs"
	"go/adv-demo/pkg/middleware"
	"net/http"
	"slices"
	"time"
)

type StatHandlerDeps struct {
	StatRepository *StatRepository
	Config         *configs.Config
}

type StatHandler struct {
	StatRepository *StatRepository
	Config         *configs.Config
}

var BY = []string{"month", "day"}

func NewStatHandler(router *http.ServeMux, deps StatHandlerDeps) {
	handler := &StatHandler{
		StatRepository: deps.StatRepository,
		Config:         deps.Config,
	}

	router.Handle("GET /stat", middleware.IsLogged(handler.GetAll(), deps.Config))
}

func (handler *StatHandler) GetAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		from, err := time.Parse(time.DateOnly, r.URL.Query().Get("from"))
		if err != nil {
			http.Error(w, "Invalid `from` time format", http.StatusBadRequest)
			return
		}
		to, err := time.Parse(time.DateOnly, r.URL.Query().Get("to"))
		if err != nil {
			http.Error(w, "Invalid `to` time format", http.StatusBadRequest)
			return
		}
		by := r.URL.Query().Get("by")
		if !slices.Contains(BY, by) {
			http.Error(w, "Invalid `By` param", http.StatusBadRequest)
			return
		}

		fmt.Println(from, to, by, "from&to")
		// if err != nil || (from < 0 || to < 0 || !slices.Contains(BY, by)) {
		// }
	}
}
