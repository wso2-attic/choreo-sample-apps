using Microsoft.AspNetCore.Mvc;
using Microsoft.Extensions.Caching.Memory;
using weather_api.Services;

namespace weather_api.Controllers;

[ApiController]
[Route("weather")]
public class WeatherForecastController : ControllerBase
{
    private static readonly string[] Summaries = new[]
    {
        "Freezing", "Bracing", "Chilly", "Cool", "Mild", "Warm", "Balmy", "Hot", "Sweltering", "Scorching"
    };

    private readonly ILogger<WeatherForecastController> _logger;
    private readonly IMemoryCache _memoryCache;
    private IWeatherService weatherService;

    private readonly string cacheKey = "weatherapiKey";


    public WeatherForecastController(ILogger<WeatherForecastController> logger,
        IMemoryCache memoryCache)
    {
        _logger = logger;
        _memoryCache = memoryCache;
        weatherService = new WeatherService();
    }

    [HttpGet]
    public IEnumerable<WeatherForecast> Get()
    {
        IEnumerable<WeatherForecast>? weatherForecastCollection = null;

        if (_memoryCache.TryGetValue(cacheKey, out weatherForecastCollection))
        {
            return weatherForecastCollection;
        }

        weatherForecastCollection = weatherService.GetWeather(Summaries);

        var cacheOptions = new MemoryCacheEntryOptions()
            .SetSlidingExpiration(TimeSpan.FromSeconds(10))
            .SetAbsoluteExpiration(TimeSpan.FromSeconds(30));

        _memoryCache.Set(cacheKey, weatherForecastCollection, cacheOptions);
        return weatherForecastCollection;
    }

}

