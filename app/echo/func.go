package main

import (
	"context"
	"fmt"
	fdk "github.com/fnproject/fdk-go"
	"io"
)

func main() {
	fdk.Handle(fdk.HandlerFunc(myHandler))
}

func myHandler(ctx context.Context, in io.Reader, out io.Writer) {
	//	fnctx := fdk.Context(ctx)
	fdk.WriteStatus(out, 200)
	fmt.Println("Receive a notification")

}
