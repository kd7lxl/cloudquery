# Changelog

## [2.0.2](https://github.com/cloudquery/cloudquery/compare/scaffold-v2.0.1...scaffold-v2.0.2) (2023-06-06)


### Bug Fixes

* **deps:** Update github.com/apache/arrow/go/v13 digest to e07e22c ([#11151](https://github.com/cloudquery/cloudquery/issues/11151)) ([5083cf7](https://github.com/cloudquery/cloudquery/commit/5083cf720f0ae98e07448ba2ae1116048e2d3a90))
* **deps:** Update github.com/cloudquery/arrow/go/v13 digest to 88d5dc2 ([#11226](https://github.com/cloudquery/cloudquery/issues/11226)) ([9f306bc](https://github.com/cloudquery/cloudquery/commit/9f306bcaf3833b4611f0df5c50277be43aa19cbb))
* **deps:** Update github.com/cloudquery/arrow/go/v13 digest to c67fb39 ([#11169](https://github.com/cloudquery/cloudquery/issues/11169)) ([dcb0f92](https://github.com/cloudquery/cloudquery/commit/dcb0f9296a770a5cc2eb6bffd6b1ee30fbccb5dc))
* **deps:** Update module github.com/cloudquery/plugin-pb-go to v1.0.9 ([#11240](https://github.com/cloudquery/cloudquery/issues/11240)) ([f92cd4b](https://github.com/cloudquery/cloudquery/commit/f92cd4bfe3c3d0088964d52ab9cd01ca4cf622e1))
* **deps:** Update module github.com/cloudquery/plugin-sdk/v3 to v3.10.3 ([#11216](https://github.com/cloudquery/cloudquery/issues/11216)) ([f7ac733](https://github.com/cloudquery/cloudquery/commit/f7ac73346690ec36f437fbc88123e1404ff7184d))
* **deps:** Update module github.com/cloudquery/plugin-sdk/v3 to v3.10.4 ([#11244](https://github.com/cloudquery/cloudquery/issues/11244)) ([8fceef6](https://github.com/cloudquery/cloudquery/commit/8fceef6f9041e173923555d8ff221cfe83b424c2))
* **deps:** Update module github.com/cloudquery/plugin-sdk/v3 to v3.7.0 ([#11113](https://github.com/cloudquery/cloudquery/issues/11113)) ([487bf87](https://github.com/cloudquery/cloudquery/commit/487bf871afe360cb8d9d592dfea48837d6e7cf27))

## [2.0.1](https://github.com/cloudquery/cloudquery/compare/scaffold-v2.0.0...scaffold-v2.0.1) (2023-05-29)


### Bug Fixes

* **deps:** Update module github.com/cloudquery/plugin-sdk/v3 to v3.6.7 ([#11043](https://github.com/cloudquery/cloudquery/issues/11043)) ([3c6d885](https://github.com/cloudquery/cloudquery/commit/3c6d885c3d201b0b39cbc1406c6e54a57ec5ed5f))

## [2.0.0](https://github.com/cloudquery/cloudquery/compare/scaffold-v1.8.0...scaffold-v2.0.0) (2023-05-26)


### ⚠ BREAKING CHANGES

* This release introduces an internal change to our type system to use [Apache Arrow](https://arrow.apache.org/). This should not have any visible breaking changes, however due to the size of the change we are introducing it under a major version bump to communicate that it might have some bugs that we weren't able to catch during our internal tests. If you encounter an issue during the upgrade, please submit a [bug report](https://github.com/cloudquery/cloudquery/issues/new/choose). You will also need to update destinations depending on which one you use:
    - Azure Blob Storage >= v3.2.0
    - BigQuery >= v3.0.0
    - ClickHouse >= v3.1.1
    - DuckDB >= v1.1.6
    - Elasticsearch >= v2.0.0
    - File >= v3.2.0
    - Firehose >= v2.0.2
    - GCS >= v3.2.0
    - Gremlin >= v2.1.10
    - Kafka >= v3.0.1
    - Meilisearch >= v2.0.1
    - Microsoft SQL Server >= v4.2.0
    - MongoDB >= v2.0.1
    - MySQL >= v2.0.2
    - Neo4j >= v3.0.0
    - PostgreSQL >= v4.2.0
    - S3 >= v4.4.0
    - Snowflake >= v2.1.1
    - SQLite >= v2.2.0

### Features

* Update to use [Apache Arrow](https://arrow.apache.org/) type system ([#10998](https://github.com/cloudquery/cloudquery/issues/10998)) ([69e387a](https://github.com/cloudquery/cloudquery/commit/69e387a35bbaf2cfdec2718a0f30976eaec06eb2))

## [1.8.0](https://github.com/cloudquery/cloudquery/compare/scaffold-v1.7.0...scaffold-v1.8.0) (2023-05-18)


### Features

* **deps:** Upgrade to Apache Arrow v13 (latest `cqmain`) ([#10605](https://github.com/cloudquery/cloudquery/issues/10605)) ([a55da3d](https://github.com/cloudquery/cloudquery/commit/a55da3dbefafdc68a6bda2d5f1d334d12dd97b97))

## [1.7.0](https://github.com/cloudquery/cloudquery/compare/scaffold-v1.6.7...scaffold-v1.7.0) (2023-05-02)


### Features

* Migrate scaffold to plugin-sdk v2 ([#10377](https://github.com/cloudquery/cloudquery/issues/10377)) ([62f26a8](https://github.com/cloudquery/cloudquery/commit/62f26a8d5c3f27eda196ca4192df23a85caf54cb))


### Bug Fixes

* **deps:** Update module github.com/cloudquery/plugin-sdk/v2 to v2.5.0 ([#10390](https://github.com/cloudquery/cloudquery/issues/10390)) ([f706688](https://github.com/cloudquery/cloudquery/commit/f706688b2f5b8393d09d57020d31fb1d280f0dbd))
* **deps:** Update module github.com/cloudquery/plugin-sdk/v2 to v2.5.1 ([#10448](https://github.com/cloudquery/cloudquery/issues/10448)) ([cc85b93](https://github.com/cloudquery/cloudquery/commit/cc85b939fe945939caf72f8c08095e1e744b9ee8))

## [1.6.7](https://github.com/cloudquery/cloudquery/compare/scaffold-v1.6.6...scaffold-v1.6.7) (2023-04-25)


### Bug Fixes

* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.45.0 ([#9863](https://github.com/cloudquery/cloudquery/issues/9863)) ([2799d62](https://github.com/cloudquery/cloudquery/commit/2799d62518283ac304beecda9478f8f2db43cdc5))

## [1.6.6](https://github.com/cloudquery/cloudquery/compare/scaffold-v1.6.5...scaffold-v1.6.6) (2023-04-04)


### Bug Fixes

* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.44.1 ([#9520](https://github.com/cloudquery/cloudquery/issues/9520)) ([202c31b](https://github.com/cloudquery/cloudquery/commit/202c31b2788c3df35b5df7d07fdc750f92e7bb23))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.44.2 ([#9661](https://github.com/cloudquery/cloudquery/issues/9661)) ([a27dc84](https://github.com/cloudquery/cloudquery/commit/a27dc84a9b67b68b5b75b04dd3afe13e2c556082))

## [1.6.5](https://github.com/cloudquery/cloudquery/compare/scaffold-v1.6.4...scaffold-v1.6.5) (2023-03-21)


### Bug Fixes

* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.44.0 ([#9167](https://github.com/cloudquery/cloudquery/issues/9167)) ([49d6477](https://github.com/cloudquery/cloudquery/commit/49d647730a85ea6fae51e97194ba61c0625d1331))

## [1.6.4](https://github.com/cloudquery/cloudquery/compare/scaffold-v1.6.3...scaffold-v1.6.4) (2023-03-14)


### Bug Fixes

* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.43.0 ([#8949](https://github.com/cloudquery/cloudquery/issues/8949)) ([31dfc63](https://github.com/cloudquery/cloudquery/commit/31dfc634850b699ba7bb7876399270a7367d6c7e))

## [1.6.3](https://github.com/cloudquery/cloudquery/compare/scaffold-v1.6.2...scaffold-v1.6.3) (2023-03-07)


### Bug Fixes

* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.41.0 ([#8682](https://github.com/cloudquery/cloudquery/issues/8682)) ([ea9d065](https://github.com/cloudquery/cloudquery/commit/ea9d065ae9f77c6dd990570974630ae6ac3f153e))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.42.0 ([#8725](https://github.com/cloudquery/cloudquery/issues/8725)) ([b83b277](https://github.com/cloudquery/cloudquery/commit/b83b277a2421d1caf46a26c3229041b27a3da148))

## [1.6.2](https://github.com/cloudquery/cloudquery/compare/scaffold-v1.6.1...scaffold-v1.6.2) (2023-02-28)


### Bug Fixes

* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.39.0 ([#8344](https://github.com/cloudquery/cloudquery/issues/8344)) ([9c57544](https://github.com/cloudquery/cloudquery/commit/9c57544d06f9a774adcc659bcabd2518a905bdaa))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.39.1 ([#8371](https://github.com/cloudquery/cloudquery/issues/8371)) ([e3274c1](https://github.com/cloudquery/cloudquery/commit/e3274c109739bc107387627d340a713470c3a3c1))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.40.0 ([#8401](https://github.com/cloudquery/cloudquery/issues/8401)) ([4cf36d6](https://github.com/cloudquery/cloudquery/commit/4cf36d68684f37c0407332930766c1ba60807a93))

## [1.6.1](https://github.com/cloudquery/cloudquery/compare/scaffold-v1.6.0...scaffold-v1.6.1) (2023-02-21)


### Bug Fixes

* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.38.2 ([#8156](https://github.com/cloudquery/cloudquery/issues/8156)) ([ac2d2d7](https://github.com/cloudquery/cloudquery/commit/ac2d2d70d5c4bc45fb8734bd4deb8a1e36074f6d))

## [1.6.0](https://github.com/cloudquery/cloudquery/compare/scaffold-v1.5.2...scaffold-v1.6.0) (2023-02-17)


### Features

* **readme:** Add release guide section ([#8149](https://github.com/cloudquery/cloudquery/issues/8149)) ([2b6275e](https://github.com/cloudquery/cloudquery/commit/2b6275ec0aa5058055c5e8088bc99e4a7664e658))

## [1.5.2](https://github.com/cloudquery/cloudquery/compare/scaffold-v1.5.1...scaffold-v1.5.2) (2023-02-14)


### Bug Fixes

* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.36.0 ([#7809](https://github.com/cloudquery/cloudquery/issues/7809)) ([c85a9cb](https://github.com/cloudquery/cloudquery/commit/c85a9cb697477520e94a1fd260c56b89da62fc87))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.36.1 ([#7930](https://github.com/cloudquery/cloudquery/issues/7930)) ([39dccc1](https://github.com/cloudquery/cloudquery/commit/39dccc1bf81f4eb02d181ba0c47b37038a4c5455))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.37.0 ([#7933](https://github.com/cloudquery/cloudquery/issues/7933)) ([dc9cffb](https://github.com/cloudquery/cloudquery/commit/dc9cffbf37bbc6fae73a20bf47e6bbf17e74d1f9))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.37.1 ([#8008](https://github.com/cloudquery/cloudquery/issues/8008)) ([c47aac0](https://github.com/cloudquery/cloudquery/commit/c47aac0b5e3190a04299713651b97e360043911f))

## [1.5.1](https://github.com/cloudquery/cloudquery/compare/scaffold-v1.5.0...scaffold-v1.5.1) (2023-02-07)


### Bug Fixes

* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.34.0 ([#7719](https://github.com/cloudquery/cloudquery/issues/7719)) ([6a33085](https://github.com/cloudquery/cloudquery/commit/6a33085c75adcf2387f7bbb5aa4f7a84ce7e2957))

## [1.5.0](https://github.com/cloudquery/cloudquery/compare/scaffold-v1.4.0...scaffold-v1.5.0) (2023-02-06)


### Features

* Add *.log to gitignore template ([#7680](https://github.com/cloudquery/cloudquery/issues/7680)) ([6f5586a](https://github.com/cloudquery/cloudquery/commit/6f5586aa323678b3b92c9aeb48ed630192e3345c))

## [1.4.0](https://github.com/cloudquery/cloudquery/compare/scaffold-v1.3.1...scaffold-v1.4.0) (2023-02-02)


### Features

* Add `.gitignore` template ([#7608](https://github.com/cloudquery/cloudquery/issues/7608)) ([76b7c09](https://github.com/cloudquery/cloudquery/commit/76b7c0942883df7d4132aade63bea4e241da1561))


### Bug Fixes

* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.33.0 ([#7595](https://github.com/cloudquery/cloudquery/issues/7595)) ([c5adc75](https://github.com/cloudquery/cloudquery/commit/c5adc750d4b0242563997c04c582f8da27913095))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.33.1 ([#7614](https://github.com/cloudquery/cloudquery/issues/7614)) ([2fe665c](https://github.com/cloudquery/cloudquery/commit/2fe665cdd80d88c5699bb203bd7accd604dfba99))
* **goreleaser:** Use `clean` instead of `rm-dist` ([#7606](https://github.com/cloudquery/cloudquery/issues/7606)) ([c7e166d](https://github.com/cloudquery/cloudquery/commit/c7e166dbb99b53ee67c4ea178cafdcdf4feddb17))

## [1.3.1](https://github.com/cloudquery/cloudquery/compare/scaffold-v1.3.0...scaffold-v1.3.1) (2023-01-31)


### Bug Fixes

* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.30.0 ([#7222](https://github.com/cloudquery/cloudquery/issues/7222)) ([73ca21c](https://github.com/cloudquery/cloudquery/commit/73ca21c4259545f7e949c9d780d8184db475d2ac))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.31.0 ([#7228](https://github.com/cloudquery/cloudquery/issues/7228)) ([36e8549](https://github.com/cloudquery/cloudquery/commit/36e8549f722658d909865723630fad1b2821db62))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.32.0 ([#7334](https://github.com/cloudquery/cloudquery/issues/7334)) ([b684122](https://github.com/cloudquery/cloudquery/commit/b68412222219f9ca160c0753290709d52de7fcd6))

## [1.3.0](https://github.com/cloudquery/cloudquery/compare/scaffold-v1.2.0...scaffold-v1.3.0) (2023-01-26)


### Features

* **scaffold-go-releaser:** Use go version from `go.mod`, add `cache: true` ([#7058](https://github.com/cloudquery/cloudquery/issues/7058)) ([946d643](https://github.com/cloudquery/cloudquery/commit/946d643476b7b741ee10650edcc7243878e65fe9))

## [1.2.0](https://github.com/cloudquery/cloudquery/compare/scaffold-v1.1.0...scaffold-v1.2.0) (2023-01-25)


### Features

* **scaffold:** Add Makefile and test workflow to scaffold, and flesh out readme ([#7118](https://github.com/cloudquery/cloudquery/issues/7118)) ([d75df29](https://github.com/cloudquery/cloudquery/commit/d75df29d468a84724229ecef6ce558d31e665707))


### Bug Fixes

* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.29.0 ([#7121](https://github.com/cloudquery/cloudquery/issues/7121)) ([b7441c9](https://github.com/cloudquery/cloudquery/commit/b7441c93c274ae3a6009474a2b28f44a172dd6dc))

## [1.1.0](https://github.com/cloudquery/cloudquery/compare/scaffold-v1.0.0...scaffold-v1.1.0) (2023-01-24)


### Features

* **scaffold:** Update SDK, add auto deps updates ([#7115](https://github.com/cloudquery/cloudquery/issues/7115)) ([afd7793](https://github.com/cloudquery/cloudquery/commit/afd77938053ab9e539fe2a9233fe1f1cc60e2f04))
* **scaffold:** Update SDK, add auto deps updates ([#7115](https://github.com/cloudquery/cloudquery/issues/7115)) ([3c529cf](https://github.com/cloudquery/cloudquery/commit/3c529cf55e258df54b7c861b525cb57a4195b354))

## 1.0.0 (2023-01-15)


### Features

* Add scaffold CLI ([#6789](https://github.com/cloudquery/cloudquery/issues/6789)) ([7fad312](https://github.com/cloudquery/cloudquery/commit/7fad312104967dbaddeb746d83282bc6e04a7eed))


### Bug Fixes

* **scaffold:** Use org and name for plugin name instead of `test` ([#6802](https://github.com/cloudquery/cloudquery/issues/6802)) ([c5e10fc](https://github.com/cloudquery/cloudquery/commit/c5e10fcc298429ec0049e5b22f0f5f12c63182da))
