package router

import (
	"github.com/Sandhya-Pratama/Ngonser-App/internal/http/handler"
	"github.com/labstack/echo/v4"
)

const (
	Admin = "Admin"
	User  = "User"
)

var (
	allRoles  = []string{Admin, User}
	onlyAdmin = []string{Admin}
	onlyUser  = []string{User}
)

type Route struct {
	Method  string
	Path    string
	Handler echo.HandlerFunc
	Roles   []string
}

func PublicRoutes(authHandler *handler.AuthHandler, ticketHandler *handler.TicketHandler) []*Route {
	return []*Route{
		{
			Method:  echo.POST,
			Path:    "/login",
			Handler: authHandler.Login,
		},
		{
			Method:  echo.POST,
			Path:    "/register",
			Handler: authHandler.Registration,
		},
		{
			Method:  echo.GET,
			Path:    "/tickets",
			Handler: ticketHandler.GetAllTicket,
		},
		{
			Method:  echo.POST,
			Path:    "/tickets",
			Handler: ticketHandler.CreateTicket,
		},
		{
			Method:  echo.PUT,
			Path:    "/tickets:id",
			Handler: ticketHandler.UpdateTicket,
		},
	}

}

func PrivateRoutes(userHandler *handler.UserHandler) []*Route {
	return []*Route{
		{
			Method:  echo.POST,
			Path:    "/users",
			Handler: userHandler.CreateUser,
			Roles:   allRoles,
		},

		{
			Method:  echo.GET,
			Path:    "/users",
			Handler: userHandler.GetAllUser,
			Roles:   onlyAdmin,
		},

		{
			Method:  echo.PUT,
			Path:    "/users/:id",
			Handler: userHandler.UpdateUser,
			Roles:   allRoles,
		},

		{
			Method:  echo.GET,
			Path:    "/users/:id",
			Handler: userHandler.GetUserByID,
			Roles:   allRoles,
		},

		{
			Method:  echo.DELETE,
			Path:    "/users/:id",
			Handler: userHandler.DeleteUser,
			Roles:   onlyAdmin,
		},

		{
			Method:  echo.PUT,
			Path:    "/users/self",
			Handler: userHandler.UpdateUserSelf,
			Roles:   onlyUser,
		},
	}
}
