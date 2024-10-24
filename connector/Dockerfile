FROM python

WORKDIR /
COPY connector/requirements.txt connector/connectorMQTT2DB.py /app/
RUN pip install -r /app/requirements.txt

CMD [ "python", "/app/connectorMQTT2DB.py" ]
