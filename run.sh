echo "1) Start server"
echo "2) Start client"
echo "3) Start playground"
echo "4) Kill running server and client"
echo "5) Run test"
echo "6) Run build"

read -p "Type: " cmd
if [[ $cmd == 1 ]]; then
    go run cmd/server/main.go
elif [[ $cmd == 2 ]]; then
    go run cmd/client/main.go
elif [[ $cmd == 3 ]]; then
    go run cmd/playground/main.go
elif [[ $cmd == 4 ]]; then
    sudo kill -9 $(sudo lsof -t -i:8888)
elif [[ $cmd == 5 ]]; then
    go test -coverpkg=github.com/nahK994/TinyCache/pkg/resp tests/resp/* -v
    go test -coverpkg=github.com/nahK994/TinyCache/pkg/handlers tests/handlers/* -v
    go test -coverpkg=github.com/nahK994/TinyCache/pkg/cache tests/cache/* -v
    go test -coverpkg=github.com/nahK994/TinyCache/connection/utils tests/utils/* -v
elif [[ $cmd == 6 ]]; then
    go build -ldflags="-s -w" -o ./bin/tinycache cmd/client/main.go
    go build -ldflags="-s -w" -o ./bin/tinycache-server cmd/server/main.go
fi
