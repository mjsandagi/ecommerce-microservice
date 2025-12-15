# Go eCommerce Microservice

A lightweight, production-ready REST API built with **Go (Golang)**, containerised with **Docker (Multi-stage)**, and orchestrated via **Kubernetes**.

This project demonstrates the transition from traditional heavy runtime environments (like Python/Java) to efficient, statically linked infrastructure suitable for high-performance platform engineering.

## Architectural Decisions

### 1. Why Go?

-   **Concurrency:** Utilises the `net/http` standard library and `chi` router for low-latency request handling.
-   **Static Linking:** The app compiles to a single binary, eliminating the need for a heavy OS runtime in production.
-   **Type Safety:** Uses interfaces (`products.Service`) to decouple the handler from the business logic, making the system testable and modular.

### 2. Docker Strategy (Distroless)

I utilised a **Multi-Stage Build** in the Dockerfile.

-   **Stage 1 (Builder):** Compiles the code using the full Golang toolchain.
-   **Stage 2 (Runtime):** Copies _only_ the binary into a `gcr.io/distroless/static` image.
-   **Result:** The final image is **~10MB** (vs ~800MB for a standard Python image), massively reducing security risks and deployment time.

### 3. Kubernetes Resiliency

The `k8s/deployment.yaml` configuration ensures:

-   **High Availability:** `replicas: 3` ensures the service can handle node failures without downtime.
-   **Self-Healing:** A `livenessProbe` checks the `/` health endpoint; if the app hangs, Kubernetes automatically restarts the pod.
-   **Resource Quotas:** CPU/Memory limits prevent resource issues on the cluster.

## Quick Start

### Prerequisites

-   Docker
-   Minikube (Kubernetes)

### 1. Build the Image

```bash
docker build -t go-ecommerce:latest .
```

### 2\. Deploy to Kubernetes

```bash
# Apply the Deployment and Service
kubectl apply -f k8s/deployment.yaml
kubectl apply -f k8s/service.yaml

# Verify the Pods are running
kubectl get pods
```

### 3\. Test the Endpoint

Forward the port to your local machine to test:

```bash
kubectl port-forward svc/go-ecommerce-svc 8080:80
```

Visit `http://localhost:8080/products` to see the JSON response.

## Project Structure

-   `cmd/`: Entry points for the application.
-   `internal/`: Private application code (Handlers, Services).
-   `k8s/`: Infrastructure as Code (IaC) manifests.
