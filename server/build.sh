export GOPATH=`pwd`

go fmt gorss/atom gorss/rss gorss/domain gorss

ops/test-clear.sh
ops/test-start.sh
go test gorss/atom gorss/rss gorss/domain gorss/ingestion gorss
ops/test-stop.sh

go install gorss
