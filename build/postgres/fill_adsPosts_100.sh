#!/bin/bash

count=101
while [ "$count" -le 199 ]; do

  title="cat__"$count
  printf "\nLoad post with title:${title}\n"
  curl -X POST -H "Content-Type: application/json" -d \
  '{ "title":"'$title'", "description":"some_description", "price":'100500', "photos":["link_1"] }' "http://127.0.0.1:80/api/v1/adsPost"; \

  count=$(( count + 1 ))
done