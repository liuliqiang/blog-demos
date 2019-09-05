PWD=$(pwd)
DST=$PWD/90-tls
kubectl create -n istio-system secret generic httpbin-credential \
--from-file=key=$DST/httpbin.example.com/3_application/private/httpbin.example.com.key.pem \
--from-file=cert=$DST/httpbin.example.com/3_application/certs/httpbin.example.com.cert.pem
