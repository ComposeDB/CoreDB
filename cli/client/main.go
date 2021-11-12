package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	pb "github.com/composeDB/coredb/internal/server/service/generated"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
)

const (
	address = "localhost:50051"
)

var createCmd = &cobra.Command{
	Use:  "create",
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
		if err != nil {
			log.Fatalf("Cannot connect: %v", err)
		}
		defer conn.Close()
		c := pb.NewCreateServiceClient(conn)
		name := args[0]
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()
		r, err := c.HandleCreate(ctx, &pb.CreateRequest{Name: name})
		if err != nil {
			log.Fatalf("could not greet: %v", err)
		}
		log.Printf("Greeting: %s", r.GetMessage())
	},
}

var rootCmd = &cobra.Command{
	Use:   "coredb",
	Short: "CoreDB: DB for Datasets",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Hello!")
	},
}

func Execute() {
	rootCmd.AddCommand(createCmd)
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func main() {
	Execute()
}
