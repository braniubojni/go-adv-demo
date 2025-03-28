package stat

import (
	"go/adv-demo/configs"
	"go/adv-demo/pkg/middleware"
	"go/adv-demo/pkg/res"
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

const (
	GroupByMonth = "month"
	GroupByDay   = "day"
)

var BY = []string{GroupByMonth, GroupByDay}

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

		stats := handler.StatRepository.GetStats(by, from, to)
		res.Json(w, stats, http.StatusOK)
	}
}
