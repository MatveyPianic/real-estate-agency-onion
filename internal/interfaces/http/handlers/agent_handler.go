package handlers

import (
	"net/http"
	"strconv"

	"real-estate-agency-onion/internal/application/dto"
	"real-estate-agency-onion/internal/application/usecases/agent"
	"real-estate-agency-onion/internal/interfaces/http/requests"
	"real-estate-agency-onion/internal/interfaces/http/responses"

	"github.com/gin-gonic/gin"
)

type AgentHandler struct {
	createUC     *agent.CreateUseCase
	getByIDUC    *agent.GetAgentByIDUseCase
	listUC       *agent.ListUseCase
	updateUC     *agent.UpdateUseCase
	softDeleteUC *agent.SoftDeleteUseCase
	deactivateUC *agent.DeactivateUseCase
}

func NewAgentHandler(
	create *agent.CreateUseCase,
	getByID *agent.GetAgentByIDUseCase,
	list *agent.ListUseCase,
	update *agent.UpdateUseCase,
	softDelete *agent.SoftDeleteUseCase,
	deactivate *agent.DeactivateUseCase,
) *AgentHandler {
	return &AgentHandler{
		createUC:     create,
		getByIDUC:    getByID,
		listUC:       list,
		updateUC:     update,
		softDeleteUC: softDelete,
		deactivateUC: deactivate,
	}
}

func (h *AgentHandler) Create(c *gin.Context) {
	var req requests.CreateAgentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, responses.ErrorResponse{Error: "bad_request", Message: err.Error()})
		return
	}

	out, err := h.createUC.Execute(c.Request.Context(), dto.CreateAgentInput{
		FirstName:  req.FirstName,
		LastName:   req.LastName,
		MiddleName: req.MiddleName,
		Phone:      req.Phone,
		Telegram:   req.Telegram,
		Whatsapp:   req.Whatsapp,
	})
	if err != nil {
		status, resp := responses.MapDomainError(err)
		c.JSON(status, resp)
		return
	}

	c.JSON(http.StatusCreated, responses.AgentResponse{
		ID:        out.ID,
		UserID:    out.UserID,
		FullName:  out.FullName,
		Phone:     out.Phone,
		IsActive:  out.IsActive,
		CreatedAt: out.CreatedAt,
	})
}

func (h *AgentHandler) GetByID(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	out, err := h.getByIDUC.Execute(c.Request.Context(), dto.GetAgentByIDInput{ID: id})
	if err != nil {
		status, resp := responses.MapDomainError(err)
		c.JSON(status, resp)
		return
	}
	c.JSON(http.StatusOK, responses.AgentResponse{
		ID:        out.ID,
		UserID:    out.UserID,
		FullName:  out.FullName,
		Phone:     out.Phone,
		IsActive:  out.IsActive,
		CreatedAt: out.CreatedAt,
	})
}

func (h *AgentHandler) List(c *gin.Context) {
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "20"))
	offset, _ := strconv.Atoi(c.DefaultQuery("offset", "0"))

	var isActive *bool
	if v := c.Query("is_active"); v != "" {
		val := v == "true"
		isActive = &val
	}

	out, err := h.listUC.Execute(c.Request.Context(), dto.ListAgentsInput{
		Limit:    limit,
		Offset:   offset,
		IsActive: isActive,
	})
	if err != nil {
		status, resp := responses.MapDomainError(err)
		c.JSON(status, resp)
		return
	}

	items := make([]responses.AgentResponse, len(out.Items))
	for i, item := range out.Items {
		items[i] = responses.AgentResponse{
			ID:        item.ID,
			UserID:    item.UserID,
			FullName:  item.FullName,
			Phone:     item.Phone,
			IsActive:  item.IsActive,
			CreatedAt: item.CreatedAt,
		}
	}
	c.JSON(http.StatusOK, responses.AgentListResponse{Data: items, Total: out.Total})
}

func (h *AgentHandler) Update(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)

	var req struct {
		FirstName  *string `json:"first_name"`
		LastName   *string `json:"last_name"`
		MiddleName *string `json:"middle_name"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, responses.ErrorResponse{Error: "bad_request", Message: err.Error()})
		return
	}

	err := h.updateUC.Execute(c.Request.Context(), dto.UpdateAgentInput{
		ID:         id,
		FirstName:  req.FirstName,
		LastName:   req.LastName,
		MiddleName: req.MiddleName,
	})
	if err != nil {
		status, resp := responses.MapDomainError(err)
		c.JSON(status, resp)
		return
	}
	c.Status(http.StatusNoContent)
}

func (h *AgentHandler) SoftDelete(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	err := h.softDeleteUC.Execute(c.Request.Context(), dto.SoftDeleteAgentInput{ID: id})
	if err != nil {
		status, resp := responses.MapDomainError(err)
		c.JSON(status, resp)
		return
	}
	c.Status(http.StatusNoContent)
}

func (h *AgentHandler) Deactivate(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	err := h.deactivateUC.Execute(c.Request.Context(), dto.DeactivateAgentInput{ID: id})
	if err != nil {
		status, resp := responses.MapDomainError(err)
		c.JSON(status, resp)
		return
	}
	c.Status(http.StatusNoContent)
}
