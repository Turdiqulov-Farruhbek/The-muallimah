package handlers

import (
	"context"
	"strconv"

	pb "muallimah-gateway/internal/pkg/genproto"

	"github.com/gin-gonic/gin"
)

// CreateTransaction creates a new transaction
// @Summary       Create Transaction
// @Description   Create a new transaction
// @Tags          Transaction
// @Accept        json
// @Produce       json
// @Security      BearerAuth
// @Param         transaction body pb.TransactionCreateReq true "Transaction data"
// @Success       200  {string}  string "Transaction created successfully"
// @Failure       400  {string}  string "Invalid request"
// @Failure       500  {string}  string "Internal server error"
// @Router        /transactions/create [post]
func (h *Handler) CreateTransaction(c *gin.Context) {
	var req pb.TransactionCreateReq
	if err := c.ShouldBindJSON(&req); err != nil {
		h.Logger.ERROR.Println("Failed to bind request:", err)
		c.JSON(400, "Invalid request: "+err.Error())
		return
	}

	_, err := h.Clients.Transaction.CreateTransaction(c, &req)
	if err != nil {
		h.Logger.ERROR.Println("Failed to create transaction:", err)
		c.JSON(500, "Internal server error: "+err.Error())
		return
	}

	c.JSON(200, "Transaction created successfully")
}

// GetTransaction retrieves a transaction by ID
// @Summary       Get Transaction
// @Description   Retrieve a transaction by its ID
// @Tags          Transaction
// @Accept        json
// @Produce       json
// @Security      BearerAuth
// @Param         id path string true "Transaction ID"
// @Success       200  {object}  pb.TransactionGetRes "Transaction details"
// @Failure       400  {string}  string "Invalid request"
// @Failure       404  {string}  string "Transaction not found"
// @Failure       500  {string}  string "Internal server error"
// @Router        /transactions/{id} [get]
func (h *Handler) GetTransaction(c *gin.Context) {
	transactionID := c.Param("id")

	req := &pb.ById{Id: transactionID}

	res, err := h.Clients.Transaction.GetTransaction(c, req)
	if err != nil {
		h.Logger.ERROR.Println("Failed to get transaction:", err)
		c.JSON(500, "Internal server error: "+err.Error())
		return
	}

	c.JSON(200, res)
}

// ListTransactions retrieves a list of transactions
// @Summary       List Transactions
// @Description   Retrieve a list of transactions
// @Tags          Transaction
// @Accept        json
// @Produce       json
// @Security      BearerAuth
// @Param         user_course_id query string false "User Course ID"
// @Param         amount_from query float false "Amount from"
// @Param         amount_to query float false "Amount to"
// @Param         type query string false "Transaction type"
// @Param         user_id query string false "User ID"
// @Param         course_id query string false "Course ID"
// @Param         limit query int false "Limit"
// @Param         offset query int false "Offset"
// @Success       200  {object}  pb.TransactionListRes "List of transactions"
// @Failure       400  {string}  string "Invalid request"
// @Failure       500  {string}  string "Internal server error"
// @Router        /transactions/list [get]
func (h *Handler) ListTransactions(c *gin.Context) {
	var req pb.TransactionListReq

	req.UserCourseId = c.Query("user_course_id")
	req.Type = c.Query("type")
	req.UserId = c.Query("user_id")
	req.CourseId = c.Query("course_id")

	if amountFrom := c.Query("amount_from"); amountFrom != "" {
		if af, err := strconv.ParseFloat(amountFrom, 64); err == nil {
			req.AmountFrom = float32(af)
		}
	}

	if amountTo := c.Query("amount_to"); amountTo != "" {
		if at, err := strconv.ParseFloat(amountTo, 64); err == nil {
			req.AmountTo = float32(at)
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

	res, err := h.Clients.Transaction.ListTransactions(context.Background(), &req)
	if err != nil {
		h.Logger.ERROR.Println("Failed to list transactions:", err)
		c.JSON(500, "Internal server error: "+err.Error())
		return
	}

	c.JSON(200, res)
}

// GetBalance retrieves the balance for a user
// @Summary       Get Balance
// @Description   Retrieve the balance and usage statistics for a user
// @Tags          Transaction
// @Accept        json
// @Produce       json
// @Security      BearerAuth
// @Param         id path string true "User ID"
// @Success       200  {object}  pb.BalanceGetRes "User balance and usage statistics"
// @Failure       400  {string}  string "Invalid request"
// @Failure       404  {string}  string "User not found"
// @Failure       500  {string}  string "Internal server error"
// @Router        /transactions/balance/{id} [get]
func (h *Handler) GetBalance(c *gin.Context) {
	userID := c.Param("id")

	req := &pb.ById{Id: userID}

	res, err := h.Clients.Transaction.GetBalance(c, req)
	if err != nil {
		h.Logger.ERROR.Println("Failed to get balance:", err)
		c.JSON(500, "Internal server error: "+err.Error())
		return
	}

	c.JSON(200, res)
}
