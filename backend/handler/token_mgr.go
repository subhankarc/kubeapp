package handler

import "github.com/smjn/ipl18/backend/auth"

var tokenManager = auth.NewTokenManager(auth.SignMethodSHA512)
