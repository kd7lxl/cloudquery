# Table: azure_appservice_deleted_web_apps

This table shows data for Azure App Service Deleted Web Apps.

https://learn.microsoft.com/en-us/rest/api/appservice/deleted-web-apps/list#deletedsite

The primary key for this table is **id**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|`utf8`|
|_cq_sync_time|`timestamp[us, tz=UTC]`|
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|subscription_id|`utf8`|
|kind|`utf8`|
|properties|`json`|
|id (PK)|`utf8`|
|name|`utf8`|
|type|`utf8`|