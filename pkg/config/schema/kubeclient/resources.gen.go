// GENERATED FILE -- DO NOT EDIT
//

package kubeclient

import (
	"context"
	"fmt"

	k8sioapiadmissionregistrationv1 "k8s.io/api/admissionregistration/v1"
	k8sioapiappsv1 "k8s.io/api/apps/v1"
	k8sioapicertificatesv1 "k8s.io/api/certificates/v1"
	k8sioapicorev1 "k8s.io/api/core/v1"
	k8sioapidiscoveryv1 "k8s.io/api/discovery/v1"
	k8sioapinetworkingv1 "k8s.io/api/networking/v1"
	k8sioapiextensionsapiserverpkgapisapiextensionsv1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/tools/cache"
	sigsk8siogatewayapiapisv1alpha2 "sigs.k8s.io/gateway-api/apis/v1alpha2"
	sigsk8siogatewayapiapisv1beta1 "sigs.k8s.io/gateway-api/apis/v1beta1"

	apiistioioapiextensionsv1alpha1 "istio.io/client-go/pkg/apis/extensions/v1alpha1"
	apiistioioapinetworkingv1alpha3 "istio.io/client-go/pkg/apis/networking/v1alpha3"
	apiistioioapinetworkingv1beta1 "istio.io/client-go/pkg/apis/networking/v1beta1"
	apiistioioapisecurityv1beta1 "istio.io/client-go/pkg/apis/security/v1beta1"
	apiistioioapitelemetryv1alpha1 "istio.io/client-go/pkg/apis/telemetry/v1alpha1"
	"istio.io/istio/pkg/config/schema/gvr"
	"istio.io/istio/pkg/kube/informerfactory"
	ktypes "istio.io/istio/pkg/kube/kubetypes"
	"istio.io/istio/pkg/ptr"
)

func GetWriteClient[T runtime.Object](c ClientGetter, namespace string) ktypes.WriteAPI[T] {
	switch any(ptr.Empty[T]()).(type) {
	case *apiistioioapisecurityv1beta1.AuthorizationPolicy:
		return c.Istio().SecurityV1beta1().AuthorizationPolicies(namespace).(ktypes.WriteAPI[T])
	case *k8sioapicertificatesv1.CertificateSigningRequest:
		return c.Kube().CertificatesV1().CertificateSigningRequests().(ktypes.WriteAPI[T])
	case *k8sioapicorev1.ConfigMap:
		return c.Kube().CoreV1().ConfigMaps(namespace).(ktypes.WriteAPI[T])
	case *k8sioapiextensionsapiserverpkgapisapiextensionsv1.CustomResourceDefinition:
		return c.Ext().ApiextensionsV1().CustomResourceDefinitions().(ktypes.WriteAPI[T])
	case *k8sioapiappsv1.Deployment:
		return c.Kube().AppsV1().Deployments(namespace).(ktypes.WriteAPI[T])
	case *apiistioioapinetworkingv1alpha3.DestinationRule:
		return c.Istio().NetworkingV1alpha3().DestinationRules(namespace).(ktypes.WriteAPI[T])
	case *k8sioapidiscoveryv1.EndpointSlice:
		return c.Kube().DiscoveryV1().EndpointSlices(namespace).(ktypes.WriteAPI[T])
	case *k8sioapicorev1.Endpoints:
		return c.Kube().CoreV1().Endpoints(namespace).(ktypes.WriteAPI[T])
	case *apiistioioapinetworkingv1alpha3.EnvoyFilter:
		return c.Istio().NetworkingV1alpha3().EnvoyFilters(namespace).(ktypes.WriteAPI[T])
	case *sigsk8siogatewayapiapisv1alpha2.GRPCRoute:
		return c.GatewayAPI().GatewayV1alpha2().GRPCRoutes(namespace).(ktypes.WriteAPI[T])
	case *apiistioioapinetworkingv1alpha3.Gateway:
		return c.Istio().NetworkingV1alpha3().Gateways(namespace).(ktypes.WriteAPI[T])
	case *sigsk8siogatewayapiapisv1beta1.GatewayClass:
		return c.GatewayAPI().GatewayV1beta1().GatewayClasses().(ktypes.WriteAPI[T])
	case *sigsk8siogatewayapiapisv1beta1.HTTPRoute:
		return c.GatewayAPI().GatewayV1beta1().HTTPRoutes(namespace).(ktypes.WriteAPI[T])
	case *k8sioapinetworkingv1.Ingress:
		return c.Kube().NetworkingV1().Ingresses(namespace).(ktypes.WriteAPI[T])
	case *k8sioapinetworkingv1.IngressClass:
		return c.Kube().NetworkingV1().IngressClasses().(ktypes.WriteAPI[T])
	case *sigsk8siogatewayapiapisv1beta1.Gateway:
		return c.GatewayAPI().GatewayV1beta1().Gateways(namespace).(ktypes.WriteAPI[T])
	case *k8sioapiadmissionregistrationv1.MutatingWebhookConfiguration:
		return c.Kube().AdmissionregistrationV1().MutatingWebhookConfigurations().(ktypes.WriteAPI[T])
	case *k8sioapicorev1.Namespace:
		return c.Kube().CoreV1().Namespaces().(ktypes.WriteAPI[T])
	case *k8sioapicorev1.Node:
		return c.Kube().CoreV1().Nodes().(ktypes.WriteAPI[T])
	case *apiistioioapisecurityv1beta1.PeerAuthentication:
		return c.Istio().SecurityV1beta1().PeerAuthentications(namespace).(ktypes.WriteAPI[T])
	case *k8sioapicorev1.Pod:
		return c.Kube().CoreV1().Pods(namespace).(ktypes.WriteAPI[T])
	case *apiistioioapinetworkingv1beta1.ProxyConfig:
		return c.Istio().NetworkingV1beta1().ProxyConfigs(namespace).(ktypes.WriteAPI[T])
	case *sigsk8siogatewayapiapisv1beta1.ReferenceGrant:
		return c.GatewayAPI().GatewayV1beta1().ReferenceGrants(namespace).(ktypes.WriteAPI[T])
	case *apiistioioapisecurityv1beta1.RequestAuthentication:
		return c.Istio().SecurityV1beta1().RequestAuthentications(namespace).(ktypes.WriteAPI[T])
	case *k8sioapicorev1.Secret:
		return c.Kube().CoreV1().Secrets(namespace).(ktypes.WriteAPI[T])
	case *k8sioapicorev1.Service:
		return c.Kube().CoreV1().Services(namespace).(ktypes.WriteAPI[T])
	case *k8sioapicorev1.ServiceAccount:
		return c.Kube().CoreV1().ServiceAccounts(namespace).(ktypes.WriteAPI[T])
	case *apiistioioapinetworkingv1alpha3.ServiceEntry:
		return c.Istio().NetworkingV1alpha3().ServiceEntries(namespace).(ktypes.WriteAPI[T])
	case *apiistioioapinetworkingv1alpha3.Sidecar:
		return c.Istio().NetworkingV1alpha3().Sidecars(namespace).(ktypes.WriteAPI[T])
	case *sigsk8siogatewayapiapisv1alpha2.TCPRoute:
		return c.GatewayAPI().GatewayV1alpha2().TCPRoutes(namespace).(ktypes.WriteAPI[T])
	case *sigsk8siogatewayapiapisv1alpha2.TLSRoute:
		return c.GatewayAPI().GatewayV1alpha2().TLSRoutes(namespace).(ktypes.WriteAPI[T])
	case *apiistioioapitelemetryv1alpha1.Telemetry:
		return c.Istio().TelemetryV1alpha1().Telemetries(namespace).(ktypes.WriteAPI[T])
	case *sigsk8siogatewayapiapisv1alpha2.UDPRoute:
		return c.GatewayAPI().GatewayV1alpha2().UDPRoutes(namespace).(ktypes.WriteAPI[T])
	case *k8sioapiadmissionregistrationv1.ValidatingWebhookConfiguration:
		return c.Kube().AdmissionregistrationV1().ValidatingWebhookConfigurations().(ktypes.WriteAPI[T])
	case *apiistioioapinetworkingv1alpha3.VirtualService:
		return c.Istio().NetworkingV1alpha3().VirtualServices(namespace).(ktypes.WriteAPI[T])
	case *apiistioioapiextensionsv1alpha1.WasmPlugin:
		return c.Istio().ExtensionsV1alpha1().WasmPlugins(namespace).(ktypes.WriteAPI[T])
	case *apiistioioapinetworkingv1alpha3.WorkloadEntry:
		return c.Istio().NetworkingV1alpha3().WorkloadEntries(namespace).(ktypes.WriteAPI[T])
	case *apiistioioapinetworkingv1alpha3.WorkloadGroup:
		return c.Istio().NetworkingV1alpha3().WorkloadGroups(namespace).(ktypes.WriteAPI[T])
	default:
		panic(fmt.Sprintf("Unknown type %T", ptr.Empty[T]()))
	}
}

func GetClient[T, TL runtime.Object](c ClientGetter, namespace string) ktypes.ReadWriteAPI[T, TL] {
	switch any(ptr.Empty[T]()).(type) {
	case *apiistioioapisecurityv1beta1.AuthorizationPolicy:
		return c.Istio().SecurityV1beta1().AuthorizationPolicies(namespace).(ktypes.ReadWriteAPI[T, TL])
	case *k8sioapicertificatesv1.CertificateSigningRequest:
		return c.Kube().CertificatesV1().CertificateSigningRequests().(ktypes.ReadWriteAPI[T, TL])
	case *k8sioapicorev1.ConfigMap:
		return c.Kube().CoreV1().ConfigMaps(namespace).(ktypes.ReadWriteAPI[T, TL])
	case *k8sioapiextensionsapiserverpkgapisapiextensionsv1.CustomResourceDefinition:
		return c.Ext().ApiextensionsV1().CustomResourceDefinitions().(ktypes.ReadWriteAPI[T, TL])
	case *k8sioapiappsv1.Deployment:
		return c.Kube().AppsV1().Deployments(namespace).(ktypes.ReadWriteAPI[T, TL])
	case *apiistioioapinetworkingv1alpha3.DestinationRule:
		return c.Istio().NetworkingV1alpha3().DestinationRules(namespace).(ktypes.ReadWriteAPI[T, TL])
	case *k8sioapidiscoveryv1.EndpointSlice:
		return c.Kube().DiscoveryV1().EndpointSlices(namespace).(ktypes.ReadWriteAPI[T, TL])
	case *k8sioapicorev1.Endpoints:
		return c.Kube().CoreV1().Endpoints(namespace).(ktypes.ReadWriteAPI[T, TL])
	case *apiistioioapinetworkingv1alpha3.EnvoyFilter:
		return c.Istio().NetworkingV1alpha3().EnvoyFilters(namespace).(ktypes.ReadWriteAPI[T, TL])
	case *sigsk8siogatewayapiapisv1alpha2.GRPCRoute:
		return c.GatewayAPI().GatewayV1alpha2().GRPCRoutes(namespace).(ktypes.ReadWriteAPI[T, TL])
	case *apiistioioapinetworkingv1alpha3.Gateway:
		return c.Istio().NetworkingV1alpha3().Gateways(namespace).(ktypes.ReadWriteAPI[T, TL])
	case *sigsk8siogatewayapiapisv1beta1.GatewayClass:
		return c.GatewayAPI().GatewayV1beta1().GatewayClasses().(ktypes.ReadWriteAPI[T, TL])
	case *sigsk8siogatewayapiapisv1beta1.HTTPRoute:
		return c.GatewayAPI().GatewayV1beta1().HTTPRoutes(namespace).(ktypes.ReadWriteAPI[T, TL])
	case *k8sioapinetworkingv1.Ingress:
		return c.Kube().NetworkingV1().Ingresses(namespace).(ktypes.ReadWriteAPI[T, TL])
	case *k8sioapinetworkingv1.IngressClass:
		return c.Kube().NetworkingV1().IngressClasses().(ktypes.ReadWriteAPI[T, TL])
	case *sigsk8siogatewayapiapisv1beta1.Gateway:
		return c.GatewayAPI().GatewayV1beta1().Gateways(namespace).(ktypes.ReadWriteAPI[T, TL])
	case *k8sioapiadmissionregistrationv1.MutatingWebhookConfiguration:
		return c.Kube().AdmissionregistrationV1().MutatingWebhookConfigurations().(ktypes.ReadWriteAPI[T, TL])
	case *k8sioapicorev1.Namespace:
		return c.Kube().CoreV1().Namespaces().(ktypes.ReadWriteAPI[T, TL])
	case *k8sioapicorev1.Node:
		return c.Kube().CoreV1().Nodes().(ktypes.ReadWriteAPI[T, TL])
	case *apiistioioapisecurityv1beta1.PeerAuthentication:
		return c.Istio().SecurityV1beta1().PeerAuthentications(namespace).(ktypes.ReadWriteAPI[T, TL])
	case *k8sioapicorev1.Pod:
		return c.Kube().CoreV1().Pods(namespace).(ktypes.ReadWriteAPI[T, TL])
	case *apiistioioapinetworkingv1beta1.ProxyConfig:
		return c.Istio().NetworkingV1beta1().ProxyConfigs(namespace).(ktypes.ReadWriteAPI[T, TL])
	case *sigsk8siogatewayapiapisv1beta1.ReferenceGrant:
		return c.GatewayAPI().GatewayV1beta1().ReferenceGrants(namespace).(ktypes.ReadWriteAPI[T, TL])
	case *apiistioioapisecurityv1beta1.RequestAuthentication:
		return c.Istio().SecurityV1beta1().RequestAuthentications(namespace).(ktypes.ReadWriteAPI[T, TL])
	case *k8sioapicorev1.Secret:
		return c.Kube().CoreV1().Secrets(namespace).(ktypes.ReadWriteAPI[T, TL])
	case *k8sioapicorev1.Service:
		return c.Kube().CoreV1().Services(namespace).(ktypes.ReadWriteAPI[T, TL])
	case *k8sioapicorev1.ServiceAccount:
		return c.Kube().CoreV1().ServiceAccounts(namespace).(ktypes.ReadWriteAPI[T, TL])
	case *apiistioioapinetworkingv1alpha3.ServiceEntry:
		return c.Istio().NetworkingV1alpha3().ServiceEntries(namespace).(ktypes.ReadWriteAPI[T, TL])
	case *apiistioioapinetworkingv1alpha3.Sidecar:
		return c.Istio().NetworkingV1alpha3().Sidecars(namespace).(ktypes.ReadWriteAPI[T, TL])
	case *sigsk8siogatewayapiapisv1alpha2.TCPRoute:
		return c.GatewayAPI().GatewayV1alpha2().TCPRoutes(namespace).(ktypes.ReadWriteAPI[T, TL])
	case *sigsk8siogatewayapiapisv1alpha2.TLSRoute:
		return c.GatewayAPI().GatewayV1alpha2().TLSRoutes(namespace).(ktypes.ReadWriteAPI[T, TL])
	case *apiistioioapitelemetryv1alpha1.Telemetry:
		return c.Istio().TelemetryV1alpha1().Telemetries(namespace).(ktypes.ReadWriteAPI[T, TL])
	case *sigsk8siogatewayapiapisv1alpha2.UDPRoute:
		return c.GatewayAPI().GatewayV1alpha2().UDPRoutes(namespace).(ktypes.ReadWriteAPI[T, TL])
	case *k8sioapiadmissionregistrationv1.ValidatingWebhookConfiguration:
		return c.Kube().AdmissionregistrationV1().ValidatingWebhookConfigurations().(ktypes.ReadWriteAPI[T, TL])
	case *apiistioioapinetworkingv1alpha3.VirtualService:
		return c.Istio().NetworkingV1alpha3().VirtualServices(namespace).(ktypes.ReadWriteAPI[T, TL])
	case *apiistioioapiextensionsv1alpha1.WasmPlugin:
		return c.Istio().ExtensionsV1alpha1().WasmPlugins(namespace).(ktypes.ReadWriteAPI[T, TL])
	case *apiistioioapinetworkingv1alpha3.WorkloadEntry:
		return c.Istio().NetworkingV1alpha3().WorkloadEntries(namespace).(ktypes.ReadWriteAPI[T, TL])
	case *apiistioioapinetworkingv1alpha3.WorkloadGroup:
		return c.Istio().NetworkingV1alpha3().WorkloadGroups(namespace).(ktypes.ReadWriteAPI[T, TL])
	default:
		panic(fmt.Sprintf("Unknown type %T", ptr.Empty[T]()))
	}
}

func gvrToObject(g schema.GroupVersionResource) runtime.Object {
	switch g {
	case gvr.AuthorizationPolicy:
		return &apiistioioapisecurityv1beta1.AuthorizationPolicy{}
	case gvr.CertificateSigningRequest:
		return &k8sioapicertificatesv1.CertificateSigningRequest{}
	case gvr.ConfigMap:
		return &k8sioapicorev1.ConfigMap{}
	case gvr.CustomResourceDefinition:
		return &k8sioapiextensionsapiserverpkgapisapiextensionsv1.CustomResourceDefinition{}
	case gvr.Deployment:
		return &k8sioapiappsv1.Deployment{}
	case gvr.DestinationRule:
		return &apiistioioapinetworkingv1alpha3.DestinationRule{}
	case gvr.EndpointSlice:
		return &k8sioapidiscoveryv1.EndpointSlice{}
	case gvr.Endpoints:
		return &k8sioapicorev1.Endpoints{}
	case gvr.EnvoyFilter:
		return &apiistioioapinetworkingv1alpha3.EnvoyFilter{}
	case gvr.GRPCRoute:
		return &sigsk8siogatewayapiapisv1alpha2.GRPCRoute{}
	case gvr.Gateway:
		return &apiistioioapinetworkingv1alpha3.Gateway{}
	case gvr.GatewayClass:
		return &sigsk8siogatewayapiapisv1beta1.GatewayClass{}
	case gvr.HTTPRoute:
		return &sigsk8siogatewayapiapisv1beta1.HTTPRoute{}
	case gvr.Ingress:
		return &k8sioapinetworkingv1.Ingress{}
	case gvr.IngressClass:
		return &k8sioapinetworkingv1.IngressClass{}
	case gvr.KubernetesGateway:
		return &sigsk8siogatewayapiapisv1beta1.Gateway{}
	case gvr.MutatingWebhookConfiguration:
		return &k8sioapiadmissionregistrationv1.MutatingWebhookConfiguration{}
	case gvr.Namespace:
		return &k8sioapicorev1.Namespace{}
	case gvr.Node:
		return &k8sioapicorev1.Node{}
	case gvr.PeerAuthentication:
		return &apiistioioapisecurityv1beta1.PeerAuthentication{}
	case gvr.Pod:
		return &k8sioapicorev1.Pod{}
	case gvr.ProxyConfig:
		return &apiistioioapinetworkingv1beta1.ProxyConfig{}
	case gvr.ReferenceGrant:
		return &sigsk8siogatewayapiapisv1beta1.ReferenceGrant{}
	case gvr.RequestAuthentication:
		return &apiistioioapisecurityv1beta1.RequestAuthentication{}
	case gvr.Secret:
		return &k8sioapicorev1.Secret{}
	case gvr.Service:
		return &k8sioapicorev1.Service{}
	case gvr.ServiceAccount:
		return &k8sioapicorev1.ServiceAccount{}
	case gvr.ServiceEntry:
		return &apiistioioapinetworkingv1alpha3.ServiceEntry{}
	case gvr.Sidecar:
		return &apiistioioapinetworkingv1alpha3.Sidecar{}
	case gvr.TCPRoute:
		return &sigsk8siogatewayapiapisv1alpha2.TCPRoute{}
	case gvr.TLSRoute:
		return &sigsk8siogatewayapiapisv1alpha2.TLSRoute{}
	case gvr.Telemetry:
		return &apiistioioapitelemetryv1alpha1.Telemetry{}
	case gvr.UDPRoute:
		return &sigsk8siogatewayapiapisv1alpha2.UDPRoute{}
	case gvr.ValidatingWebhookConfiguration:
		return &k8sioapiadmissionregistrationv1.ValidatingWebhookConfiguration{}
	case gvr.VirtualService:
		return &apiistioioapinetworkingv1alpha3.VirtualService{}
	case gvr.WasmPlugin:
		return &apiistioioapiextensionsv1alpha1.WasmPlugin{}
	case gvr.WorkloadEntry:
		return &apiistioioapinetworkingv1alpha3.WorkloadEntry{}
	case gvr.WorkloadGroup:
		return &apiistioioapinetworkingv1alpha3.WorkloadGroup{}
	default:
		panic(fmt.Sprintf("Unknown type %v", g))
	}
}

func getInformerFiltered(c ClientGetter, opts ktypes.InformerOptions, g schema.GroupVersionResource) informerfactory.StartableInformer {
	var l func(options metav1.ListOptions) (runtime.Object, error)
	var w func(options metav1.ListOptions) (watch.Interface, error)

	switch g {
	case gvr.AuthorizationPolicy:
		l = func(options metav1.ListOptions) (runtime.Object, error) {
			return c.Istio().SecurityV1beta1().AuthorizationPolicies(opts.Namespace).List(context.Background(), options)
		}
		w = func(options metav1.ListOptions) (watch.Interface, error) {
			return c.Istio().SecurityV1beta1().AuthorizationPolicies(opts.Namespace).Watch(context.Background(), options)
		}
	case gvr.CertificateSigningRequest:
		l = func(options metav1.ListOptions) (runtime.Object, error) {
			return c.Kube().CertificatesV1().CertificateSigningRequests().List(context.Background(), options)
		}
		w = func(options metav1.ListOptions) (watch.Interface, error) {
			return c.Kube().CertificatesV1().CertificateSigningRequests().Watch(context.Background(), options)
		}
	case gvr.ConfigMap:
		l = func(options metav1.ListOptions) (runtime.Object, error) {
			return c.Kube().CoreV1().ConfigMaps(opts.Namespace).List(context.Background(), options)
		}
		w = func(options metav1.ListOptions) (watch.Interface, error) {
			return c.Kube().CoreV1().ConfigMaps(opts.Namespace).Watch(context.Background(), options)
		}
	case gvr.CustomResourceDefinition:
		l = func(options metav1.ListOptions) (runtime.Object, error) {
			return c.Ext().ApiextensionsV1().CustomResourceDefinitions().List(context.Background(), options)
		}
		w = func(options metav1.ListOptions) (watch.Interface, error) {
			return c.Ext().ApiextensionsV1().CustomResourceDefinitions().Watch(context.Background(), options)
		}
	case gvr.Deployment:
		l = func(options metav1.ListOptions) (runtime.Object, error) {
			return c.Kube().AppsV1().Deployments(opts.Namespace).List(context.Background(), options)
		}
		w = func(options metav1.ListOptions) (watch.Interface, error) {
			return c.Kube().AppsV1().Deployments(opts.Namespace).Watch(context.Background(), options)
		}
	case gvr.DestinationRule:
		l = func(options metav1.ListOptions) (runtime.Object, error) {
			return c.Istio().NetworkingV1alpha3().DestinationRules(opts.Namespace).List(context.Background(), options)
		}
		w = func(options metav1.ListOptions) (watch.Interface, error) {
			return c.Istio().NetworkingV1alpha3().DestinationRules(opts.Namespace).Watch(context.Background(), options)
		}
	case gvr.EndpointSlice:
		l = func(options metav1.ListOptions) (runtime.Object, error) {
			return c.Kube().DiscoveryV1().EndpointSlices(opts.Namespace).List(context.Background(), options)
		}
		w = func(options metav1.ListOptions) (watch.Interface, error) {
			return c.Kube().DiscoveryV1().EndpointSlices(opts.Namespace).Watch(context.Background(), options)
		}
	case gvr.Endpoints:
		l = func(options metav1.ListOptions) (runtime.Object, error) {
			return c.Kube().CoreV1().Endpoints(opts.Namespace).List(context.Background(), options)
		}
		w = func(options metav1.ListOptions) (watch.Interface, error) {
			return c.Kube().CoreV1().Endpoints(opts.Namespace).Watch(context.Background(), options)
		}
	case gvr.EnvoyFilter:
		l = func(options metav1.ListOptions) (runtime.Object, error) {
			return c.Istio().NetworkingV1alpha3().EnvoyFilters(opts.Namespace).List(context.Background(), options)
		}
		w = func(options metav1.ListOptions) (watch.Interface, error) {
			return c.Istio().NetworkingV1alpha3().EnvoyFilters(opts.Namespace).Watch(context.Background(), options)
		}
	case gvr.GRPCRoute:
		l = func(options metav1.ListOptions) (runtime.Object, error) {
			return c.GatewayAPI().GatewayV1alpha2().GRPCRoutes(opts.Namespace).List(context.Background(), options)
		}
		w = func(options metav1.ListOptions) (watch.Interface, error) {
			return c.GatewayAPI().GatewayV1alpha2().GRPCRoutes(opts.Namespace).Watch(context.Background(), options)
		}
	case gvr.Gateway:
		l = func(options metav1.ListOptions) (runtime.Object, error) {
			return c.Istio().NetworkingV1alpha3().Gateways(opts.Namespace).List(context.Background(), options)
		}
		w = func(options metav1.ListOptions) (watch.Interface, error) {
			return c.Istio().NetworkingV1alpha3().Gateways(opts.Namespace).Watch(context.Background(), options)
		}
	case gvr.GatewayClass:
		l = func(options metav1.ListOptions) (runtime.Object, error) {
			return c.GatewayAPI().GatewayV1beta1().GatewayClasses().List(context.Background(), options)
		}
		w = func(options metav1.ListOptions) (watch.Interface, error) {
			return c.GatewayAPI().GatewayV1beta1().GatewayClasses().Watch(context.Background(), options)
		}
	case gvr.HTTPRoute:
		l = func(options metav1.ListOptions) (runtime.Object, error) {
			return c.GatewayAPI().GatewayV1beta1().HTTPRoutes(opts.Namespace).List(context.Background(), options)
		}
		w = func(options metav1.ListOptions) (watch.Interface, error) {
			return c.GatewayAPI().GatewayV1beta1().HTTPRoutes(opts.Namespace).Watch(context.Background(), options)
		}
	case gvr.Ingress:
		l = func(options metav1.ListOptions) (runtime.Object, error) {
			return c.Kube().NetworkingV1().Ingresses(opts.Namespace).List(context.Background(), options)
		}
		w = func(options metav1.ListOptions) (watch.Interface, error) {
			return c.Kube().NetworkingV1().Ingresses(opts.Namespace).Watch(context.Background(), options)
		}
	case gvr.IngressClass:
		l = func(options metav1.ListOptions) (runtime.Object, error) {
			return c.Kube().NetworkingV1().IngressClasses().List(context.Background(), options)
		}
		w = func(options metav1.ListOptions) (watch.Interface, error) {
			return c.Kube().NetworkingV1().IngressClasses().Watch(context.Background(), options)
		}
	case gvr.KubernetesGateway:
		l = func(options metav1.ListOptions) (runtime.Object, error) {
			return c.GatewayAPI().GatewayV1beta1().Gateways(opts.Namespace).List(context.Background(), options)
		}
		w = func(options metav1.ListOptions) (watch.Interface, error) {
			return c.GatewayAPI().GatewayV1beta1().Gateways(opts.Namespace).Watch(context.Background(), options)
		}
	case gvr.MutatingWebhookConfiguration:
		l = func(options metav1.ListOptions) (runtime.Object, error) {
			return c.Kube().AdmissionregistrationV1().MutatingWebhookConfigurations().List(context.Background(), options)
		}
		w = func(options metav1.ListOptions) (watch.Interface, error) {
			return c.Kube().AdmissionregistrationV1().MutatingWebhookConfigurations().Watch(context.Background(), options)
		}
	case gvr.Namespace:
		l = func(options metav1.ListOptions) (runtime.Object, error) {
			return c.Kube().CoreV1().Namespaces().List(context.Background(), options)
		}
		w = func(options metav1.ListOptions) (watch.Interface, error) {
			return c.Kube().CoreV1().Namespaces().Watch(context.Background(), options)
		}
	case gvr.Node:
		l = func(options metav1.ListOptions) (runtime.Object, error) {
			return c.Kube().CoreV1().Nodes().List(context.Background(), options)
		}
		w = func(options metav1.ListOptions) (watch.Interface, error) {
			return c.Kube().CoreV1().Nodes().Watch(context.Background(), options)
		}
	case gvr.PeerAuthentication:
		l = func(options metav1.ListOptions) (runtime.Object, error) {
			return c.Istio().SecurityV1beta1().PeerAuthentications(opts.Namespace).List(context.Background(), options)
		}
		w = func(options metav1.ListOptions) (watch.Interface, error) {
			return c.Istio().SecurityV1beta1().PeerAuthentications(opts.Namespace).Watch(context.Background(), options)
		}
	case gvr.Pod:
		l = func(options metav1.ListOptions) (runtime.Object, error) {
			return c.Kube().CoreV1().Pods(opts.Namespace).List(context.Background(), options)
		}
		w = func(options metav1.ListOptions) (watch.Interface, error) {
			return c.Kube().CoreV1().Pods(opts.Namespace).Watch(context.Background(), options)
		}
	case gvr.ProxyConfig:
		l = func(options metav1.ListOptions) (runtime.Object, error) {
			return c.Istio().NetworkingV1beta1().ProxyConfigs(opts.Namespace).List(context.Background(), options)
		}
		w = func(options metav1.ListOptions) (watch.Interface, error) {
			return c.Istio().NetworkingV1beta1().ProxyConfigs(opts.Namespace).Watch(context.Background(), options)
		}
	case gvr.ReferenceGrant:
		l = func(options metav1.ListOptions) (runtime.Object, error) {
			return c.GatewayAPI().GatewayV1beta1().ReferenceGrants(opts.Namespace).List(context.Background(), options)
		}
		w = func(options metav1.ListOptions) (watch.Interface, error) {
			return c.GatewayAPI().GatewayV1beta1().ReferenceGrants(opts.Namespace).Watch(context.Background(), options)
		}
	case gvr.RequestAuthentication:
		l = func(options metav1.ListOptions) (runtime.Object, error) {
			return c.Istio().SecurityV1beta1().RequestAuthentications(opts.Namespace).List(context.Background(), options)
		}
		w = func(options metav1.ListOptions) (watch.Interface, error) {
			return c.Istio().SecurityV1beta1().RequestAuthentications(opts.Namespace).Watch(context.Background(), options)
		}
	case gvr.Secret:
		l = func(options metav1.ListOptions) (runtime.Object, error) {
			return c.Kube().CoreV1().Secrets(opts.Namespace).List(context.Background(), options)
		}
		w = func(options metav1.ListOptions) (watch.Interface, error) {
			return c.Kube().CoreV1().Secrets(opts.Namespace).Watch(context.Background(), options)
		}
	case gvr.Service:
		l = func(options metav1.ListOptions) (runtime.Object, error) {
			return c.Kube().CoreV1().Services(opts.Namespace).List(context.Background(), options)
		}
		w = func(options metav1.ListOptions) (watch.Interface, error) {
			return c.Kube().CoreV1().Services(opts.Namespace).Watch(context.Background(), options)
		}
	case gvr.ServiceAccount:
		l = func(options metav1.ListOptions) (runtime.Object, error) {
			return c.Kube().CoreV1().ServiceAccounts(opts.Namespace).List(context.Background(), options)
		}
		w = func(options metav1.ListOptions) (watch.Interface, error) {
			return c.Kube().CoreV1().ServiceAccounts(opts.Namespace).Watch(context.Background(), options)
		}
	case gvr.ServiceEntry:
		l = func(options metav1.ListOptions) (runtime.Object, error) {
			return c.Istio().NetworkingV1alpha3().ServiceEntries(opts.Namespace).List(context.Background(), options)
		}
		w = func(options metav1.ListOptions) (watch.Interface, error) {
			return c.Istio().NetworkingV1alpha3().ServiceEntries(opts.Namespace).Watch(context.Background(), options)
		}
	case gvr.Sidecar:
		l = func(options metav1.ListOptions) (runtime.Object, error) {
			return c.Istio().NetworkingV1alpha3().Sidecars(opts.Namespace).List(context.Background(), options)
		}
		w = func(options metav1.ListOptions) (watch.Interface, error) {
			return c.Istio().NetworkingV1alpha3().Sidecars(opts.Namespace).Watch(context.Background(), options)
		}
	case gvr.TCPRoute:
		l = func(options metav1.ListOptions) (runtime.Object, error) {
			return c.GatewayAPI().GatewayV1alpha2().TCPRoutes(opts.Namespace).List(context.Background(), options)
		}
		w = func(options metav1.ListOptions) (watch.Interface, error) {
			return c.GatewayAPI().GatewayV1alpha2().TCPRoutes(opts.Namespace).Watch(context.Background(), options)
		}
	case gvr.TLSRoute:
		l = func(options metav1.ListOptions) (runtime.Object, error) {
			return c.GatewayAPI().GatewayV1alpha2().TLSRoutes(opts.Namespace).List(context.Background(), options)
		}
		w = func(options metav1.ListOptions) (watch.Interface, error) {
			return c.GatewayAPI().GatewayV1alpha2().TLSRoutes(opts.Namespace).Watch(context.Background(), options)
		}
	case gvr.Telemetry:
		l = func(options metav1.ListOptions) (runtime.Object, error) {
			return c.Istio().TelemetryV1alpha1().Telemetries(opts.Namespace).List(context.Background(), options)
		}
		w = func(options metav1.ListOptions) (watch.Interface, error) {
			return c.Istio().TelemetryV1alpha1().Telemetries(opts.Namespace).Watch(context.Background(), options)
		}
	case gvr.UDPRoute:
		l = func(options metav1.ListOptions) (runtime.Object, error) {
			return c.GatewayAPI().GatewayV1alpha2().UDPRoutes(opts.Namespace).List(context.Background(), options)
		}
		w = func(options metav1.ListOptions) (watch.Interface, error) {
			return c.GatewayAPI().GatewayV1alpha2().UDPRoutes(opts.Namespace).Watch(context.Background(), options)
		}
	case gvr.ValidatingWebhookConfiguration:
		l = func(options metav1.ListOptions) (runtime.Object, error) {
			return c.Kube().AdmissionregistrationV1().ValidatingWebhookConfigurations().List(context.Background(), options)
		}
		w = func(options metav1.ListOptions) (watch.Interface, error) {
			return c.Kube().AdmissionregistrationV1().ValidatingWebhookConfigurations().Watch(context.Background(), options)
		}
	case gvr.VirtualService:
		l = func(options metav1.ListOptions) (runtime.Object, error) {
			return c.Istio().NetworkingV1alpha3().VirtualServices(opts.Namespace).List(context.Background(), options)
		}
		w = func(options metav1.ListOptions) (watch.Interface, error) {
			return c.Istio().NetworkingV1alpha3().VirtualServices(opts.Namespace).Watch(context.Background(), options)
		}
	case gvr.WasmPlugin:
		l = func(options metav1.ListOptions) (runtime.Object, error) {
			return c.Istio().ExtensionsV1alpha1().WasmPlugins(opts.Namespace).List(context.Background(), options)
		}
		w = func(options metav1.ListOptions) (watch.Interface, error) {
			return c.Istio().ExtensionsV1alpha1().WasmPlugins(opts.Namespace).Watch(context.Background(), options)
		}
	case gvr.WorkloadEntry:
		l = func(options metav1.ListOptions) (runtime.Object, error) {
			return c.Istio().NetworkingV1alpha3().WorkloadEntries(opts.Namespace).List(context.Background(), options)
		}
		w = func(options metav1.ListOptions) (watch.Interface, error) {
			return c.Istio().NetworkingV1alpha3().WorkloadEntries(opts.Namespace).Watch(context.Background(), options)
		}
	case gvr.WorkloadGroup:
		l = func(options metav1.ListOptions) (runtime.Object, error) {
			return c.Istio().NetworkingV1alpha3().WorkloadGroups(opts.Namespace).List(context.Background(), options)
		}
		w = func(options metav1.ListOptions) (watch.Interface, error) {
			return c.Istio().NetworkingV1alpha3().WorkloadGroups(opts.Namespace).Watch(context.Background(), options)
		}
	default:
		panic(fmt.Sprintf("Unknown type %v", g))
	}
	return c.Informers().InformerFor(g, opts, func() cache.SharedIndexInformer {
		inf := cache.NewSharedIndexInformer(
			&cache.ListWatch{
				ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
					options.FieldSelector = opts.FieldSelector
					options.LabelSelector = opts.LabelSelector
					return l(options)
				},
				WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
					options.FieldSelector = opts.FieldSelector
					options.LabelSelector = opts.LabelSelector
					return w(options)
				},
			},
			gvrToObject(g),
			0,
			cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc},
		)
		setupInformer(opts, inf)
		return inf
	})
}
