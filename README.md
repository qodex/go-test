# go-test
simple go REST API service

to build run 'go build'

The code uses in-memory storage by default but there is a dynamdb based dao implemetation.
To switch to dynamodb replace articlesDao = new(articledao.ArticleDAOInMem) with articlesDao = new(articledao.ArticleDAODynamoDB) at line 56.

Using dynamodb requires working aws credentials and a table named nine-test-dev

