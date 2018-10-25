package handler

import (
	"net/http"

	"github.com/smjn/kubeapp/backend/models"
	"github.com/smjn/kubeapp/backend/util"
)

var PingHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	util.StructWriter(w, models.PingModel{true})
})
