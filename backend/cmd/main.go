package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/tsunakit99/sp-web2/backend/infra"
	"github.com/tsunakit99/sp-web2/backend/internal/handler"
	custommw "github.com/tsunakit99/sp-web2/backend/internal/middleware"
	postgres "github.com/tsunakit99/sp-web2/backend/internal/repository"
	"github.com/tsunakit99/sp-web2/backend/internal/usecase"
)

func main() {
	godotenv.Load()
	// Echo server instance
	e := echo.New()

	// CORS（開発中フロントと接続しやすくする）
	e.Use(middleware.CORS())

	// DB接続
	db, err := infra.NewDB()
	if err != nil {
		log.Fatalf("❌ DB接続失敗: %v", err)
	}
	defer db.Close()

	// DI構成（Repository → Usecase → Handler）
	taskRepo := postgres.NewTaskRepository(db)
	taskUC := usecase.NewTaskUsecase(taskRepo)
	taskHandler := handler.NewTaskHandler(taskUC)

	// 認証グループ：JWTミドルウェア
	authGroup := e.Group("")
	authGroup.Use(custommw.SupabaseAuthMiddleware)

	// /tasks エンドポイント登録
	handler.RegisterTaskRoutes(authGroup, taskHandler)

	// ポート設定
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	e.Logger.Infof("✅ Server is listening on port %s", port)
	e.Logger.Fatal(e.Start(":" + port))
}
