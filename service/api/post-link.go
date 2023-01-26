package api

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/notherealmarco/overLinks/service/api/helpers"
	"github.com/notherealmarco/overLinks/service/api/reqcontext"
	"github.com/notherealmarco/overLinks/service/structures"
)

func (rt *_router) postLink(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	// we still have no auth

	//// check if the user is changing his own username
	//if !authorization.SendAuthorizationError(ctx.Auth.UserAuthorized, uid, rt.db, w, rt.baseLogger, http.StatusNotFound) {
	//	return
	//}

	// decode request body
	var req structures.OverLinkInput

	if !helpers.DecodeJsonOrBadRequest(r.Body, w, &req, rt.baseLogger) {
		return
	}

	// here we should do some validity checks on the input

	err := rt.db.AddLink(&req)

	if err != nil {
		helpers.SendInternalError(err, "Database error: UpdateUsername", w, rt.baseLogger)
		return
	}

	helpers.SendStatus(http.StatusCreated, w, "Resource created", rt.baseLogger)
}
