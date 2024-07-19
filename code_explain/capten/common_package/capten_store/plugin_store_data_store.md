# This code snippet deals with managing Plugin Store Data in the store. Here's a breakdown of the functions:

*1. UpsertPluginStoreData(gitProjectID string, pluginData pluginstorepb.PluginData) error

This function performs an upsert operation on a Plugin Store Data record. It takes the gitProjectID (string) and a pointer to a pluginstorepb.PluginData struct containing the plugin details as input.

Check for Existing Record:

It creates a new PluginStoreData struct instance (empty).
It attempts to find a record with the matching StoreType (converted to int), GitProjectID (converted from string using uuid.MustParse), and PluginName using dbClient.FindFirst.
It checks the error type using gerrors.GetErrorType.
If the error type is postgresdb.ObjectNotExist, it means no record was found (recordFound is set to false). The error is then set to nil.
If the error is something else, the error is returned directly (considered a fetch error).
Populate PluginStoreData Struct:

The struct is populated with data from the input message (pluginData):
StoreType (converted to int)
GitProjectID (converted from string using uuid.MustParse)
PluginName
Category
Versions
Icon
Description
LastUpdateTime is set to the current time.
Insert or Update Data:

Based on the recordFound flag:
If recordFound is false (no record found), a new record is inserted using dbClient.Create.
If recordFound is true (record found), the existing record is updated using dbClient.Update. The update criteria is set to a composite key of StoreType, GitProjectID, and PluginName.
*2. GetPluginStoreData(storeType pluginstorepb.StoreType, gitProjectId, pluginName string) (pluginstorepb.PluginData, error)

This function retrieves a specific Plugin Store Data record based on a composite key of StoreType, gitProjectId, and pluginName. It takes these three parameters as input and returns a pointer to a pluginstorepb.PluginData struct containing the data or an error.

Find Record:

It creates a new PluginStoreData struct instance (empty).
It attempts to find a record with the matching StoreType (converted to int), GitProjectID (converted from string using uuid.MustParse), and PluginName using dbClient.FindFirst.
Return Data (if found):

If a record is found, a new pluginstorepb.PluginData struct is created and populated with the data from the retrieved PluginStoreData struct.
The function returns the populated pluginstorepb.PluginData struct and nil error.
Return Error (if not found):

If an error occurs during retrieval or no record is found, the function directly returns the error.
*3. GetPlugins(gitProjectId string) ([]pluginstorepb.Plugin, error)

This function retrieves all Plugin Store Data records for a specific Git Project identified by its gitProjectId. It takes the gitProjectId (string) as input and returns a slice of pointers to pluginstorepb.Plugin structs containing the data or an error.

Find Records:

It retrieves all records of type PluginStoreData with the matching GitProjectID (converted from string using uuid.MustParse) from the database using dbClient.Find.
Iterate and Process:

It iterates over the retrieved pluginDataSet slice.
For each record, a new pluginstorepb.Plugin struct is created and populated with the data from the corresponding PluginStoreData struct.
The populated pluginstorepb.Plugin struct is then appended to a slice plugins.
Return Data:

The function returns the slice plugins containing pointers to all pluginstorepb.Plugin structs and nil error.
4. DeletePluginStoreData(storeType pluginstorepb.StoreType, gitProjectId, pluginName string) error

This function deletes a Plugin Store Data record based on a composite key of StoreType, gitProjectId, and pluginName. It takes these three parameters as input and returns an