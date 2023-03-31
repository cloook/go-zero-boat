#!/bin/bash

# target project
target=user
table=user

# model generate     
goctl model mysql ddl -src ./${target}/model/${table}.sql -dir ./${target}/model -cache=false  --style=goZero --home ../deploy/goctl/1.5.0 -c 

