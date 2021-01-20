/*
Copyright 2021.

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

package v1alpha1

import (
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/webhook"
)

// log is for logging in this package.
var logstashlog = logf.Log.WithName("logstash-resource")

func (r *Logstash) SetupWebhookWithManager(mgr ctrl.Manager) error {
	return ctrl.NewWebhookManagedBy(mgr).
		For(r).
		Complete()
}

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!

//+kubebuilder:webhook:path=/mutate-apps-russellgao-cn-v1alpha1-logstash,mutating=true,failurePolicy=fail,groups=apps.russellgao.cn,resources=logstashes,verbs=create;update,versions=v1alpha1,name=mlogstash.kb.io

var _ webhook.Defaulter = &Logstash{}

// Default implements webhook.Defaulter so a webhook will be registered for the type
func (r *Logstash) Default() {
	logstashlog.Info("default", "name", r.Name)

	// TODO(user): fill in your defaulting logic.
}

// TODO(user): change verbs to "verbs=create;update;delete" if you want to enable deletion validation.
//+kubebuilder:webhook:verbs=create;update,path=/validate-apps-russellgao-cn-v1alpha1-logstash,mutating=false,failurePolicy=fail,groups=apps.russellgao.cn,resources=logstashes,versions=v1alpha1,name=vlogstash.kb.io

var _ webhook.Validator = &Logstash{}

// ValidateCreate implements webhook.Validator so a webhook will be registered for the type
func (r *Logstash) ValidateCreate() error {
	logstashlog.Info("validate create", "name", r.Name)

	// TODO(user): fill in your validation logic upon object creation.
	return nil
}

// ValidateUpdate implements webhook.Validator so a webhook will be registered for the type
func (r *Logstash) ValidateUpdate(old runtime.Object) error {
	logstashlog.Info("validate update", "name", r.Name)

	// TODO(user): fill in your validation logic upon object update.
	return nil
}

// ValidateDelete implements webhook.Validator so a webhook will be registered for the type
func (r *Logstash) ValidateDelete() error {
	logstashlog.Info("validate delete", "name", r.Name)

	// TODO(user): fill in your validation logic upon object deletion.
	return nil
}
