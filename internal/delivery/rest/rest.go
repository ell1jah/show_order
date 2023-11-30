package rest

import (
	"encoding/json"
	"net/http"

	"github.com/ell1jah/show_order/internal/logic"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
)

type OrderRest struct {
	Logic  logic.Logic
	Logger zap.Logger
}

func (or *OrderRest) GetOrder(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	uid, ok := vars["ORDER_UID"]
	if !ok {
		or.Logger.Info("order rest error: no id")
		http.Error(w, "unknown error", http.StatusInternalServerError)
		return
	}

	order, err := or.Logic.GetByUID(uid)
	if err != nil {
		or.Logger.Sugar().Infof("order logic error: %w", err)
		http.Error(w, "can`t find order", http.StatusNotFound)
		return
	}

	resp, err := json.Marshal(order)
	if err != nil {
		or.Logger.Sugar().Infof("json marshal error: %w", err)
		http.Error(w, "can`t find order", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	_, err = w.Write(resp)
	if err != nil {
		or.Logger.Sugar().Infof("response write error: %w", err)
		http.Error(w, "can`t write response", http.StatusInternalServerError)
		return
	}
}
