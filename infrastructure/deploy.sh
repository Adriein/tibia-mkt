#!/usr/bin/env bash

# Check for required environment variables
if [[ -z "$DATABASE_NAME" || -z "$DATABASE_PASSWORD" || -z "$DATABASE_USER" || -z "$SERVER_PORT" || -z "$TIBIA_MKT_API_KEY" ]]; then
  # Prompt user to enter missing environment variables
  echo "Some environment variables are missing. Please enter the following:"
  read -p "ENVIRONMENT: " ENVIRONMENT
  read -p "DATABASE_NAME: " DATABASE_NAME
  read -p "DATABASE_PASSWORD: " DATABASE_PASSWORD
  read -p "DATABASE_USER: " DATABASE_USER
  read -p "SERVER_PORT: " SERVER_PORT
  read -p "TIBIA_MKT_API_KEY: " TIBIA_MKT_API_KEY

  echo "Deleting actual .env"
  rm -rf .env

  echo "Creating new .env"
  echo "ENVIRONMENT=$ENVIRONMENT" >> .env
  echo "DATABASE_NAME=$DATABASE_NAME" >> .env
  echo "DATABASE_PASSWORD=$DATABASE_PASSWORD" >> .env
  echo "DATABASE_USER=$DATABASE_USER" >> .env
  echo "SERVER_PORT=$SERVER_PORT" >> .env
  echo "TIBIA_MKT_API_KEY=$TIBIA_MKT_API_KEY" >> .env
fi