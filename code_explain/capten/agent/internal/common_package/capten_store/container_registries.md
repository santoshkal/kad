This code snippet deals with managing Container Registry configurations in the store. Here's a breakdown of the functions:

*1. AddContainerRegistry(config captenpluginspb.ContainerRegistry) error

This function adds a new Container Registry configuration to the store. It takes a pointer to a captenpluginspb.ContainerRegistry struct containing the registry details as input.

Create ContainerRegistry Struct:
A new ContainerRegistry struct is created.
Set Fields:
The ID field is set using uuid.MustParse to convert the Id field from the input message to a valid UUID type.
The RegistryURL, RegistryType, Labels, and LastUpdateTime fields are directly copied from the input message.
Store Data:
The newly created ContainerRegistry struct is inserted into the database using dbClient.Create.
*2. UpsertContainerRegistry(config captenpluginspb.ContainerRegistry) error

This function performs an upsert operation (insert or update) on a Container Registry configuration. It takes a pointer to a captenpluginspb.ContainerRegistry struct containing the registry details as input.

Check for Existing ID:
It checks if the Id field in the input message is empty.
Create New ContainerRegistry (if ID is empty):
If the Id is empty, a new ContainerRegistry struct is created with a generated UUID using uuid.New for the ID field. Other fields are populated from the input message.
The newly created ContainerRegistry struct is then inserted into the database using dbClient.Create.
Update Existing ContainerRegistry (if ID is present):
If the Id is present, a new ContainerRegistry struct is created with the parsed ID using uuid.MustParse and populated with other fields from the input message.
The existing Container Registry record with the matching ID is then updated using dbClient.Update.
*3. GetContainerRegistryForID(id string) (captenpluginspb.ContainerRegistry, error)

This function retrieves a specific Container Registry configuration based on its ID. It takes the id (string) as input and returns a pointer to a captenpluginspb.ContainerRegistry struct containing the data or an error.

Find Record:

It creates a new ContainerRegistry struct instance.
It attempts to find a record in the database with the matching ID using dbClient.FindFirst. The ID is converted to a UUID type before searching.
Error Handling:

If an error occurs during retrieval, it directly returns the error.
Populate and Return Data:

If a record is found, a new captenpluginspb.ContainerRegistry struct is created and populated with the data from the ContainerRegistry struct retrieved from the database.
The LastUpdateTime field is converted to a string format before setting it in the struct.
Finally, the populated captenpluginspb.ContainerRegistry struct is returned.
*4. GetContainerRegistries() ([]captenpluginspb.ContainerRegistry, error)

This function retrieves all Container Registry configurations stored in the database. It returns a slice of pointers to captenpluginspb.ContainerRegistry structs containing the data or an error.

Find All Records:

It retrieves all records of type ContainerRegistry from the database using dbClient.Find.
Iterate and Process:

It iterates over the retrieved ContainerRegistry slice.
For each record, a new captenpluginspb.ContainerRegistry struct is created and populated with the data from the corresponding ContainerRegistry struct.
The LastUpdateTime field is converted to a string format before setting it in the struct.
The populated captenpluginspb.ContainerRegistry struct is then appended to a slice.
Return Data:

Finally, the slice containing pointers to all captenpluginspb.ContainerRegistry structs is returned.
*5. GetContainerRegistriesByLabels(searchLabels []string) ([]captenpluginspb.ContainerRegistry, error)

This function retrieves Container Registry configurations based on search labels. It takes a slice of strings (searchLabels) representing the labels to search for as input.

Search with GQL:

It utilizes GQL syntax within the dbClient.Find to perform the search based on labels.
Process Results:

Similar to GetContainerRegistries,