# Basic Go gRPC Server and Client

This project, demonstrating simple, server-side streaming, client-side streaming, and bidirectional streaming RPC functionalities with Go and gRPC, offers a robust foundation for building high-performance, scalable, and interoperable microservices.

## Author
- Yan-Sheng (Arsene) Kuo &nbsp;<a href="https://www.linkedin.com/in/yanshengkuo/"><img src="https://upload.wikimedia.org/wikipedia/commons/c/ca/LinkedIn_logo_initials.png" alt="LinkedIn" style="height: 1em; width:auto;"/></a> &nbsp; <a href="https://github.com/YSKuo"> <img src="https://upload.wikimedia.org/wikipedia/commons/9/91/Octicons-mark-github.svg" alt="GitHub" style="height: 1em; width: auto;"/></a>

## Business use cases

1. Real-time Data Streaming and Analytics: financial services company that needs to provide real-time stock quotes and market data to its trading platform and analytical tools.

2. High-Performance API Gateway for Microservices: a large e-commerce platform built with numerous microservices (e.g., product catalog, order processing, inventory management) needs a central API gateway to handle client requests and orchestrate communication between these services.

3. Internet of Things (IoT) Data Ingestion and Control: a smart home platform needs to collect sensor data from numerous IoT devices (temperature sensors, smart locks, cameras) and send control commands back to them.

4. Real-time Chat and Collaboration Applications: A collaborative document editing tool or a real-time chat application requires instant updates between multiple users.

5. Batch Processing and Task Management: a data processing pipeline needs to distribute large batch processing tasks to multiple worker nodes and monitor their progress.

6. Cross-Platform Mobile Backend Communication: a mobile application needs to communicate with a backend server built using a different technology stack.

## Setting up a gRPC-Go project
1. Create a new directory for your project and cd into it

```bash
mkdir go-grpc
cd go-grpc
mkdir client server proto
```

2. Installing the gRPC Go plugin

```bash
go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28

go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2

export PATH="$PATH:$(go env GOPATH)/bin"
```

3. Initialize a Go module

```bash
go mod init github.com/your_username/go-grpc

go mod tidy
```

4. Create the proto file with the required services and messages in the proto directory

5. Generate .pb.go files from the proto file

depending on what path you mention in your greet.proto file, you will either run this - 

```bash
protoc --go_out=. --go-grpc_out=. proto/greet.proto
```

6. Create the server and client directories and create the main.go files with necessary controllers and services


## Running the application

1. Install the dependencies

```bash
go mod tidy
```

2. Run the server

```bash
go run server/main.go
```

3. Run the client

```bash
go run client/main.go
```

## conclusion

The Go gRPC server and client project provides a versatile foundation for building modern, high-performance applications across various domains. Its support for different RPC types makes it adaptable to a wide range of communication patterns, offering significant advantages in terms of speed, efficiency, and interoperability.