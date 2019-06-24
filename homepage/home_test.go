package homepage

import (
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	log "github.com/sirupsen/logrus"
)

// Table test
func TestHandlers_Handler(t *testing.T) {
	tests := []struct {
		name           string
		in             *http.Request
		out            *httptest.ResponseRecorder
		expectedStatus int
		expectedBody   string
	}{
		{
			"good",
			httptest.NewRequest("GET", "/", nil),
			httptest.NewRecorder(),
			http.StatusOK,
			message,
		}, // TODO: test cases
	}
	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			h := NewHandlers(nil)
			h.Home(test.out, test.in)
			if test.out.Code != test.expectedStatus {
				t.Logf("Expected: %d, got: %d", test.expectedStatus, test.out.Code)
				t.Fail()
			}

			if body := test.out.Body.String(); body != test.expectedBody {
				t.Logf("Expected: %q, got: %q", test.expectedBody, body)
				t.Fail()
			}

		})
	}
}

func TestHandlers_Home(t *testing.T) {
	type fields struct {
		logger *log.Logger
	}
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &Handlers{
				logger: tt.fields.logger,
			}
			h.Home(tt.args.w, tt.args.r)
		})
	}
}

func TestHandlers_Logger(t *testing.T) {
	type fields struct {
		logger *log.Logger
	}
	type args struct {
		next http.HandlerFunc
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   http.HandlerFunc
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &Handlers{
				logger: tt.fields.logger,
			}
			if got := h.Logger(tt.args.next); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Handlers.Logger() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHandlers_SetupRoutes(t *testing.T) {
	type fields struct {
		logger *log.Logger
	}
	type args struct {
		mux *http.ServeMux
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &Handlers{
				logger: tt.fields.logger,
			}
			h.SetupRoutes(tt.args.mux)
		})
	}
}

func TestNewHandlers(t *testing.T) {
	type args struct {
		logger *log.Logger
	}
	tests := []struct {
		name string
		args args
		want *Handlers
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewHandlers(tt.args.logger); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewHandlers() = %v, want %v", got, tt.want)
			}
		})
	}
}
