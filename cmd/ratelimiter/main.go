package main

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"

	"github.com/fbonareis/goexpert-ratelimiter/internal/infra/configs"
	"github.com/fbonareis/goexpert-ratelimiter/internal/infra/ratelimit"
	"github.com/fbonareis/goexpert-ratelimiter/internal/infra/ratelimit/limiter"
	"github.com/fbonareis/goexpert-ratelimiter/internal/infra/web"
	"github.com/fbonareis/goexpert-ratelimiter/internal/infra/web/webserver"
	"github.com/fbonareis/goexpert-ratelimiter/internal/infra/web/webserver/middleware"
)

func main() {
	done := make(chan struct{})

	configs, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}

	rdb, err := redisClient(configs.RedisHost, configs.RedisPassword)
	if err != nil {
		panic(err)
	}
	defer rdb.Close()

	webserver := webserver.NewWebServer(":8080")
	webSampleHandler := web.NewWebSampleHandler()

	limiter := limiter.NewRedisLimiter(rdb)
	ratelimit := ratelimit.NewRateLimit(limiter, configs.RateLimitRequestPerSecondsIP, configs.RateLimitRequestPerSecondsToken)

	rl := middleware.NewWebServerRateLimiter(ratelimit)
	webserver.AddMiddleware("rate_limiter", rl.Handle)

	webserver.AddHandler("/sample", webSampleHandler.Handle)
	fmt.Println("Starting web server on port", 8080)

	webserver.Start()
	<-done
}

func redisClient(host, password string) (*redis.Client, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     host,
		Password: password,
		DB:       0,
	})
	if _, err := rdb.Ping(context.Background()).Result(); err != nil {
		return nil, err
	}
	return rdb, nil
}
