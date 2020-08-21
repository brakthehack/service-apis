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

package v1alpha1pre1

import (
	internalinterfaces "github.com/vmware-tanzu/service-apis/pkg/client/informers/externalversions/internalinterfaces"
)

// Interface provides access to all the informers in this group version.
type Interface interface {
	// Gateways returns a GatewayInformer.
	Gateways() GatewayInformer
	// GatewayClasses returns a GatewayClassInformer.
	GatewayClasses() GatewayClassInformer
	// HTTPRoutes returns a HTTPRouteInformer.
	HTTPRoutes() HTTPRouteInformer
	// TCPRoutes returns a TCPRouteInformer.
	TCPRoutes() TCPRouteInformer
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

// Gateways returns a GatewayInformer.
func (v *version) Gateways() GatewayInformer {
	return &gatewayInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// GatewayClasses returns a GatewayClassInformer.
func (v *version) GatewayClasses() GatewayClassInformer {
	return &gatewayClassInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}

// HTTPRoutes returns a HTTPRouteInformer.
func (v *version) HTTPRoutes() HTTPRouteInformer {
	return &hTTPRouteInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// TCPRoutes returns a TCPRouteInformer.
func (v *version) TCPRoutes() TCPRouteInformer {
	return &tCPRouteInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}
