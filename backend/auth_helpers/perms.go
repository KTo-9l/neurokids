package auth_helpers

type Perms int

// HardCode for permissions
// Данные настройки выносить на уровень приложения
const (
	NoPerms  Perms = 1 << 0 // бан
	AllPerms Perms = 1 << 1 // все разрешения (админ)
	OnePerms Perms = 1 << 2 // какое-то разрешение
)
