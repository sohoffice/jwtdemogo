Quick start
-----------

1. `go mod download` to update the required libraries
2. Generate RSA key pair. Run the below in files directory.
   ```
   openssl genrsa -out privatekey.pem 2048
   openssl rsa -in privatekey.pem -outform PEM -pubout -out publickey.pem
   ```
3. `go run main.go` to start the application in project folder. 
