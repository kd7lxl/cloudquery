# Table: datadog_notebooks

This table shows data for Datadog Notebooks.

The composite primary key for this table is (**account_name**, **id**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|`utf8`|
|_cq_sync_time|`timestamp[us, tz=UTC]`|
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_name (PK)|`utf8`|
|attributes|`json`|
|id (PK)|`int64`|
|type|`utf8`|
|additional_properties|`json`|