using System;
namespace weather_api.Services
{
    public interface IWeatherService
    {
        IEnumerable<WeatherForecast> GetWeather(string[] summaries);
    }
}

