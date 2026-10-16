package runway_test

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"

	"github.com/psyb0t/go-runway"
)

// This is the README usage example. A local server stands in for
// https://api.dev.runwayml.com so the example runs without credentials.
func Example() {
	srv := httptest.NewServer(http.HandlerFunc(
		func(w http.ResponseWriter, _ *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(
				`{"creditBalance":1500,"tier":{},"usage":{}}`,
			))
		},
	))
	defer srv.Close()

	auth := runway.WithRequestEditorFn(
		func(_ context.Context, req *http.Request) error {
			req.Header.Set(
				"Authorization",
				"Bearer "+os.Getenv("RUNWAYML_API_SECRET"),
			)
			req.Header.Set("X-Runway-Version", "2024-11-06")

			return nil
		},
	)

	client, err := runway.NewClientWithResponses(srv.URL, auth)
	if err != nil {
		panic(err)
	}

	res, err := client.GetV1OrganizationWithResponse(context.Background())
	if err != nil {
		panic(err)
	}

	fmt.Println(res.JSON200.CreditBalance)
	// Output: 1500
}
