package auth_helpers

import (
	"bytes"
	"encoding/json"
	"errors"
	"mime"
	"mime/multipart"
	"net/url"
	"strconv"
	"strings"

	"github.com/big-larry/suckhttp"
)

func GetPerms[T ~int](r *suckhttp.Request) (*AuthResultWithPerms[T], error) {
	if result := r.GetHeader("x-perm"); result != "" {
		var r *AuthResultWithPerms[T]
		if err := json.Unmarshal([]byte(result), &r); err != nil {
			return nil, err
		}
		return r, nil
	}
	return nil, errors.New("Now found permissions in request")
}

func ParseForm(r *suckhttp.Request) (url.Values, error) {
	return url.ParseQuery(string(r.Body))
}

func ParseMultipartForm(r *suckhttp.Request) (*multipart.Reader, error) {
	v := r.GetHeader(suckhttp.Content_Type)
	if v == "" {
		return nil, errors.New("ErrNotMultipart")
	}
	allowMixed := true
	d, params, err := mime.ParseMediaType(v)
	if err != nil || !(d == "multipart/form-data" || allowMixed && d == "multipart/mixed") {
		return nil, errors.New("ErrNotMultipart")
	}
	boundary, ok := params["boundary"]
	if !ok {
		return nil, errors.New("ErrMissingBoundary")
	}
	body := bytes.NewReader(r.Body)
	return multipart.NewReader(body, boundary), nil
}

func TryGetFormRawValue(form map[string][]string, key string) (string, bool) {
	if len(form) == 0 {
		return "", false
	}
	if v, ok := form[key]; ok {
		s := ""
		if len(v) > 0 {
			s = strings.TrimSpace(v[0])
		}
		return s, true
	}
	return "", false
}

func GetIntValue(form map[string][]string, key string) (int, error) {
	if v, ok := TryGetFormRawValue(form, key); ok {
		return strconv.Atoi(v)
	}
	return 0, errors.New("Empty")
}

func GetOriginalUri(r *suckhttp.Request) (*url.URL, error) {
	if result := r.GetHeader("x-original-uri"); result != "" {
		return url.Parse(result)
	}
	return &r.Uri, nil
}

func GetRequestId(r *suckhttp.Request) (string, error) {
	if result := r.GetHeader("x-request-id"); result != "" {
		return result, nil
	}
	return "", errors.New("RequestId not set in request")
}
