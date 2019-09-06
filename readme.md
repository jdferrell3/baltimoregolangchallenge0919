# Challenge: Consume Maryland's Road Closures

Challenge conceived of by [jboursiquot](https://gist.github.com/jboursiquot/bdda4c1faad9f4b22a3910cdb885b4de)

The state of MD exposes road closure data for you to do with whatever you'd like. Today, you'd like to fetch it and save it locally in a relational database (maybe sqlite?).

Head over to the [site](https://opendata.maryland.gov/Transportation/Maryland-Road-Closures/nigh-m2sg) to explore the data.

Your mission (and success criteria):

1. Consume the JSON endpoint for that data: `https://opendata.maryland.gov/resource/nigh-m2sg.json`
2. Successfully parse each "closure" including its lat/long data, timestamps, etc.
3. Save each closure as a record in a local relational database.
4. Retrieve all of the data from storage to display on the console.

You may produce one (or two) command line tools to do the saving and displaying.

Learning objectives:

- Use `net/http` to retrieve external data
- Use `encoding/json` to unmarshal data
- Use the appropriate interfaces to help with JSON unmarshaling
- Use the `database/sql` package and a driver of your choice to store and retrieve data
- Use the `testing` package to test the retrieval code locally (without actually hitting the JSON endpoint)
