## 这个例子用于说明如何访问 service mesh 之外的应用

### 1. 启动两个客户端(curl)

#### 文件配置

- 00-no-injected-curl-pods.yaml
- 01-injected-curl-pods.yaml

#### 描述

- 这两个文件用于启动两个 curl pod：
    - 00-no-injected-curl-pods.yaml： 这个 pod 是没有 istio-proxy 注入的
    - 01-injected-curl-pods.yaml：这个 pod 是有 istio-proxy 注入的
    ```
    [root@liqiang.io]# kubectl get pods
    NAME                       READY   STATUS    RESTARTS   AGE
    injected-curl              2/2     Running   0          12m
    no-injected-curl           1/1     Running   0          8m38s
    ```

### 2. 在 Pod 中直接用 Pod 内部的 Envoy 访问外部服务

#### 文件配置

- 10-virtual-service.yaml
- 11-destionation-rule.yaml

#### 描述

- 这两个文件用于将一个 https 服务以 http 的形式添加到 service mesh 中
    - 00-no-injected-curl-pods.yaml：因为没有 sidecar，所以无法正常使用
        ```
        [root@liqiang.io]# kubectl exec -it no-injected-curl /bin/sh
        / # curl -I http://liqiang.io
        HTTP/1.1 301 Moved Permanently
        Server: nginx/1.12.2
        Date: Tue, 08 Oct 2019 08:09:36 GMT
        Content-Type: text/html
        Content-Length: 185
        Connection: keep-alive
        Location: https://liqiang.io/
        ```
    - 01-injected-curl-pods.yaml：因为存在 sidecar，所以可以直接通过 http 访问 https
        ```
        [root@liqiang.io]# kubectl exec -it injected-curl /bin/sh
        / # curl -I http://liqiang.io
        HTTP/1.1 200 OK
        server: envoy
        date: Tue, 08 Oct 2019 08:11:38 GMT
        content-type: text/html; charset=utf-8
        content-length: 15354
        x-envoy-upstream-service-time: 737
        ```

### 3. 通过 egress 访问外部服务

#### 文件配置

- 20-egress-gateway.yaml
- 21-virtual-service.yaml
- 22-destination-rule.yaml

#### 描述

- 这 3 个文件通过 egress gateway 的方式通过 http 访问 https 的 liqiang.io
- 但是在这个版本中，no-injected-curl 还是无法使用
    - 20-egress-gateway.yaml：在 service mesh 中添加一个 egress 的监听
    - 21-virtual-service.yaml：定义了如何路由 liqiang.io 的流量，这里转了两层
    - 22-destination-rule.yaml：定义了如何和真正的 liqiang.io 连接，这里使用的是 SIMPLE 的 TLS 认证

### 4. 通过 ingress 访问外部服务

#### 文件配置

- 30-get-ingress-info.sh
- 31-service-for-ingress.yaml
- 32-liqiangio-ingress-gateway.yaml
- 33-liqiangio-ingress-virtualservice.yaml
- 34-liqiangio-ingress-destination-rule.yaml

#### 描述

- 这个 5 个文件通过 ingress gateway 的方式，以 http 的形式访问 https 的 liqiang.io
- 在这个版本中，未对 injected-curl 做设置，所以 injected-curl 无法使用
	- 30-get-ingress-info.sh：获取 ingress 的 host 和 port，主要是用于本地调试
	- 31-service-for-ingress.yaml：不带 sidecar 的客户端访问 ingress 使用
	- 32-liqiangio-ingress-gateway.yaml：ingress 接收 80 端口的请求
	- 33-liqiangio-ingress-virtualservice.yaml：ingress 直接将请求路由到 liqiang.io，并且修改 Host Header
	- 34-liqiangio-ingress-destination-rule.yaml：定义如何和外部服务连接

### 5. 通过 ingress 和 egress 暴露外部服务

#### 文件配置

- 40-get-ingress-info.sh
- 41-service-for-ingress.yaml
- 42-liqiangio-ingress-gateway.yaml
- 43-liqiangio-ingress-virtualservice.yaml
- 44-liqiangio-egress-gateway.yaml
- 45-liqiangio-egress-virtual-service.yaml
- 46-liqiangio-egress-destination-rule.yaml

#### 描述

- 这 7 个文件通过 ingress 的接受 http 请求，然后转发给 egress，egress 通过 TLS 的形式与外部服务对接
- 在这个版本中，无论是带 sidecar 还是不带 sidecar 的客户端都可以使用
	- 40-get-ingress-info.sh：获取 ingress 的 host 和 port，主要是用于本地调试
	- 41-service-for-ingress.yaml：不带 sidecar 的客户端访问 ingress 使用
	- 42-liqiangio-ingress-gateway.yaml：ingress 接收 80 端口的请求
	- 43-liqiangio-ingress-virtualservice.yaml：ingress 路由 http 请求给 egress，并且修改 Host Header
	- 44-liqiangio-egress-gateway.yaml：egress 开放对外 80 端口
	- 45-liqiangio-egress-virtual-service.yaml：egress 路由 http 请求到外部服务
	- 46-liqiangio-egress-destination-rule.yaml：定义如何和外部服务连接
