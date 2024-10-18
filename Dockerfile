FROM dpage/pgadmin4:latest as pgadmin

USER root

RUN mkdir -p /var/lib/pgadmin

RUN chown -R 5050:5050 /var/lib/pgadmin

USER pgadmin