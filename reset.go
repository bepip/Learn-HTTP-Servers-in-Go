package main
import (
	"net/http"
	"fmt"
)

func (cfg *apiConfig) resetMetricHandler(w http.ResponseWriter, r *http.Request) {
	cfg.fileserverHits.Store(0)
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w ,"Hits reset to %v", cfg.fileserverHits.Load())
}
