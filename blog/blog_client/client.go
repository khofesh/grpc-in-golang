package main

import (
	"blog/blogpb"
	"context"
	"flag"
	"fmt"
	"io"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	addr = flag.String("addr", "localhost:50051", "the address to connect to")
)

func main() {
	opts := grpc.WithTransportCredentials(insecure.NewCredentials())

	cc, err := grpc.Dial(*addr, opts)
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer cc.Close()

	c := blogpb.NewBlogServiceClient(cc)

	// creating a blog
	fmt.Println("creating the blog")
	blog := &blogpb.Blog{
		AuthorId: "Fahmi",
		Title:    "My First Blog",
		Content:  "This the content",
	}
	createBlogRes, err := c.CreateBlog(context.Background(), &blogpb.CreateBlogRequest{Blog: blog})
	if err != nil {
		log.Fatalf("Unexpected error: %v", err)
	}
	fmt.Printf("blog has been created: %v\n", createBlogRes)
	blogId := createBlogRes.GetBlog().GetId()

	// read blog
	fmt.Println("reading the blog")
	_, err2 := c.ReadBlog(context.Background(), &blogpb.ReadBlogRequest{
		BlogId: "ssdfwe",
	})
	if err2 != nil {
		fmt.Printf("Error happened while reading: %v\n", err2)
	}

	readBlogReq := &blogpb.ReadBlogRequest{BlogId: blogId}
	readBlogRes, readBlogErr := c.ReadBlog(context.Background(), readBlogReq)
	if readBlogErr != nil {
		fmt.Printf("Error happened while reading: %v\n", readBlogErr)
	}

	fmt.Printf("Blog was read: %v\n", readBlogRes)

	// update blog
	fmt.Println("update the blog")
	newBlog := &blogpb.Blog{
		Id:       blogId,
		AuthorId: "Changed author",
		Title:    "My Second Blog",
		Content:  "This the content, lksdflksoiwer lksjdflskdjf",
	}

	updateRes, updateErr := c.UpdateBlog(context.Background(), &blogpb.UpdateBlogRequest{Blog: newBlog})
	if updateErr != nil {
		fmt.Printf("Error happened while updating: %v\n", updateErr)
	}

	fmt.Printf("Blog was read: %v\n", updateRes)

	// delete blog
	deleteRes, deleteErr := c.DeleteBlog(context.Background(), &blogpb.DeleteBlogRequest{BlogId: blogId})
	if deleteErr != nil {
		fmt.Printf("Error happened while deleting: %v\n", deleteErr)
	}

	fmt.Printf("Blog was deleted: %v\n", deleteRes)

	// list blog
	stream, err := c.ListBlog(context.Background(), &blogpb.ListBlogRequest{})
	if err != nil {
		log.Fatalf("error while calling Sum RPC: %v", err)
	}

	for {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Something happened: %v", err)
		}
		fmt.Println(res.GetBlog())
	}
}
