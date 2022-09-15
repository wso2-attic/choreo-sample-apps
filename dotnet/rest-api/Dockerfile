#Build Stage

FROM mcr.microsoft.com/dotnet/sdk:6.0-focal AS build
WORKDIR /source
COPY . .
RUN dotnet restore
RUN dotnet publish -c release -o /app --no-restore

#Serve Stage
FROM mcr.microsoft.com/dotnet/aspnet:6.0-focal
ENV ASPNETCORE_URLS=http://+:5000
WORKDIR /app
COPY --from=build /app ./
EXPOSE 5000

ENTRYPOINT ["dotnet", "byoc-api.dll"]
