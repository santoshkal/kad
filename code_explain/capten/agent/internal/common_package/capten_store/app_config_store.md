Detailed explanation of the provided functions and methods:

*1. UpsertAppConfig(agentpb.SyncAppData) error

This function is responsible for upserting (insert or update) application configuration data in the store. It takes a pointer to an agentpb.SyncAppData struct as input.

Validation: It first checks if the ReleaseName field in the agentpb.AppConfig within agentpb.SyncAppData is empty. An error is returned if it's empty.

Find Existing Record:

It creates a new ClusterAppConfig struct pointer (appConfig).
It attempts to find a record in the database with the same ReleaseName as the provided data using the dbClient.FindFirst function.
If an error occurs during retrieval, it checks the error type using gerrors.GetErrorType. If the error type is not postgresdb.ObjectNotExist, it prepares a custom error using the prepareError function (not shown in the code snippet) with details about the operation (Fetch), release name, and the original error. Otherwise, it sets the recordFound flag to false indicating no record found.
If a record is found but the ReleaseName field is empty, it also sets recordFound to false.
Update or Create Data:

Based on the recordFound flag:
If no record is found (recordFound is false), a new ClusterAppConfig struct is created with the data from agentpb.SyncAppData and then inserted using dbClient.Create.
If a record is found (recordFound is true), the existing ClusterAppConfig struct is updated with the new data using dbClient.Update.
Encoding App Values:

The OverrideValues, LaunchUIValues, and TemplateValues fields within agentpb.AppValues are encoded using base64 encoding before storing them in the database using base64.StdEncoding.EncodeToString.
LastUpdateTime:

The LastUpdateTime field in ClusterAppConfig is set to the current time using time.Now().
*2. GetAppConfig(appReleaseName string) (agentpb.SyncAppData, error)

This function retrieves application configuration data for a specific release name. It takes the appReleaseName (string) as input and returns a pointer to an agentpb.SyncAppData struct containing the data or an error.

Find Record:

It creates a new ClusterAppConfig struct pointer (appConfig).
It tries to find a record in the database with the matching ReleaseName using dbClient.FindFirst.
Error Handling:

If an error occurs during retrieval, it directly returns the error.
Decode App Values:

If a record is found, the encoded OverrideValues, LaunchUIValues, and TemplateValues fields in ClusterAppConfig are decoded back to byte slices using base64.StdEncoding.DecodeString.
Populate and Return Data:

A new agentpb.SyncAppData struct is created and populated with the decoded data and other fields from ClusterAppConfig.
The LastUpdateTime field in the agentpb.AppConfig is formatted to RFC3339 format before setting it in the struct.
Finally, the populated agentpb.SyncAppData struct is returned.
*3. GetAllApps() ([]agentpb.SyncAppData, error)

This function retrieves all application configuration data stored in the store. It returns a slice of pointers to agentpb.SyncAppData structs containing the data or an error.

Find All Records:

It retrieves all records of type ClusterAppConfig from the database using dbClient.Find.
Iterate and Process:

It iterates over the retrieved ClusterAppConfig slice.
For each record, it follows similar logic to the GetAppConfig function to decode the stored app values and populate a new agentpb.SyncAppData struct.
The populated agentpb.SyncAppData struct is then appended to a slice.
Return Data:

Finally, the slice containing pointers to all agentpb.SyncAppData structs is returned.
4. DeleteAppConfigByReleaseName(releaseName string) error

This function deletes application configuration data for a specific release name. It takes the releaseName (string) as input and returns an error if any occurs during deletion.