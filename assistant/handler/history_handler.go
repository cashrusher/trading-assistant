package handler

import (
	"net/http"
	"derbysoft.com/derbysoft-rpc-go/log"
	"encoding/json"
)

func historyHandler(w http.ResponseWriter, r *http.Request) {
	history, err := assistant.getHistory()
	if err != nil {
		log.Error(err)
		w.Write(createErrorResponse(err))
	}
	h, err := json.Marshal(history)
	if err != nil {
		log.Error(err)
		w.Write(createErrorResponse(err))
	}
	w.Write(h)
}

func createErrorResponse(err error) []byte {
	return []byte(`{
		"status":"failed",
		"message":"` + err.Error() + `"
	}`)
}
