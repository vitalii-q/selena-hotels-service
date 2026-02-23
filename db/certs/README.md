# Generating a client certificate for the hotels_user user
cd /hotels-service/secure/certs/
cockroach cert create-client hotels_user --certs-dir=. --ca-key=ca.key
# It will make two certification files:
# client.hotels_user.crt AND client.hotels_user.key
# client.root.crt AND client.root.key (for making db - hotels_db and user - hotels_user in /_docker/entrypoint.sh)


# Microservice (client) /db/certs/ must have files - client certifications:
# client.hotels_user.crt
# client.hotels_user.key
# ca.crt

# DB (server) /secure/certs must have files:
# ca.crt
# ca.key         # Main 
# node.crt
# node.key


# Manual DB check
cd db/certs/
cockroach sql \
  --certs-dir=. \
  --host=localhost \
  --port=26258 \
  --user=hotels_user \
  --database=hotels_db