package manifests

import (
	"fmt"
	"io"
	"os"
	"path/filepath"

	v1Hive "github.com/openshift/hive/apis/hive/v1"
	v1HiveAgent "github.com/openshift/hive/apis/hive/v1/agent"

	//v1Assisted "github.com/openshift/assisted-service/api/v1beta" Assisted-service needs to fix its go.mod

	k8scorev1 "k8s.io/api/core/v1"
	k8smetav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	jsonSerializer "k8s.io/apimachinery/pkg/runtime/serializer/json"
)

func sampleClusterDeployment() *v1Hive.ClusterDeployment {
	return &v1Hive.ClusterDeployment{
		TypeMeta: k8smetav1.TypeMeta{
			Kind: "ClusterDeployment",
		},
		ObjectMeta: k8smetav1.ObjectMeta{
			Namespace: "mycluster",
			Name:      "first",
		},
		Spec: v1Hive.ClusterDeploymentSpec{
			ClusterName: "mycluster",
			BaseDomain:  "example.com",
			Platform: v1Hive.Platform{
				AgentBareMetal: &v1HiveAgent.BareMetalPlatform{
					AgentSelector: k8smetav1.LabelSelector{},
				},
			},
			PullSecretRef:      &k8scorev1.LocalObjectReference{},
			Ingress:            []v1Hive.ClusterIngress{},
			CertificateBundles: []v1Hive.CertificateBundleSpec{},
			ClusterMetadata:    &v1Hive.ClusterMetadata{},
			ControlPlaneConfig: v1Hive.ControlPlaneConfigSpec{
				ServingCertificates: v1Hive.ControlPlaneServingCertificateSpec{
					Additional: []v1Hive.ControlPlaneAdditionalCertificate{},
					Default:    "",
				},
				APIURLOverride: "",
			},
			Installed:    false,
			Provisioning: &v1Hive.Provisioning{},
			ClusterInstallRef: &v1Hive.ClusterInstallLocalReference{
				Group:   "extensions.hive.openshift.io",
				Version: "v1beta1",
				Kind:    "AgentClusterInstall",
				Name:    "mycluster-install",
			},
		},
	}
}

func sampleClusterImageSet() *v1Hive.ClusterImageSet {
	return &v1Hive.ClusterImageSet{
		TypeMeta: k8smetav1.TypeMeta{
			Kind: "ClusterImageSet",
		},
		ObjectMeta: k8smetav1.ObjectMeta{
			Namespace: "mycluster",
			Name:      "openshift-4.11.0",
		},
		Spec: v1Hive.ClusterImageSetSpec{
			ReleaseImage: "quay.io/openshift-release-dev/ocp-release:4.11.0-x86_64",
		},
	}
}

// func sampleInfraEnv() *v1Assisted.InfraEnv {
// 	return &v1Assisted.InfraEnv{}
// }

func GenerateManifests(path string) error {
	var manifests = make(map[string]runtime.Object)
	manifests["cluster-deployment.yaml"] = sampleClusterDeployment()
	manifests["cluster-image-set.yaml"] = sampleClusterImageSet()
	serializer := jsonSerializer.NewSerializerWithOptions(
		jsonSerializer.DefaultMetaFactory, nil, nil,
		jsonSerializer.SerializerOptions{
			Yaml:   true,
			Pretty: true,
			Strict: true,
		},
	)

	for name, object := range manifests {
		var writer io.Writer
		if path == "" {
			fmt.Printf("--- #%s\n", name)
			writer = os.Stdout
		} else {
			destFile := filepath.Join(path, name)
			manifestFile, err := os.OpenFile(destFile, os.O_CREATE|os.O_WRONLY|os.O_EXCL, 0644)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Failed to create manifest file '%s'. Error: %s\n", destFile, err)
				return nil
			}
			defer manifestFile.Close()
			writer = manifestFile
		}
		serializer.Encode(object, writer)
	}
	return nil
}
