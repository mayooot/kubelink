package FluxApplicationService

import (
	"k8s.io/apimachinery/pkg/runtime/schema"
)

const (
	FluxKustomizationGroup     = "kustomize.toolkit.fluxcd.io"
	FluxAppKustomizationKind   = "Kustomization"
	FluxKustomizationVersionV1 = "v1"
	AllNamespaces              = ""
	FluxHelmReleaseGroup       = "helm.toolkit.fluxcd.io"
	FluxAppHelmReleaseKind     = "HelmRelease"
	FluxHelmReleaseVersionV2   = "v2"
	FluxLabel                  = "labels"
	KustomizeNameLabel         = "kustomize.toolkit.fluxcd.io/name"
	KustomizeNamespaceLabel    = "kustomize.toolkit.fluxcd.io/namespace"
)

var GvkForKustomizationFluxApp = schema.GroupVersionKind{
	Group:   FluxKustomizationGroup,
	Kind:    FluxAppKustomizationKind,
	Version: FluxKustomizationVersionV1,
}

var GvkForHelmreleaseFluxApp = schema.GroupVersionKind{
	Group:   FluxHelmReleaseGroup,
	Kind:    FluxAppHelmReleaseKind,
	Version: FluxHelmReleaseVersionV2,
}

type FluxApplicationDto struct {
	Name                  string             `json:"appName"`
	HealthStatus          string             `json:"appStatus"`
	SyncStatus            string             `json:"syncStatus"`
	EnvironmentDetails    *EnvironmentDetail `json:"environmentDetail"`
	FluxAppDeploymentType FluxAppType        `json:"fluxAppDeploymentType"`
}
type EnvironmentDetail struct {
	ClusterId   int    `json:"clusterId"`
	ClusterName string `json:"clusterName"`
	Namespace   string `json:"namespace"`
}

type FluxAppType string

const (
	NameKey       = "Name"
	StatusKey     = "Status"
	ReadyKey      = "Ready"
	ColumnNameKey = "name"
)
