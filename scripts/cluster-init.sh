cmd () {
    echo "> $@"
    "$@" >/dev/null
}

cmd k3sup app install openfaas --set=faasIdler.dryRun=false
cmd k3sup app install nginx-ingress
cmd k3sup app install cert-manager

kubectl rollout status -n openfaas deploy/gateway
cmd kubectl port-forward -n openfaas svc/gateway 8080:8080 &
sleep 1

echo "*** Logging into OpenFaaS Gateway"
kubectl get secret -n openfaas basic-auth -o jsonpath="{.data.basic-auth-password}" | base64 --decode > /tmp/faas-pw
cat /tmp/faas-pw | faas login --password-stdin

cmd kubectl apply -f ./k8s/ingress-operator