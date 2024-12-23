# Use the official Python base image with Alpine for minimal size
FROM python:3.11-alpine

# Disable Python output buffering
ENV PYTHONUNBUFFERED=1 

# Flask app entry point
ENV FLASK_APP=meme_site.py

# Default port for Flask  
ENV PORT=5000            

# Default refresh URL
ENV REFRESH_URL=/       

# Install system dependencies for Python and Flask
RUN apk add --no-cache gcc musl-dev libffi-dev python3-dev git

# Create a working directory
WORKDIR /app

# Copy requirements first to leverage Docker layer caching
COPY requirements.txt ./

# Install Python dependencies
RUN pip install --no-cache-dir -r requirements.txt

# Copy the rest of the application code into the container
COPY . .

# Command to run the Flask app
CMD ["sh", "-c", "flask run --host=0.0.0.0 --port=$PORT"]
