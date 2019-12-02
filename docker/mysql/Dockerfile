FROM mysql
ENV MYSQL_DATABASE snippetreviews
# All scripts in docker-entrypoint-initdb.d/ are automatically
# executed during container startup
COPY ./sql-scripts/ /docker-entrypoint-initdb.d/