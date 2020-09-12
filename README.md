# Go Wrapper for Census.gov Geocoder API

This Go module wraps the REST API provided by [census.gov geocoder](https://geocoding.geo.census.gov/geocoder).

## Methods

Two types of data can be returned, 'locations' or 'geographies'. Geographies are an optional extension of 'locations' return type.

Locations require a 'benchmark' parameter.

```go
SearchOneLineAddressLocations(ctx context, onelineaddress, benchmark string)
SearchAddressLocations(ctx context, street, city, state, benchmark string)
SearchByCoordinateLocations(ctx context, x,y float64, benchmark string)
```

Geographies require a 'benchmark' and 'vintage' parameter.
```go
SearchOneLineAddressGeographies(ctx context, onelineaddress, benchmark, vintage string)
SearchAddressGeographies(ctx context, street, city, state, benchmark, vintage string)
SearchByCoordinateGeographies(ctx context, x,y float64, benchmark, vintage string)
```

### Technical Details

For more information about the returned fields, see the [official Census.gov documentation](https://geocoding.geo.census.gov/geocoder/Geocoding_Services_API.pdf).
