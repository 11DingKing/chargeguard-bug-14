package httpapi

import (
	"chargeguard/internal/charging"
	"encoding/json"
	"net/http"
)

func TaskHTTPHandler(w http.ResponseWriter, r *http.Request) {
	station := charging.RestoreLegacyStation("legacy-1")
	station.Assign("rectification", "operator-a")
	_ = json.NewEncoder(w).Encode(map[string]string{"owner": station.Owner("rectification")})
}
