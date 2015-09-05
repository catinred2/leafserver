GOPATH=$GOPATH:/Users/ming/Documents/go/src/github.com/name5566/leafserver
cd /Users/ming/Documents/go/src/github.com/name5566/leafserver/src/client
echo "Clean And Build"
go clean
go build
cd /Users/ming/Documents/go/src/github.com/name5566/leafserver/bin
../src/client/client -id=$1

