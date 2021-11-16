package api

import (
	"net/http"
	"os"
	"runtime"
	"strconv"

	"github.com/sirupsen/logrus"
)

func infoHandler(w http.ResponseWriter, r *http.Request) {
	logger.WithFields(logrus.Fields{"api": "info"}).Infof("handling request")

	host, err := os.Hostname()
	if err != nil {
		logger.Errorf("can not retrieve host")
	}

	data := struct {
		Hostname     string `json:"hostname"`
		GOOS         string `json:"goos"`
		GOARCH       string `json:"goarch"`
		Runtime      string `json:"runtime"`
		NumGoroutine string `json:"num_goroutine"`
		NumCPU       string `json:"num_cpu"`
		MagicValue   string `json:"magic_value"`
	}{
		Hostname:     host,
		GOOS:         runtime.GOOS,
		GOARCH:       runtime.GOARCH,
		Runtime:      runtime.Version(),
		NumGoroutine: strconv.FormatInt(int64(runtime.NumGoroutine()), 10),
		NumCPU:       strconv.FormatInt(int64(runtime.NumCPU()), 10),
		MagicValue:   os.Getenv("MAGIC_VALUE"),
	}

	JSONResponse(w, r, data)
}
