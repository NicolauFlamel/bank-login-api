package dtos

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type ClientReq struct {
  SessionId string `json:"session_id"`
  Sequence string `json:"sequence"`
}

type CustomClaims struct {
    SessionID string `json:"session_id"`
    jwt.RegisteredClaims
}

type Session struct {
  Id string
  Layout string
  IsValid bool
  CreatedAt time.Time
}

type User struct {
  Id int
  Name string
  Digit1 string
  Digit2 string
  Digit3 string
  Digit4 string
  Salt string
}
