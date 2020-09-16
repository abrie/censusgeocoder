# CLI and Go Wrapper for Census.gov Geocoder 

This repository contains a Go module for interfacing with [Census.gov Geocoder](https://geocoding.geo.census.gov/geocoder) REST API. It also includes a CLI utility for interacting with the geocoder.

## CLI

- Build the CLI using Make: `make build`
- Invoke the CLI without parameters for help: `./bin/cli`

Results are written to stdout as JSON.

## API

Two types of results are available: 'locations' or 'geographies'. Both require a 'benchmark' parameter.
Geographies takes two additional parameters: 'vintage' parameters, and an (optionally empty) 'layers' parameter.

### Locations
```go
SearchOneLineAddressLocations(ctx context, onelineaddress, benchmark string)
SearchAddressLocations(ctx context, street, city, state, benchmark string)
SearchByCoordinateLocations(ctx context, x,y float64, benchmark string)
```
### Geographies

Geographies require a 'benchmark', 'vintage' and 'layers' parameter. Leave layers as an empty array for default behaviour.
```go
SearchOneLineAddressGeographies(ctx context, onelineaddress, benchmark, vintage string, layers []string)
SearchAddressGeographies(ctx context, street, city, state, benchmark, vintage string, layers []string)
SearchByCoordinateGeographies(ctx context, x,y float64, benchmark, vintage string, layers[]string)
```

## Technical Details

For more information about the returned fields, see the [official Census.gov documentation](https://geocoding.geo.census.gov/geocoder/Geocoding_Services_API.pdf).
