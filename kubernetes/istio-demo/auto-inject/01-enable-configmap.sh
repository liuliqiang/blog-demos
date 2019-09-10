kubectl -n istio-system get configmap istio-sidecar-injector -o yaml | sed 's/policy: disabled/policy: enabled/g' | kubectl replace -n istio-system -f -
