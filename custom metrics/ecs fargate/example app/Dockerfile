#!/bin/sh

# Use Python image
FROM python:3.9-slim

# Set working directory
WORKDIR /app

# Copy app and requirements
COPY app.py .
COPY requirements.txt .

# Install dependencies
RUN pip install --no-cache-dir -r requirements.txt

# Expose port for statsd (optional)
EXPOSE 8125

# Run the app
CMD ["python", "app.py"]
