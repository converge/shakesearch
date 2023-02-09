# Database resources

This directory contains database resources. The **migrations** folder contains SQL files that are managed by 
[migrate](https://github.com/golang-migrate/migrate) tool. 

Migrate enables migration management, automating repetitive tasks and to give developers confidence to handle database
changes.

## Generating SQL content based on Shakespeare eBook

The folder `cmd/cli` contains a CLI tool that generates SQL content based on the Shakespeare eBook. The tool is called 
via:

```bash
go run cmd/cli/main.go
```

This is what the CLI does:
- Reads the Shakespeare eBook from the **data** folder
- Parse the eBook into a list of chapters, grouped by id
- Generates SQL content based on the title and chapters
- Stores the SQL statements into the **migrations** folder (check db/migrations/000002_shakesearch_data.up.sql)

**Next, when migration is called, the SQL statements are executed and the database is populated with the Shakespeare 
eBook.**

## Running migrations

The migrations are run via the CLI tool, and [Task](https://taskfile.dev) can be used to call the CLI and run the 
migrations.

```bash
task db-migrate-up
```

## Reverting migrations

The migrations can be reverted via the CLI tool, and [Task](https://taskfile.dev) can be used to call the CLI and the
migrations.

```bash
task db-migrate-down
```

*note: a confirmation step is needed for reverting the migration, or forced with the **-f** flag.* 
