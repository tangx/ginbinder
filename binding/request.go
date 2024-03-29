package binding

import (
	"bytes"
	"errors"
	"io"
	"net/http"
	"reflect"
)

var ErrInvalidTagInRequestBody = errors.New("body struct should not contain tag `query`, `header`, `cookie`, `uri` in binding request api")

type requestBinding struct{}

func (requestBinding) Name() string {
	return "request"
}

func (b requestBinding) Bind(obj interface{}, req *http.Request, form map[string][]string) error {
	if err := b.BindOnly(obj, req, form); err != nil {
		return err
	}

	return validate(obj)
}

// requestClone to clone request as a new one, so that we can read data multiple times
// https://stackoverflow.com/a/62017757
func requestClone(r *http.Request) (*http.Request, error) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}

	r2 := r.Clone(r.Context())
	r.Body = io.NopCloser(bytes.NewBuffer(body))
	r2.Body = io.NopCloser(bytes.NewBuffer(body))

	err = r.ParseForm()
	if err != nil {
		return nil, err
	}

	return r2, nil
}

func (b requestBinding) BindOnly(obj interface{}, req *http.Request, uriMap map[string][]string) error {

	// r2 := req.Clone(req.Context())
	req, err := requestClone(req)
	if err != nil {
		return err
	}

	if err := Uri.BindOnly(uriMap, obj); err != nil {
		return err
	}

	// 以 path 为主
	if err := Path.BindOnly(uriMap, obj); err != nil {
		return err
	}

	if err := b.bindingQuery(req, obj); err != nil {
		return err
	}

	binders := []Binding{Header, Cookie}
	for _, binder := range binders {
		if err := binder.BindOnly(req, obj); err != nil {
			return err
		}
	}

	// body decode
	mime, bodyObj := extractBody(obj)
	if bodyObj == nil {
		return nil
	}

	// get Binding base on mime tag first,
	// if nil, use content-type default
	bb := MimeBinding(mime)

	if bb == nil {
		// default json
		contentType := req.Header.Get("Content-Type")
		if contentType == "" {
			contentType = MIMEJSON
		}
		bb = Default(req.Method, contentType)
	}

	err = bb.BindOnly(req, bodyObj)
	if err == nil || err == io.EOF {
		return nil
	}

	return err

}

func (b requestBinding) bindingQuery(req *http.Request, obj interface{}) error {
	values := req.URL.Query()
	return mapFormByTag(obj, values, "query")
}

// extractBody return body object
func extractBody(obj interface{}) (mime string, body interface{}) {

	// pre-check obj
	rv := reflect.ValueOf(obj)
	rv = reflect.Indirect(rv)
	if rv.Kind() != reflect.Struct {
		return "", nil
	}

	return extract(rv)
}

func extract(rv reflect.Value) (mime string, body interface{}) {

	typ := rv.Type()
	for i := 0; i < rv.NumField(); i++ {
		tf := typ.Field(i)
		vf := rv.Field(i)

		_, ok := tf.Tag.Lookup("body")
		if !ok {
			continue
		}

		// get mime tag
		mime := tf.Tag.Get("mime")
		// find body struct or map
		if vf := reflect.Indirect(vf); vf.Kind() == reflect.Struct ||
			vf.Kind() == reflect.Map {

			// body MUST NOT include the following tags
			for _, name := range []string{"query", "header", "cookie", "uri", "path"} {
				if hasTag(vf, name) {
					panic(ErrInvalidTagInRequestBody)
				}
			}

			return mime, vf.Addr().Interface()
		}

	}

	return mime, nil
}

func hasTag(rv reflect.Value, tag string) bool {
	rv = reflect.Indirect(rv)
	if rv.Kind() != reflect.Struct {
		return false
	}

	typ := rv.Type()
	for i := 0; i < typ.NumField(); i++ {
		_, ok := typ.Field(i).Tag.Lookup(tag)
		if ok {
			return true
		}
	}

	return false
}

// MimeBinding returns the appropriate Binding instance based on mime value in body tag
func MimeBinding(mime string) Binding {
	switch mime {
	case "json":
		return JSON
	case "xml":
		return XML
	case "yaml", "yml":
		return YAML
	case "form":
		return Form
	}
	return nil
}
