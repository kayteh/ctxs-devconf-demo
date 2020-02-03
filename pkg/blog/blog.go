package blog

import (
	"cloud.google.com/go/firestore"
	"golang.org/x/net/context"
	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
	"k8s.io/klog"
)

type BlogMixin struct {
	db *firestore.Client
}

func NewBlogMixin() BlogMixin {
	opt := option.WithCredentialsFile("/var/openfaas/secrets/firebase-config")
	ctx := context.Background()
	app, err := firebase.NewApp(ctx, nil, opt)
	if err != nil {
		klog.Fatal(err)
	}

	db, err := app.Firestore(ctx)
	if err != nil {
		klog.Fatal(err)
	}

	return BlogMixin{
		db: db,
	}
}
