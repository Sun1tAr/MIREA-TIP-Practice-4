package task

import (
	"encoding/json"
	"net/http"
	"strconv"

	"example.com/pz4-todo/pkg/middleware"
	"github.com/go-chi/chi/v5"
)

type Handler struct {
	repo *Repo
}

func NewHandler(repo *Repo) *Handler {
	return &Handler{repo: repo}
}

func (h *Handler) Routes() chi.Router {
	r := chi.NewRouter()
	r.Get("/", h.list)          // GET /tasks
	r.Post("/", h.create)       // POST /tasks
	r.Get("/{id}", h.get)       // GET /tasks/{id}
	r.Put("/{id}", h.update)    // PUT /tasks/{id}
	r.Delete("/{id}", h.delete) // DELETE /tasks/{id}
	r.Get("/health", h.Health)  // GET /health -- состояние на уровне версий
	return r
}

func (h *Handler) list(w http.ResponseWriter, r *http.Request) {
	// Получаем параметры из query string (?done=true&page=1&limit=10)
	query := r.URL.Query()

	// Обрабатываем параметр done
	if doneStr := query.Get("done"); doneStr != "" {
		isDone, err := strconv.ParseBool(doneStr)
		if err != nil {
			httpError(w, http.StatusBadRequest, "invalid done parameter")
			return
		}
		WriteJSON(w, http.StatusOK, h.repo.DoneList(isDone))
		return
	}

	// Обрабатываем пагинацию
	if pageStr := query.Get("page"); pageStr != "" {
		page, err1 := strconv.ParseInt(pageStr, 10, 64)
		limit, err2 := strconv.ParseInt(query.Get("limit"), 10, 64)
		if err1 != nil || err2 != nil || page <= 0 || limit <= 0 {
			httpError(w, http.StatusBadRequest, "invalid pagination parameters")
			return
		}
		WriteJSON(w, http.StatusOK, h.repo.PaginatedList(page, limit))
		return
	}

	// Если нет параметров - возвращаем все задачи
	WriteJSON(w, http.StatusOK, h.repo.List())
}

func (h *Handler) get(w http.ResponseWriter, r *http.Request) {
	id, bad := parseID(w, r)
	if bad {
		return
	}
	t, err := h.repo.Get(id)
	if err != nil {
		httpError(w, http.StatusNotFound, err.Error())
		return
	}
	WriteJSON(w, http.StatusOK, t)
}

type createReq struct {
	Title string `json:"title"`
}

func (h *Handler) create(w http.ResponseWriter, r *http.Request) {
	var req createReq
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil || req.Title == "" {
		httpError(w, http.StatusBadRequest, "invalid json: require non-empty title")
		return
	}
	if !middleware.TitleValidation(req.Title) {
		httpError(w, http.StatusBadRequest, "invalid title")
		return
	}
	t := h.repo.Create(req.Title)
	WriteJSON(w, http.StatusCreated, t)
}

type updateReq struct {
	Title string `json:"title"`
	Done  bool   `json:"done"`
}

func (h *Handler) update(w http.ResponseWriter, r *http.Request) {
	id, bad := parseID(w, r)
	if bad {
		return
	}
	var req updateReq
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil || req.Title == "" {
		httpError(w, http.StatusBadRequest, "invalid json: require non-empty title")
		return
	}
	t, err := h.repo.Update(id, req.Title, req.Done)
	if err != nil {
		httpError(w, http.StatusNotFound, err.Error())
		return
	}
	WriteJSON(w, http.StatusOK, t)
}

func (h *Handler) delete(w http.ResponseWriter, r *http.Request) {
	id, bad := parseID(w, r)
	if bad {
		return
	}
	if err := h.repo.Delete(id); err != nil {
		httpError(w, http.StatusNotFound, err.Error())
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func (h *Handler) Health(w http.ResponseWriter, r *http.Request) {
	WriteJSON(w, http.StatusOK, map[string]string{"status": "ok"})
}

// helpers

func parseID(w http.ResponseWriter, r *http.Request) (int64, bool) {
	raw := chi.URLParam(r, "id")
	id, err := strconv.ParseInt(raw, 10, 64)
	bad := false
	if err != nil || id <= 0 {
		httpError(w, http.StatusBadRequest, "invalid id")
		id = 0
		bad = true
	}
	return id, bad
}

func WriteJSON(w http.ResponseWriter, code int, v any) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(code)
	_ = json.NewEncoder(w).Encode(v)
}

func httpError(w http.ResponseWriter, code int, msg string) {
	WriteJSON(w, code, map[string]string{"error": msg})
}
