CREATE DATABASE blog;

CREATE USER 'exporter1'@'%' IDENTIFIED BY 'password' WITH MAX_USER_CONNECTIONS 3;

GRANT PROCESS, REPLICATION CLIENT ON *.* TO 'exporter1'@'%';

GRANT SELECT ON performance_schema.* TO 'exporter1'@'%';