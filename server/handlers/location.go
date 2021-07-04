package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"morrah77.org/location_tracker/common"
	"morrah77.org/location_tracker/domain"
	"morrah77.org/location_tracker/server/utils"
	"net/http"
)

type LocationHandler struct {
	Storage common.Storage
	PathPrefix string
	logger *log.Logger
}

func NewLocationHandler(storage common.Storage, path string, logger *log.Logger) *LocationHandler {
	return &LocationHandler{
		Storage: storage,
		PathPrefix:    path,
		logger: logger,
	};
}

func(handler *LocationHandler) ServeHTTP(rw http.ResponseWriter, req *http.Request)  {
	handler.logger.Printf(`%v %v %v`, req.Method, req.URL.Path, req.URL.Query())
	switch req.Method {
	case http.MethodGet:
		handler.HandleGet(rw, req)
		break
	case http.MethodPut:
		handler.HandlePut(rw, req)
		break
	case http.MethodDelete:
		handler.HandleDelete(rw, req)
		break
	default:
		respondWithError(http.StatusNotFound, ``, rw)
	}
}

func(handler *LocationHandler) HandleGet(rw http.ResponseWriter, req *http.Request) {
	orderId, err := utils.GetOrderId(req.URL.Path, handler.PathPrefix);
	if err != nil {
		respondWithError(http.StatusBadRequest, err.Error(), rw)
		return
	}
	depth, err := utils.GetDepth(req.URL, `max`);
	if err != nil {
		respondWithError(http.StatusBadRequest, err.Error(), rw)
		return
	}
	handler.logger.Printf(`orderId: %v, depth: %v\n`, orderId, depth)
	resp, err := handler.Storage.Fetch(orderId, depth)
	handler.logger.Printf(`Response: %#v, error: %#v`, resp, err)
	if err != nil {
		respondWithError(http.StatusNotFound, err.Error(), rw)
		return
	}

	var response domain.OrderHistory = domain.OrderHistory{
		OrderId: orderId,
		History: resp.([]*domain.Location),
	}
	b, err := json.Marshal(response)
	if err != nil {
		respondWithError(http.StatusInternalServerError, ``, rw)
		return
	}
	rw.WriteHeader(200);
	rw.Write(b)
}

func(handler *LocationHandler) HandlePut(rw http.ResponseWriter, req *http.Request) {
	orderId, err := utils.GetOrderId(req.URL.Path, handler.PathPrefix);
	if err != nil {
		respondWithError(http.StatusBadRequest, err.Error(), rw)
		return
	}
	handler.logger.Printf(`orderId: %v\n`, orderId)
	b, err := ioutil.ReadAll(req.Body)
	if err != nil {
		respondWithError(http.StatusInternalServerError, ``, rw)
		return
	}
	entity := &domain.Location{};
	err = json.Unmarshal(b, entity)
	if err != nil {
		respondWithError(http.StatusBadRequest, `Wrong location!`, rw)
		return
	}
	err = handler.Storage.Store(orderId, entity)

	if err != nil {
		respondWithError(http.StatusInternalServerError, ``, rw)
		return
	}

	rw.WriteHeader(200);

}

func(handler *LocationHandler) HandleDelete(rw http.ResponseWriter, req *http.Request) {
	orderId, err := utils.GetOrderId(req.URL.Path, handler.PathPrefix);
	if err != nil {
		respondWithError(http.StatusBadRequest, err.Error(), rw)
		return
	}
	handler.logger.Printf(`orderId: %v\n`, orderId)
	err = handler.Storage.Delete(orderId)

	if err != nil {
		respondWithError(http.StatusNotFound, err.Error(), rw)
		return
	}

	rw.WriteHeader(200);

}

func respondWithError(code int, message string, rw http.ResponseWriter)  {
	switch code {
	case http.StatusInternalServerError:
		rw.WriteHeader(code);
	case http.StatusBadRequest:
		rw.WriteHeader(code);
	case http.StatusNotFound:
		rw.WriteHeader(code);
	default:
		rw.WriteHeader(http.StatusInternalServerError);
	}
	if message == `` {
		rw.Write([]byte(`{"error":"Something went wrong!"}`))
		return
	}
	rw.Write([]byte(fmt.Sprintf(`{"error":"%s"}`, message)))
}
