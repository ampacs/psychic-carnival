FROM python:3.8-alpine

WORKDIR /app

# Copy and install dependencies
COPY requirements.txt ./
RUN pip install -r requirements.txt

# Copy app
COPY *.py ./

CMD ["python", "feed.py"]
