package router

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/sirupsen/logrus"
	"gitlab.com/learnProto/configs"
	"gitlab.com/learnProto/pkg/v1/utils/constants"
	"google.golang.org/grpc"

	hlpb "gitlab.com/learnProto/proto/v1/health"
	lspb "gitlab.com/learnProto/proto/v1/luas"
)

// NewHTTPServer creates the http server serve mux.
func NewHTTPServer(configs *configs.Configs, logger *logrus.Logger) error {
	// runtime.HTTPError = errors.CustomHTTPError

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	//opts := httplog.NewOpts()

	// Connect to the GRPC server
	conn, err := grpc.Dial(fmt.Sprintf("0.0.0.0:%v", configs.Config.Server.Grpc.Port), grpc.WithInsecure())
	if err != nil {
		return err
	}
	defer conn.Close()

	// Create new grpc-gateway
	rmux := runtime.NewServeMux()

	// register gateway endpoints
	for _, f := range []func(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error{
		// register grpc service handler
		hlpb.RegisterHealthServiceHandler,
		lspb.RegisterLuasServiceHandler,
	} {
		if err = f(ctx, rmux, conn); err != nil {
			return err
		}
	}

	// create http server mux
	mux := http.NewServeMux()
	mux.Handle("/", rmux)

	// run swagger server
	if configs.Config.Env != constants.EnvProduction {
		CreateSwagger(mux)
	}

	// Where ORIGIN_ALLOWED is like `scheme://dns[:port]`, or `*` (insecure)
	headersOk := handlers.AllowedHeaders([]string{"Accept", "Accept-Language", "Content-Language", "Content-Type", "X-Requested-With", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization", "Timezone-Offset"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})

	// running rest http server
	log.Println("[SERVER] REST HTTP server is ready")

	err = http.ListenAndServe(fmt.Sprintf("0.0.0.0:%v", configs.Config.Server.Rest.Port), handlers.CORS(headersOk, originsOk, methodsOk)(mux))
	if err != nil {
		return err
	}

	return nil
}

// CreateSwagger creates the swagger server serve mux.
func CreateSwagger(gwmux *http.ServeMux) {
	// register swagger service server
	gwmux.HandleFunc("/api/go-grpc-learning/health/docs.json", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "swagger/docs.json")
	})

	// load swagger-ui file
	fs := http.FileServer(http.Dir("swagger/swagger-ui"))
	gwmux.Handle("/api/go-grpc-learning/health/docs/", http.StripPrefix("/api/go-grpc-learning/health/docs", fs))
}
