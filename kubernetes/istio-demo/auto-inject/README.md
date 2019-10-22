## 这个示例用于控制 Istio 的 Sidecar 注入

- 打开 namespace 的自动注入
    - ```$ kubectl label namespace default istio-injection=enabled```
- 00-disable-configmap.sh
	- 关闭 configmap 中自动注入的配置
- 01-enable-configmap.sh
	- 打开 configmap 中自动注入的配置
- 10-no-injected-pod.yaml
	- 在关闭 cm 之后，不会被自动注入的测试例子
- 11-injected-pod.yaml
	- 在关闭 cm 之后，会被自动注入的测试例子
