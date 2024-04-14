package tests

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

type testTable struct {
	Name           string
	Method         string
	Route          string
	Body           string
	Headers        [][2]string
	ExpectedStatus int
	ParseResp      bool
	RespBody       interface{}
}

type auth struct {
	Token string `json:"token"`
}

func createRequest(method, route, body string, headers [][2]string) (*http.Request, error) {
	host := os.Getenv("HTTP_HOST")
	if len(host) == 0 {
		return nil, errors.New("http host not found")
	}

	port := os.Getenv("HTTP_PORT")
	if len(port) == 0 {
		return nil, errors.New("http port not found")
	}

	url := fmt.Sprintf("http://%s:%s%s", host, port, route)

	req, err := http.NewRequest(method, url, strings.NewReader(body))
	if err != nil {
		return nil, err
	}

	for _, header := range headers {
		req.Header.Add(header[0], header[1])
	}

	return req, nil
}

func sendRequest(t *testing.T, client *http.Client, req *http.Request, expectedStatus int, parseResp bool, respBody interface{}) {
	t.Helper()
	resp, err := client.Do(req)
	require.NoError(t, err)

	require.Equal(t, expectedStatus, resp.StatusCode)
	if parseResp {
		err = json.NewDecoder(resp.Body).Decode(respBody)
		require.NoError(t, err)
	}
	defer resp.Body.Close() //nolint
}

func TestCreateUser(t *testing.T) {
	client := http.Client{}

	var authData auth

	tests := []testTable{
		{
			Name:           "Create User no Content-Type header",
			Method:         http.MethodPost,
			Route:          "/register",
			Body:           "{\"name\": \"Sarasti\",\"password\": \"password\",\"role\": \"ADMIN\"}",
			Headers:        [][2]string{},
			ExpectedStatus: http.StatusBadRequest,
			ParseResp:      false,
			RespBody:       &authData,
		},
		{
			Name:           "Create User no invalid Content-Type",
			Method:         http.MethodPost,
			Route:          "/register",
			Body:           "{\"name\": \"Sarasti\",\"password\": \"password\",\"role\": \"ADMIN\"}",
			Headers:        [][2]string{{"Content-Type", "invalidType"}},
			ExpectedStatus: http.StatusBadRequest,
			ParseResp:      false,
			RespBody:       &authData,
		},
		{
			Name:           "Create User without name",
			Method:         http.MethodPost,
			Route:          "/register",
			Body:           "{\"password\": \"password\",\"role\": \"ADMIN\"}",
			Headers:        [][2]string{{"Content-Type", "application/json"}},
			ExpectedStatus: http.StatusBadRequest,
			ParseResp:      false,
			RespBody:       &authData,
		},
		{
			Name:           "Create User without password",
			Method:         http.MethodPost,
			Route:          "/register",
			Body:           "{\"name\": \"Sarasti\",\"role\": \"ADMIN\"}",
			Headers:        [][2]string{{"Content-Type", "application/json"}},
			ExpectedStatus: http.StatusBadRequest,
			ParseResp:      false,
			RespBody:       &authData,
		},
		{
			Name:           "Create User no Role",
			Method:         http.MethodPost,
			Route:          "/register",
			Body:           "{\"name\": \"Sarasti\", \"password\": \"password\"}",
			Headers:        [][2]string{{"Content-Type", "application/json"}},
			ExpectedStatus: http.StatusBadRequest,
			ParseResp:      false,
			RespBody:       &authData,
		},
		{
			Name:           "Create User incorrect Role",
			Method:         http.MethodPost,
			Route:          "/register",
			Body:           "{\"name\": \"Sarasti\", \"password\": \"password\", \"role\": \"UNKNOWN\"}",
			Headers:        [][2]string{{"Content-Type", "application/json"}},
			ExpectedStatus: http.StatusBadRequest,
			ParseResp:      false,
			RespBody:       &authData,
		},
		{
			Name:           "Create User Success",
			Method:         http.MethodPost,
			Route:          "/register",
			Body:           "{\"name\": \"User_test\", \"password\": \"password\", \"role\": \"USER\"}",
			Headers:        [][2]string{{"Content-Type", "application/json"}},
			ExpectedStatus: http.StatusCreated,
			ParseResp:      true,
			RespBody:       &authData,
		},
		{
			Name:           "Create User Already Exists",
			Method:         http.MethodPost,
			Route:          "/register",
			Body:           "{\"name\": \"User_test\", \"password\": \"password\", \"role\": \"USER\"}",
			Headers:        [][2]string{{"Content-Type", "application/json"}},
			ExpectedStatus: http.StatusBadRequest,
			ParseResp:      false,
			RespBody:       &authData,
		},
		{
			Name:           "Create Admin Success",
			Method:         http.MethodPost,
			Route:          "/register",
			Body:           "{\"name\": \"Admin_test\", \"password\": \"password\", \"role\": \"ADMIN\"}",
			Headers:        [][2]string{{"Content-Type", "application/json"}},
			ExpectedStatus: http.StatusCreated,
			ParseResp:      true,
			RespBody:       &authData,
		},
		{
			Name:           "Create User Already Exists",
			Method:         http.MethodPost,
			Route:          "/register",
			Body:           "{\"name\": \"Admin_test\", \"password\": \"password\", \"role\": \"ADMIN\"}",
			Headers:        [][2]string{{"Content-Type", "application/json"}},
			ExpectedStatus: http.StatusBadRequest,
			ParseResp:      false,
			RespBody:       &authData,
		},
	}

	for _, tc := range tests {
		t.Run(tc.Name, func(t *testing.T) {
			req, err := createRequest(tc.Method, tc.Route, tc.Body, tc.Headers)
			require.NoError(t, err)

			sendRequest(t, &client, req, tc.ExpectedStatus, tc.ParseResp, tc.RespBody)
			if tc.ParseResp {
				require.NotEmpty(t, len(authData.Token))
			}
		})
	}

}

func TestLogin(t *testing.T) {
	// cleanUp(t)
	client := http.Client{}

	var authData auth

	tests := []testTable{
		{
			Name:           "Login not registered",
			Method:         http.MethodPost,
			Route:          "/login",
			Body:           "{\"name\": \"Admin_2_test\", \"password\": \"password\", \"role\": \"ADMIN\"}",
			Headers:        [][2]string{{"Content-Type", "application/json"}},
			ExpectedStatus: http.StatusNotFound,
			ParseResp:      false,
			RespBody:       &authData,
		},
		{
			Name:           "Login no Content-type",
			Method:         http.MethodPost,
			Route:          "/login",
			Body:           "{\"name\": \"Admin_2_test\", \"password\": \"password\"}",
			Headers:        [][2]string{},
			ExpectedStatus: http.StatusBadRequest,
			ParseResp:      false,
			RespBody:       &authData,
		},
		{
			Name:           "Login empty name",
			Method:         http.MethodPost,
			Route:          "/login",
			Body:           "{\"password\": \"password\"}",
			Headers:        [][2]string{{"Content-Type", "application/json"}},
			ExpectedStatus: http.StatusBadRequest,
			ParseResp:      false,
			RespBody:       &authData,
		},
		{
			Name:           "Login empty password",
			Method:         http.MethodPost,
			Route:          "/login",
			Body:           "{\"name\": \"Admin_2_test\"}",
			Headers:        [][2]string{{"Content-Type", "application/json"}},
			ExpectedStatus: http.StatusBadRequest,
			ParseResp:      false,
			RespBody:       &authData,
		},
		{
			Name:           "Login empty body",
			Method:         http.MethodPost,
			Route:          "/login",
			Body:           "{}",
			Headers:        [][2]string{{"Content-Type", "application/json"}},
			ExpectedStatus: http.StatusBadRequest,
			ParseResp:      false,
			RespBody:       &authData,
		},
		{
			Name:           "Create User Success",
			Method:         http.MethodPost,
			Route:          "/register",
			Body:           "{\"name\": \"User_2_test\", \"password\": \"password\", \"role\": \"USER\"}",
			Headers:        [][2]string{{"Content-Type", "application/json"}},
			ExpectedStatus: http.StatusCreated,
			ParseResp:      true,
			RespBody:       &authData,
		},
		{
			Name:           "Login Success",
			Method:         http.MethodPost,
			Route:          "/login",
			Body:           "{\"name\": \"User_2_test\", \"password\": \"password\", \"role\": \"USER\"}",
			Headers:        [][2]string{{"Content-Type", "application/json"}},
			ExpectedStatus: http.StatusOK,
			ParseResp:      true,
			RespBody:       &authData,
		},
	}

	for _, tc := range tests {
		t.Log(tc.Name)

		req, err := createRequest(tc.Method, tc.Route, tc.Body, tc.Headers)
		require.NoError(t, err)

		sendRequest(t, &client, req, tc.ExpectedStatus, tc.ParseResp, tc.RespBody)
		if tc.ParseResp {
			require.NotEmpty(t, len(authData.Token))
		}
	}

}

func TestUserGetBanner(t *testing.T) {
	// cleanUp(t)
	client := http.Client{}

	var adminAuthData, userAuthData auth
	var banner, updatedBanner, cachedBanner json.RawMessage

	tests := []testTable{
		{
			Name:           "Get Banner no Auth",
			Method:         http.MethodGet,
			Route:          "/user_banner",
			Body:           "",
			Headers:        [][2]string{{"Content-Type", "application/json"}},
			ExpectedStatus: http.StatusUnauthorized,
			ParseResp:      false,
			RespBody:       nil,
		},
		{
			Name:           "Create Admin",
			Method:         http.MethodPost,
			Route:          "/register",
			Body:           "{\"name\": \"Admin_3_test\", \"password\": \"password\", \"role\": \"ADMIN\"}",
			Headers:        [][2]string{{"Content-Type", "application/json"}},
			ExpectedStatus: http.StatusCreated,
			ParseResp:      true,
			RespBody:       &adminAuthData,
		},
		{
			Name:           "Create User",
			Method:         http.MethodPost,
			Route:          "/register",
			Body:           "{\"name\": \"User_3_test\", \"password\": \"password\", \"role\": \"USER\"}",
			Headers:        [][2]string{{"Content-Type", "application/json"}},
			ExpectedStatus: http.StatusCreated,
			ParseResp:      true,
			RespBody:       &userAuthData,
		},
		{
			Name:           "Get no tag_id",
			Method:         http.MethodGet,
			Route:          "/user_banner?feature_id=1",
			Body:           "",
			Headers:        [][2]string{{"Content-Type", "application/json"}},
			ExpectedStatus: http.StatusBadRequest,
			ParseResp:      false,
			RespBody:       nil,
		},
		{
			Name:           "Get no feature_id",
			Method:         http.MethodGet,
			Route:          "/user_banner?tag_id=1",
			Body:           "",
			Headers:        [][2]string{{"Content-Type", "application/json"}},
			ExpectedStatus: http.StatusBadRequest,
			ParseResp:      false,
			RespBody:       nil,
		},
		{
			Name:           "Get empty query",
			Method:         http.MethodGet,
			Route:          "/user_banner",
			Body:           "",
			Headers:        [][2]string{{"Content-Type", "application/json"}},
			ExpectedStatus: http.StatusBadRequest,
			ParseResp:      false,
			RespBody:       nil,
		},
		{
			Name:           "Get tag_id is not a number",
			Method:         http.MethodGet,
			Route:          "/user_banner?tag_id=wasd",
			Body:           "",
			Headers:        [][2]string{{"Content-Type", "application/json"}},
			ExpectedStatus: http.StatusBadRequest,
			ParseResp:      false,
			RespBody:       nil,
		},
		{
			Name:           "Get feature_id is not a number",
			Method:         http.MethodGet,
			Route:          "/user_banner?feature_id=wasd",
			Body:           "",
			Headers:        [][2]string{{"Content-Type", "application/json"}},
			ExpectedStatus: http.StatusBadRequest,
			ParseResp:      false,
			RespBody:       nil,
		},
		{
			Name:           "Get banner not found",
			Method:         http.MethodGet,
			Route:          "/user_banner?feature_id=1&tag_id=1",
			Body:           "",
			Headers:        [][2]string{{"Content-Type", "application/json"}},
			ExpectedStatus: http.StatusNotFound,
			ParseResp:      false,
			RespBody:       nil,
		},
		{
			Name:           "Create banner (Required ADMIN)",
			Method:         http.MethodPost,
			Route:          "/banner",
			Body:           "{\"tag_ids\": [1,2,3],\"content\": {\"title\": \"banner title\",\"text\": \"text\",\"url\": \"https://example.com/banner\"},\"feature_id\": 1,\"is_active\": true}",
			Headers:        [][2]string{{"Content-Type", "application/json"}},
			ExpectedStatus: http.StatusCreated,
			ParseResp:      false,
			RespBody:       nil,
		},
		{
			Name:           "Get banner ok",
			Method:         http.MethodGet,
			Route:          "/user_banner?feature_id=1&tag_id=1",
			Body:           "",
			Headers:        [][2]string{{"Content-Type", "application/json"}},
			ExpectedStatus: http.StatusOK,
			ParseResp:      true,
			RespBody:       &banner,
		},
		{
			Name:           "Create  inactive banner (Required ADMIN)",
			Method:         http.MethodPost,
			Route:          "/banner",
			Body:           "{\"tag_ids\": [1,2,3],\"content\": {\"title\": \"banner title\",\"text\": \"text\",\"url\": \"https://example.com/banner\"},\"feature_id\": 2,\"is_active\": false}",
			Headers:        [][2]string{{"Content-Type", "application/json"}},
			ExpectedStatus: http.StatusCreated,
			ParseResp:      false,
			RespBody:       nil,
		},
		{
			Name:           "Get inactive banner",
			Method:         http.MethodGet,
			Route:          "/user_banner?feature_id=2&tag_id=1",
			Body:           "",
			Headers:        [][2]string{{"Content-Type", "application/json"}},
			ExpectedStatus: http.StatusNotFound,
			ParseResp:      false,
			RespBody:       nil,
		},
		{
			Name:           "Update banner (Required ADMIN)",
			Method:         http.MethodPatch,
			Route:          "/banner/1",
			Body:           "{\"content\": {\"title\": \"updated banner title\"}}",
			Headers:        [][2]string{{"Content-Type", "application/json"}},
			ExpectedStatus: http.StatusOK,
			ParseResp:      false,
			RespBody:       nil,
		},
		{
			Name:           "Get cached banner",
			Method:         http.MethodGet,
			Route:          "/user_banner?feature_id=1&tag_id=1",
			Body:           "",
			Headers:        [][2]string{{"Content-Type", "application/json"}},
			ExpectedStatus: http.StatusOK,
			ParseResp:      true,
			RespBody:       &cachedBanner,
		},
		{
			Name:           "Get updated banner (Required Updated Banner)",
			Method:         http.MethodGet,
			Route:          "/user_banner?feature_id=1&tag_id=1&use_last_revision=true",
			Body:           "",
			Headers:        [][2]string{{"Content-Type", "application/json"}},
			ExpectedStatus: http.StatusOK,
			ParseResp:      true,
			RespBody:       &updatedBanner,
		},
		{
			Name:           "Update banner (Required ADMIN)",
			Method:         http.MethodPatch,
			Route:          "/banner/1",
			Body:           "{\"content\": {\"title\": \"updated banner title\",\"text\": \"updated text\"}}",
			Headers:        [][2]string{{"Content-Type", "application/json"}},
			ExpectedStatus: http.StatusOK,
			ParseResp:      false,
			RespBody:       nil,
		},
		{
			Name:           "Get older revision (Required Updated Banner)",
			Method:         http.MethodGet,
			Route:          "/user_banner?feature_id=1&tag_id=1&revision_id=1",
			Body:           "",
			Headers:        [][2]string{{"Content-Type", "application/json"}},
			ExpectedStatus: http.StatusOK,
			ParseResp:      true,
			RespBody:       &cachedBanner,
		},
	}

	var token string
	for _, tc := range tests {
		t.Log(tc.Name)

		if strings.Contains(tc.Name, "(Required ADMIN)") {
			token = adminAuthData.Token
		} else {
			token = userAuthData.Token
		}

		req, err := createRequest(tc.Method, tc.Route, tc.Body, append(tc.Headers, [2]string{"token", token}))
		if err != nil {
			t.Log(err.Error())
		}
		require.NoError(t, err)

		sendRequest(t, &client, req, tc.ExpectedStatus, tc.ParseResp, tc.RespBody)
		if strings.Contains(tc.Name, "(Required Updated Banner)") {
			require.NotEqual(t, len(cachedBanner), len(updatedBanner))
		}
	}
}
