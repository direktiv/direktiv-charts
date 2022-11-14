package main

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	appsv1 "k8s.io/api/apps/v1"
	batchv1 "k8s.io/api/batch/v1"
	corev1 "k8s.io/api/core/v1"
	rbacv1 "k8s.io/api/rbac/v1"
	apiextv1beta1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/serializer"
	"k8s.io/cli-runtime/pkg/printers"
	"k8s.io/client-go/kubernetes/scheme"
)

var (
	dl           = "https://github.com/knative/net-contour/releases/download/knative-v%s/contour.yaml"
	targetFolder = ""
)

const (
	relaseName      = "{{ .Release.Name }}"
	relaseNameSpace = "{{ .Release.Namespace }}"
	helmLabels      = "AQ-{{- include \"knative-instance.labels\" . | nindent 4 }}"
)

func main() {

	if len(os.Args) != 3 {
		log.Fatalf("contour version, e.g. 1.8.0 and target folder required")
	}

	dl = fmt.Sprintf(dl, os.Args[1])
	fmt.Printf("downloading %s\n", dl)

	targetFolder = os.Args[2]
	fmt.Printf("storing in %s\n", targetFolder)

	yaml := downloadYAML(dl)

	decoder := createDecoder()

	y := printers.YAMLPrinter{}
	var buf bytes.Buffer

	for _, resourceYAML := range strings.Split(yaml, "\n---\n") {

		// skip empty
		if len(resourceYAML) <= 1 {
			continue
		}

		obj, gvk, err := decoder.Decode([]byte(resourceYAML), nil, nil)
		if err != nil {
			log.Print(err)
			continue
		}

		switch gvk.Kind {
		case "Namespace":
			{
				md := &obj.(*corev1.Namespace).ObjectMeta
				if md.Name == "contour-external" {
					continue
				}
				addHelmLabels(&obj.(*corev1.Namespace).ObjectMeta, false)
				// we always set it to true. if installed its ok, otherwise no effect
				obj.(*corev1.Namespace).ObjectMeta.Annotations = make(map[string]string)
				obj.(*corev1.Namespace).ObjectMeta.Annotations["linkerd.io/inject"] = "disabled"
			}
		case "CustomResourceDefinition":
			addHelmLabels(&obj.(*apiextv1beta1.CustomResourceDefinition).ObjectMeta, false)
		case "ServiceAccount":
			md := &obj.(*corev1.ServiceAccount).ObjectMeta
			if md.Namespace == "contour-external" {
				continue
			}
			addHelmLabels(&obj.(*corev1.ServiceAccount).ObjectMeta, false)
		case "Role":
			md := &obj.(*rbacv1.Role).ObjectMeta
			if md.Namespace == "contour-external" {
				continue
			}
			addHelmLabels(&obj.(*rbacv1.Role).ObjectMeta, false)
		case "RoleBinding":
			md := &obj.(*rbacv1.RoleBinding).ObjectMeta
			if md.Namespace == "contour-external" {
				continue
			}
			addHelmLabels(&obj.(*rbacv1.RoleBinding).ObjectMeta, false)
		case "ClusterRoleBinding":
			md := &obj.(*rbacv1.ClusterRoleBinding).ObjectMeta
			if md.Namespace == "contour-external" {
				continue
			}
			if md.Name == "knative-contour-external" {
				continue
			}
			addHelmLabels(&obj.(*rbacv1.ClusterRoleBinding).ObjectMeta, false)
		case "ClusterRole":
			md := &obj.(*rbacv1.ClusterRole).ObjectMeta
			if md.Namespace == "contour-external" {
				continue
			}
			addHelmLabels(&obj.(*rbacv1.ClusterRole).ObjectMeta, false)
		case "Deployment":
			md := &obj.(*appsv1.Deployment).ObjectMeta
			md.Annotations = make(map[string]string)
			md.Annotations["linkerd.io/inject"] = "enabled"
			if md.Namespace == "contour-external" {
				continue
			}
			addHelmLabels(&obj.(*appsv1.Deployment).ObjectMeta, false)
		case "ConfigMap":
			md := &obj.(*corev1.ConfigMap).ObjectMeta
			if md.Namespace == "contour-external" {
				continue
			}
			addHelmLabels(&obj.(*corev1.ConfigMap).ObjectMeta, false)
		case "Service":
			md := &obj.(*corev1.Service).ObjectMeta
			if md.Namespace == "contour-external" {
				continue
			}
			addHelmLabels(&obj.(*corev1.Service).ObjectMeta, false)
		case "Job":
			md := &obj.(*batchv1.Job).ObjectMeta
			if md.Namespace == "contour-external" {
				continue
			}
			addHelmLabels(&obj.(*batchv1.Job).ObjectMeta, false)
			obj.(*batchv1.Job).Spec.Template.ObjectMeta.Annotations = make(map[string]string)
			obj.(*batchv1.Job).Spec.Template.ObjectMeta.Annotations["linkerd.io/inject"] = "disabled"
		case "DaemonSet":
			md := &obj.(*appsv1.DaemonSet).ObjectMeta
			md.Annotations = make(map[string]string)
			md.Annotations["linkerd.io/inject"] = "enabled"
			if md.Namespace == "contour-external" {
				continue
			}
			addHelmLabels(&obj.(*appsv1.DaemonSet).ObjectMeta, false)
		default:
			{
				log.Fatalf("Kind %v\n", gvk.Kind)
			}
		}

		y.PrintObj(obj, &buf)

	}

	s := buf.String()
	s = strings.ReplaceAll(s, "LABELREMOVE: ", "")
	s = strings.ReplaceAll(s, "AQ-", "")
	buf.Reset()

	err := os.WriteFile(filepath.Join(targetFolder, "contour.yaml"), []byte(s), 0644)
	if err != nil {
		log.Fatalf("can not save changed yaml: %v", err)
	}

}

func addHelmLabels(meta *metav1.ObjectMeta, ns bool) {
	meta.Labels["app.kubernetes.io/name"] = relaseName
	meta.Labels["LABELREMOVE"] = helmLabels
	if ns {
		meta.Namespace = relaseNameSpace
	}
}

func downloadYAML(url string) string {

	fileName := fmt.Sprintf("%s.contour.yaml", md5enc(url))

	if _, err := os.Stat(fileName); errors.Is(err, os.ErrNotExist) {

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

		err = os.WriteFile(fileName, b, 0644)
		if err != nil {
			log.Fatalf("can not store yaml: %s", err)
		}

		return string(b)
	}

	dat, err := os.ReadFile(fileName)
	if err != nil {
		log.Fatalf("can not read yaml: %s", err)
	}

	return string(dat)
}

func md5enc(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}

func createDecoder() runtime.Decoder {
	sch := runtime.NewScheme()
	_ = scheme.AddToScheme(sch)
	_ = apiextv1beta1.AddToScheme(sch)
	return serializer.NewCodecFactory(sch).UniversalDeserializer()
}
