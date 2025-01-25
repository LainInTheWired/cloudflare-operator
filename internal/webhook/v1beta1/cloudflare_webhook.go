/*
Copyright 2025.

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

package v1beta1

import (
	"context"
	"fmt"

	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/webhook"
	"sigs.k8s.io/controller-runtime/pkg/webhook/admission"

	cloudflarev1beta1 "github.com/laininthewired/cloudflare-ingress-controller/api/v1beta1"
)

// nolint:unused
// log is for logging in this package.
var cloudflarelog = logf.Log.WithName("cloudflare-resource")

// SetupCloudflareWebhookWithManager registers the webhook for Cloudflare in the manager.
func SetupCloudflareWebhookWithManager(mgr ctrl.Manager) error {
	return ctrl.NewWebhookManagedBy(mgr).For(&cloudflarev1beta1.Cloudflare{}).
		WithValidator(&CloudflareCustomValidator{}).
		WithDefaulter(&CloudflareCustomDefaulter{}).
		Complete()
}

// TODO(user): EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!

// +kubebuilder:webhook:path=/mutate-cloudflare-laininthewired-github-io-v1beta1-cloudflare,mutating=true,failurePolicy=fail,sideEffects=None,groups=cloudflare.laininthewired.github.io,resources=cloudflares,verbs=create;update,versions=v1beta1,name=mcloudflare-v1beta1.kb.io,admissionReviewVersions=v1

// CloudflareCustomDefaulter struct is responsible for setting default values on the custom resource of the
// Kind Cloudflare when those are created or updated.
//
// NOTE: The +kubebuilder:object:generate=false marker prevents controller-gen from generating DeepCopy methods,
// as it is used only for temporary operations and does not need to be deeply copied.
type CloudflareCustomDefaulter struct {
	// TODO(user): Add more fields as needed for defaulting
}

var _ webhook.CustomDefaulter = &CloudflareCustomDefaulter{}

// Default implements webhook.CustomDefaulter so a webhook will be registered for the Kind Cloudflare.
func (d *CloudflareCustomDefaulter) Default(ctx context.Context, obj runtime.Object) error {
	cloudflare, ok := obj.(*cloudflarev1beta1.Cloudflare)

	if !ok {
		return fmt.Errorf("expected an Cloudflare object but got %T", obj)
	}
	cloudflarelog.Info("Defaulting for Cloudflare", "name", cloudflare.GetName())

	// TODO(user): fill in your defaulting logic.

	return nil
}

// TODO(user): change verbs to "verbs=create;update;delete" if you want to enable deletion validation.
// NOTE: The 'path' attribute must follow a specific pattern and should not be modified directly here.
// Modifying the path for an invalid path can cause API server errors; failing to locate the webhook.
// +kubebuilder:webhook:path=/validate-cloudflare-laininthewired-github-io-v1beta1-cloudflare,mutating=false,failurePolicy=fail,sideEffects=None,groups=cloudflare.laininthewired.github.io,resources=cloudflares,verbs=create;update,versions=v1beta1,name=vcloudflare-v1beta1.kb.io,admissionReviewVersions=v1

// CloudflareCustomValidator struct is responsible for validating the Cloudflare resource
// when it is created, updated, or deleted.
//
// NOTE: The +kubebuilder:object:generate=false marker prevents controller-gen from generating DeepCopy methods,
// as this struct is used only for temporary operations and does not need to be deeply copied.
type CloudflareCustomValidator struct {
	// TODO(user): Add more fields as needed for validation
}

var _ webhook.CustomValidator = &CloudflareCustomValidator{}

// ValidateCreate implements webhook.CustomValidator so a webhook will be registered for the type Cloudflare.
func (v *CloudflareCustomValidator) ValidateCreate(ctx context.Context, obj runtime.Object) (admission.Warnings, error) {
	cloudflare, ok := obj.(*cloudflarev1beta1.Cloudflare)
	if !ok {
		return nil, fmt.Errorf("expected a Cloudflare object but got %T", obj)
	}
	cloudflarelog.Info("Validation for Cloudflare upon creation", "name", cloudflare.GetName())

	// TODO(user): fill in your validation logic upon object creation.

	return nil, nil
}

// ValidateUpdate implements webhook.CustomValidator so a webhook will be registered for the type Cloudflare.
func (v *CloudflareCustomValidator) ValidateUpdate(ctx context.Context, oldObj, newObj runtime.Object) (admission.Warnings, error) {
	cloudflare, ok := newObj.(*cloudflarev1beta1.Cloudflare)
	if !ok {
		return nil, fmt.Errorf("expected a Cloudflare object for the newObj but got %T", newObj)
	}
	cloudflarelog.Info("Validation for Cloudflare upon update", "name", cloudflare.GetName())

	// TODO(user): fill in your validation logic upon object update.

	return nil, nil
}

// ValidateDelete implements webhook.CustomValidator so a webhook will be registered for the type Cloudflare.
func (v *CloudflareCustomValidator) ValidateDelete(ctx context.Context, obj runtime.Object) (admission.Warnings, error) {
	cloudflare, ok := obj.(*cloudflarev1beta1.Cloudflare)
	if !ok {
		return nil, fmt.Errorf("expected a Cloudflare object but got %T", obj)
	}
	cloudflarelog.Info("Validation for Cloudflare upon deletion", "name", cloudflare.GetName())

	// TODO(user): fill in your validation logic upon object deletion.

	return nil, nil
}
