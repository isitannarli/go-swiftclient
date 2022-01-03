# Go Swift Client

### Args

```text
-auth-key string
    Auth key/password
    
-auth-url string
    Auth url
    
-auth-user string
    Auth username
    
-container-name string
    Container name (example: frontend/assets)
    
-path string
    File or directory path (example: ./test.txt OR ./directory/sub-directory)
```

### Usage

```shell
./build/go-swiftclient-darwin-amd64 -auth-url=AUTH_URL -auth-user=AUTH_USERNAME -auth-key=AUTH_KEY -container-name=frontend -path=./test-directory
```

### Build

```shell
make
```