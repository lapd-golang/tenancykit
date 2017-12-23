# Tokens
Tokens provide a package centered around th idea of a multirecord types which have a identifying key and differing values that 
can not have duplicates. It allows us build a backend that acts like a set where a given key and it's non-repetive values are laid
flat within a chosen backend but can be treated like a set.


## Backends
Tokens currently has 2 backends build for storing such sets:

- MongoDB ([MGOTokenset](./tokenset/mongo.go)]
- Sql ([SQLTokenset](./tokenset/sql.go))
- InMemory (Planned)
