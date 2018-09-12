package handler

import (
	"net/http"

	"github.com/smjn/ipl18/backend/models"
	"github.com/smjn/ipl18/backend/util"
)

var PingHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	util.StructWriter(w, models.PingModel{true})
})
