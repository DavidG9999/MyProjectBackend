package base

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"reflect"
)

func DecodeBody(r *http.Request, dest interface{}) error {
	expected := map[string]struct{}{}
	elem := reflect.ValueOf(dest).Elem()
	for i := 0; i < elem.NumField(); i++ {
		expected[string(elem.Type().Field(i).Tag.Get("json"))] = struct{}{}
	}
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return err
	}
	var jsonKeys map[string]interface{}
	if err := json.Unmarshal(body, &jsonKeys); err != nil {
		return err
	}
	for key := range expected {
		if _, ok := jsonKeys[key]; !ok {
			return fmt.Errorf("пропущенное поле: %s", key)
		}
	}
	for key := range jsonKeys {
		if _, ok := expected[key]; !ok {
			return fmt.Errorf("дополнительное поле %s", key)
		}
	}
	r.Body.Close()
	err = json.Unmarshal(body, &dest)
	return err
}
func EncodeErrorRespornse(_ context.Context, err error, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusInternalServerError)
	json.NewEncoder(w).Encode(map[string]string{
		"error": err.Error(),
	})
}

func EncodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(w).Encode(response)

}
