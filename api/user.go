package api

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mniudanri/go-auth-paseto/api/payload"
	"github.com/mniudanri/go-auth-paseto/api/response"
	db "github.com/mniudanri/go-auth-paseto/db/sqlc"
	"github.com/mniudanri/go-auth-paseto/util"
)

func NewUserResponse(user db.User) response.UserResponse {
	return response.UserResponse{
		Username:          user.Username,
		FullName:          user.FullName,
		Email:             user.Email,
		PasswordChangedAt: user.PasswordChangedAt,
		CreatedAt:         user.CreatedAt,
	}
}

func (server *Server) CreateUser(ctx *gin.Context) {
	var req payload.CreateUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, response.GenerateErrResponse(err))
		return
	}

	hashedPassword, err := util.GenerateHashPassword(req.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.GenerateErrResponse(err))
		return
	}

	arg := db.CreateUserTxParams{
		CreateUserParams: db.CreateUserParams{
			Username:       req.Username,
			HashedPassword: hashedPassword,
			FullName:       req.FullName,
			Email:          req.Email,
		},
	}

	user, err := server.Store.CreateUserTx(ctx, arg)
	if err != nil {
		if db.ErrorCode(err) == db.UniqueViolation {
			ctx.JSON(http.StatusForbidden, response.GenerateErrResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, response.GenerateErrResponse(err))
		return
	}

	rsp := NewUserResponse(user.User)
	ctx.JSON(http.StatusOK, rsp)
}

func (server *Server) LoginUser(ctx *gin.Context) {
	var req payload.LoginUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, response.GenerateErrResponse(err))
		return
	}

	user, err := server.Store.GetUser(ctx, req.Username)
	if err != nil {
		if errors.Is(err, db.ErrRecordNotFound) {
			ctx.JSON(http.StatusNotFound, response.GenerateErrResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, response.GenerateErrResponse(err))
		return
	}

	err = util.CheckPassword(req.Password, user.HashedPassword)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, response.GenerateErrResponse(err))
		return
	}

	accessToken, accessPayload, err := server.TokenMaker.CreateToken(
		user.Username,
		server.Config.AccessTokenDuration,
	)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.GenerateErrResponse(err))
		return
	}

	refreshToken, refreshPayload, err := server.TokenMaker.CreateToken(
		user.Username,
		server.Config.RefreshTokenDuration,
	)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.GenerateErrResponse(err))
		return
	}

	session, err := server.Store.CreateSession(ctx, db.CreateSessionParams{
		ID:           refreshPayload.ID,
		Username:     user.Username,
		RefreshToken: refreshToken,
		UserAgent:    ctx.Request.UserAgent(),
		ClientIp:     ctx.ClientIP(),
		IsBlocked:    false,
		ExpiresAt:    refreshPayload.ExpiredAt,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.GenerateErrResponse(err))
		return
	}

	rsp := response.LoginUserResponse{
		SessionID:             session.ID,
		AccessToken:           accessToken,
		AccessTokenExpiresAt:  accessPayload.ExpiredAt,
		RefreshToken:          refreshToken,
		RefreshTokenExpiresAt: refreshPayload.ExpiredAt,
		User:                  NewUserResponse(user),
	}
	ctx.JSON(http.StatusOK, rsp)
}
