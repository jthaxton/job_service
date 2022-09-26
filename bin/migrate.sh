export POSTGRESQL_URL="postgres://root:example@localhost:5432/job_db?sslmode=disable"
migrate -database ${POSTGRESQL_URL} -path db/migrations up