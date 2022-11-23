# Ginger Beer

## Boilerplate generation

This tool uses a model struct file as input and build :
- an OpenAPI 3.0 specification contract
- choice of :  
  - a `clean` architecture repository and file structure
      - a controller
      - a use case
      - a domain
      - a repository
      - a factory
  - a `basic` controller for CRUD operations 

## Usage

Display the help :

```bash
$ ginger-beer -h
   Usage of ginger-beer:
    -f string
      Path to the model file (default "model.go")
    -o string
      Path to the output directory (default ".")
    -t string
      Type of boilerplate to generate (default "clean")
```

Generate your boilerplate :

```bash
$ ginger-beer -f model.go -o ./output -t clean
```
