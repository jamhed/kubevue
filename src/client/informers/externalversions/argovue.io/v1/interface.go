/*
Copyright The Kubernetes Authors.

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

// Code generated by informer-gen. DO NOT EDIT.

package v1

import (
	internalinterfaces "argovue/client/informers/externalversions/internalinterfaces"
)

// Interface provides access to all the informers in this group version.
type Interface interface {
	// AppConfigs returns a AppConfigInformer.
	AppConfigs() AppConfigInformer
	// Datasources returns a DatasourceInformer.
	Datasources() DatasourceInformer
	// Services returns a ServiceInformer.
	Services() ServiceInformer
	// Tokens returns a TokenInformer.
	Tokens() TokenInformer
}

type version struct {
	factory          internalinterfaces.SharedInformerFactory
	namespace        string
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// New returns a new Interface.
func New(f internalinterfaces.SharedInformerFactory, namespace string, tweakListOptions internalinterfaces.TweakListOptionsFunc) Interface {
	return &version{factory: f, namespace: namespace, tweakListOptions: tweakListOptions}
}

// AppConfigs returns a AppConfigInformer.
func (v *version) AppConfigs() AppConfigInformer {
	return &appConfigInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// Datasources returns a DatasourceInformer.
func (v *version) Datasources() DatasourceInformer {
	return &datasourceInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// Services returns a ServiceInformer.
func (v *version) Services() ServiceInformer {
	return &serviceInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// Tokens returns a TokenInformer.
func (v *version) Tokens() TokenInformer {
	return &tokenInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}
