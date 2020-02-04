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
	service := &BlogService{
		firestore: pkg.NewFirestore(),
	}

	jEnc := json.NewEncoder(w)

	posts, err := service.ListPosts()
	if err != nil {
		klog.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		jEnc.Encode(map[string]string{"error": "something terrible happened"})
		return
	}

	w.WriteHeader(http.StatusOK)
	jEnc.Encode(posts)
}

// ListPosts from Firestore
func (b *BlogService) ListPosts() ([]BlogPost, error) {
	rootCollection := b.firestore.Collection("/devconf_blogposts")
	if rootCollection == nil {
		return []BlogPost{}, nil
	}

	documentIterator := rootCollection.Documents(context.Background())
	if documentIterator == nil {
		return []BlogPost{}, nil
	}

	documents, err := documentIterator.GetAll()
	if err != nil {
		return []BlogPost{}, err
	}

	posts := make([]BlogPost, len(documents))

	for idx, document := range documents {
		post := BlogPost{
			ID: document.Ref.ID,
		}

		document.DataTo(&post)

		posts[idx] = post
	}

	return posts, nil
}
