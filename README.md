> **Part of the [QorVertex live demo](https://qorvertex.com).** A small sample application that shows a gRPC client calling the services of Google's open-source microservices-demo, in another repository. It is intentionally minimal and not production code.

# grpc-order-reconciler

A **cross-repo gRPC consumer** sample application for the QorVertex
dependency catalog. It is a standalone Go service (its own repository) that calls
the **hipstershop** gRPC microservices defined in a *separate* repo,
[`GoogleCloudPlatform/microservices-demo`](https://github.com/GoogleCloudPlatform/microservices-demo).

It exists to validate that the scanner + catalog resolve gRPC client calls
**across repository boundaries**:

- `pb.NewCartServiceClient(conn).GetCart(...)` → `hipstershop.CartService.GetCart`
- ...and the same for ProductCatalog / Currency / Shipping / Payment services.

After scanning both repos, microservices-demo's gRPC Services should list
`grpc-order-reconciler` under **"Consumed by"**, and each rpc-method endpoint
should show it under **"Called by"** — cross-repo.

> Fixture only — the proto/genproto imports are illustrative; it is not meant to
> `go build`. It is scanned as source to exercise cross-repo gRPC resolution.
