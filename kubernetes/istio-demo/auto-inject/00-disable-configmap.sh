kubectl -n istio-system get configmap istio-sidecar-injector -o yaml | sed 's/policy: enabled/policy: disabled/g' | kubectl replace -n istio-system -f -
