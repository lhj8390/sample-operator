/*
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
*/

package v1

import (
	batchv1 "k8s.io/api/batch/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// ConcurrencyPolicy Job 이 어떻게 처리될지 설명. 기본값 AllowConcurrent
// +kubebuilder:validation:Enum=Allow;Forbid;Replace
type ConcurrencyPolicy string

const (
	// AllowConcurrent CronJobs 이 동시에 실행되도록 허용
	AllowConcurrent ConcurrencyPolicy = "Allow"

	// ForbidConcurrent 동시 실행을 금지하며, 이전 작업이 아직 끝나지 않았다면 다음 실행을 건너뛴다
	ForbidConcurrent ConcurrencyPolicy = "Forbid"

	// ReplaceConcurrent 현재 실행 중인 작업을 취소하고 새로운 작업으로 대체
	ReplaceConcurrent ConcurrencyPolicy = "Replace"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// CronJobSpec defines the desired state of CronJob
type CronJobSpec struct {
	// controller 로 들어오는 모든 "input" 은 spec 에 속한다.
	// +kubebuilder:validation:MinLength=0

	// Cron 형식의 일정
	Schedule string `json:"schedule"`

	// +kubebuilder:validation:Minimum=0

	// 예정된 시간에 작업을 시작하지 못한 경우 다시 시작하기 위한 제한 시간. 놓친 작업은 실패한 것으로 간주
	// +optional
	StartingDeadlineSeconds *int64 `json:"startingDeadlineSeconds,omitempty"`

	// Job 동시 실행 처리 방법 (Allow - 기본값, Forbid, Replace)
	// +optional
	ConcurrencyPolicy ConcurrencyPolicy `json:"concurrencyPolicy,omitempty"`

	// 컨트롤러에게 후속 실행을 중단하라고 지시. 이미 시작된 실행에는 적용되지 않음. 기본값 false
	// +optional
	Suspend *bool `json:"suspend,omitempty"`

	// CronJob을 실행할 때 생성될 Job 지정
	JobTemplate batchv1.JobTemplateSpec `json:"jobTemplate"`

	//+kubebuilder:validation:Minimum=0

	// History 에 유지할 성공적으로 완료된 Job 수.
	// +optional
	SuccessfulJobsHistoryLimit *int32 `json:"successfulJobsHistoryLimit,omitempty"`

	//+kubebuilder:validation:Minimum=0

	// +optional
	FailedJobsHistoryLimit *int32 `json:"failedJobsHistoryLimit,omitempty"`
}

// CronJobStatus defines the observed state of CronJob
type CronJobStatus struct {
	// 사용자나 다른 컨트롤러가 쉽게 얻을 수 있는 모든 정보를 포함하여 상태 정의

	// 현재 실행 중인 작업 목록.
	// +optional
	Active []corev1.ObjectReference `json:"active,omitempty"`

	// 작업이 마지막으로 성공적으로 예약된 시간에 대한 정보.
	// +optional
	LastScheduleTime *metav1.Time `json:"lastScheduleTime,omitempty"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// CronJob is the Schema for the cronjobs API
type CronJob struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   CronJobSpec   `json:"spec,omitempty"`
	Status CronJobStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// CronJobList contains a list of CronJob
type CronJobList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CronJob `json:"items"`
}

func init() {
	SchemeBuilder.Register(&CronJob{}, &CronJobList{})
}
