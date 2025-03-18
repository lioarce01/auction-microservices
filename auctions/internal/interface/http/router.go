package http

import (
	"github.com/gin-gonic/gin"
	"github.com/lioarce01/auction-microservices/internal/application/usecase/auction"
	"github.com/lioarce01/auction-microservices/internal/infrastructure/database"
	"github.com/lioarce01/auction-microservices/internal/infrastructure/db_repositories"
	"github.com/lioarce01/auction-microservices/internal/interface/handler"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	db, err := database.NewDBConnection()
	if err != nil {
		panic("Error connecting to the database: " + err.Error())
	}

	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "ok",
		})
	})

	//create repository
	auctionRepo := db_repositories.NewMongoAuctionRepository(db)

	//create usecase
	listAuctionsUseCase := auction.NewListAuctionsUseCase(auctionRepo)
	createAuctionUseCase := auction.NewCreateAuctionUseCase(auctionRepo)
	getOneAuctionUseCase := auction.NewGetOneAuctionUseCase(auctionRepo)
	updateAuctionUseCase := auction.NewUpdateAuctionUseCase(auctionRepo)
	deleteAuctionUseCase := auction.NewDeleteAuctionUseCase(auctionRepo)

	//create handler
	auctionHandler := handler.NewAuctionHandler(
		listAuctionsUseCase,
		getOneAuctionUseCase,
		createAuctionUseCase,
		updateAuctionUseCase,
		deleteAuctionUseCase,
	)

	//create routes
	auctionRoutes := r.Group("/auctions")
	{
		auctionRoutes.GET("/", auctionHandler.ListAllAuctions)
		auctionRoutes.POST("/", auctionHandler.CreateAuction)
		auctionRoutes.GET("/:id", auctionHandler.GetAuction)
		auctionRoutes.PUT("/:id", auctionHandler.UpdateAuction)
		auctionRoutes.DELETE("/:id", auctionHandler.DeleteAuction)
	}

	return r
}
