DB_URL=postgres://postgres:secret@localhost:5432/test?sslmode=disable
MIGRATION_DIR=your_migration_dir
SEEDERS_DIR=your_seeders_dir
# Default number of migration(s)
N = 1
# Migration schema version
V = 1

.PHONY: test
test:
	@$(foreach package,$(packages), \
	  set -e; \
	  go test -coverprofile $(package)/cover.out -covermode=count $(package);)

migrate-create:
	migrate create -ext sql -dir $(MIGRATION_DIR) -seq -digits 5 $(NAME)

migrate-force:
	migrate -source file://$(MIGRATION_DIR) -database $(DB_URL) force $(V)

migrate-up:
	migrate -source file://$(MIGRATION_DIR) -database $(DB_URL) up

migrate-down:
	migrate -source file://$(MIGRATION_DIR) -database $(DB_URL) down $(N)

migrate-test-up:
	migrate -source file://$(MIGRATION_DIR) -database $(TEST_DB_URL) up

migrate-test-down:
	migrate -source file://$(MIGRATION_DIR) -database $(TEST_DB_URL) down $(N)

seeders-create:
	migrate create -ext sql -dir $(SEEDERS_DIR) -seq -digits 5 $(NAME)

seeders-up:
	migrate -source file://$(SEEDERS_DIR) -database $(DB_URL) up

seeders-force:
	migrate -source file://$(SEEDERS_DIR) -database $(DB_URL) force $(V)

seeders-down:
	migrate -source file://$(SEEDERS_DIR) -database $(DB_URL) down
