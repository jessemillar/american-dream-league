#!/usr/bin/env bash

echo "Stopping running application"
ssh $DEPLOY_USERNAME@$DEPLOY_HOST 'docker stop american-dream-league'
ssh $DEPLOY_USERNAME@$DEPLOY_HOST 'docker rm american-dream-league'

echo "Pulling latest version"
ssh $DEPLOY_USERNAME@$DEPLOY_HOST 'docker pull jessemillar/american-dream-league:latest'

echo "Starting the new version"
ssh $DEPLOY_USERNAME@$DEPLOY_HOST 'docker run -d -e DATABASE_USERNAME="'$DATABASE_USERNAME'" -e DATABASE_PASSWORD="'$DATABASE_PASSWORD'" -e DATABASE_HOST="'$DATABASE_HOST'" -e DATABASE_PORT="'$DATABASE_PORT'" -e DATABASE_NAME="'$DATABASE_NAME'" --restart=always --name american-dream-league -p 9999:9999 jessemillar/american-dream-league:latest'

echo "Success!"

exit 0
