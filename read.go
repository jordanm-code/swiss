package swiss

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"reflect"
)

var ErrReaderNotFound = errors.New("reader could not be found")

// Read will read from [io.Reader] and return the content as a string.
//
// When a struct is provided, the first [io.Reader] that is encountered will
// be read.  [io.ReadCloser] will be closed after reading.
//
// When a struct pointer is provided, [io.Reader] and [io.ReadCloser] will
// automatically rewind by swapping out the reader with a new [io.ReadCloser] with
// a nopCloser or a [bytes.Reader] that contains the same content.
func Read(r interface{}) (string, error) {
	var b []byte
	var err error

	if reader, ok := r.(io.Reader); ok {
		b, err = io.ReadAll(reader)
		if err != nil {
			return "", fmt.Errorf("read error: %v", err)
		}
	}

	if rc, ok := r.(io.ReadCloser); ok {
		if err := rc.Close(); err != nil {
			return "", fmt.Errorf("close error: %v", err)
		}
	}

	if b != nil {
		return string(b), nil
	}

	var isPtr bool
	v := reflect.ValueOf(r)
	t := reflect.TypeOf(r)
	if v.Kind() == reflect.Ptr {
		isPtr = true
		v = v.Elem()
		t = t.Elem()
	}
	if v.Kind() == reflect.Struct {
		for i := 0; i < v.NumField(); i++ {
			if tf := t.Field(i); !tf.IsExported() {
				continue
			}

			f := v.Field(i)
			if f.Kind() == reflect.Interface {
				if reader, ok := f.Interface().(io.Reader); ok {
					s, err := Read(reader)
					if err != nil {
						return "", fmt.Errorf("nested read error: %v", err)
					}

					if isPtr {
						// Rewind the reader
						switch f.Type().String() {
						case "io.Reader":
							f.Set(reflect.ValueOf(bytes.NewReader([]byte(s))))
						case "io.ReadCloser":
							rr := io.NopCloser(bytes.NewBufferString(s))
							f.Set(reflect.ValueOf(rr))
						}
					}

					return s, nil
				}
			}
		}
	}

	return "", ErrReaderNotFound
}
