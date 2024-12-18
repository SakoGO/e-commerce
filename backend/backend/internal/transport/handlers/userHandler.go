package handlers

type UserService interface {
	//	UserFindByID(userID int) (*model.User, error) //
	//	UserDelete(userID int) error                  //
}

/* //// GOPOTA
func (h *Handler) UserFindByUsername(w http.ResponseWriter, r *http.Request) {
	// Извлекаем username из URL-параметра
	username := chi.URLParam(r, "username")
	if username == "" {
		http.Error(w, "Username is required", http.StatusBadRequest)
		return
	}

	// Ищем пользователя по имени
	user, err := h.UserService.UserFindByUsername(username)
	if err != nil {
		// Если ошибка поиска (например, пользователь не найден)
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	// Кодируем пользователя в JSON и отправляем в ответ
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(user)
	if err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}
*/
