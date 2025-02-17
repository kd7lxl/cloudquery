# Table: aws_rds_events

This table shows data for Amazon Relational Database Service (RDS) Events.

https://docs.aws.amazon.com/AmazonRDS/latest/APIReference/API_DescribeEvents.html

The primary key for this table is **_cq_id**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|`utf8`|
|_cq_sync_time|`timestamp[us, tz=UTC]`|
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|date|`timestamp[us, tz=UTC]`|
|event_categories|`list<item: utf8, nullable>`|
|message|`utf8`|
|source_arn|`utf8`|
|source_identifier|`utf8`|
|source_type|`utf8`|