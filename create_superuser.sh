#!/bin/bash

# Get CSRF token from Pulsar Manager
CSRF_TOKEN=$(curl -s http://localhost:7750/pulsar-manager/csrf-token)

# Check if CSRF token was successfully retrieved
if [ -z "$CSRF_TOKEN" ]; then
    echo "Failed to retrieve CSRF token."
    exit 1
fi

# Create a superuser using the CSRF token
curl -s \
    -H "X-XSRF-TOKEN: $CSRF_TOKEN" \
    -H "Cookie: XSRF-TOKEN=$CSRF_TOKEN;" \
    -H 'Content-Type: application/json' \
    -X PUT http://localhost:7750/pulsar-manager/users/superuser \
    -d '{"name": "admin", "password": "apachepulsar", "description": "test", "email": "username@test.org"}'

# Check the result of the user creation
if [ $? -eq 0 ]; then
    echo "Superuser 'admin' created successfully."
else
    echo "Failed to create superuser."
    exit 1
fi
