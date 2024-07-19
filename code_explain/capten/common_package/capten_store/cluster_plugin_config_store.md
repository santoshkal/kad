### Cluster Plugin configurations, incorporating a more comprehensive view and avoiding repetition.

Cluster Plugin Configuration Management

These functions handle the lifecycle of plugin configurations within the cluster. They provide functionalities for storing, retrieving, and deleting these configurations using a database as the backend storage.

1. Upserting Plugin Configurations (UpsertClusterPluginConfig):

This function takes a clusterpluginspb.Plugin protobuf message as input, representing the desired plugin configuration.
It performs basic validation by checking if the PluginName field is empty. An error is returned if the name is missing.
The function attempts to find an existing record in the database with the same PluginName. This leverages the dbClient and likely utilizes a database query mechanism.
If no record is found (indicating a new plugin configuration), a new ClusterPluginConfig struct is created. This struct maps the data structure of the stored configuration in the database.
The newly created ClusterPluginConfig struct is then populated with the data from the provided clusterpluginspb.Plugin message.
This populated ClusterPluginConfig struct is ultimately inserted into the database using the dbClient functionality.
If a record is already present (indicating an update scenario), the existing ClusterPluginConfig struct retrieved from the database is updated with the new data from the input message. The update operation likely utilizes the dbClient as well.
Before storing the configuration data, two specific fields, Values and OverrideValues, are encoded using base64 encoding. These fields likely contain additional configuration details in a format that needs to be persisted in the database.
Finally, the LastUpdateTime field within the ClusterPluginConfig struct is set to the current timestamp, indicating when the configuration was last modified.
2. Deleting Plugin Configurations (DeleteClusterPluginConfig):

This function takes a pluginName (string) as input, which identifies the specific plugin configuration to be deleted.
It attempts to delete the record with the matching PluginName from the database. This likely involves formulating a database query that targets the record based on the provided name and utilizes the dbClient for execution.
In case of any errors during the deletion process, a custom error is prepared with details about the operation (Delete), the plugin name, and the original error that occurred. This provides informative error handling for debugging purposes.
3. Retrieving a Specific Plugin Configuration (GetClusterPluginConfig):

This function retrieves a plugin configuration for a specific plugin based on its pluginName (string) provided as input.
It searches the database for a record with the matching PluginName. This likely involves a database query using the dbClient and filtering based on the provided name.
If a record is found, the encoded Values and OverrideValues fields within the retrieved ClusterPluginConfig struct are decoded back to their original format using base64 decoding.
A new clusterpluginspb.Plugin struct, which mirrors the message format used for exchanging plugin configuration data, is created.
This newly created clusterpluginspb.Plugin struct is then populated with the data retrieved from the database and the decoded Values and OverrideValues fields.
Finally, the populated clusterpluginspb.Plugin struct containing the retrieved configuration details for the specified plugin is returned.
4. Retrieving All Plugin Configurations (GetAllClusterPluginConfigs):

This function retrieves all Cluster Plugin configurations currently stored in the database.
It fetches all records of type ClusterPluginConfig from the database. This likely involves a general query that retrieves all entries from the relevant table using the dbClient.
The function then iterates over each retrieved ClusterPluginConfig record.
For each record, a new clusterpluginspb.Plugin struct is created, similar to the GetClusterPluginConfig function.
The newly created clusterpluginspb.Plugin struct is populated with the data from the corresponding ClusterPluginConfig struct.
The encoded Values and OverrideValues fields within the ClusterPluginConfig struct are decoded back to their original format using base64 decoding.
Finally, a slice containing pointers to all the populated clusterpluginspb.Plugin structs, representing the retrieved configurations for all plugins, is returned.
In summary, these functions provide a comprehensive mechanism for managing the lifecycle of Cluster Plugin configurations within the system. They handle upserting (insert or update), deleting, and retrieving individual or all plugin configurations, interacting with the database as the persistence layer.