using System;

namespace weather_api.Services
{
    public class WeatherService : IWeatherService
    {
        private static readonly string[] locations = new[]
   {
        "Colombo", "Galle", "Kandy", "Jaffna", "Anuradhapura", "Trincomalee", "Negombo", "Ratnapura", 
    };

        public IEnumerable<WeatherForecast> GetWeather(string[] summaries)
        {
            var range = new Random();
            return Enumerable.Range(1, 5).Select(index => new WeatherForecast
            {
                Date = DateTime.Now.AddDays(index),
                TemperatureC = range.Next(-20, 55),
                Summary = summaries[range.Next(summaries.Length)],
                Location = locations[range.Next(locations.Length)]
            })
        .ToArray();
        }
    }
}

