package function

import (
	"context"

	"cloud.google.com/go/firestore"
	empty "github.com/golang/protobuf/ptypes/empty"
	pkg "github.com/kayteh/ctxs-devconf-demo/pkg/blog"
	proto "github.com/kayteh/ctxs-devconf-rpc/blog"

	"google.golang.org/grpc"
)

// BlogService is a skeleton RPC type for this function
type BlogService struct {
	proto.UnimplementedBlogServer
	firestore *firestore.Client
}

// Register the function
func Register(server *grpc.Server) {
	proto.RegisterBlogServer(server, &BlogService{
		firestore: pkg.NewFirestore(),
	})
}

// ListPosts from Firestore
func (b *BlogService) ListPosts(ctx context.Context, req *empty.Empty) (*proto.PostList, error) {
	rootCollection := b.firestore.Collection("/devconf_blogposts")

	documents, err := rootCollection.Documents(context.Background()).GetAll()
	if err != nil {
		return nil, err
	}

	posts := make([]*proto.Post, len(documents))

	for idx, document := range documents {
		post := &proto.Post{
			Id: document.Ref.ID,
		}

		document.DataTo(post)

		posts[idx] = post
	}

	return &proto.PostList{
		Posts: posts,
	}, nil
}
