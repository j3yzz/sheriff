#!/bin/bash

grep -E '^database:' -A 1000 config.yml | grep -E '^\s{2}\w+:' | while read -r line; do
  key=$(echo "$line" | cut -d ':' -f 1 | tr '[:lower:]' '[:upper:]' | xargs)
  value=$(echo "$line" | cut -d ':' -f 2- | xargs)
  echo "$key=$value" >> .env
done
