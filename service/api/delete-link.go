package api

import (
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
	"github.com/notherealmarco/overLinks/service/api/helpers"
	"github.com/notherealmarco/overLinks/service/api/reqcontext"
)

func (rt *_router) deleteLink(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	// we still have no auth

	//// check if the user is changing his own username
	//if !authorization.SendAuthorizationError(ctx.Auth.UserAuthorized, uid, rt.db, w, rt.baseLogger, http.StatusNotFound) {
	//	return
	//}

	id_str := ps.ByName("id")
	id, err := strconv.ParseInt(id_str, 10, 64)

	if err != nil {
		helpers.SendBadRequestError(err, "Invalid id", w, rt.baseLogger)
		return
	}

	// here we should do some validity checks on the input

	err = rt.db.DeleteLink(id)

	if err != nil {
		helpers.SendInternalError(err, "Database error: UpdateUsername", w, rt.baseLogger)
		return
	}

	helpers.SendStatus(http.StatusOK, w, "Resource deleted", rt.baseLogger)
}
