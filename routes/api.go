package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/shiv122/go-todo/app/controllers"
	authController "github.com/shiv122/go-todo/app/controllers/auth"
	"github.com/shiv122/go-todo/app/middlewares"
	"github.com/shiv122/go-todo/app/requests"
	authRequest "github.com/shiv122/go-todo/app/requests/auth"
)

func SetupApiRoute(app *fiber.App) {

	route := app.Group("/api")

	route.Post("login",
		new(authRequest.UserLoginRequest).Validate,
		new(authController.LoginController).Login,
	)

	route.Post("sign-up",
		new(authRequest.UserSignUpRequest).Validate,
		new(authController.SignUpController).SignUp,
	)

	userAuthMiddleware := new(middlewares.UserAuthMiddleware).Auth
	usersRoute := route.Group("users", userAuthMiddleware)

	userController := new(controllers.UserController)

	usersRoute.Get("/", userController.GetUsers)
	usersRoute.Get("/profile", userController.GetProfile)

	todoController := new(controllers.TodoController)
	todoRoute := usersRoute.Group("todos")

	todoRoute.Get("/", todoController.GetList)
	todoRoute.Post("create", new(requests.StoreTodoRequest).Validate, todoController.StoreTodo)

}
