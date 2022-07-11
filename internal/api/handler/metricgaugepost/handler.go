package metricgaugepost

import (
	"net/http"
	"strconv"

	"github.com/IgorAleksandroff/yp-musthave-devops/internal/pkg/metricscollection"
	"github.com/gorilla/mux"
)

type handler struct {
	metricsUC metricscollection.Usecase
}

func New(
	metricsUC metricscollection.Usecase,
) *handler {
	return &handler{
		metricsUC: metricsUC,
	}
}

func (h *handler) Handle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	metricName := vars["NAME"]
	metricValue, err := strconv.ParseFloat(vars["VALUE"], 64)
	if err != nil {
		http.Error(w, "can't parse a float64. internal error", http.StatusBadRequest)
		return
	}

	h.metricsUC.SaveGaugeMetric(metricName, metricValue)

	w.WriteHeader(http.StatusOK)
}
