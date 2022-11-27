# Ginger Beer

## Boilerplate generation

This tool uses a database connection and build :
- an OpenAPI 3.0 specification contract in format :
  - JSON 
  - YAML

## Usage

### Create the database

```bash
createdb ginger-beer 
```

### Create your table(s)

```sql
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL,
    password VARCHAR(255) NOT NULL
);
```

You can use any way you want to build your tables (SQL script, database designer, already existing database, etc).

### Generate your boilerplate :

```bash
$ ginger-beer -c "postgresql://localhost/ginger-beer" -o ./output -f yaml
```

-c : database connection string (default : postgresql://localhost/postgres)
-o : output directory (optional, default is your current directory)
-f : format (optional, default is json)

### Other commands

#### Display the help :

```bash
$ ginger-beer -h
   Usage of ginger-beer:
    -f string
      Path to the model file (default "model.go")
    -o string
      Path to the output directory (default ".")
    -t string
      Type of boilerplate to generate (default "basic")
    -format string
      Format of the OpenAPI 3.0 specification (default "json")
    -v    Display version
```