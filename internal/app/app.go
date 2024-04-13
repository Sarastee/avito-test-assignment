package app

import (
	"context"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/sarastee/avito-test-assignment/internal/config"
	"github.com/sarastee/platform_common/pkg/closer"
)

// App ..
type App struct {
	serviceProvider *serviceProvider
	httpServer      *http.Server
	configPath      string
}

// NewApp ..
func NewApp(ctx context.Context, configPath string) (*App, error) {
	a := &App{configPath: configPath}

	if err := a.initDeps(ctx); err != nil {
		return nil, err
	}

	return a, nil
}

// Run ..
func (a *App) Run() error {
	defer func() {
		closer.CloseAll()
		closer.Wait()
	}()

	wg := sync.WaitGroup{}
	wg.Add(1)

	go func() {
		defer wg.Done()

		err := a.runHTTPServer()
		if err != nil {
			log.Fatalf("failure while running HTTP server")
		}
	}()

	wg.Wait()

	return nil
}

func (a *App) initDeps(ctx context.Context) error {
	initDepFunctions := []func(context.Context) error{
		a.initConfig,
		a.initServiceProvider,
		a.initHTTPServer,
	}

	for _, f := range initDepFunctions {
		if err := f(ctx); err != nil {
			return err
		}
	}

	return nil
}

func (a *App) initConfig(_ context.Context) error {
	err := config.Load(a.configPath)
	if err != nil {
		return err
	}

	return nil
}

func (a *App) initServiceProvider(_ context.Context) error {
	a.serviceProvider = newServiceProvider()
	return nil
}

func (a *App) initHTTPServer(ctx context.Context) error {
	i := a.serviceProvider.BannerImpl(ctx)
	ai := a.serviceProvider.AuthImpl(ctx)
	mw := a.serviceProvider.Middleware()

	mux := http.NewServeMux()

	mux.Handle("POST /register", http.HandlerFunc(ai.CreateUser))
	mux.Handle("POST /login", http.HandlerFunc(ai.LogIn))

	mux.Handle("GET /user_banner", mw.AuthRequired(i.GetUserBanner))

	mux.Handle("POST /banner", mw.AdminRequired(http.HandlerFunc(i.CreateBanner)))
	// mux.Handle("GET /banner", mw.AdminRequired(http.HandlerFunc(i.GetAdminBanners())))
	mux.Handle("DELETE /banner/{id}", mw.AdminRequired(http.HandlerFunc(i.DeleteBanner)))

	a.httpServer = &http.Server{
		Addr:              a.serviceProvider.HTTPConfig().Address(),
		Handler:           mux,
		ReadHeaderTimeout: 2 * time.Second,
	}

	return nil
}

func (a *App) runHTTPServer() error {
	log.Printf("HTTP started at %s", a.serviceProvider.HTTPConfig().Address())

	err := a.httpServer.ListenAndServe()
	if err != nil {
		return err
	}

	return nil
}
