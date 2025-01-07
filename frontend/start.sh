#!/usr/bin/env sh

# Build the frontend
npm run build

# Start the server
node ./.output/server/index.mjs