package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

type StubPlayerStore struct {
	scores map[string]int 
	winCalls []string
}

func (s *StubPlayerStore) GetPlayerScoreByName(name string) int {
	return s.scores[name]
}

func (s *StubPlayerStore) RecordWin(name string) {
	s.winCalls = append(s.winCalls, name)
}

func TestGetPlayer(t *testing.T) {
	playerScores := map[string]int{
		"Pepper": 20,
		"Floyd": 40,
	}

	store := &StubPlayerStore{scores: playerScores}
	server := &PlayerServer{store: store}
	t.Run("returns Pepper's score", func(t *testing.T) {
		req, res := newGetScoreRequest("Pepper")

		server.ServeHTTP(res, req)
		got := res.Body.String()
		want := "20"

		assertStatusCode(t, res.Code, http.StatusOK)
		assertResponseBody(t, got, want)
	})

	t.Run("returns Floyd's score", func(t *testing.T) {
		req, res := newGetScoreRequest("Floyd")

		server.ServeHTTP(res, req)
		got := res.Body.String()
		want := "40"

		assertStatusCode(t, res.Code, http.StatusOK)
		assertResponseBody(t, got, want)
	})

	t.Run("returns 404 status code for missing players", func(t *testing.T) {
		req, res := newGetScoreRequest("Aryan")

		server.ServeHTTP(res, req)

		assertStatusCode(t, res.Code, http.StatusNotFound)
	})
}

func TestPlayerWin(t *testing.T) {
	playerScore := map[string]int{}
	store := StubPlayerStore{scores: playerScore, winCalls: nil}
	server := &PlayerServer{&store}

	t.Run("records win on POST", func(t *testing.T) {
		req, res := newPostWinRequest("Floyd")

		server.ServeHTTP(res, req)

		assertStatusCode(t, res.Code, http.StatusAccepted)

		if len(store.winCalls) != 1 {
			t.Errorf("got %d calls to RecordWin, wanted %d", len(store.winCalls), 1)
		}
	})	
}

func newPostWinRequest(name string) (*http.Request, *httptest.ResponseRecorder) {
	request, _ := http.NewRequest(http.MethodPost, fmt.Sprintf("/player/%s", name), nil)
	response := httptest.NewRecorder()
	return request, response
}

func newGetScoreRequest(name string) (*http.Request, *httptest.ResponseRecorder) {
	request, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/player/%s", name), nil)
	response := httptest.NewRecorder()
	return request, response
}

func assertStatusCode(t testing.TB, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("incorrect status code recieved, got %v, wanted %v", got, want)
	}
}

func assertResponseBody(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}
