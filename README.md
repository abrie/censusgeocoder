# CLI and Go Wrapper for Census.gov Geocoder 

This repository contains a Go module for interfacing with [Census.gov Geocoder](https://geocoding.geo.census.gov/geocoder) REST API. It also includes a CLI utility for interacting with the geocoder. You may run the CLI via the pre-built Docker image or build it yourself using the Go toolchain.

## Use CLI from Docker Image

- `docker run docker.pkg.github.com/abrie/censusgeocoder/cli:0.0.1`

## Build CLI Locally

- Build the CLI using Make: `make build`
- Invoke the CLI without parameters for help: `./bin/cli`

Results are written to stdout as JSON.

## Use the API programmatically

- `import ("github.com/abrie/censusgeocoder")`
Two types of results are available: 'locations' or 'geographies'. Both require a 'benchmark' parameter.
Geographies takes two additional parameters: 'vintage' parameters, and an (optionally empty) 'layers' parameter.

### Locations
```go
censusgeocoder.SearchOneLineAddressLocations(ctx context, onelineaddress, benchmark string)
censusgeocoder.SearchAddressLocations(ctx context, street, city, state, benchmark string)
censusgeocoder.SearchByCoordinateLocations(ctx context, x,y float64, benchmark string)
```
### Geographies

Geographies require a 'benchmark', 'vintage' and 'layers' parameter. Leave layers as an empty array for default behaviour.
```go
censusgeocoder.SearchOneLineAddressGeographies(ctx context, onelineaddress, benchmark, vintage string, layers []string)
censusgeocoder.SearchAddressGeographies(ctx context, street, city, state, benchmark, vintage string, layers []string)
censusgeocoder.SearchByCoordinateGeographies(ctx context, x,y float64, benchmark, vintage string, layers[]string)
```

### Benchmarks and Vintages

A list of available Benchmarks and/or vintages are available to the API:
```go
censusgeocoder.GetBenchmarks()
censusgeocoder.GetVintages(benchmark string)
```

## Technical Details

For more information about the returned fields, see the [official Census.gov documentation](https://geocoding.geo.census.gov/geocoder/Geocoding_Services_API.pdf).
