cd /secure/certs-cloud

# 1. Create CA files
cockroach cert create-ca --certs-dir=. --ca-key=ca.key

# 2. Create a node certificate with the correct SAN
cockroach cert create-node \
  10.0.2.50 \
  ip-10-0-2-50.eu-central-1.compute.internal \
  localhost \
  127.0.0.1 \
  --certs-dir=. \
  --ca-key=ca.key

# 3. Create a client certificate for hotels_user
cockroach cert create-client hotels_user \
  --certs-dir=. \
  --ca-key=ca.key

# 4. Copy certs to directories /infrastructure/certs/hotels_db/ and /infrastructure/certs/hotels_service/ 

# Set rights:
chmod 600 client.hotels_user.key
chmod 644 ca.crt client.hotels_user.crt node.crt


# Check if the certificates match using a hash:
# sha256sum ca.crt
# sha256sum node.crt
# sha256sum node.key
# sha256sum client.hotels_user.crt
# sha256sum client.hotels_user.key
