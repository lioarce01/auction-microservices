package handler

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/lioarce01/auction-microservices/internal/application/usecase/auction"
	"github.com/lioarce01/auction-microservices/internal/domain/entities"
)

type AuctionHandler struct {
	CreateAuctionUseCase *auction.CreateAuctionUseCase
	ListAuctionsUseCase  *auction.ListAuctionsUseCase
	GetOneAuctionUseCase *auction.GetOneAuctionUseCase
	UpdateAuctionUseCase *auction.UpdateAuctionUseCase
	DeleteAuctionUseCase *auction.DeleteAuctionUseCase
}

func NewAuctionHandler(
	listAuctionsUseCase *auction.ListAuctionsUseCase,
	getOneAuctionUseCase *auction.GetOneAuctionUseCase,
	createAuctionUseCase *auction.CreateAuctionUseCase,
	updateAuctionUseCase *auction.UpdateAuctionUseCase,
	deleteAuctionUseCase *auction.DeleteAuctionUseCase,
) *AuctionHandler {
	return &AuctionHandler{
		ListAuctionsUseCase:  listAuctionsUseCase,
		GetOneAuctionUseCase: getOneAuctionUseCase,
		CreateAuctionUseCase: createAuctionUseCase,
		UpdateAuctionUseCase: updateAuctionUseCase,
		DeleteAuctionUseCase: deleteAuctionUseCase,
	}
}

func (h *AuctionHandler) ListAllAuctions(c *gin.Context) {
	auctions, err := h.ListAuctionsUseCase.Execute()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, auctions)
}

func (h *AuctionHandler) GetAuction(c *gin.Context) {
	id := c.Param("id")

	auction, err := h.GetOneAuctionUseCase.Execute(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, auction)
}

func (h *AuctionHandler) CreateAuction(c *gin.Context) {
	// 🔹 Obtener el token del encabezado Authorization
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		fmt.Println("❌ [Handler] Falta el token de autorización")
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization token missing"})
		return
	}

	// Remover el prefijo "Bearer " del token
	token := strings.TrimPrefix(authHeader, "Bearer ")
	if token == "" {
		fmt.Println("❌ [Handler] Token en formato inválido")
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token format"})
		return
	}

	fmt.Println("✅ [Handler] Token obtenido correctamente")

	bodyBytes, _ := io.ReadAll(c.Request.Body)
	fmt.Println("📦 [Handler] Raw request body:", string(bodyBytes))

	// 🔹 Restaurar el cuerpo para que pueda ser leído nuevamente por ShouldBindJSON
	c.Request.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))

	// 🔹 Parsear el JSON
	var req entities.AuctionRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		fmt.Println("❌ [Handler] Error al parsear JSON:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Mostrar el JSON parseado para depuración
	fmt.Println("✅ [Handler] JSON parseado correctamente:", req)

	// 🔹 Convertir la fecha de tipo string a time.Time
	endDate, err := time.Parse(time.RFC3339, req.EndDate)
	if err != nil {
		// Si la fecha tiene un formato incorrecto, retornar un error
		fmt.Println("❌ [Handler] Error al parsear fecha:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid date format"})
		return
	}

	// 🔹 Crear la estructura de la subasta
	auction := entities.Auction{
		Title:        req.Title,
		Description:  req.Description,
		CurrentPrice: req.CurrentPrice,
		EndDate:      endDate,
		Status:       req.Status,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	// Mostrar los datos de la subasta para depuración
	fmt.Println("📌 [Handler] Ejecutando el caso de uso con datos:", auction)

	// 🔹 Ejecutar el caso de uso de la creación de subasta
	createdAuction, err := h.CreateAuctionUseCase.Execute(token, auction)
	if err != nil {
		// Si hay un error en el caso de uso, devolver un error
		fmt.Println("❌ [Handler] Error en UseCase:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// 🔹 Responder con la subasta creada
	fmt.Println("✅ [Handler] Subasta creada con éxito:", createdAuction)
	c.JSON(http.StatusCreated, createdAuction)
}

func (h *AuctionHandler) UpdateAuction(c *gin.Context) {
	var auction entities.Auction
	id := c.GetString("id")

	if err := c.ShouldBindJSON(&auction); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	auction, err := h.UpdateAuctionUseCase.Execute(id, auction)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, auction)
}

func (h *AuctionHandler) DeleteAuction(c *gin.Context) {
	id := c.Param("id")

	creatorID := c.GetString("userID")

	err := h.DeleteAuctionUseCase.Execute(creatorID, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Auction deleted successfully"})
}
