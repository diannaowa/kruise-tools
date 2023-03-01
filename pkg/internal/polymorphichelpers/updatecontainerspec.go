/*
Copyright 2023 The Kubernetes Authors.

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

package polymorphichelpers

import (
	"fmt"

	kruiseappsv1alpha1 "github.com/openkruise/kruise-api/apps/v1alpha1"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

func updateContainerSpecForObject(obj runtime.Object, fn func(container *v1.Container) error) (bool, error) {
	switch t := obj.(type) {
	case *kruiseappsv1alpha1.SidecarSet:
		for i, c := range t.Spec.Containers {
			err := fn(&c.Container)
			if err != nil {
				return false, err
			}
			t.Spec.Containers[i] = c
		}
		return true, nil
	default:
		return false, fmt.Errorf("the object is not a container or does not have a container spec: %T", t)
	}
}
