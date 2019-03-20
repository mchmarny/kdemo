# Assumes following env vars set
#  * GCP_PROJECT - ID of your project
#  * KDEMO_OAUTH_CLIENT_ID - Google OAuth2 Client ID
#  * KDEMO_OAUTH_CLIENT_SECRET - Google OAuth2 Client Secret

test:
	go test ./... -v

deps:
	go mod tidy

image:
	gcloud builds submit \
		--project ${GCP_PROJECT} \
		--tag gcr.io/${GCP_PROJECT}/kdemo

secrets:
	# kubectl delete secret kdemo -n demo
	kubectl create secret generic kdemo -n demo \
		--from-literal=OAUTH_CLIENT_ID=$(DEMO_OAUTH_CLIENT_ID) \
		--from-literal=OAUTH_CLIENT_SECRET=$(DEMO_OAUTH_CLIENT_SECRET)

config:
	kubectl apply -f deployment/config.yaml -n demo

service:
	kubectl apply -f deployment/service.yaml -n demo

