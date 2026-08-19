package httpapi

import (
	"chargeguard/internal/charging"
	"encoding/json"
	"net/http"
)

func TaskHTTPHandler(w http.ResponseWriter, r *http.Request) {
	station := charging.RestoreLegacyStation("legacy-1")
	if station == nil {
		http.Error(w, "station missing", http.StatusNotFound)
		return
	}
	station.Assign("rectification", "operator-a")
	owner := station.Owner("rectification")
	if owner == "" {
		http.Error(w, "assignment missing", http.StatusConflict)
		return
	}
	_ = json.NewEncoder(w).Encode(map[string]string{"owner": owner})
}
