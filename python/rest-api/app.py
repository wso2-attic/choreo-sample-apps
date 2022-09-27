import os
import random
from flask import Flask
from datetime import datetime, date, timedelta


app = Flask(__name__)

summaries = ["Freezing", "Bracing", "Chilly", "Cool", "Mild", "Warm", "Balmy", "Hot", "Sweltering"]
location = [ "Colombo", "Galle", "Kandy", "Jaffna", "Anuradhapura", "Trincomalee", "Negombo", "Ratnapura", "Matale"]

@app.route('/')
def getWeatherApi():
    return {"app":"weather-api", "time":datetime.now()};

@app.route('/healthz')
def healthz():
    return {"app":"weather-api", "time":datetime.now(), "status":"running"};

@app.route('/weather')
def getWeatherForcast():
    weatherResults = []
    current_date = datetime.today()

    for idx, s in enumerate(summaries):
        tempC = random.randint(-20, 55)
        print(idx)
        weatherResults.append({
            "date": current_date + timedelta(idx),
            "temperatureC":tempC,
            "temperatureF":32 + (int)(tempC / 0.5556),
            "summary":s,
            "location":location[idx]
        })
    
    return weatherResults;

if __name__ == "__main__":
    port = int(os.environ.get('PORT', 5000))
    app.run(debug=True, host='0.0.0.0', port=port)
