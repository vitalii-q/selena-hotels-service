# log in as root user
cd /secure/certs
cockroach sql \ 
  --certs-dir=. \
  --host=localhost \
  --port=9264 \
  --user=root