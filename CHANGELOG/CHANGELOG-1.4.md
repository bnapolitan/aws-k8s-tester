

<hr>


## [v1.4.0](https://github.com/aws/aws-k8s-tester/releases/tag/v1.4.0) (2020-06)

See [code changes](https://github.com/aws/aws-k8s-tester/compare/v1.3.9...v1.4.0).

### `ec2-utils`

- [`ec2-utils --auto-path` now automatically uses the generated cluster name for local file paths, instead of random string](https://github.com/aws/aws-k8s-tester/commit/53b51d38b1aa4e6ea1454cc631c9511dcbe4267a).

### `aws-k8s-tester`

- [`aws-k8s-tester --auto-path` now automatically uses the generated cluster name for local file paths, instead of random string](https://github.com/aws/aws-k8s-tester/commit/53b51d38b1aa4e6ea1454cc631c9511dcbe4267a).

### `ec2`

- Clean up [`colorstring` printf](https://github.com/aws/aws-k8s-tester/pull/101).

### `eks`

- Publish [performance data to CloudWatch](https://github.com/aws/aws-k8s-tester/commit/038fd83e6a180d5a98287b508d243661b23a356a).
- Add [compare tests for all stressing tests, useful for regression tests](https://github.com/aws/aws-k8s-tester/commit/28939738fd0ca8aeaf512fef43f706472650ab13).
- Improve [`eks/configmaps`, `eks/csrs`, `eks/secrets` results collect with S3](https://github.com/aws/aws-k8s-tester/commit/a8500fbf1b9218ca587d265daed6a6845b3ebfcb).
- Add [`eks/tester.Tester.Name` method](https://github.com/aws/aws-k8s-tester/commit/2f8f08595d53f18abe77c47a6f43c6e734127f53).
- Fix [`eks/stresser` collect metrics](https://github.com/aws/aws-k8s-tester/commit/2f8f08595d53f18abe77c47a6f43c6e734127f53).
- Clean up [`colorstring` printf](https://github.com/aws/aws-k8s-tester/pull/101).
- Clean up [polling operation error handling](https://github.com/aws/aws-k8s-tester/commit/26627f14f4dbbcc8dd64d6307ed6e58c0b809f52).
  - Rename [`eks/cluster/version-upgrade.Poll` to `eks/cluster/wait.PollUpdate`](https://github.com/aws/aws-k8s-tester/commit/a6eeea26a7ab3c7069a4278026b56de87707c9b1).
- Discard [HTTP download progress for URL checks](https://github.com/aws/aws-k8s-tester/commit/d54e2c7b125d22779b014fb0eb0ac72e165b2350).
- Increase [`cluster-loader` timeout and error if output is not expected](https://github.com/aws/aws-k8s-tester/commit/13ff01fad653249435770138069ef600b0c873fa).
- Run [`AWS_K8S_TESTER_EKS_ADD_ON_CLUSTER_LOADER` before `AWS_K8S_TESTER_EKS_ADD_ON_HOLLOW_NODES`](https://github.com/aws/aws-k8s-tester/commit/c0377d39377c47da86677028209585b854046e1a).
  - *TODO*: Handle `23 resource_gather_worker.go:63] error while reading data from hollowdreamprimehaze8iw5ul: the server is currently unable to handle the request (get nodes hollowdreamprimehaze8iw5ul:10250)`.
- Improve [`eks/cluster-loader/remote` result download, use `"scp"`](https://github.com/aws/aws-k8s-tester/commit/a3c9d7d5e3382c378de686fe0faec6bdeb47f027).
- Store [`kube-apiserver` `/metrics` output](https://github.com/aws/aws-k8s-tester/commit/9e7985fe8ffc948866e792d0984faafbf4e57c59).
- Add [`eks/cluster-loader.ParsePodStartupLatency`](https://github.com/aws/aws-k8s-tester/commit/322cd88e94e879157f6b409f9c604fdbbc95e465).
- Add [`eks/cluster-loader.MergePodStartupLatency`](https://github.com/aws/aws-k8s-tester/commit/322cd88e94e879157f6b409f9c604fdbbc95e465).
- Record [`PodStartupLatency` in `eksconfig.AddOnClusterLoaderLocal` via `eks/cluster-loader/local`](https://github.com/aws/aws-k8s-tester/commit/8d4cb87b7bd798ad7f1b5d2de22d0deb26c4c75e).
- Record [`PodStartupLatency` in `eksconfig.AddOnClusterLoaderRemote` via `eks/cluster-loader/remote`](https://github.com/aws/aws-k8s-tester/commit/8d4cb87b7bd798ad7f1b5d2de22d0deb26c4c75e).
- Fix [`eks/cluster-loader` error handling](https://github.com/aws/aws-k8s-tester/commit/4a9d982929c32efbdfad820e0cece67721d53034).
- Add [S3 access policies to worker node roles](https://github.com/aws/aws-k8s-tester/commit/bcf0b1da501fc9a1bcf1a7691e690e729ee95b59).
- Improve [`eks/stresser/remote` results fetch](https://github.com/aws/aws-k8s-tester/commit/a982c1b484d8133b113bfa1f22df6698411898b7).
- Fix [multiple `eks/cluster-loader` runs](https://github.com/aws/aws-k8s-tester/commit/a7d6ebc79d76782d5bbff533183d9baa05bd663e).
- Add [extra namespace force-deletion function to `eks/stresser/remote`](https://github.com/aws/aws-k8s-tester/commit/dc6ef6849a57d2236bc23a0a89413a7b377a211c).

- [`eks/mng/scale` added to scale mngs up and down](https://github.com/aws/aws-k8s-tester/pull/106)
  - See https://docs.aws.amazon.com/cli/latest/reference/eks/update-nodegroup-config.html
### `eksconfig`

 - Add [`ClusterAutoscaler` add-on per node-group using `AWS_K8S_TESTER_EKS_ADD_ON_NODE_GROUPS_ASGS={"GetRef.Name-...":{..."cluster-autoscaler":{"enable":false}...}}`](https://github.com/aws/aws-k8s-tester/pull/99).
- Fix [typo in `eksconfig.AddOnManagedNodeGroups.LogsTarGzPath`](https://github.com/aws/aws-k8s-tester/commit/7b60047ca4d6fad281db512d4de905a27b80303a).
- Add [`Status.PrivateDNSToSSHConfig` for node SSH access](https://github.com/aws/aws-k8s-tester/commit/a3c9d7d5e3382c378de686fe0faec6bdeb47f027).
- Add [`Status.ClusterMetricsRawOutputDir` for `kube-apiserver` `/metrics`](https://github.com/aws/aws-k8s-tester/commit/3ee7554e14f53feae7c5b8ebb1ee4d50b71e0bd7).
- Record [`PodStartupLatency` in `eksconfig.AddOnClusterLoaderLocal` via `eks/cluster-loader/local`](https://github.com/aws/aws-k8s-tester/commit/8d4cb87b7bd798ad7f1b5d2de22d0deb26c4c75e).
- Record [`PodStartupLatency` in `eksconfig.AddOnClusterLoaderRemote` via `eks/cluster-loader/remote`](https://github.com/aws/aws-k8s-tester/commit/8d4cb87b7bd798ad7f1b5d2de22d0deb26c4c75e).
- Add [`RequestsWritesSummaryS3Dir` and `RequestsReadsSummaryS3Dir`](https://github.com/aws/aws-k8s-tester/commit/e559fae84787e7936fd167cd7da9a893c691e856).

### `ssh`

- Move [`"scp"` executable binary path check before creating context timeouts](https://github.com/aws/aws-k8s-tester/commit/4c3950e6582745684a9d628c5c0ea355e3f7edc1).
- Fix and [improve retries](https://github.com/aws/aws-k8s-tester/commit/949cc1ea63131ce7d27808d7fc12d6e988d07978).

### `pkg`

- Add [`pkg/aws/s3`](https://github.com/aws/aws-k8s-tester/commit/a982c1b484d8133b113bfa1f22df6698411898b7).
- Add [`pkg/k8s-client.EKSConfig.MetricsRawOutputDir` to store `kube-apiserver` `/metrics` output](https://github.com/aws/aws-k8s-tester/commit/9e7985fe8ffc948866e792d0984faafbf4e57c59).
- Add [`pkg/k8s-client.WithForceDelete` option for `DeleteNamespaceAndWait`](https://github.com/aws/aws-k8s-tester/commit/803ba2d263227adea026fcf1bb5262ebb2abd230).
  - Fix https://github.com/aws/aws-k8s-tester/issues/100.
  - See [`kubernetes/kubernetes#60807`](https://github.com/kubernetes/kubernetes/issues/60807).
- Add [`pkg/metrics.RequestsSummaryCompare`](https://github.com/aws/aws-k8s-tester/commit/00b7c5c922f77db2243fb8d5c26c0e0f9fd9d484).

### Dependency

- Upgrade [`github.com/aws/aws-sdk-go`](https://github.com/aws/aws-sdk-go/releases) from [`v1.32.2`](https://github.com/aws/aws-sdk-go/releases/tag/v1.32.2) to [`v1.32.7`](https://github.com/aws/aws-sdk-go/releases/tag/v1.32.7).
- Upgrade [`github.com/kubernetes/kubernetes`](https://github.com/kubernetes/kubernetes/releases) from [`v0.18.2`](https://github.com/kubernetes/clienthttps://github.com/kubernetes/kubernetes/releases/tag/v0.18.3) to [`v0.18.4`](https://github.com/kubernetes/kubernetes/releases/tag/v0.18.4).
- Upgrade [`github.com/kubernetes/client-go`](https://github.com/kubernetes/client-go/releases) from [`v0.18.3`](https://github.com/kubernetes/clienthttps://github.com/kubernetes/client-go/releases/tag/v0.18.3) to [`v0.18.4`](https://github.com/kubernetes/client-go/releases/tag/v0.18.4).
  - See [commit `fc93a579` for all the changes](https://github.com/aws/aws-k8s-tester/commit/fc93a5792c7334fc099e18ad4a4de394f8c2a35c).
- Add [`k8s.io/perf-tests`](https://github.com/kubernetes/perf-tests/releases).
  - See [`1aea23d3` for commit](https://github.com/aws/aws-k8s-tester/commit/1aea23d3259794307b45d344d3a953238c394efb).


<hr>

