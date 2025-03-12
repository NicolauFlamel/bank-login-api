package models

import "github.com/golang-jwt/jwt/v5"

type CustomClaims struct {
    SessionID string `json:"session_id"`
    jwt.RegisteredClaims
}

type  LayoutRequest struct {
  User string `json:"user"`
}

type LayoutResponse struct {
	Layout      Layout `json:"layout"`
	SessionHash string   `json:"session_hash"`
}

type Layout [5][2]int

type User struct {
  Id int
  Name string
  Digit1 string
  Digit2 string
  Digit3 string
  Digit4 string
  Salt string
}
