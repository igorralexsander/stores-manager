export CREATE_KEYSPACE_PATH := $(shell pwd)/db/migration/create_keyspace.cql

createkeyspace:
	docker exec -i marketplace-analytics-scylla-1\
 		cqlsh -u cassandra -p cassandra\
 		-e "$(shell cat ${CREATE_KEYSPACE_PATH})"

migrationup:
	migrate -path db/migration -database "cassandra://127.0.0.1:9042/marketplace_analytics_stores?username=cassandra&password=cassandra"  -verbose up

.PHONY: migrationup