package handlers

import (
	"errors"
	"net/http"
	"v3/pkg/httpcore"
	"v3/pkg/models"
)

// CreateGPS godoc
// @Summary Create a new gps
// @Description Create a new gps with device data and gps values
// @Accept  json
// @Produce  json
// @Param   gps  body   models.GPS  true  "GPS data"
// @Success 201 {object} models.GPS
// @Failure 400 {object} httpcore.ApiError
// @Router /telemetry/gps [post]
func (tc *ApiController) CreateGPS(w http.ResponseWriter, r *http.Request) (any, int) {
	newGPS, err := httpcore.DecodeBody[models.GPS](w, r)
	if err != nil {
		return httpcore.ErrBadRequest.With(err), http.StatusBadRequest
	}

	if newGPS.DeviceData == nil {
		return httpcore.ErrBadRequest.With(errors.New("device cannot be nil")), http.StatusBadRequest
	}

	g, err := models.NewGPS(
		newGPS.DeviceData,
		newGPS.Latitude,
		newGPS.Longitude,
	)

	if err != nil {
		return httpcore.ErrBadRequest.With(err), http.StatusBadRequest
	}

	if tc.db == nil {
		tc.db = make([]DataModel, 0)
	}

	tc.db = append(tc.db, g)

	return g, http.StatusCreated
}
