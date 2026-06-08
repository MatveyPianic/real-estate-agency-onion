package auth

import (
    "fmt"
    "time"

    "github.com/golang-jwt/jwt/v5"
    appinterfaces "real-estate-agency-onion/internal/application/interfaces"
)

type JWTTokenManager struct {
    secret     []byte
    expiration time.Duration
}

func NewJWTTokenManager(secret string, expiration time.Duration) *JWTTokenManager {
    return &JWTTokenManager{
        secret:     []byte(secret),
        expiration: expiration,
    }
}

func (m *JWTTokenManager) Generate(userID int64, roles []string) (string, error) {
    claims := jwt.MapClaims{
        "user_id": userID,
        "roles":   roles,
        "exp":     time.Now().Add(m.expiration).Unix(),
        "iat":     time.Now().Unix(),
    }
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    return token.SignedString(m.secret)
}

func (m *JWTTokenManager) Validate(tokenString string) (*appinterfaces.TokenClaims, error) {
    token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
        if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
            return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
        }
        return m.secret, nil
    })
    if err != nil {
        return nil, err
    }

    claims, ok := token.Claims.(jwt.MapClaims)
    if !ok || !token.Valid {
        return nil, fmt.Errorf("invalid token")
    }

    userIDFloat, ok := claims["user_id"].(float64)
    if !ok {
        return nil, fmt.Errorf("invalid user_id in token")
    }

    rolesRaw, ok := claims["roles"].([]interface{})
    if !ok {
        return nil, fmt.Errorf("invalid roles in token")
    }

    roles := make([]string, len(rolesRaw))
    for i, r := range rolesRaw {
        roles[i], ok = r.(string)
        if !ok {
            return nil, fmt.Errorf("invalid role in token")
        }
    }

    return &appinterfaces.TokenClaims{
        UserID: int64(userIDFloat),
        Roles:  roles,
    }, nil
}