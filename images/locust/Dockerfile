FROM python:slim

COPY pip.conf /root/.pip/pip.conf

# Install locust
RUN pip install locust

COPY ./startup.sh /usr/bin/startup.sh

ENTRYPOINT [ "/usr/bin/startup.sh" ]
