#!/usr/bin/env bash

echo "Stopping running application"
ssh $DEPLOY_USERNAME@$DEPLOY_HOST 'docker stop american-dream-league'
ssh $DEPLOY_USERNAME@$DEPLOY_HOST 'docker rm american-dream-league'

echo "Pulling latest version"
ssh $DEPLOY_USERNAME@$DEPLOY_HOST 'docker pull jessemillar/american-dream-league:latest'

echo "Starting the new version"
ssh $DEPLOY_USERNAME@$DEPLOY_HOST 'docker run -d -e AUTH0_CALLBACK="'$AUTH0_CALLBACK'" -e AUTH0_CLIENT_ID="'$AUTH0_CLIENT_ID'" -e AUTH0_CLIENT_SECRET="'$AUTH0_CLIENT_SECRET'" -e DATABASE_USERNAME="'$DATABASE_USERNAME'" -e DATABASE_PASSWORD="'$DATABASE_PASSWORD'" -e DATABASE_HOST="'$DATABASE_HOST'" -e DATABASE_PORT="'$DATABASE_PORT'" -e DATABASE_NAME="'$DATABASE_NAME'" --restart=always --name american-dream-league -p 8000:8000 jessemillar/american-dream-league:latest'

echo "Success!"

exit 0
