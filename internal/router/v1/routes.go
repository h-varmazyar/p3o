package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/h-varmazyar/p3o/internal/controllers/auth"
	"github.com/h-varmazyar/p3o/internal/controllers/link"
	"go.uber.org/fx"
)

type Router struct {
	ginRouter      *gin.RouterGroup
	authController *auth.Controller
	linkController *link.Controller
	//userController *user.Controller
}

type Params struct {
	fx.In

	AuthController *auth.Controller
	LinkController *link.Controller
}

type Result struct {
	fx.Out

	Router *Router
}

func New(params Params) Result {
	router := &Router{
		authController: params.AuthController,
		linkController: params.LinkController,
	}

	return Result{Router: router}
}

func (r *Router) RegisterRoutes(ginRouter *gin.RouterGroup) {
	v1Router := ginRouter.Group("/v1")

	{
		authRouter := v1Router.Group("/auth")
		authRouter.POST("/login", r.authController.Login)
		authRouter.PUT("/verify", r.authController.Verify)
		authRouter.DELETE("/logout", r.authController.Logout)
	}
	{
		linkRouter := v1Router.Group("/links")
		linkRouter.POST("", r.linkController.Create)
		linkRouter.GET("/count", r.linkController.Counts)
		linkRouter.GET("/visits", r.linkController.Visits)
	}
	{
		//userRouter := v1Router.Group("/users")
		//userRouter.GET("/count", r.userController.Counts)
	}
}

curl \
-X POST https://generativelanguage.googleapis.com/v1beta/models/gemini-1.5-flash:generateContent?key=AIzaSyBhjxODqEPQS5ZbYjchps_z7PJH4WmJQd4 \
-H 'Content-Type: application/json' \
-d @<(echo '{
"contents": [
{
"role": "user",
"parts": [
{
"text": "as a stock specialist, please specify the polarity of below twits:\n1324- #های_وب با سلام خدمت دوستان صبور\nشاخص کل داره پائین میاد ،به نظرتون داره اصلاح میکنه ؟\nممنون میشم اگه راهنمائی بفرمائید\n28363- #غگل پیش خور به چی میگی به سهم غذایی ۲۷۰ تومنی بعد مجمع اصلا همچین سهمی تو غذایی ها هست\n338931-#وپاسار بزودی رشد خوبی خواهد کرد قیمت واقعی سهم این نیست. حقوقی هم خوب داره می خره. تا 500 خواهد رفت.\nthe response must be follow below pattern without any extra information:\noutput patern:\n1324- positive\n28363- negative\n338931- natural"
}
]
}
],
"generationConfig": {
"temperature": 0,
"topK": 64,
"topP": 0.95,
"maxOutputTokens": 8192,
"responseMimeType": "text/plain"
},
"safetySettings": [
{
"category": "HARM_CATEGORY_HARASSMENT",
"threshold": "BLOCK_MEDIUM_AND_ABOVE"
},
{
"category": "HARM_CATEGORY_HATE_SPEECH",
"threshold": "BLOCK_MEDIUM_AND_ABOVE"
},
{
"category": "HARM_CATEGORY_SEXUALLY_EXPLICIT",
"threshold": "BLOCK_MEDIUM_AND_ABOVE"
},
{
"category": "HARM_CATEGORY_DANGEROUS_CONTENT",
"threshold": "BLOCK_MEDIUM_AND_ABOVE"
}
]
}')
