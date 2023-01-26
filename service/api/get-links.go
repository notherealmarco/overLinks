package api

import (
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/notherealmarco/overLinks/service/api/helpers"
	"github.com/notherealmarco/overLinks/service/api/reqcontext"
)

func (rt *_router) getLinks(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	// we still have no auth

	//if !authorization.SendErrorIfNotLoggedIn(ctx.Auth.Authorized, rt.db, w, rt.baseLogger) ||
	//	!helpers.SendNotFoundIfBanned(rt.db, ctx.Auth.GetUserID(), uid, w, rt.baseLogger) {
	//	return
	//}

	// Get user profile
	links, err := rt.db.GetLinks()

	if err != nil {
		helpers.SendInternalError(err, "Database error: GetLinks", w, rt.baseLogger)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(links)

	if err != nil {
		helpers.SendInternalError(err, "Error encoding json", w, rt.baseLogger)
		return
	}
}
