package server

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"html/template"
	"io"
	"net/http"
	root "skynet/pkg"
	"strconv"
)

type claimRouter struct {
	claimService root.ClaimService
	port         string
}

// NewRecordRouter create the router for Record schema
func NewClaimRouter(claim root.ClaimService, router *mux.Router, port string) *mux.Router {
	claimrt := claimRouter{claim, port}

	router.HandleFunc("/createClaimDefn", claimrt.createClaimDefnHandler).Methods("POST")
	router.HandleFunc("/createClaim", claimrt.createClaimHandler).Methods("POST")
	router.HandleFunc("/getClaimDefn", claimrt.getClaimDefn)
	router.HandleFunc("/displayAllClaimDefns", claimrt.displayAllClaimDefns)
	router.HandleFunc("/displayAllClaims", claimrt.displayAllClaims)

	return router
}

func (claim *claimRouter) createClaimDefnHandler(w http.ResponseWriter, r *http.Request) {

	result := make(map[string]string)

	for i := 1; i < 4; i++ {
		result[r.FormValue("attr"+strconv.Itoa(i))] = r.FormValue("type" + strconv.Itoa(i))
	}

	err := claim.claimService.CreateClaimDefn(result, r.FormValue("username"), r.FormValue("cname"))
	if err != nil {
		Error(w, http.StatusInternalServerError, err.Error())
		return
	}

	http.Redirect(w, r, "/display", 302)
}

func (claim *claimRouter) createClaimHandler(w http.ResponseWriter, r *http.Request) {

	var newClaim root.Claim

	/*	err := json.NewDecoder(r.Body).Decode(&newClaim)
		if err != nil {
			Error(w, http.StatusInternalServerError, err.Error())
		}
	*/

	attrToType := make(map[string]string)
	r.ParseForm()

	for key, value := range r.Form {

		switch key {
		case "username":
			newClaim.UserName = value[0]
		case "endpoint":
			newClaim.Endpoint = value[0]
		case "commonname":
			newClaim.CommonName = value[0]
		case "issuername":
			newClaim.IssuerName = value[0]
		default:
			attrToType[key] = value[0]
		}
	}
	newClaim.HashedData = attrToType
	err := claim.claimService.CreateClaim(&newClaim)
	if err != nil {
		Error(w, http.StatusInternalServerError, err.Error())
		return
	}
	http.Redirect(w, r, "/display", 302)
	//_, _ := http.Get("http://localhost" + claim.port + "/claim/enterData")

}

func (claimrt *claimRouter) getClaimDefn(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		renderTemplate(w, "accessAttributes", nil)
		return
	}

	IssuerName := r.FormValue("IssuerName")
	CommonName := r.FormValue("CommonName")
	claimDef, _ := claimrt.claimService.GetClaimDefnByCommonName(IssuerName, CommonName)
	tmpl, err := template.ParseFiles("createClaim.html")
	tmpl.Execute(w, claimDef.AttributesToType)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (claimrt *claimRouter) displayAllClaimDefns(w http.ResponseWriter, r *http.Request) {

	results, err := claimrt.claimService.GetAllClaimDefns()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	bytes, err := json.MarshalIndent(results, "", "")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	io.WriteString(w, string(bytes))
}

func (claimrt *claimRouter) displayAllClaims(w http.ResponseWriter, r *http.Request) {

	results, err := claimrt.claimService.GetAllClaims()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	bytes, err := json.MarshalIndent(results, "", "")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	io.WriteString(w, string(bytes))
}
