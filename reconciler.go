// grpc-order-reconciler is a standalone service (its OWN repo) that reconciles
// orders by calling the hipstershop gRPC microservices defined in another repo
// (GoogleCloudPlatform/microservices-demo). Cross-repo gRPC consumer.
package main

import (
	"context"

	pb "github.com/GoogleCloudPlatform/microservices-demo/src/frontend/genproto"
	"google.golang.org/grpc"
)

type reconciler struct {
	cartConn     *grpc.ClientConn
	catalogConn  *grpc.ClientConn
	currencyConn *grpc.ClientConn
	shippingConn *grpc.ClientConn
	paymentConn  *grpc.ClientConn
}

func (r *reconciler) fetchCart(ctx context.Context, userID string) (*pb.Cart, error) {
	return pb.NewCartServiceClient(r.cartConn).GetCart(ctx, &pb.GetCartRequest{UserId: userID})
}

func (r *reconciler) listCatalog(ctx context.Context) (*pb.ListProductsResponse, error) {
	return pb.NewProductCatalogServiceClient(r.catalogConn).ListProducts(ctx, &pb.Empty{})
}

func (r *reconciler) getProduct(ctx context.Context, id string) (*pb.Product, error) {
	return pb.NewProductCatalogServiceClient(r.catalogConn).GetProduct(ctx, &pb.GetProductRequest{Id: id})
}

func (r *reconciler) convert(ctx context.Context, from *pb.Money) (*pb.Money, error) {
	return pb.NewCurrencyServiceClient(r.currencyConn).Convert(ctx, &pb.CurrencyConversionRequest{From: from})
}

func (r *reconciler) quote(ctx context.Context) (*pb.GetQuoteResponse, error) {
	return pb.NewShippingServiceClient(r.shippingConn).GetQuote(ctx, &pb.GetQuoteRequest{})
}

func (r *reconciler) charge(ctx context.Context, amt *pb.Money) (*pb.ChargeResponse, error) {
	return pb.NewPaymentServiceClient(r.paymentConn).Charge(ctx, &pb.ChargeRequest{Amount: amt})
}

func main() {}
