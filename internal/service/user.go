package service

// UserService отвечает за работу с пользователями
type UserService struct{}

// IsAdmin проверяет, является ли пользователь админом
// Публичный метод (заглавная) — может использоваться в handler
func (s *UserService) IsAdmin(userID int64) bool {
	// Здесь могла быть проверка в БД, кэше, конфиге
	// Для примера — хардкод
	return userID == 123456789
}

// GreetUser возвращает персональное приветствие
func (s *UserService) GreetUser(userID int64, name string) string {
	if s.IsAdmin(userID) {
		return "👑 Привет, админ " + name + "!"
	}
	return "👋 Добро пожаловать, " + name + "!"
}