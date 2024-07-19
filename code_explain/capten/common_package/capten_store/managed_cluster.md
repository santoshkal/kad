## This code snippet deals with managing Managed Cluster configurations in the store. Here's a breakdown of the functions:

*1. AddManagedCluster(managedCluster captenpluginspb.ManagedCluster) error

This function adds a new Managed Cluster configuration to the store. It takes a pointer to a captenpluginspb.ManagedCluster struct containing the cluster details as input.

Create ManagedCluster Struct:
A new ManagedCluster struct is created.
Set Fields:
The ID field is set using uuid.MustParse to convert the Id field from the input message to a valid UUID type.
The ClusterName, ClusterEndpoint, ClusterDeployStatus, AppDeployStatus, and LastUpdateTime fields are directly copied from the input message.
Store Data:
The newly created ManagedCluster struct is inserted into the database using dbClient.Create.
*2. UpsertManagedCluster(managedCluster captenpluginspb.ManagedCluster) error

This function performs an upsert operation on a Managed Cluster configuration. It takes a pointer to a captenpluginspb.ManagedCluster struct containing the cluster details as input.

Check for Existing ID:
It checks if the Id field in the input message is empty.
Create New ManagedCluster (if ID is empty):
If the Id is empty, a new ManagedCluster struct is created with a generated UUID using uuid.New for the ID field. Other fields are populated from the input message.
The newly created ManagedCluster struct is then inserted into the database using dbClient.Create.
Update Existing ManagedCluster (if ID is present):
If the Id is present, a new ManagedCluster struct is created with other fields populated from the input message. The ID field is set using uuid.MustParse to convert the parsed ID from the message.
The existing Managed Cluster record with the matching ID is then updated using dbClient.Update.
3. DeleteManagedClusterById(id string) error

This function deletes a Managed Cluster configuration based on its ID. It takes the id (string) as input and returns an error if deletion fails.

Delete Record:
It attempts to delete the record with the matching ID (converted to a UUID type) from the database using dbClient.Delete.
Error Handling:
In case of errors during deletion, a custom error is prepared with details about the operation (Delete), the ID, and the original error.
*4. GetManagedClusterForID(id string) (captenpluginspb.ManagedCluster, error)

This function retrieves a specific Managed Cluster configuration based on its ID. It takes the id (string) as input and returns a pointer to a captenpluginspb.ManagedCluster struct containing the data or an error.

Find Record:

It creates a new ManagedCluster struct instance.
It attempts to find a record in the database with the matching ID using dbClient.FindFirst. The ID is converted to a UUID type before searching.
Error Handling:

If an error occurs during retrieval, it directly returns the error.
Populate and Return Data:

If a record is found, a new captenpluginspb.ManagedCluster struct is created and populated with the data from the ManagedCluster struct retrieved from the database.
The LastUpdateTime field is converted to a string format before setting it in the struct.
Finally, the populated captenpluginspb.ManagedCluster struct is returned.
*5. GetManagedClusters() ([]captenpluginspb.ManagedCluster, error)

This function retrieves all Managed Cluster configurations stored in the database. It returns a slice of pointers to captenpluginspb.ManagedCluster structs containing the data or an error.

Find All Records:

It retrieves all records of type ManagedCluster from the database using dbClient.Find.
Iterate and Process:

It iterates over the retrieved ManagedCluster slice.
For each record, a new captenpluginspb.ManagedCluster struct is created and populated with the data from the corresponding ManagedCluster struct.
The LastUpdateTime field is converted to a string format before setting it in the struct.
The populated captenpluginspb.ManagedCluster struct is then appended to a slice.

