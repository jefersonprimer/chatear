package api

import (
	"context"
	"errors"

	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
	"github.com/jefersonprimer/chatear/backend/config"
	"github.com/jefersonprimer/chatear/backend/graph"
	"github.com/jefersonprimer/chatear/backend/infrastructure"
	"github.com/jefersonprimer/chatear/backend/application/usecases"
	userApp "github.com/jefersonprimer/chatear/backend/internal/user/application"
	userInfra "github.com/jefersonprimer/chatear/backend/internal/user/infrastructure"
	userSvc "github.com/jefersonprimer/chatear/backend/internal/user/services"
	userPres "github.com/jefersonprimer/chatear/backend/internal/user/presentation"
	"github.com/jefersonprimer/chatear/backend/presentation/http"
	"github.com/jefersonprimer/chatear/backend/presentation/middleware"
	"github.com/jefersonprimer/chatear/backend/shared/auth"
)

func SetupServer(cfg *config.Config) (*gin.Engine, error) {

	infra, err := infrastructure.NewInfrastructure(cfg.SupabaseConnectionString, cfg.RedisURL, cfg.NatsURL)
	if err != nil {
		return nil, err
	}

	// Initialize repositories
	userRepo := userInfra.NewPostgresUserRepository(infra.DB)
	blacklistRepo := userInfra.NewRedisBlacklistRepository(infra.Redis)
	refreshTokenRepo := userInfra.NewPostgresRefreshTokenRepository(infra.DB)
	emailLimiter := userInfra.NewRedisEmailLimiter(infra.Redis, cfg)
	userDeletionRepo := userInfra.NewPostgresUserDeletionRepository(infra.DB)
	

	// Initialize event bus (NATS for example)
	eventBus := userInfra.NewNATSEventBus(infra.NatsConn)

	// Initialize shared services
	tokenService := auth.NewTokenService(refreshTokenRepo, cfg.JwtSecret, cfg)
	oneTimeTokenService := userInfra.NewRedisOneTimeTokenService(infra.Redis, cfg)
	

	
		// Initialize user application services
		registerUserUseCase := userApp.NewRegisterUser(userRepo, eventBus, oneTimeTokenService, emailLimiter)
		loginUseCase := userApp.NewLogin(userRepo, tokenService, refreshTokenRepo)
		verifyEmailUseCase := userApp.NewVerifyEmail(userRepo, oneTimeTokenService)
		logoutUser := userApp.NewLogoutUser(refreshTokenRepo, blacklistRepo, tokenService)
		passwordRecovery := userApp.NewPasswordRecovery(userRepo, oneTimeTokenService, eventBus, emailLimiter, cfg.AppURL)
		deleteUser := userApp.NewDeleteUser(userRepo, oneTimeTokenService, eventBus, userDeletionRepo, cfg.AppURL)
		recoverAccount := userApp.NewRecoverAccount(userRepo, nil, oneTimeTokenService)
		refreshToken := userApp.NewRefreshToken(refreshTokenRepo, tokenService, userRepo)
		getUsersUseCase := usecases.NewUserUseCases(userRepo)
		verifyTokenAndResetPasswordUseCase := userApp.NewVerifyTokenAndResetPassword(userRepo, oneTimeTokenService)
		cloudinaryService, err := userSvc.NewCloudinaryService(cfg.CloudinaryURL)
		if err != nil {
			return nil, err
		}
		avatarUsecases := usecases.NewAvatarUsecases(userRepo, cloudinaryService)
	
			
		// Initialize HTTP handlers
		r := gin.Default()
		r.Use(cors.New(cors.Config{
			AllowOrigins:     []string{"http://localhost:3000"},
			AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
			AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
			ExposeHeaders:    []string{"Content-Length"},
			AllowCredentials: true,
			AllowOriginFunc: func(origin string) bool {
				return origin == "http://localhost:3000"
			},
		}))
	err = r.SetTrustedProxies([]string{"127.0.0.1", "::1"})
	if err != nil {
		return nil, err
	}

		userPres.NewUserHandlers(
			r.Group("/api/v1"),
			registerUserUseCase,
			loginUseCase,
			verifyEmailUseCase,
			logoutUser,
			passwordRecovery,
			verifyTokenAndResetPasswordUseCase,
			recoverAccount,
			deleteUser,
			refreshToken,
			oneTimeTokenService,
			tokenService,
			blacklistRepo,
			cfg.FrontendURL,
		)
	
		// Health check routes
		publicRoutes := r.Group("/api/v1")
		healthHandler := http.NewHealthHandler(infra, cfg)
		publicRoutes.GET("/healthz", healthHandler.Healthz)
		publicRoutes.GET("/readyz", healthHandler.Readyz)
	
			// GraphQL setup
			c := graph.Config{
				Resolvers: &graph.Resolver{
					RegisterUserUseCase: registerUserUseCase,
					LoginUseCase:        loginUseCase,
					VerifyEmailUseCase:  verifyEmailUseCase,
					LogoutUser:          logoutUser,
					RecoverPassword:     passwordRecovery,
					DeleteUser:          deleteUser,
					RecoverAccount:      recoverAccount,
					RefreshToken:        refreshToken,
					GetUsersUseCase:     getUsersUseCase,
					TokenService:        tokenService,
					OneTimeTokenService: oneTimeTokenService,
					EmailRateLimiter:    emailLimiter,
					EventBus:            eventBus,
					UserRepository:      userRepo,
					AvatarUsecases:      avatarUsecases,
				},
			}
		
			c.Directives.IsAuthenticated = func(ctx context.Context, obj interface{}, next graphql.Resolver) (res interface{}, err error) {
				_, err = auth.GetUserIDFromContext(ctx)
				if err != nil {
					return nil, errors.New("Access denied: User not authenticated.")
				}
				return next(ctx)
			}
		
			srv := handler.NewDefaultServer(graph.NewExecutableSchema(c))
			graphqlHandler := gin.WrapH(srv)
	r.POST("/graphql", auth.OptionalAuthMiddleware(tokenService, blacklistRepo), middleware.GinContextToContextMiddleware(), graphqlHandler)

	r.GET("/playground", gin.WrapH(playground.Handler("GraphQL playground", "/graphql")))

	return r, nil
}
