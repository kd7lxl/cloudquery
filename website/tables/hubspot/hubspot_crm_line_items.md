# Table: hubspot_crm_line_items

This table shows data for HubSpot CRM Line Items.

https://developers.hubspot.com/docs/api/crm/line-items

The primary key for this table is **id**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|`utf8`|
|_cq_sync_time|`timestamp[us, tz=UTC]`|
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|id (PK)|`utf8`|
|properties|`json`|
|created_at|`timestamp[us, tz=UTC]`|
|updated_at|`timestamp[us, tz=UTC]`|
|archived|`bool`|
|archived_at|`timestamp[us, tz=UTC]`|
|associations|`json`|