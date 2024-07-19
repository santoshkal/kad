## This code defines two functions within the K8SClient struct for managing resources related to the External Secrets Operator:

1. CreateOrUpdateSecretStore:

This function creates or updates a ClusterSecretStore resource in the Kubernetes cluster.
It takes arguments for:
secretStoreName: Name for the SecretStore resource.
namespace: Namespace where the SecretStore resource should be created.
vaultAddr: Address of the Vault server.
tokenSecretName: Name of the Kubernetes Secret containing the Vault token.
tokenSecretKey: Key within the token Secret that holds the actual Vault token value.
It constructs a SecretStore object with the provided details:
API version for the External Secrets Operator.
Kind set to ClusterSecretStore indicating a cluster-wide secret store.
Metadata with the desired name for the SecretStore resource.
Spec section defines the configuration for the secret store:
Refresh interval set to 10 seconds (optional).
Provider section specifies Vault as the backend:
Server address of the Vault instance.
Path within Vault to access secrets (defaults to "secret").
Version of the Vault API used (defaults to "v2").
Authentication details using a token stored in a Kubernetes Secret:
Reference to the Secret containing the Vault token.
Key within the Secret that holds the actual token value.
Namespace where the token Secret resides.
The function marshals the SecretStore object into YAML format.
It then calls DynamicClient.CreateResource to create or update the SecretStore resource in the cluster using the dynamic client.
2. CreateOrUpdateExternalSecret:

This function creates or updates an ExternalSecret resource in a specific namespace.
It takes arguments for:
externalSecretName: Name for the ExternalSecret resource.
namespace: Namespace where the ExternalSecret resource should be created.
secretStoreRefName: Name of the existing ClusterSecretStore resource that provides secrets.
secretName: Name of the Kubernetes Secret where the retrieved secrets will be stored.
secretType: Type of the Secret (e.g., Opaque, kubernetes.io/tls).
vaultKeyPathdata: Map containing key-value pairs where keys are desired secret keys in the final Secret and values are paths within Vault to retrieve the secret values.
It constructs an ExternalSecret object with the provided details:
API version for the External Secrets Operator.
Kind set to ExternalSecret indicating a regular secret referencing external data.
Metadata with the desired name and namespace for the ExternalSecret resource.
Spec section defines how to fetch and store the secrets:
Refresh interval set to 10 seconds (as a string, optional).
Target section defines where to store the retrieved secrets:
Name of the Kubernetes Secret where the secrets will be placed.
Type of the Secret to be created.
SecretStoreRef section specifies the ClusterSecretStore resource to use for secrets:
Name of the existing ClusterSecretStore resource.
Kind set to ClusterSecretStore to reference the cluster-wide secret store.
Data section defines the mapping between desired secret keys and their corresponding paths within Vault:
Loops through the provided vaultKeyPathdata map.
For each key-value pair, it creates an ExternalSecretData object:
SecretKey: The desired name of the secret key in the final Kubernetes Secret.
RemoteRef: Reference to the secret data within Vault:
Key: Path within Vault to retrieve the secret value.
Property: Name of the secret key in the final Kubernetes Secret (defaults to the key from the map).
The function marshals the ExternalSecret object into YAML format.
It then calls DynamicClient.CreateResource to create or update the ExternalSecret resource in the specified namespace using the dynamic client.

Overall, these functions simplify managing secrets stored in Vault by leveraging the External Secrets Operator in Kubernetes.