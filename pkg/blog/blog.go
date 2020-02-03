package blog

import (
	"cloud.google.com/go/firestore"
	"golang.org/x/net/context"
	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
	"k8s.io/klog"
)

func NewFirestore() *firestore.Client {
	opt := option.WithCredentialsFile("/var/openfaas/secrets/firebase-config")
	ctx := context.Background()
	app, err := firebase.NewApp(ctx, nil, opt)
	if err != nil {
		klog.Fatal(err)
	}

	firestore, err := app.Firestore(ctx)
	if err != nil {
		klog.Fatal(err)
	}

	return firestore
}
