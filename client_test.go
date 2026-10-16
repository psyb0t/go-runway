package runway_test

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/google/uuid"
	"github.com/psyb0t/go-runway"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const (
	testAPIKey     = "test-key"
	testAPIVersion = "2024-11-06"
	testTaskID     = "4b3c2a1d-0e9f-4a8b-9c7d-6e5f4a3b2c1d"
)

func newTestClient(
	t *testing.T,
	handler http.HandlerFunc,
) *runway.ClientWithResponses {
	t.Helper()

	srv := httptest.NewServer(handler)
	t.Cleanup(srv.Close)

	auth := runway.WithRequestEditorFn(
		func(_ context.Context, req *http.Request) error {
			req.Header.Set("Authorization", "Bearer "+testAPIKey)
			req.Header.Set("X-Runway-Version", testAPIVersion)

			return nil
		},
	)

	client, err := runway.NewClientWithResponses(srv.URL, auth)
	require.NoError(t, err)

	return client
}

func assertHeaders(t *testing.T, r *http.Request) {
	t.Helper()

	assert.Equal(t, "Bearer "+testAPIKey, r.Header.Get("Authorization"))
	assert.Equal(t, testAPIVersion, r.Header.Get("X-Runway-Version"))
}

func TestGetOrganization(t *testing.T) {
	t.Parallel()

	client := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, "/v1/organization", r.URL.Path)
		assertHeaders(t, r)

		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(
			`{"creditBalance":1500,` +
				`"tier":{"maxMonthlyCreditSpend":10000,"models":{}},` +
				`"usage":{"models":{}}}`,
		))
	})

	res, err := client.GetV1OrganizationWithResponse(context.Background())
	require.NoError(t, err)
	require.Equal(t, http.StatusOK, res.StatusCode())
	require.NotNil(t, res.JSON200)

	assert.Equal(t, 1500, res.JSON200.CreditBalance)
	assert.Equal(t, 10000, res.JSON200.Tier.MaxMonthlyCreditSpend)
}

func TestGetTask(t *testing.T) {
	t.Parallel()

	client := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, "/v1/tasks/"+testTaskID, r.URL.Path)
		assertHeaders(t, r)

		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(
			`{"id":"` + testTaskID + `","status":"PENDING",` +
				`"createdAt":"2026-01-02T03:04:05Z",` +
				`"estimatedCost":{"credits":12.5}}`,
		))
	})

	res, err := client.GetV1TasksIdWithResponse(
		context.Background(),
		uuid.MustParse(testTaskID),
	)
	require.NoError(t, err)
	require.Equal(t, http.StatusOK, res.StatusCode())
	require.NotNil(t, res.JSON200)

	task, err := res.JSON200.AsGetV1TasksId200JSONResponseBody0()
	require.NoError(t, err)

	assert.Equal(t, testTaskID, task.Id.String())
	assert.Equal(t, "PENDING", string(task.Status))
	assert.InDelta(t, 12.5, task.EstimatedCost.Credits, 0)
}

// Union request bodies go through the WithBody method as marshaled JSON; the
// typed JSONRequestBody method would send `{}` (see the README).
func TestPostTextToImageUnionBody(t *testing.T) {
	t.Parallel()

	client := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method)
		assert.Equal(t, "/v1/text_to_image", r.URL.Path)
		assert.Equal(t, "application/json", r.Header.Get("Content-Type"))
		assertHeaders(t, r)

		var got map[string]any

		assert.NoError(t, json.NewDecoder(r.Body).Decode(&got))
		assert.Equal(t, "gen4_image", got["model"])
		assert.Equal(t, "a test prompt", got["promptText"])
		assert.Equal(t, "1080:1080", got["ratio"])

		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(
			`{"id":"` + testTaskID + `","estimatedCost":{"credits":5}}`,
		))
	})

	var body runway.PostV1TextToImageJSONBody

	require.NoError(t, body.FromPostV1TextToImageJSONBody1(
		runway.PostV1TextToImageJSONBody1{
			Model:      runway.PostV1TextToImageJSONBody1ModelGen4Image,
			PromptText: "a test prompt",
			Ratio:      runway.PostV1TextToImageJSONBody1RatioN10801080,
		},
	))

	payload, err := json.Marshal(body)
	require.NoError(t, err)

	res, err := client.PostV1TextToImageWithBodyWithResponse(
		context.Background(),
		"application/json",
		bytes.NewReader(payload),
	)
	require.NoError(t, err)
	require.Equal(t, http.StatusOK, res.StatusCode())
	require.NotNil(t, res.JSON200)

	assert.Equal(t, testTaskID, res.JSON200.Id.String())
	assert.InDelta(t, 5.0, res.JSON200.EstimatedCost.Credits, 0)
}
