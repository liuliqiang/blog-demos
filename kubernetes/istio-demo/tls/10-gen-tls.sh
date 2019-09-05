PWD=$(pwd)
DST=$PWD/90-tls
echo $DST
git clone https://github.com/nicholasjackson/mtls-go-example $DST/mtls-go-example
pushd $DST/mtls-go-example
$DST/mtls-go-example/generate.sh httpbin.example.com password
popd
