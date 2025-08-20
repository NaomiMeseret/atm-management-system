package delivery

import (
	"atm-management-system/atm_api/Delivery/routers"
	"atm-management-system/atm_api/Repositories/database"
	"log"
	"os"

	"github.com/joho/godotenv"
)
func main() {
  // Load environment variables
  if err := godotenv.Load(); err != nil {
    log.Printf("Warning: Error loading .env file: %v", err)
  }

  // Connect to MongoDB
  db, err := database.ConnectToMongoDB()
  if err != nil {
    log.Fatal("Failed to connect to MongoDB:", err)
  }
  defer func() {
    if client := db.Client(); client != nil {
      if err := database.CloseMongoDBConnection(client); err != nil {
        log.Printf("Error closing MongoDB connection: %v", err)
      }
    }
  }()
  // Setup router
  router := routers.SetupRouter()

  // Get port from environment variable or use default
  port := os.Getenv("PORT")
  if port == "" {
    port = "8080"
  }

  // Start server
  log.Printf("Server starting on port %s...", port)
  if err := router.Run(":" + port); err != nil {
    log.Fatal("Failed to start server:", err)
  }
}
