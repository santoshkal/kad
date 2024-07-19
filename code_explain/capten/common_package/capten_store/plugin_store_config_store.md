## This code snippet deals with managing Plugin Store configurations in the store. Here's a breakdown of the functions:

*1. UpsertPluginStoreConfig(config pluginstorepb.PluginStoreConfig) error

This function performs an upsert operation on a Plugin Store configuration. It takes a pointer to a pluginstorepb.PluginStoreConfig struct containing the configuration details as input.

Check for Existing Record:

It creates a new PluginStoreConfig struct instance (empty).
It attempts to find a record with the matching StoreType (converted to int) using dbClient.FindFirst.
It checks the error type using gerrors.GetErrorType.
If the error type is postgresdb.ObjectNotExist, it means no record was found (recordFound is set to false). The error is then set to nil.
If the error is something else, the error is returned directly (considered a fetch error).
Populate PluginStoreConfig Struct:

The StoreType, GitProjectID (converted from string using uuid.MustParse), GitProjectURL, and LastUpdateTime fields are set in the pluginStoreConfig struct with the corresponding values from the input message.
Insert or Update Data:

Based on the recordFound flag:
If recordFound is false (no record found), a new record is inserted using dbClient.Create.
If recordFound is true (record found), the existing record is updated using dbClient.Update. The update criteria is set to the StoreType.
*2. GetPluginStoreConfig(storeType pluginstorepb.StoreType) (pluginstorepb.PluginStoreConfig, error)

This function retrieves a specific Plugin Store configuration based on its StoreType. It takes the storeType as input and returns a pointer to a pluginstorepb.PluginStoreConfig struct containing the data or an error.

Find Record:

It creates a new PluginStoreConfig struct instance (empty).
It attempts to find a record with the matching StoreType (converted to int) using dbClient.FindFirst.
Return Data (if found):

If a record is found, a new pluginstorepb.PluginStoreConfig struct is created and populated with the data from the retrieved PluginStoreConfig struct.
The function returns the populated pluginstorepb.PluginStoreConfig struct and nil error.
Return Error (if not found):

If an error occurs during retrieval or no record is found, the function directly returns the error.
3. DeletePluginStoreConfig(storeType pluginstorepb.StoreType) error

This function deletes a Plugin Store configuration based on its StoreType. It takes the storeType as input and returns an error if deletion fails.

Delete Record:

It attempts to delete the record with the matching StoreType (converted to int) from the database using dbClient.Delete.
Return Error (if deletion fails):

In case of errors during deletion, the error is directly returned.