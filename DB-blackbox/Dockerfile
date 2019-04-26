FROM postgres:9.6

RUN apt-get update \
 && apt-get install -y --no-install-recommends \
    python3 postgresql-plpython3-9.6
COPY . /docker-entrypoint-initdb.d/ 

# Add VOLUMEs to allow backup of config, logs and databases
VOLUME  ["/etc/postgresql", "/var/log/postgresql", "/var/lib/postgresql"]
