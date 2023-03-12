#Build Stage

FROM mcr.microsoft.com/dotnet/sdk:6.0-focal AS build
WORKDIR /source
COPY . .
RUN dotnet restore
RUN dotnet publish -c release -o /app --no-restore
RUN useradd -u 10015 -m choreouser

#Serve Stage
FROM mcr.microsoft.com/dotnet/aspnet:6.0-focal
ENV ASPNETCORE_URLS=http://+:5000
ENV COMPlus_EnableDiagnostics=0
WORKDIR /app
COPY --from=build /app ./
EXPOSE 5000
USER 10015
ENTRYPOINT ["dotnet", "weather-api.dll"]
