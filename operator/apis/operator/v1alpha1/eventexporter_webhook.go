// Licensed to Apache Software Foundation (ASF) under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Apache Software Foundation (ASF) licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

package v1alpha1

import (
	"fmt"

	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/webhook"
	"sigs.k8s.io/controller-runtime/pkg/webhook/admission"
)

const latestVersion = "latest"
const image = "apache/skywalking-kubernetes-event-exporter"

// log is for logging in this package.
var eventexporterlog = logf.Log.WithName("eventexporter-resource")

func (r *EventExporter) SetupWebhookWithManager(mgr ctrl.Manager) error {
	return ctrl.NewWebhookManagedBy(mgr).
		For(r).
		Complete()
}

// nolint: lll
//+kubebuilder:webhook:path=/mutate-operator-skywalking-apache-org-v1alpha1-eventexporter,mutating=true,failurePolicy=fail,sideEffects=None,groups=operator.skywalking.apache.org,resources=eventexporters,verbs=create;update,versions=v1alpha1,name=meventexporter.kb.io,admissionReviewVersions=v1

var _ webhook.Defaulter = &EventExporter{}

// Default implements webhook.Defaulter so a webhook will be registered for the type
func (r *EventExporter) Default() {
	eventexporterlog.Info("default", "name", r.Name)

	if r.Spec.Version == "" {
		r.Spec.Version = latestVersion
	}

	if r.Spec.Image == "" {
		r.Spec.Image = fmt.Sprintf("%s:%s", image, r.Spec.Version)
	}

	if r.Spec.Replicas == 0 {
		r.Spec.Replicas = 1
	}
}

// nolint: lll
// +kubebuilder:webhook:admissionReviewVersions=v1,sideEffects=None,path=/mutate-operator-skywalking-apache-org-v1alpha1-eventexporter,mutating=true,failurePolicy=fail,groups=operator.skywalking.apache.org,resources=eventexporters,verbs=create;update,versions=v1alpha1,name=meventexporter.kb.io

var _ webhook.Validator = &EventExporter{}

// ValidateCreate implements webhook.Validator so a webhook will be registered for the type
func (r *EventExporter) ValidateCreate() (admission.Warnings, error) {
	eventexporterlog.Info("validate create", "name", r.Name)

	return nil, r.validate()
}

// ValidateUpdate implements webhook.Validator so a webhook will be registered for the type
func (r *EventExporter) ValidateUpdate(_ runtime.Object) (admission.Warnings, error) {
	eventexporterlog.Info("validate update", "name", r.Name)

	return nil, r.validate()
}

// ValidateDelete implements webhook.Validator so a webhook will be registered for the type
func (r *EventExporter) ValidateDelete() (admission.Warnings, error) {
	eventexporterlog.Info("validate delete", "name", r.Name)

	return nil, nil
}

func (r *EventExporter) validate() error {
	if r.Spec.Image == "" {
		return fmt.Errorf("image is absent")
	}
	return nil
}
