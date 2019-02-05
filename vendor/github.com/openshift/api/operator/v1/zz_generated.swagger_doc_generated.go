package v1

// This file contains a collection of methods that can be used from go-restful to
// generate Swagger API documentation for its models. Please read this PR for more
// information on the implementation: https://github.com/emicklei/go-restful/pull/215
//
// TODOs are ignored from the parser (e.g. TODO(andronat):... || TODO:...) if and only if
// they are on one line! For multiple line or blocks that you want to ignore use ---.
// Any context after a --- is ignored.
//
// Those methods can be generated by using hack/update-swagger-docs.sh

// AUTO-GENERATED FUNCTIONS START HERE
var map_GenerationStatus = map[string]string{
	"":               "GenerationStatus keeps track of the generation for a given resource so that decisions about forced updates can be made.",
	"group":          "group is the group of the thing you're tracking",
	"resource":       "resource is the resource type of the thing you're tracking",
	"namespace":      "namespace is where the thing you're tracking is",
	"name":           "name is the name of the thing you're tracking",
	"lastGeneration": "lastGeneration is the last generation of the workload controller involved",
	"hash":           "hash is an optional field set for resources without generation that are content sensitive like secrets and configmaps",
}

func (GenerationStatus) SwaggerDoc() map[string]string {
	return map_GenerationStatus
}

var map_MyOperatorResource = map[string]string{
	"": "MyOperatorResource is an example operator configuration type",
}

func (MyOperatorResource) SwaggerDoc() map[string]string {
	return map_MyOperatorResource
}

var map_NodeStatus = map[string]string{
	"":                         "NodeStatus provides information about the current state of a particular node managed by this operator.",
	"nodeName":                 "nodeName is the name of the node",
	"currentRevision":          "currentRevision is the generation of the most recently successful deployment",
	"targetRevision":           "targetRevision is the generation of the deployment we're trying to apply",
	"lastFailedRevision":       "lastFailedRevision is the generation of the deployment we tried and failed to deploy.",
	"lastFailedRevisionErrors": "lastFailedRevisionErrors is a list of the errors during the failed deployment referenced in lastFailedRevision",
}

func (NodeStatus) SwaggerDoc() map[string]string {
	return map_NodeStatus
}

var map_OperandContainerSpec = map[string]string{
	"name":      "name is the name of the container to modify",
	"resources": "resources are the requests and limits to place in the container.  Nil means to accept the defaults.",
}

func (OperandContainerSpec) SwaggerDoc() map[string]string {
	return map_OperandContainerSpec
}

var map_OperandSpec = map[string]string{
	"":                           "OperandSpec holds information for customization of a particular functional unit - logically maps to a workload",
	"name":                       "name is the name of this unit.  The operator must be aware of it.",
	"operandContainerSpecs":      "operandContainerSpecs are per-container options",
	"unsupportedResourcePatches": "unsupportedResourcePatches are applied to the workload resource for this unit. This is an unsupported workaround if anything needs to be modified on the workload that is not otherwise configurable.",
}

func (OperandSpec) SwaggerDoc() map[string]string {
	return map_OperandSpec
}

var map_OperatorCondition = map[string]string{
	"": "OperatorCondition is just the standard condition fields.",
}

func (OperatorCondition) SwaggerDoc() map[string]string {
	return map_OperatorCondition
}

var map_OperatorSpec = map[string]string{
	"":                           "OperatorSpec contains common fields operators need.  It is intended to be anonymous included inside of the Spec struct for your particular operator.",
	"managementState":            "managementState indicates whether and how the operator should manage the component",
	"logLevel":                   "logLevel is an intent based logging for an overall component.  It does not give fine grained control, but it is a simple way to manage coarse grained logging choices that operators have to interpret for their operands.",
	"operandSpecs":               "operandSpecs provide customization for functional units within the component",
	"unsupportedConfigOverrides": "unsupportedConfigOverrides holds a sparse config that will override any previously set options.  It only needs to be the fields to override it will end up overlaying in the following order: 1. hardcoded defaults 2. observedConfig 3. unsupportedConfigOverrides",
	"observedConfig":             "observedConfig holds a sparse config that controller has observed from the cluster state.  It exists in spec because it is an input to the level for the operator",
}

func (OperatorSpec) SwaggerDoc() map[string]string {
	return map_OperatorSpec
}

var map_OperatorStatus = map[string]string{
	"observedGeneration": "observedGeneration is the last generation change you've dealt with",
	"conditions":         "conditions is a list of conditions and their status",
	"version":            "version is the level this availability applies to",
	"readyReplicas":      "readyReplicas indicates how many replicas are ready and at the desired state",
	"generations":        "generations are used to determine when an item needs to be reconciled or has changed in a way that needs a reaction.",
}

func (OperatorStatus) SwaggerDoc() map[string]string {
	return map_OperatorStatus
}

var map_ResourcePatch = map[string]string{
	"":      "ResourcePatch is a way to represent the patch you would issue to `kubectl patch` in the API",
	"type":  "type is the type of patch to apply: jsonmerge, strategicmerge",
	"patch": "patch the patch itself",
}

func (ResourcePatch) SwaggerDoc() map[string]string {
	return map_ResourcePatch
}

var map_StaticPodOperatorSpec = map[string]string{
	"": "StaticPodOperatorSpec is spec for controllers that manage static pods.",
	"failedRevisionLimit":    "failedRevisionLimit is the number of failed static pod installer revisions to keep on disk and in the api -1 = unlimited, 0 or unset = 5 (default)",
	"succeededRevisionLimit": "succeededRevisionLimit is the number of successful static pod installer revisions to keep on disk and in the api -1 = unlimited, 0 or unset = 5 (default)",
}

func (StaticPodOperatorSpec) SwaggerDoc() map[string]string {
	return map_StaticPodOperatorSpec
}

var map_StaticPodOperatorStatus = map[string]string{
	"": "StaticPodOperatorStatus is status for controllers that manage static pods.  There are different needs because individual node status must be tracked.",
	"latestAvailableRevision": "latestAvailableRevision is the deploymentID of the most recent deployment",
	"nodeStatuses":            "nodeStatuses track the deployment values and errors across individual nodes",
}

func (StaticPodOperatorStatus) SwaggerDoc() map[string]string {
	return map_StaticPodOperatorStatus
}

var map_Authentication = map[string]string{
	"": "Authentication provides information to configure an operator to manage authentication.",
}

func (Authentication) SwaggerDoc() map[string]string {
	return map_Authentication
}

var map_AuthenticationList = map[string]string{
	"": "AuthenticationList is a collection of items",
}

func (AuthenticationList) SwaggerDoc() map[string]string {
	return map_AuthenticationList
}

var map_ConsoleCustomization = map[string]string{
	"brand":                "brand is the default branding of the web console which can be overridden by providing the brand field.  There is a limited set of specific brand options. This field controls elements of the console such as the logo. Invalid value will prevent a console rollout.",
	"documentationBaseURL": "documentationBaseURL links to external documentation are shown in various sections of the web console.  Providing documentationBaseURL will override the default documentation URL. Invalid value will prevent a console rollout.",
}

func (ConsoleCustomization) SwaggerDoc() map[string]string {
	return map_ConsoleCustomization
}

var map_ConsoleSpec = map[string]string{
	"customization": "customization is used to optionally provide a small set of customization options to the web console.",
}

func (ConsoleSpec) SwaggerDoc() map[string]string {
	return map_ConsoleSpec
}

var map_Etcd = map[string]string{
	"": "Etcd provides information to configure an operator to manage kube-apiserver.",
}

func (Etcd) SwaggerDoc() map[string]string {
	return map_Etcd
}

var map_EtcdList = map[string]string{
	"":         "KubeAPISOperatorConfigList is a collection of items",
	"metadata": "Standard object's metadata.",
	"items":    "Items contains the items",
}

func (EtcdList) SwaggerDoc() map[string]string {
	return map_EtcdList
}

var map_EtcdSpec = map[string]string{
	"forceRedeploymentReason": "forceRedeploymentReason can be used to force the redeployment of the kube-apiserver by providing a unique string. This provides a mechanism to kick a previously failed deployment and provide a reason why you think it will work this time instead of failing again on the same config.",
}

func (EtcdSpec) SwaggerDoc() map[string]string {
	return map_EtcdSpec
}

var map_KubeAPIServer = map[string]string{
	"": "KubeAPIServer provides information to configure an operator to manage kube-apiserver.",
}

func (KubeAPIServer) SwaggerDoc() map[string]string {
	return map_KubeAPIServer
}

var map_KubeAPIServerList = map[string]string{
	"":         "KubeAPIServerList is a collection of items",
	"metadata": "Standard object's metadata.",
	"items":    "Items contains the items",
}

func (KubeAPIServerList) SwaggerDoc() map[string]string {
	return map_KubeAPIServerList
}

var map_KubeAPIServerSpec = map[string]string{
	"forceRedeploymentReason": "forceRedeploymentReason can be used to force the redeployment of the kube-apiserver by providing a unique string. This provides a mechanism to kick a previously failed deployment and provide a reason why you think it will work this time instead of failing again on the same config.",
}

func (KubeAPIServerSpec) SwaggerDoc() map[string]string {
	return map_KubeAPIServerSpec
}

var map_KubeControllerManager = map[string]string{
	"": "KubeControllerManager provides information to configure an operator to manage kube-controller-manager.",
}

func (KubeControllerManager) SwaggerDoc() map[string]string {
	return map_KubeControllerManager
}

var map_KubeControllerManagerList = map[string]string{
	"":         "KubeControllerManagerList is a collection of items",
	"metadata": "Standard object's metadata.",
	"items":    "Items contains the items",
}

func (KubeControllerManagerList) SwaggerDoc() map[string]string {
	return map_KubeControllerManagerList
}

var map_KubeControllerManagerSpec = map[string]string{
	"forceRedeploymentReason": "forceRedeploymentReason can be used to force the redeployment of the kube-controller-manager by providing a unique string. This provides a mechanism to kick a previously failed deployment and provide a reason why you think it will work this time instead of failing again on the same config.",
}

func (KubeControllerManagerSpec) SwaggerDoc() map[string]string {
	return map_KubeControllerManagerSpec
}

var map_OpenShiftAPIServer = map[string]string{
	"": "OpenShiftAPIServer provides information to configure an operator to manage openshift-apiserver.",
}

func (OpenShiftAPIServer) SwaggerDoc() map[string]string {
	return map_OpenShiftAPIServer
}

var map_OpenShiftAPIServerList = map[string]string{
	"":         "OpenShiftAPIServerList is a collection of items",
	"metadata": "Standard object's metadata.",
	"items":    "Items contains the items",
}

func (OpenShiftAPIServerList) SwaggerDoc() map[string]string {
	return map_OpenShiftAPIServerList
}

var map_OpenShiftControllerManager = map[string]string{
	"": "OpenShiftControllerManager provides information to configure an operator to manage openshift-controller-manager.",
}

func (OpenShiftControllerManager) SwaggerDoc() map[string]string {
	return map_OpenShiftControllerManager
}

var map_OpenShiftControllerManagerList = map[string]string{
	"":         "OpenShiftControllerManagerList is a collection of items",
	"metadata": "Standard object's metadata.",
	"items":    "Items contains the items",
}

func (OpenShiftControllerManagerList) SwaggerDoc() map[string]string {
	return map_OpenShiftControllerManagerList
}

var map_KubeScheduler = map[string]string{
	"": "KubeScheduler provides information to configure an operator to manage scheduler.",
}

func (KubeScheduler) SwaggerDoc() map[string]string {
	return map_KubeScheduler
}

var map_KubeSchedulerList = map[string]string{
	"":         "KubeSchedulerList is a collection of items",
	"metadata": "Standard object's metadata.",
	"items":    "Items contains the items",
}

func (KubeSchedulerList) SwaggerDoc() map[string]string {
	return map_KubeSchedulerList
}

var map_KubeSchedulerSpec = map[string]string{
	"forceRedeploymentReason": "forceRedeploymentReason can be used to force the redeployment of the kube-scheduler by providing a unique string. This provides a mechanism to kick a previously failed deployment and provide a reason why you think it will work this time instead of failing again on the same config.",
}

func (KubeSchedulerSpec) SwaggerDoc() map[string]string {
	return map_KubeSchedulerSpec
}

// AUTO-GENERATED FUNCTIONS END HERE
