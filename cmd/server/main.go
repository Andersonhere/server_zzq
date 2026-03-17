package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"server_zzq/internal/config"
	"server_zzq/internal/handlers"
	"server_zzq/internal/middleware"
	"server_zzq/internal/utils"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func main() {
	// 初始化配置
	if err := config.Init(); err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}
	cfg := config.GetConfig()

	// 初始化日志
	logger, err := initLogger(cfg)
	if err != nil {
		log.Fatalf("Failed to init logger: %v", err)
	}
	defer logger.Sync()

	// 初始化数据库（可选）
	if err := utils.InitDB(&cfg.Database); err != nil {
		logger.Warn("Database connection failed, running without database", zap.Error(err))
	} else {
		// 自动迁移数据库表
		if err := migrateDatabase(); err != nil {
			logger.Warn("Database migration failed", zap.Error(err))
		}
	}

	// 初始化 Gin
	gin.SetMode(cfg.Server.Mode)
	r := gin.New()

	// 使用中间件
	r.Use(gin.Recovery())
	r.Use(middleware.CORS())
	r.Use(middleware.Logger(logger))

	// 注册路由
	setupRoutes(r, cfg)

	// 启动服务器
	addr := fmt.Sprintf(":%d", cfg.Server.Port)
	srv := &http.Server{
		Addr:    addr,
		Handler: r,
	}

	// 优雅关闭
	go func() {
		quit := make(chan os.Signal, 1)
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
		<-quit

		logger.Info("Shutting down server...")
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		if err := srv.Shutdown(ctx); err != nil {
			logger.Fatal("Server forced to shutdown", zap.Error(err))
		}
	}()

	logger.Info("Server starting", zap.String("addr", addr), zap.Bool("https", cfg.Server.HTTPS.Enabled))
	
	// 根据配置选择 HTTP 或 HTTPS 启动
	if cfg.Server.HTTPS.Enabled {
		logger.Info("HTTPS enabled", zap.String("cert", cfg.Server.HTTPS.CertFile), zap.String("key", cfg.Server.HTTPS.KeyFile))
		if err := srv.ListenAndServeTLS(cfg.Server.HTTPS.CertFile, cfg.Server.HTTPS.KeyFile); err != nil && err != http.ErrServerClosed {
			logger.Fatal("Failed to start HTTPS server", zap.Error(err))
		}
	} else {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Fatal("Failed to start server", zap.Error(err))
		}
	}

	logger.Info("Server stopped")
}

// initLogger 初始化日志
func initLogger(cfg *config.Config) (*zap.Logger, error) {
	logConfig := zap.NewProductionConfig()
	
	// 如果配置了日志文件路径且不是 stdout，则同时输出到文件和控制台
	if cfg.Log.Filename != "" && cfg.Log.Filename != "stdout" {
		logConfig.OutputPaths = []string{cfg.Log.Filename, "stdout"}
	} else {
		// 只输出到控制台
		logConfig.OutputPaths = []string{"stdout"}
	}
	
	logConfig.EncoderConfig.TimeKey = "time"
	logConfig.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder

	return logConfig.Build()
}

// migrateDatabase 自动迁移数据库表
func migrateDatabase() error {
	db := utils.GetDB()
	if db == nil {
		return fmt.Errorf("database not initialized")
	}

	// TODO: 添加需要迁移的模型
	// return db.AutoMigrate(&models.Shop{}, &models.Product{}, &models.Category{}, &models.Order{}, &models.OrderItem{}, &models.AfterSale{})
	return nil
}

// setupRoutes 设置路由
func setupRoutes(r *gin.Engine, cfg *config.Config) {
	v1 := r.Group("/api/v1/manager")

	// 认证路由（无需鉴权）
	auth := v1.Group("/auth")
	{
		auth.POST("/login", handlers.Login)
		// auth.POST("/refresh", handlers.RefreshToken) // 需要鉴权
		// auth.POST("/logout", handlers.Logout)        // 需要鉴权
	}

	// 需要鉴权的路由
	authorized := v1.Group("")
	authorized.Use(middleware.Auth())
	{
		// 认证相关
		authorized.POST("/auth/profile", handlers.GetProfile)

		// 商品管理
		products := authorized.Group("/products")
		{
			products.POST("", handlers.ListProducts)
			products.POST("/:id", handlers.GetProduct)
			products.POST("/create", handlers.CreateProduct)
			products.POST("/:id/update", handlers.UpdateProduct)
			products.POST("/:id/delete", handlers.DeleteProduct)
			products.POST("/:id/status", handlers.UpdateProductStatus)
			products.POST("/:id/discount", handlers.SetDiscount)
			products.POST("/:id/presale", handlers.SetPresale)
			products.POST("/discount/batch", handlers.BatchSetDiscount)
			products.POST("/presale/batch", handlers.BatchSetPresale)
		}

		// 分类管理
		categories := authorized.Group("/categories")
		{
			categories.POST("", handlers.ListCategories)
			categories.POST("/:id", handlers.GetCategory)
			categories.POST("/create", handlers.CreateCategory)
			categories.POST("/:id/update", handlers.UpdateCategory)
			categories.POST("/:id/delete", handlers.DeleteCategory)
		}

		// 库存管理
		inventory := authorized.Group("/inventory")
		{
			inventory.POST("", handlers.ListInventory)
			inventory.POST("/:productId", handlers.GetInventory)
			inventory.POST("/:productId/adjust", handlers.AdjustInventory)
			inventory.POST("/batch-adjust", handlers.BatchAdjustInventory)
		}

		// 订单管理
		orders := authorized.Group("/orders")
		{
			orders.POST("", handlers.ListOrders)
			orders.POST("/:id", handlers.GetOrder)
			orders.POST("/:id/ship", handlers.ShipOrder)
		}

		// 售后管理
		afterSales := authorized.Group("/after-sales")
		{
			afterSales.POST("", handlers.ListAfterSales)
			afterSales.POST("/:id/handle", handlers.HandleAfterSale)
		}

		// 数据看板
		dashboard := authorized.Group("/dashboard")
		{
			dashboard.POST("/overview", handlers.DashboardOverview)
			dashboard.POST("/compare", handlers.DashboardCompare)
			dashboard.POST("/trend", handlers.DashboardTrend)
			dashboard.POST("/pending", handlers.DashboardPending)
		}
	}
}
