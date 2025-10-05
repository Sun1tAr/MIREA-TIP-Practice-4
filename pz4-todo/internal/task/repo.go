package task

import (
	"encoding/json"
	"errors"
	"log"
	"os"
	"sync"
	"time"
)

var ErrNotFound = errors.New("task not found")

type Repo struct {
	mu    sync.RWMutex
	seq   int64
	items map[int64]*Task
	path  string
}

func NewRepo() *Repo {
	return &Repo{items: make(map[int64]*Task)}
}

func NewFileRepo(path string) *Repo {
	repo := &Repo{items: make(map[int64]*Task), path: path}
	err := repo.Load()
	if err != nil {
		log.Fatal(err)
		return NewRepo()
	}
	log.Printf("InFileRepo was loaded")
	return repo
}

func (r *Repo) List() []*Task {
	r.mu.RLock()
	defer r.mu.RUnlock()
	out := make([]*Task, 0, len(r.items))
	for _, t := range r.items {
		out = append(out, t)
	}
	return out
}

// DoneList Список задач с конкретным флагом выполнения
func (r *Repo) DoneList(isDone bool) []*Task {
	r.mu.RLock()
	defer r.mu.RUnlock()
	out := make([]*Task, 0, len(r.items))
	for _, t := range r.items {
		if t.Done == isDone {
			out = append(out, t)
		}
	}
	return out
}

func (r *Repo) PaginatedList(page int64, limit int64) []*Task {
	r.mu.RLock()
	defer r.mu.RUnlock()

	// Если некорректные параметры - возвращаем пустой список
	if page <= 0 || limit <= 0 {
		return []*Task{}
	}

	// Преобразуем map в slice для пагинации
	allTasks := make([]*Task, 0, len(r.items))
	for _, t := range r.items {
		allTasks = append(allTasks, t)
	}

	// Если лимит больше общего количества - возвращаем все
	if int(limit) >= len(allTasks) {
		return allTasks
	}

	// Пагинация
	start := (page - 1) * limit
	if start >= int64(len(allTasks)) {
		return []*Task{} // Страница за пределами
	}

	end := start + limit
	if end > int64(len(allTasks)) {
		end = int64(len(allTasks))
	}

	return allTasks[start:end]
}

func (r *Repo) Get(id int64) (*Task, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	t, ok := r.items[id]
	if !ok {
		return nil, ErrNotFound
	}
	return t, nil
}

func (r *Repo) Create(title string) *Task {
	r.mu.Lock()
	r.seq++
	now := time.Now()
	t := &Task{ID: r.seq, Title: title, CreatedAt: now, UpdatedAt: now, Done: false}
	r.items[t.ID] = t
	r.mu.Unlock()

	err := r.Save()
	if err != nil {
		log.Printf("Save failed when create: %v", err)
	}
	return t
}

func (r *Repo) Update(id int64, title string, done bool) (*Task, error) {
	r.mu.Lock()
	t, ok := r.items[id]
	if !ok {
		r.mu.Unlock()
		return nil, ErrNotFound
	}
	t.Title = title
	t.Done = done
	t.UpdatedAt = time.Now()

	r.mu.Unlock()
	err := r.Save()
	if err != nil {
		log.Fatal("Save failed when update: ", err)
	}
	return t, nil
}

func (r *Repo) Delete(id int64) error {
	r.mu.Lock()
	if _, ok := r.items[id]; !ok {
		r.mu.Unlock()
		return ErrNotFound
	}
	delete(r.items, id)

	r.mu.Unlock()
	err := r.Save()
	if err != nil {
		log.Fatal("Save failed when delete: ", err)
	}
	return nil
}

// Save сохраняет задачи в файл
func (r *Repo) Save() error {
	r.mu.RLock()
	defer r.mu.RUnlock()

	if r.path == "" {
		return nil // Нет пути - ничего не делаем
	}

	file, err := os.Create(r.path)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	return encoder.Encode(r.items)
}

// Load загружает задачи из файла
func (r *Repo) Load() error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if r.path == "" {
		return nil // Нет пути - ничего не делаем
	}

	file, err := os.Open(r.path)
	if err != nil {
		if os.IsNotExist(err) {
			return nil // Файла нет - это нормально
		}
		return err
	}
	defer file.Close()

	var items map[int64]*Task
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&items); err != nil {
		return err
	}

	r.items = items
	// Восстанавливаем seq как максимальный ID
	r.seq = 0
	for id := range r.items {
		if id > r.seq {
			r.seq = id
		}
	}

	return nil
}
