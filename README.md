# gin-gonic practice
## Download source code from git
```
git clone git@github.com:k2star0118/practice-gin-gonic.git
```
## Config from intellij
1. Import: "Project From Existing Source" -> "Gradle"
2. Preferences Setting: "Preferences" -> "Laugnages & Frameworks -> Go"
   * GOROOT: version >= 1.11
   * GOPAYH: Select a folder, you can create a folder anywhere
   * GO Module (vgo): Check "Enable Go Modules (vgo) integration"
   
## Build your own docker
Please Follow the steps
```
# Generate vendor folder
$ go mod vendor

# Build
$ docker build --no-cache -t test/gin-gonic .
```

## Run your container
```
# Run the image
$ docker run -d -p 8888:8888 test/gin-gonic

# Test
$ curl http://localhost:8888
{"message":"Successfully to query get example"}
```