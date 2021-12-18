package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"gopkg.in/yaml.v2"
	admissionregistrationv1 "k8s.io/api/admissionregistration/v1"
	appsv1 "k8s.io/api/apps/v1"
	scalev2beta2 "k8s.io/api/autoscaling/v2beta2"
	corev1 "k8s.io/api/core/v1"
	policyv1beta1 "k8s.io/api/policy/v1beta1"
	rbacv1 "k8s.io/api/rbac/v1"
	apiextv1beta1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/serializer"
	"k8s.io/cli-runtime/pkg/printers"
	"k8s.io/client-go/kubernetes/scheme"
	cachingv1alpha1 "knative.dev/caching/pkg/apis/caching/v1alpha1"
)

const (
	relaseName = "knative-serving-{{ .Release.Name }}"
	helmLabels = "AQ-{{- include \"knative.labels\" . | nindent 4 }}"
)

var (
	baseCRD     = "https://github.com/knative/serving/releases/download/knative-%s/serving-crds.yaml"
	baseCtrl    = "https://github.com/knative/serving/releases/download/knative-%s/serving-core.yaml"
	baseKourier = "https://github.com/knative/net-kourier/releases/download/knative-%s/kourier.yaml"
)

func main() {

	if len(os.Args) < 3 {
		log.Fatalf("knative version and component required, e.g. v1.1.0 crds")
	}

	version := os.Args[1]

	log.Printf("building knative helm for %s\n", version)

	switch os.Args[2] {
	case "kourier":
		prepareKnativeKourier(version)
	case "crds":
		prepareKnativeCRDS(version)
	default:
		prepareKnativeServing(version)
	}

}

func createDecoder() runtime.Decoder {
	sch := runtime.NewScheme()
	_ = scheme.AddToScheme(sch)
	_ = apiextv1beta1.AddToScheme(sch)
	_ = cachingv1alpha1.AddToScheme(sch)
	return serializer.NewCodecFactory(sch).UniversalDeserializer()
}

func downloadYAML(url string) string {

	resp, err := http.Get(url)
	if err != nil {
		log.Fatalf("can not download yaml: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Fatalf("bad status: %s", resp.Status)
	}

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("can not read body: %s", err)
	}

	return string(b)

}

func addHelmLabels(meta metav1.ObjectMeta, ns bool) {
	meta.Labels["app.kubernetes.io/name"] = relaseName
	meta.Labels["LABELREMOVE"] = helmLabels
	if ns {
		meta.Namespace = relaseName
	}
}

func prepareKnativeKourier(version string) {
	yamlData := downloadYAML(fmt.Sprintf(baseKourier, version))

	decoder := createDecoder()
	y := printers.YAMLPrinter{}

	var buf bytes.Buffer

	for _, resourceYAML := range strings.Split(yamlData, "\n---\n") {
		// skip empty
		if len(resourceYAML) <= 1 {
			continue
		}

		obj, gvk, err := decoder.Decode([]byte(resourceYAML), nil, nil)
		if err != nil {
			log.Print(err)
			continue
		}

		print := true

		switch gvk.Kind {
		case "Namespace":
			// remove
			print = false
		case "ConfigMap":
			addHelmLabels(obj.(*corev1.ConfigMap).ObjectMeta, true)
		case "ServiceAccount":
			addHelmLabels(obj.(*corev1.ServiceAccount).ObjectMeta, true)
		case "ClusterRole":
			addHelmLabels(obj.(*rbacv1.ClusterRole).ObjectMeta, false)
		case "ClusterRoleBinding":
			addHelmLabels(obj.(*rbacv1.ClusterRoleBinding).ObjectMeta, false)
		case "Deployment":
			addHelmLabels(obj.(*appsv1.Deployment).ObjectMeta, true)

			depl := obj.(*appsv1.Deployment)
			for i := range depl.Spec.Template.Spec.Containers[0].Env {
				if depl.Spec.Template.Spec.Containers[0].Env[i].Name == "KOURIER_GATEWAY_NAMESPACE" {
					depl.Spec.Template.Spec.Containers[0].Env[i].Value = relaseName
				}
			}
		case "Service":
			addHelmLabels(obj.(*corev1.Service).ObjectMeta, true)
		default:
			log.Fatalf("unknown kind: %v", gvk.Kind)
		}

		if print {
			y.PrintObj(obj, &buf)
		}

	}

	s := buf.String()
	s = strings.Replace(s, "address: \"net-kourier-controller.knative-serving\"",
		fmt.Sprintf("address: \"net-kourier-controller.%s\"", relaseName), 1)
	s = strings.ReplaceAll(s, "LABELREMOVE: ", "")
	s = strings.ReplaceAll(s, "AQ-", "")

	err := os.WriteFile("/tmp/templates/kourier.yaml", []byte(s), 0644)
	if err != nil {
		log.Fatalf("can not write kourier: %v", err)
	}

}

func prepareKnativeCRDS(version string) {
	yamlData := downloadYAML(fmt.Sprintf(baseCRD, version))
	err := os.WriteFile("/tmp/crds/serving-crds.yaml", []byte(yamlData), 0644)
	if err != nil {
		log.Fatalf("can not write crds: %v", err)
	}
}

func prepareKnativeServing(version string) {

	yamlData := downloadYAML(fmt.Sprintf(baseCtrl, version))
	decoder := createDecoder()
	y := printers.YAMLPrinter{}

	var buf bytes.Buffer

	for _, resourceYAML := range strings.Split(yamlData, "\n---\n") {
		// skip empty
		if len(resourceYAML) <= 1 {
			continue
		}

		obj, gvk, err := decoder.Decode([]byte(resourceYAML), nil, nil)
		if err != nil {
			log.Print(err)
			continue
		}

		print := true

		switch gvk.Kind {
		case "Namespace":
			// remove
			print = false
		case "HorizontalPodAutoscaler":
			addHelmLabels(obj.(*scalev2beta2.HorizontalPodAutoscaler).ObjectMeta, true)
		case "Service":
			addHelmLabels(obj.(*corev1.Service).ObjectMeta, true)
		case "Deployment":
			addHelmLabels(obj.(*appsv1.Deployment).ObjectMeta, true)
			depl := obj.(*appsv1.Deployment)

			var j int32 = 11223344
			if depl.ObjectMeta.Name != "activator" {
				depl.Spec.Replicas = &j
			}

			if depl.ObjectMeta.Name == "controller" {

				e := depl.Spec.Template.Spec.Containers[0].Env
				e = append(e, corev1.EnvVar{
					Name:  "HTTPS_PROXY",
					Value: "AQ-{{ .Values.https_proxy }}",
				})
				e = append(e, corev1.EnvVar{
					Name:  "HTTP_PROXY",
					Value: "AQ-{{ .Values.http_proxy }}",
				})
				e = append(e, corev1.EnvVar{
					Name:  "NO_PROXY",
					Value: "AQ-{{ .Values.no_proxy }}",
				})
				depl.Spec.Template.Spec.Containers[0].Env = e
			}
		case "CustomResourceDefinition":
			addHelmLabels(obj.(*apiextv1beta1.CustomResourceDefinition).ObjectMeta, false)
		case "ClusterRoleBinding":
			addHelmLabels(obj.(*rbacv1.ClusterRoleBinding).ObjectMeta, false)
		case "ClusterRole":
			addHelmLabels(obj.(*rbacv1.ClusterRole).ObjectMeta, false)
		case "ValidatingWebhookConfiguration":
			addHelmLabels(obj.(*admissionregistrationv1.ValidatingWebhookConfiguration).ObjectMeta, false)
		case "MutatingWebhookConfiguration":
			addHelmLabels(obj.(*admissionregistrationv1.MutatingWebhookConfiguration).ObjectMeta, false)
		case "ServiceAccount":
			addHelmLabels(obj.(*corev1.ServiceAccount).ObjectMeta, true)
		case "PodDisruptionBudget":
			addHelmLabels(obj.(*policyv1beta1.PodDisruptionBudget).ObjectMeta, true)
		case "Secret":
			addHelmLabels(obj.(*corev1.Secret).ObjectMeta, true)
		case "Image":
			addHelmLabels(obj.(*cachingv1alpha1.Image).ObjectMeta, true)
		case "ConfigMap":
			obj = updateConfigMaps(obj)
		default:
			log.Fatalf("unknown kind: %v", gvk.Kind)
		}

		if print {
			y.PrintObj(obj, &buf)
		}
	}

	s := buf.String()
	s = strings.ReplaceAll(s, "LABELREMOVE: ", "")
	s = strings.ReplaceAll(s, "AQ-", "")
	s = strings.ReplaceAll(s, "11223344", "{{ .Values.replicas }}")

	err := os.WriteFile("/tmp/templates/serving-core.yaml", []byte(s), 0644)
	if err != nil {
		log.Fatalf("can not write core: %v", err)
	}

}

func updateConfigMaps(obj runtime.Object) *corev1.ConfigMap {

	cm := obj.(*corev1.ConfigMap)
	cm.ObjectMeta.Namespace = relaseName
	cm.ObjectMeta.Labels["app.kubernetes.io/name"] = relaseName
	cm.ObjectMeta.Labels["LABELREMOVE"] = helmLabels

	var data map[string]string
	err := yaml.Unmarshal([]byte(cm.Data["_example"]), &data)
	if err != nil {
		log.Fatalf("can not unmarshall default cm: %v", err)
	}

	if cm.ObjectMeta.Name == "config-defaults" {
		data["revision-timeout-seconds"] = "AQ-\"{{ .Values.defaults.timeout_seconds }}\""
		data["max-revision-timeout-seconds"] = "AQ-\"{{ .Values.defaults.max_timeout_seconds }}\""
		cm.Data = data
	}

	if cm.ObjectMeta.Name == "config-autoscaler" {
		data["scale-to-zero-grace-period"] = "AQ-\"{{ .Values.autoscaler.grace_period }}\""
		data["scale-to-zero-pod-retention-period"] = "AQ-\"{{ .Values.autoscaler.retention_period }}\""
		data["max-scale-limit"] = "AQ-\"{{ .Values.autoscaler.max-scale_limit }}\""
		cm.Data = data
	}

	if cm.ObjectMeta.Name == "config-deployment" {
		data["registries-skipping-tag-resolving"] = "AQ-{{ .Values.deployment.skip_tag }}"
		cm.Data = data
	}

	if cm.ObjectMeta.Name == "config-features" {
		data["kubernetes.podspec-runtimeclassname"] = "enabled"
		data["kubernetes.podspec-securitycontext"] = "enabled"
		data["tag-header-based-routing"] = "enabled"
		data["kubernetes.podspec-volumes-emptydir"] = "enabled"
		data["kubernetes.podspec-init-containers"] = "enabled"
		cm.Data = data
	}

	if cm.ObjectMeta.Name == "config-network" {
		data["ingress-class"] = "kourier.ingress.networking.knative.dev"
		data["rollout-duration"] = "60"
	}

	return cm
}
