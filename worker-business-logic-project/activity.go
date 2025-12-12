package main

import (
	"context"
	"fmt"
)

func SayHelloActivity(ctx context.Context, name string) (string, error) {
	return fmt.Sprintf("Hello, %s! Welcome to Temporal.", name), nil
}
