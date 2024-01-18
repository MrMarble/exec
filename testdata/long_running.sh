#!/bin/bash

echo "outputA"

# Spawn a background process
sleep 3 &

echo "outputB"
exit 0
```