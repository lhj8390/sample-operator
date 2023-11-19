# sample-operator
// TODO(user): Add simple overview of use/purpose

## Description
// TODO(user): An in-depth paragraph about your project and overview of use

## Getting Started

### Prerequisites
- go version v1.20.0+
- docker version 17.03+.
- kubectl version v1.11.3+.
- Access to a Kubernetes v1.11.3+ cluster.

### To Deploy on the cluster
**Build and push your image to the location specified by `IMG`:**

```sh
make docker-build docker-push IMG=<some-registry>/sample-operator:tag
```

**NOTE:** This image ought to be published in the personal registry you specified. 
And it is required to have access to pull the image from the working environment. 
Make sure you have the proper permission to the registry if the above commands don’t work.

**Install the CRDs into the cluster:**

```sh
make install
```

**Deploy the Manager to the cluster with the image specified by `IMG`:**

```sh
make deploy IMG=<some-registry>/sample-operator:tag
```

> **NOTE**: If you encounter RBAC errors, you may need to grant yourself cluster-admin 
privileges or be logged in as admin.

**Create instances of your solution**
You can apply the samples (examples) from the config/sample:

```sh
kubectl apply -k config/samples/
```

>**NOTE**: Ensure that the samples has default values to test it out.

### To Uninstall
**Delete the instances (CRs) from the cluster:**

```sh
kubectl delete -k config/samples/
```

**Delete the APIs(CRDs) from the cluster:**

```sh
make uninstall
```

**UnDeploy the controller from the cluster:**

```sh
make undeploy
```

## Creating API

```shell
kubebuilder create api --group <그룹 이름> --version <버전 명> --kind <생성할 kind 이름>
```
프로젝트 하위에 api/<버전 명> 디렉터리가 생성된다.
- <kind 이름>_types.go : kind 에 대한 추가 메타데이터를 직접 정의할 수 있다.
- groupversion_info.go : group-version 에 대한 공통 메타데이터 포함
  - `+kubebuilder:object:generate=true` : 객체 생성기가 사용하는 패키지 수준의 마커
  - `+groupName=batch.tutorial.kubebuilder.io` : CRD들에 대한 올바른 메타데이터를 생성하기 위해 CRD 생성기에 의해 사용되는 마커
- zz_generated.deepcopy.go : runtime.Object 인터페이스가 사용하는 deepcopy 메서드.

## Implementing a controller

### CronJob Controller example

CronJob Controller의 기본 로직은 다음과 같다.
1. CronJob 로드
2. 실행중인 Job 목록을 가져오고, 상태를 업데이트
3. history 노출 개수에 따라 오래된 Job 정리
4. Job이 중지되었는지 확인 (중지 상태라면 다른 Job을 실행하지 않음)
5. 예약된 다음 Job 가져오기
6. deadline을 넘기지 않고, concurrency 정책에 의해 차단되지 않을 경우 새 Job 실행
7. 실행 중인 Job을 보거나 다음 Job의 실행 시간이 되었을 때 Requeue.

이러한 로직을 internal/controller 디렉터리 내부의 `cronjob_controller.go`에서 구현한다.

## Contributing
// TODO(user): Add detailed information on how you would like others to contribute to this project

**NOTE:** Run `make --help` for more information on all potential `make` targets

More information can be found via the [Kubebuilder Documentation](https://book.kubebuilder.io/introduction.html)

## License

Copyright 2023.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.

