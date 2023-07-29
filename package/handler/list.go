package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-exes/todo-serv"
)

type GetAllListsResponse struct {
	Data []todo.TodoList `json:"data"`
}

func (h *Handler) createList(ctx *gin.Context) {
	userId, err := h.getUserId(ctx)
	if err != nil {
		return
	}

	var input todo.TodoList
	err = ctx.BindJSON(&input)
	if err != nil {
		NewErrorResponse(ctx, http.StatusBadRequest, "Request is incorrect type")
	}

	id, err := h.services.TodoList.CreateList(input, userId)
	if err != nil {
		NewErrorResponse(ctx, http.StatusInternalServerError, err.Error())
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) getAllLists(ctx *gin.Context) {
	userId, err := h.getUserId(ctx)
	if err != nil {
		return
	}

	lists, err := h.services.TodoList.GetAllLists(userId)
	if err != nil {
		NewErrorResponse(ctx, http.StatusInternalServerError, err.Error())
	}

	ctx.JSON(http.StatusOK, GetAllListsResponse{
		Data: lists,
	})
}

func (h *Handler) getListById(ctx *gin.Context) {
	userId, err := h.getUserId(ctx)
	if err != nil {
		return
	}

	id, err := h.getUrlParam(ctx, "id")
	if err != nil {
		return
	}

	list, err := h.services.TodoList.GetListById(userId, id)
	if err != nil {
		NewErrorResponse(ctx, http.StatusInternalServerError, err.Error())
	}

	ctx.JSON(http.StatusOK, list)
}

func (h *Handler) deleteList(ctx *gin.Context) {
	id, err := h.getUrlParam(ctx, "id")
	if err != nil {
		return
	}

	err = h.services.TodoList.DeleteList(id)
	if err != nil {
		NewErrorResponse(ctx, http.StatusInternalServerError, err.Error())
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{
		"Status": http.StatusOK,
	})
}

func (h *Handler) updateList(ctx *gin.Context) {
	id, err := h.getUrlParam(ctx, "id")
	if err != nil {
		return
	}
	var input todo.UpdateListRequest
	err = ctx.BindJSON(&input)
	if err != nil {
		NewErrorResponse(ctx, http.StatusBadRequest, "Request is incorrect type")
	}

	result, err := h.services.TodoList.UpdateList(input, id)
	if err != nil {
		NewErrorResponse(ctx, http.StatusInternalServerError, err.Error())
	}

	ctx.JSON(http.StatusOK, result)
}
