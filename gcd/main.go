package main

import (
	"log"
	"net"

	pb "github.com/lisuizhe/go-examples/gcd/gcd.pb"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
    "google.golang.org/grpc/reflection"
)