package handlers

import (
	"context"
	"strconv"

	pb "muallimah-gateway/internal/pkg/genproto"

	"github.com/gin-gonic/gin"
)

// CreateEvaluation creates a new evaluation
// @Summary 		Create Evaluation
// @Description 	Create a new evaluation
// @Tags 			Evaluation
// @Accept  		json
// @Produce  		json
// @Security  		BearerAuth
// @Param     		evaluation body    pb.EvaluationCreate  true   "Evaluation data"
// @Success 		200  {string}  string "Evaluation created successfully"
// @Failure 		400  {string}  string "Invalid request"
// @Failure 		500  {string}  string "Internal server error"
// @Router 			/evaluations/create [post]
func (h *Handler) CreateEvaluation(c *gin.Context) {
	var req pb.EvaluationCreate
	if err := c.ShouldBindJSON(&req); err != nil {
		h.Logger.ERROR.Println("Failed to bind request:", err)
		c.JSON(400, "Invalid request"+err.Error())
		return
	}

	_, err := h.Clients.Evaluation.CreateEvaluation(c, &req)
	if err != nil {
		h.Logger.ERROR.Println("Failed to create evaluation:", err)
		c.JSON(500, "Internal server error"+err.Error())
		return
	}

	c.JSON(200, "Evaluation created successfully")
}

// GetEvaluation retrieves an evaluation by ID
// @Summary 		Get Evaluation
// @Description 	Retrieve an evaluation by its ID
// @Tags 			Evaluation
// @Accept  		json
// @Produce  		json
// @Security  		BearerAuth
// @Param     		id path    string       true   "Evaluation ID"
// @Success 		200  {object}  pb.EvaluationGet "Evaluation details"
// @Failure 		400  {string}  string "Invalid request"
// @Failure 		404  {string}  string "Evaluation not found"
// @Failure 		500  {string}  string "Internal server error"
// @Router 			/evaluations/{id} [get]
func (h *Handler) GetEvaluation(c *gin.Context) {
	evaluationID := c.Param("id")

	req := &pb.ById{Id: evaluationID}

	res, err := h.Clients.Evaluation.GetEvaluation(c, req)
	if err != nil {
		h.Logger.ERROR.Println("Failed to get evaluation:", err)
		c.JSON(500, "Internal server error"+err.Error())
		return
	}

	c.JSON(200, res)
}

// DeleteEvaluation deletes an evaluation by ID
// @Summary 		Delete Evaluation
// @Description 	Delete an evaluation by its ID
// @Tags 			Evaluation
// @Accept  		json
// @Produce  		json
// @Security  		BearerAuth
// @Param     		id path    string       true   "Evaluation ID"
// @Success 		200  {string}  string "Evaluation deleted successfully"
// @Failure 		400  {string}  string "Invalid request"
// @Failure 		404  {string}  string "Evaluation not found"
// @Failure 		500  {string}  string "Internal server error"
// @Router 			/evaluations/{id} [delete]
func (h *Handler) DeleteEvaluation(c *gin.Context) {
	evaluationID := c.Param("id")

	req := &pb.ById{Id: evaluationID}

	_, err := h.Clients.Evaluation.DeleteEvaluation(c, req)
	if err != nil {
		h.Logger.ERROR.Println("Failed to delete evaluation:", err)
		c.JSON(500, "Internal server error"+err.Error())
		return
	}

	c.JSON(200, "Evaluation deleted successfully")
}

// ListEvaluations retrieves a list of evaluations
// @Summary 		List Evaluations
// @Description 	Retrieve a list of evaluations
// @Tags 			Evaluation
// @Accept  		json
// @Produce  		json
// @Security  		BearerAuth
// @Param 			course_id  query    string  false  "Course ID"
// @Param 			user_id    query    string  false  "User ID"
// @Param 			score_from query    int     false  "Score From"
// @Param 			score_to   query    int     false  "Score To"
// @Param 			limit      query    int     false  "Limit"
// @Param 			offset     query    int     false  "Offset"
// @Success 		200  {object}  pb.EvaluationList "List of evaluations"
// @Failure 		400  {string}  string "Invalid request"
// @Failure 		500  {string}  string "Internal server error"
// @Router 			/evaluations/list [get]
func (h *Handler) ListEvaluations(c *gin.Context) {
	var req pb.EvaluationFilter
	req.Filter = &pb.Pagination{}

	req.CourseId = c.Query("course_id")
	req.UserId = c.Query("user_id")

	if scoreFrom := c.Query("score_from"); scoreFrom != "" {
		if sf, err := strconv.Atoi(scoreFrom); err == nil {
			req.ScoreFrom = int32(sf)
		}
	}

	if scoreTo := c.Query("score_to"); scoreTo != "" {
		if st, err := strconv.Atoi(scoreTo); err == nil {
			req.ScoreTo = int32(st)
		}
	}

	if limit := c.Query("limit"); limit != "" {
		if l, err := strconv.Atoi(limit); err == nil {
			req.Filter.Limit = int64(l)
		}
	}

	if offset := c.Query("offset"); offset != "" {
		if o, err := strconv.Atoi(offset); err == nil {
			req.Filter.Offset = int64(o)
		}
	}

	res, err := h.Clients.Evaluation.ListEvaluations(context.Background(), &req)
	if err != nil {
		h.Logger.ERROR.Println("Failed to list evaluations:", err)
		c.JSON(500, "Internal server error"+err.Error())
		return
	}

	c.JSON(200, res)
}
