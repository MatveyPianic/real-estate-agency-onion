package interfaces

type TokenManager interface {
    Generate(userID int64, roles []string) (string, error)
    Validate(tokenString string) (*TokenClaims, error)
}

type TokenClaims struct {
    UserID int64
    Roles  []string
}