# go-goat

Dumb vulnerable Go app created to test vulnerability scanners with Go container images.
There are several versions, newer versions "fix" several vulnerabilities with respect the previous one.


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

Edit deploy-k8s.yaml to change the tag of the go-goat image you want to deploy.
Deploy it into a kubernetes cluster using kubectl. 

````
kubectl apply -f deploy-k8s.yaml
````

# go-goat vulnerabilities

## 1.0
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
PullString: manuelbcd/go-goat:1.2

4 vulnerabilities found
0 Critical (0 fixable)
3 High (3 fixable)
1 Medium (1 fixable)
0 Low (0 fixable)
0 Negligible (0 fixable)

          PACKAGE            TYPE   VERSION  SUGGESTED FIX  CRITICAL  HIGH  MEDIUM  LOW  NEGLIGIBLE  EXPLOIT  
  github.com/gin-gonic/gin  golang  v1.6.0      v1.7.7         0       2      1      0       0          0     
  github.com/gogo/protobuf  golang  v1.3.0      v1.3.2         0       1      0      0       0          0     
