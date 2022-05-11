package app

import (
	"encoding/json"
	"net/http"
	"strconv"

	"go.uber.org/zap"

	"github.com/Yandex-Practicum/pelemennaya/internal/logger"
	"github.com/Yandex-Practicum/pelemennaya/internal/store/dumplings"
)

func (i *Instance) ListDumplingsController(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	packs, err := i.store.ListProducts(ctx)
	if err != nil {
		logger.Log.Error("cannot fetch packs list", zap.Error(err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if len(packs) == 0 {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).
		Encode(map[string][]dumplings.Product{
			"results": packs,
		})

	for _, pack := range packs {
		i.dumplingsListingCounter.
			With(map[string]string{"id": strconv.Itoa(int(pack.ID))}).
			Inc()
	}
}