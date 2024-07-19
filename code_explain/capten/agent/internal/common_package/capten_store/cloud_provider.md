These functions and methods deal with managing Cloud Provider information in the store.

*1. AddCloudProvider(config captenpluginspb.CloudProvider) error

This function adds a new Cloud Provider configuration to the store. It takes a pointer to a captenpluginspb.CloudProvider struct containing the cloud provider details as input.

Create CloudProvider Struct:
A new CloudProvider struct is created.
Set Fields:
The ID field is set using uuid.MustParse to convert the Id field from the input message to a valid UUID type.
The CloudType, Labels, and LastUpdateTime fields are directly copied from the input message.
Store Data:
The newly created CloudProvider struct is inserted into the database using dbClient.Create.
*2. UpsertCloudProvider(config captenpluginspb.CloudProvider) error

This function performs an upsert operation (insert or update) on a Cloud Provider configuration. It takes a pointer to a captenpluginspb.CloudProvider struct containing the cloud provider details as input.

Check for Existing ID:
It checks if the Id field in the input message is empty.
Create New CloudProvider (if ID is empty):
If the Id is empty, a new CloudProvider struct is created with a generated UUID using uuid.New for the ID field. Other fields are populated from the input message.
The newly created CloudProvider struct is then inserted into the database using dbClient.Create.
Update Existing CloudProvider (if ID is present):
If the Id is present, a new CloudProvider struct is created with the parsed ID using uuid.MustParse and populated with other fields from the input message.
The existing Cloud Provider record with the matching ID is then updated using dbClient.Update.
*3. GetCloudProviderForID(id string) (captenpluginspb.CloudProvider, error)

This function retrieves a specific Cloud Provider configuration based on its ID. It takes the id (string) as input and returns a pointer to a captenpluginspb.CloudProvider struct containing the data or an error.

Find Record:

It creates a new CloudProvider struct instance.
It attempts to find a record in the database with the matching ID using dbClient.FindFirst. The ID is converted to a UUID type before searching.
Error Handling:

If an error occurs during retrieval, it directly returns the error.
Populate and Return Data:

If a record is found, a new captenpluginspb.CloudProvider struct is created and populated with the data from the CloudProvider struct retrieved from the database.
The LastUpdateTime field is converted to a string format before setting it in the struct.
Finally, the populated captenpluginspb.CloudProvider struct is returned.
*4. GetCloudProviders() ([]captenpluginspb.CloudProvider, error)

This function retrieves all Cloud Provider configurations stored in the database. It returns a slice of pointers to captenpluginspb.CloudProvider structs containing the data or an error.

Find All Records:

It retrieves all records of type CloudProvider from the database using dbClient.Find.
Iterate and Process:

It iterates over the retrieved CloudProvider slice.
For each record, a new captenpluginspb.CloudProvider struct is created and populated with the data from the corresponding CloudProvider struct.
The LastUpdateTime field is converted to a string format before setting it in the struct.
The populated captenpluginspb.CloudProvider struct is then appended to a slice.
Return Data:

Finally, the slice containing pointers to all captenpluginspb.CloudProvider structs is returned.
*5. GetCloudProvidersByLabelsAndCloudType(searchLabels []string, cloudType string) ([]captenpluginspb.CloudProvider, error)

This function retrieves Cloud Provider configurations based on a combination of search labels and cloud type. It takes two arguments:

searchLabels: A slice of strings representing the labels to search for.
cloudType: A string representing the cloud type to filter by.
Search with GQL:
It utilizes GQL syntax within the dbClient.Session to perform the