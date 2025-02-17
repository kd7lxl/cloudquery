# Table: aws_codebuild_source_credentials

This table shows data for AWS CodeBuild Source Credentials.

https://docs.aws.amazon.com/codebuild/latest/APIReference/API_SourceCredentialsInfo.html

The primary key for this table is **arn**.

## Relations

The following tables depend on aws_codebuild_source_credentials:
  - [aws_codebuild_builds](aws_codebuild_builds)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|`utf8`|
|_cq_sync_time|`timestamp[us, tz=UTC]`|
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|arn (PK)|`utf8`|
|auth_type|`utf8`|
|server_type|`utf8`|