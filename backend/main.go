package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-contrib/pprof"

	"github.com/gerodp/simpleBlogApp/controller"
	"github.com/gerodp/simpleBlogApp/model"
	"github.com/gerodp/simpleBlogApp/repository"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func buildHandler(promRespTimeDuration *prometheus.HistogramVec, method string, path string, handler func(c *gin.Context)) func(c *gin.Context) {
	return func(c *gin.Context) {
		start := time.Now()
		handler(c)
		duration := time.Since(start).Seconds()
		promRespTimeDuration.With(prometheus.Labels{"method": method, "path": path}).Observe(duration)
	}
}

func createTestUser(userRepo model.UserRepository) error {

	_, err := userRepo.FindByUsername("testint1")

	if err == nil {
		return nil
	} else {
		user := model.User{
			ID:        1,
			Username:  "testint1",
			Password:  "$2a$14$oemuupbL/xOA3d.jS3CBOeLf0SfCt.cWrqWBqr4jC71CxuF.mrrCa",
			Email:     "test@integration.int",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}

		_, err2 := userRepo.Save(&user)
		return err2
	}
}

func main() {
	log.SetFlags(log.LstdFlags | log.Lmicroseconds) // include the timestamp in the output

	gin.DefaultWriter = log.Writer()
	gin.DefaultErrorWriter = log.Writer()
	log.Println("Backend server started!")

	promRespTimeDuration := prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Name:    "http_request_duration_seconds",
		Help:    "A histogram of the HTTP request durations in seconds.",
		Buckets: prometheus.ExponentialBuckets(0.001, 2, 6),
	}, []string{"method", "path"})

	prometheus.MustRegister(promRespTimeDuration)

	addr := fmt.Sprintf("%s:%s", os.Getenv("DATABASE_HOST"), os.Getenv("DATABASE_PORT"))
	log.Printf("DB address =%s\n", addr)

	repo, err := repository.NewRepository(os.Getenv("DATABASE_USER"), os.Getenv("DATABASE_PASSWORD"), addr, os.Getenv("DATABASE_NAME"))

	if err != nil {
		log.Fatalln("Error creating repository", err)
		os.Exit(1)
	}

	errUser := createTestUser(repo.Users)

	if errUser != nil {
		log.Fatalln("Error creating test user", err)
		os.Exit(1)
	}

	userCon := controller.NewUserController(repo.Users)

	postCon := controller.NewPostController(repo.Posts)

	authMdw, amErr := controller.NewAuthMiddleware(repo.Users)

	if amErr != nil {
		log.Fatalln("Error creating Auth Middleware. Exiting...")
		os.Exit(1)
	}

	r := gin.Default()
	pprof.Register(r)

	r.GET("/metrics", func(c *gin.Context) {
		promhttp.Handler().ServeHTTP(c.Writer, c.Request)
	})

	r.GET("/health_check", func(c *gin.Context) {
		c.JSON(http.StatusOK, "OK")
	})

	r.POST("/login", buildHandler(promRespTimeDuration, "POST", "/login", authMdw.LoginHandler))

	auth := r.Group("/auth")

	r.POST("/refresh_token", buildHandler(promRespTimeDuration, "POST", "/refresh_token", authMdw.RefreshHandler))

	auth.Use(authMdw.MiddlewareFunc())
	{
		auth.POST("/user", buildHandler(promRespTimeDuration, "POST", "/auth/user", userCon.CreateUser))

		auth.GET("/user", buildHandler(promRespTimeDuration, "GET", "/auth/user", userCon.Find))

		auth.POST("/post", buildHandler(promRespTimeDuration, "POST", "/auth/post", postCon.CreatePost))

		auth.GET("/post", buildHandler(promRespTimeDuration, "GET", "/auth/post", postCon.Find))

		auth.DELETE("/post/:id", buildHandler(promRespTimeDuration, "DELETE", "auth/post/:id", postCon.DeletePost))
	}

	r.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"message": "Endpoint not found"})
	})

	s := &http.Server{
		Addr:     os.Getenv("BACKEND_ADDRESS"),
		Handler:  r,
		ErrorLog: log.Default(),
	}
	s.ListenAndServe()
}
