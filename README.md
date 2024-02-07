# go-goat


# Build

General command 

docker build -t <your-namespace>/go-goat .

If you are using a Mac and want to build for Linux

docker buildx build --platform linux/amd64 -t <your-namespace>>/py-goat:1.0 .

Example

````
docker buildx build --platform linux/amd64 -t manuelbcd/go-goat:1.0 .
````



# Deploy
