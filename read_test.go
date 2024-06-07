package swiss

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"testing"
)

func TestRead(t *testing.T) {
	type structWithReader struct {
		R io.Reader
	}

	type structWithUnexportedReader struct {
		r io.Reader
	}

	tests := []struct {
		name       string
		r          interface{}
		want       string
		wantErr    error
		wantRewind bool
	}{
		{
			name:    "Read string",
			r:       "read string",
			wantErr: ErrReaderNotFound,
		},
		{
			name:    "*structWithUnexportedReader",
			r:       &structWithUnexportedReader{r: bytes.NewReader([]byte("read struct with unexported reader"))},
			wantErr: ErrReaderNotFound,
		},
		{
			name: "io.ReadCloser",
			r:    io.NopCloser(bytes.NewReader([]byte("read io read closer"))),
			want: "read io read closer",
		},
		{
			name: "bytes.Reader",
			r:    bytes.NewReader([]byte("read bytes reader")),
			want: "read bytes reader",
		},
		{
			name:       "*http.Response",
			r:          &http.Response{Body: io.NopCloser(bytes.NewReader([]byte("read http response")))},
			want:       "read http response",
			wantRewind: true,
		},
		{
			name:       "*http.Request",
			r:          &http.Request{Body: io.NopCloser(bytes.NewReader([]byte("read http request")))},
			want:       "read http request",
			wantRewind: true,
		},
		{
			name:       "*structWithReader",
			r:          &structWithReader{R: bytes.NewReader([]byte("read struct with reader"))},
			want:       "read struct with reader",
			wantRewind: true,
		},
		{
			name: "http.Response",
			r:    http.Response{Body: io.NopCloser(bytes.NewReader([]byte("read http response")))},
			want: "read http response",
		},
		{
			name: "http.Request",
			r:    http.Request{Body: io.NopCloser(bytes.NewReader([]byte("read http request")))},
			want: "read http request",
		},
		{
			name: "structWithReader",
			r:    structWithReader{R: bytes.NewReader([]byte("read struct with reader"))},
			want: "read struct with reader",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got, err := Read(test.r)
			if err != test.wantErr {
				t.Errorf("Read() error = %v, wantErr %v", err, test.wantErr)
			}

			if got != test.want {
				t.Errorf("Read() = %v, want %v", got, test.want)
			}

			// re-read to test rewind
			got, err = Read(test.r)
			if err != test.wantErr {
				t.Fatalf("rewind: Read() unexpected error = %v", err)
			}

			if test.wantRewind && got != test.want {
				t.Errorf("rewind: Read() = %v, want %v", got, test.want)
			}

			if !test.wantRewind && got != "" {
				t.Errorf("rewind: Read() = %v, want %v", got, "")
			}
		})
	}
}

func ExampleRead() {
	request, err := http.NewRequest(http.MethodGet, "https://example.com", nil)
	if err != nil {
		panic(err)
	}
	request.Body = io.NopCloser(bytes.NewReader([]byte("<html></html>")))

	body, err := Read(request) // pass in *http.Request; will rewind
	if err != nil {
		panic(err)
	}
	fmt.Println(body)

	body, err = Read(request.Body) // pass in io.ReadCloser; will not rewind
	if err != nil {
		panic(err)
	}
	fmt.Println(body)

	body, err = Read(request) // pass in request; body is empty
	if err != nil {
		panic(err)
	}
	fmt.Println(body)

	// Output:
	// <html></html>
	// <html></html>
}
