FROM dpage/pgadmin4:latest AS pgadmin

USER root

RUN mkdir -p /var/lib/pgadmin

COPY ./pgadmin/pgadmin4.db /var/lib/pgadmin/pgadmin4.db

COPY ./pgadmin/config_local.py /pgadmin4/config_local.py

RUN chown -R 5050:5050 /var/lib/pgadmin

USER pgadmin
