# Setup
You'll need to have `docker-compose` and `go` installed and functional on your machine before starting. Once those are installed, navigate to the project directory in your terminal and run `docker-compose build && docker-compose up -d`. This will spin up Docker containers and link them based on the project files.

# Database Population
Editing the `database/database.sql` file will change how MySQL implements the schema/data dump upon each `docker-compose up` command.
