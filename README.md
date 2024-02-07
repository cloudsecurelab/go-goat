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


# go-goat vulnerabilities

## 1.0

ImageID: sha256:499908ea84d54d04b15928df8baf4bf2691d83296f767386fd1cd686266ffc1b
Digest: sha256:8df1802e1b8189f223428d2514794bef9b8166d4faea1d49ac912315ccb14c8e
PullString: manuelbcd/go-goat:1.0

14 vulnerabilities found
0 Critical (0 fixable)
11 High (10 fixable)
3 Medium (3 fixable)
0 Low (0 fixable)
0 Negligible (0 fixable)

          PACKAGE            TYPE   VERSION  SUGGESTED FIX  CRITICAL  HIGH  MEDIUM  LOW  NEGLIGIBLE  EXPLOIT  
  github.com/tidwall/gjson  golang  v1.3.0      v1.9.3         0       5      0      0       0          0     
  github.com/gin-gonic/gin  golang  v1.4.0      v1.6.0         0       3      1      0       0          0     
  gopkg.in/yaml.v2          golang  v2.2.2      v2.2.4         0       1      2      0       0          0     
  github.com/gogo/protobuf  golang  v1.3.0      v1.3.2         0       1      0      0       0          0     

## 1.1

ImageID: sha256:b33847c7e91611b09eeee9e241ec10e85f76c304525a6fbbd7a47f53d969a2cd
Digest: sha256:096a80c93ff0a59eda9b12e7723fae370ea38d4288e23322fce1610838f1b393
PullString: manuelbcd/go-goat:1.1

9 vulnerabilities found
0 Critical (0 fixable)
6 High (5 fixable)
3 Medium (3 fixable)
0 Low (0 fixable)
0 Negligible (0 fixable)

          PACKAGE            TYPE   VERSION  SUGGESTED FIX  CRITICAL  HIGH  MEDIUM  LOW  NEGLIGIBLE  EXPLOIT  
  github.com/gin-gonic/gin  golang  v1.4.0      v1.6.0         0       3      1      0       0          0     
  gopkg.in/yaml.v2          golang  v2.2.2      v2.2.4         0       1      2      0       0          0     
  github.com/gogo/protobuf  golang  v1.3.0      v1.3.2         0       1      0      0       0          0   


## 1.2
