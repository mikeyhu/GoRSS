export GOPATH=`pwd`

packages=(
	"gorss/atom"
	"gorss/rss"
	"gorss/domain"
	"gorss/state"
	"gorss/controllers"
	"gorss/collector"
	"gorss"
)

go fmt "${packages[@]}"

ops/test-clear.sh
ops/test-start.sh
go test "${packages[@]}"
ops/test-stop.sh

go install gorss
