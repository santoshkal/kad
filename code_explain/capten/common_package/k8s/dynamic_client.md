## This code focuses on extending the K8SClient functionalities for managing Kubernetes resources using the dynamic client. Here's a detailed breakdown:

DynamicClientSet:

This struct represents a wrapper around the underlying dynamic client interface.
YAML/JSON Conversion:

ConvertYamlToJson and ConvertJsonToYaml functions provide utilities for converting between YAML and JSON formats for resource data.
Resource Identification:

GetNameNamespace function extracts the namespace and name of a resource from its JSON representation.
Extracting Group Version Kind (GVK):

getGVK function uses the unstructured.Unstructured type to decode YAML/JSON data and identify the Group, Version, and Kind (GVK) of the resource.
Resource Management:

CreateResourceFromFile function reads a YAML file representing a Kubernetes resource definition and creates the resource in the cluster.
CreateResource function takes raw resource data (YAML or JSON) and creates the resource in the cluster. It handles both cluster-scoped resources (without a namespace) and namespace-scoped resources. It checks for existing resources before creation to avoid duplicates (unless forced).
GetResource function retrieves an existing resource based on its definition in a YAML file.
ListNamespaceResource function retrieves a list of resources of a specific GVK within a particular namespace.
ListAllNamespaceResource retrieves a list of resources of a specific GVK across all namespaces in the cluster.
Error Handling:

Similar to the core K8SClient, these functions handle potential errors during resource interactions with the Kubernetes API server using k8serrors.IsNotFound for checking resource existence.
Overall, this code provides a dynamic client wrapper for the K8SClient that enables working with various Kubernetes resources without requiring pre-defined Go structs for each resource type.