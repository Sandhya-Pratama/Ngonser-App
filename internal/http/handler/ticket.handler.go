package handler

import (
	"net/http"
	"time"

	"github.com/Sandhya-Pratama/Ngonser-App/entity"
	"github.com/Sandhya-Pratama/Ngonser-App/internal/http/validator"
	"github.com/Sandhya-Pratama/Ngonser-App/internal/service"
	"github.com/labstack/echo/v4"
)

type TicketHandler struct {
	ticketService service.TicketUseCase
}

func NewTicketHandler(ticketService service.TicketUseCase) *TicketHandler {
	return &TicketHandler{ticketService}
}

func (h *TicketHandler) GetAllTicket(ctx echo.Context) error {
	tickets, err := h.ticketService.GetAllTicket(ctx.Request().Context())
	if err != nil {
		return ctx.JSON(http.StatusUnprocessableEntity, err)
	}
	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"data": tickets,
	})
}

func (h *TicketHandler) CreateTicket(ctx echo.Context) error {
	var input struct {
		Title       string    `json:"title" validate:"required"`
		Description string    `json:"description" validate:"required"`
		Image       string    `json:"image"`
		Location    string    `json:"location"`
		Date        time.Time `json:"date"`
		Status      string    `json:"status"`
		Price       float64   `json:"price"`
		Quota       int       `json:"quota"`
		Terjual     int       `json:"terjual"`
		Category    string    `json:"category"`
	}

	if err := ctx.Bind(&input); err != nil {
		return ctx.JSON(http.StatusBadRequest, validator.ValidatorErrors(err))
	}

	dateStr := input.Date.Format("2006-01-02T15:04:05Z")

	ticket := entity.Ticket{
		Title:       input.Title,
		Description: input.Description,
		Image:       input.Image,
		Location:    input.Location,
		Date:        dateStr,
		Status:      input.Status,
		Price:       int64(input.Price),
		Quota:       int64(input.Quota),
		Terjual:     int64(input.Terjual),
		Category:    input.Category,
		CreatedAt:   time.Now(),
	}

	err := h.ticketService.CreateTicket(ctx.Request().Context(), &ticket)

	if err != nil {
		return ctx.JSON(http.StatusUnprocessableEntity, err)
	}

	// kalau retrun nya kaya gini akan tampil pesan "Ticket created successfully"
	return ctx.JSON(http.StatusCreated, "Ticket created successfully")
}

func (h *TicketHandler) UpdateTicket(ctx echo.Context) error {
	var input struct {
		ID          int       `json:"id" validate:"required"`
		Title       string    `json:"title" validate:"required"`
		Description string    `json:"description" validate:"required"`
		Image       string    `json:"image"`
		Location    string    `json:"location"`
		Date        time.Time `json:"date"`
		Price       float64   `json:"price"`
		Quota       int       `json:"quota"`
		Category    string    `json:"category"`
	}

	if err := ctx.Bind(&input); err != nil {
		return ctx.JSON(http.StatusBadRequest, validator.ValidatorErrors(err))
	}

	dateStr := input.Date.Format("2006-01-02T15:04:05Z")

	ticket := entity.Ticket{
		ID:          int64(input.ID),
		Title:       input.Title,
		Description: input.Description,
		Image:       input.Image,
		Location:    input.Location,
		Date:        dateStr,
		Price:       int64(input.Price),
		Quota:       int64(input.Quota),
		Category:    input.Category,
	}

	err := h.ticketService.UpdateTicket(ctx.Request().Context(), &ticket)

	if err != nil {
		return ctx.JSON(http.StatusUnprocessableEntity, err)
	}
	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"message": "Ticket updated successfully",
		"data": map[string]interface{}{
			"id":          ticket.ID,
			"title":       ticket.Title,
			"description": ticket.Description,
			"image":       ticket.Image,
			"location":    ticket.Location,
			"date":        ticket.Date,
			"price":       ticket.Price,
			"quota":       ticket.Quota,
			"category":    ticket.Category,
			"update":      ticket.UpdatedAt,
		},
	})
}

func (h *TicketHandler) DeleteTicket(c echo.Context) error {
	var input struct {
		ID int64 `param:"id" validate:"required"`
	}

	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, validator.ValidatorErrors(err))
	}

	err := h.ticketService.DeleteTicket(c.Request().Context(), input.ID)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err)
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Ticket deleted successfully",
	})
}
