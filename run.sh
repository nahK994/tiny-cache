echo "1) start process"
echo "2) kill process"

read -p "Type: " cmd
if [[ $cmd == 1 ]]; then
   go run cmd/main.go
elif [[ $cmd == 2 ]]; then
    sudo kill -9 $(sudo lsof -t -i:8000)
fi
