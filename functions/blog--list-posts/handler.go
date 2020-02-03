package function

import (
	"context"

	proto "github.com/kayteh/ctxs-devconf-rpc/blog"
	pkg "github.com/kayteh/ctxs-devconf-demo/pkg/blog"
	"google.golang.org/grpc"
)

// Skeleton Blog Service
type BlogService struct {
	proto.UnimplementedBlogServer
	pkg.BlogMixin
}

func Register(server *grpc.Server) {
	proto.RegisterBlogServer(server, &BlogService{
		pkg.NewBlogMixin()
	})
}

func (b *BlogService) ListPosts(ctx context.Context, req *empty.Empty) (*proto.PostList, error) {
	return nil, nil
}
