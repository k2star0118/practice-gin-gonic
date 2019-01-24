# gin-gonic practice
## Download source code from git
``
git clone git@github.com:k2star0118/practice-gin-gonic.git
``
## Config from intellij
1. Import: "Project From Existing Source" -> "Gradle"
2. Preferences Setting: "Preferences" -> "Laugnages & Frameworks -> Go"
   * GOROOT: version >= 1.11
   * GOPAYH: Select a folder, you can create a folder anywhere
   * GO Module (vgo): Check "Enable Go Modules (vgo) integration"
   
## Run via debug mode
1. Find "/gin-gonic/main.go", and click mouse right key "Debug 'go build main.go'"
2. Edit the "go build main.go", change "Working folder" to your project path ".../gin-gonic" and save
3. Click "debug" to run

## Try http server
```
curl -X GET http://localhost:8888/example
```
You will get the response with json format
```
{"message":"Successfully to query get example"}
```