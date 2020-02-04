package function

import (
	"context"
	"encoding/json"
	"net/http"

	"cloud.google.com/go/firestore"
	pkg "github.com/kayteh/ctxs-devconf-demo/pkg/blog"
	"k8s.io/klog"
)

// BlogService is a skeleton RPC type for this function
type BlogService struct {
	firestore *firestore.Client
}

type BlogPost struct {
	ID    string `json:"id"`
	Title string `json:"title"`
	Body  string `json:"body"`
}

func Handle(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[1:]

	service := &BlogService{
		firestore: pkg.NewFirestore(),
	}

	jEnc := json.NewEncoder(w)

	post, err := service.GetPost(id)
	if err != nil {
		klog.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		jEnc.Encode(map[string]string{"error": "something terrible happened"})
		return
	}

	w.WriteHeader(http.StatusOK)
	jEnc.Encode(post)
}

// GetPost from Firestore
func (b *BlogService) GetPost(id string) (*BlogPost, error) {
	doc, err := b.firestore.Doc("devconf_blogposts/" + id).Get(context.Background())
	if err != nil {
		return nil, nil
	}

	post := &BlogPost{
		ID: id,
	}

	doc.DataTo(post)
	return post, nil
}
