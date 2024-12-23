#!/bin/python

from flask import Flask, render_template
import requests
import os #Import os for environment variables

app = Flask(__name__)

def get_meme():
    # Uncomment these 2 lines and comment out the other URL for a specific subreddit
    # sr = "/wholesomememes"
    # url = "https://meme-api.com/gimme" + sr
    url = "https://meme-api.com/gimme"  # Corrected typo in the URL
    try:
        response = requests.get(url)
        response.raise_for_status()  # Ensure we got a successful response
        response_json = response.json()  # Parse the JSON response
        meme_large = response_json["preview"][-2]  # Get the second-to-last image
        subreddit = response_json["subreddit"]  # Get the subreddit name
        return meme_large, subreddit
    except requests.exceptions.RequestException as e:
        print(f"Error fetching meme: {e}")
        return None, None  # Return None if there's an error

@app.route("/")
def index():
    meme_pic, subreddit = get_meme()
    if meme_pic is None:
        meme_pic = "https://imgflip.com/memegenerator/264254566/Image-Not-Available"  # Placeholder image
        subreddit = "Unknown"
    refresh_url = os.getenv("REFRESH_URL", "/")  # Defaults to "/"
    return render_template("meme_index.html", meme_pic=meme_pic, subreddit=subreddit)

# Start the Flask app with configurable port
if __name__ == "__main__":
    port = int(os.getenv("PORT", 80))  # Default to port 80 if not set
    app.run(host="0.0.0.0", port=port)