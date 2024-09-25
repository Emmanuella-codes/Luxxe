package webapi

import (
	"context"
	"log"
	"time"

	auth_api "github.com/Emmanuella-codes/Luxxe/luxxe-auth/api"
	config "github.com/Emmanuella-codes/Luxxe/luxxe-config"
	entities "github.com/Emmanuella-codes/Luxxe/luxxe-entities"
	shared_api "github.com/Emmanuella-codes/Luxxe/luxxe-shared/api"
	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"go.mongodb.org/mongo-driver/bson"
)

func GenerateApp() *fiber.App {
	app := fiber.New(fiber.Config{
		JSONEncoder: json.Marshal,
		JSONDecoder: json.Unmarshal,
		BodyLimit:   500 * 1024 * 1024, //500mb
	})

	app.Use(cors.New())

	// Custom middleware to conditionally skip logging for a specific route
	app.Use(func(ctx *fiber.Ctx) error {
		if ctx.Path() == "/health-check" || ctx.Path() == "/docs" {
			// Skip logging for the health check route
			return ctx.Next()
		}
		// Invoke the logger middleware for all other routes
		return logger.New()(ctx)
	})

	// Middleware to save request details to the database
	app.Use(func(ctx *fiber.Ctx) error {
		if ctx.Path() == "/health-check" || ctx.Path() == "/docs" {
			// Skip logging for the health check route
			return ctx.Next()
		}

		if config.EnvConfig.ENV == config.ServerEnvironmentProduction {
			requestData := entities.AuditLog{
				RequestIP:   ctx.IP(), // Capture the request IP
				QueryParams: ctx.Queries(),
				OriginalURL: ctx.OriginalURL(),
				CreatedAt:   time.Now(),
			}

			go func() {
				_, err := entities.AuditLogModel.Create(
					context.Background(), // use a background context here
					&bson.M{
						"requestIP":   requestData.RequestIP,
						"queryParams": requestData.QueryParams,
						"originalURL": requestData.OriginalURL,
						"createdAt":   requestData.CreatedAt,
					},
				)
				if err != nil {
					log.Println("Error inserting request data:", err)
				}
			}()
		}

		return ctx.Next()
	})

	// create health check route
	app.Get("/health-check", func(ctx *fiber.Ctx) error {
		return ctx.JSON(map[string]string{"check": "LoveWall server is live!. 📦 🧧 💪🏾"})
	})

	app.Get("/docs", func(ctx *fiber.Ctx) error {
		externalURL := config.EnvConfig.API_DOCUMENTATION_URL
		return ctx.Redirect(externalURL, fiber.StatusMovedPermanently)
	})

	authGroup := app.Group("/auth")
	shared_api.BaseRouter(authGroup, auth_api.AuthRoutes)

	// userGroup := app.Group("/user")
	// shared_api.BaseRouter(userGroup, user_api.UserRoutes)

	

	return app
}