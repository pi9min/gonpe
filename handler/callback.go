package handler

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/golang/protobuf/proto"
	"google.golang.org/grpc/status"
)

func withErrorLog(_ context.Context, w http.ResponseWriter, r *http.Request, _, _ proto.Message, err error) {
	if err != nil {
		log.Printf("[ERROR] %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		p := status.New(status.Code(err), err.Error()).Proto()
		switch r.Header.Get("Content-Type") {
		case "application/protobuf", "application/x-protobuf":
			buf, err := proto.Marshal(p)
			if err != nil {
				return
			}
			if _, err := io.Copy(w, bytes.NewBuffer(buf)); err != nil {
				return
			}
		case "application/json":
			if err := json.NewEncoder(w).Encode(p); err != nil {
				return
			}
		default:
		}
	}
}
