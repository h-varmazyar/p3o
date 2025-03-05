package auth

import(
	"github.com/gin-gonic/gin"
	 "time"
)

type LoginReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResp struct {
	Token        string    `json:"token"`
	ExpireAt     time.Time `json:"expire_at"`
	VerifiedUser bool      `json:"verified_user"`
}

func (c *Controller) Login(ctx *gin.Context) {
	loginReq := new(LoginReq)

	if err := ctx.ShouldBindJSON(&loginReq); err != nil {
		utils.JsonHttpResponse(ctx, nil, ErrLoginFailed.AddOriginalError(err), false)
		return
	}

	user, found, err := c.fetchUser(ctx, loginReq.Username)
	if err != nil {
		utils.JsonHttpResponse(ctx, nil, err, false)
		return
	}

	if found {
		if err = utils.CompareHashPassword(loginReq.Password, user.HashedPassword); err != nil {
			log.WithError(err).Error("failed to generage hashed password")
			utils.JsonHttpResponse(ctx, nil, ErrInvalidUsernamePassword.AddOriginalError(err), false)
			return
		}
	} else {
		user.HashedPassword, err = utils.GenerateHashPassword(loginReq.Password)
		if err != nil {
			utils.JsonHttpResponse(ctx, nil, ErrPasswordHashingFailed.AddOriginalError(err), false)
			return
		}
		err = c.userModel.Create(ctx, user)
		if err != nil {
			utils.JsonHttpResponse(ctx, nil, err, false)
			return
		}
	}

	expirationTime := time.Now().Add(30 * 24 * time.Hour)

	claims := &entities.Claims{
		Role: user.Role.ToString(),
		StandardClaims: jwt.StandardClaims{
			Subject:   loginReq.Username,
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS512, claims)

	tokenString, err := token.SignedString(signKey)
	if err != nil {
		utils.JsonHttpResponse(ctx, nil, ErrLoginFailed.AddOriginalError(err), false)
		return
	}

	resp := &LoginResp{
		Token:        tokenString,
		ExpireAt:     expirationTime,
		VerifiedUser: user.VerifiedAt != nil,
	}

	utils.JsonHttpResponse(ctx, resp, nil, true)
}