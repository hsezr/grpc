package main

import (
	"context"
	"log"

	pb "github.com/go-programming-tour-book/tag-service/proto"
	"google.golang.org/grpc"
)

type Auth struct {
	AppKey    string
	AppSrcret string
}

func (a *Auth) GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error) {
	return map[string]string{"app_key": a.AppKey, "app_secret": a.AppSrcret}, nil
}

func (a *Auth) RequireTransportSecurity() bool {
	return false
}

func main() {
	auth := Auth{
		AppKey:    "go-programming-tour-book",
		AppSrcret: "Hsezr",
	}
	ctx := context.Background()
	opts := []grpc.DialOption{grpc.WithPerRPCCredentials(&auth)}
	clientConn, err := GetClientConn(ctx, "localhost:8004", opts)
	if err != nil {
		log.Fatalf("err: %v", err)
	}
	defer clientConn.Close()

	tagServiceClient := pb.NewTagServiceClient(clientConn)
	resp, _ := tagServiceClient.GetTagList(ctx, &pb.GetTagListRequest{Name: "C"})
	log.Printf("resp: %v", resp)
}

func GetClientConn(ctx context.Context, target string, opts []grpc.DialOption) (*grpc.ClientConn, error) {
	opts = append(opts, grpc.WithInsecure())
	return grpc.DialContext(ctx, target, opts...)
}
