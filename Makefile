oapigen:
	oapi-codegen -generate types,chi-server -package api -o server/api/petstore.gen.go petstore-expanded.yaml
