# Configuration Dictionary

| Data Type | Configuration JSON                                                     | Description                                                      | Additional Parameters                                                                                                                         | Required Parameters | Default Values                       |
|-----------|------------------------------------------------------------------------|------------------------------------------------------------------|-----------------------------------------------------------------------------------------------------------------------------------------------|---------------------|--------------------------------------|
| string    | `"ColumnName": { "type": "string", "options": ["opt1", "opt2", ...] }` | Generates random string values from provided options.            | -                                                                                                                                             | -                   | -                                    |
| int       | `"ColumnName": { "type": "int", "min": minVal, "max": maxVal }`        | Generates random integer values between min and max (inclusive). | `min`: Minimum value for generated integers. <br> `max`: Maximum value for generated integers.                                                | `min`, `max`        | -                                    |
| float     | `"ColumnName": { "type": "float", "min": minVal, ... }`                | Generates random floating-point values within min/max range.     | `min`: Minimum value for generated floats. <br> `max`: Maximum value for generated floats. <br> `precision`: Decimal places for float values. | `min`, `max`        | `precision` = 2                      |
| bool      | `"ColumnName": "bool"`                                                 | Generates random boolean values (true or false).                 | -                                                                                                                                             | -                   | -                                    |
| timestamp | `"ColumnName": { "type": "timestamp", ... }`                           | Generates timestamps in specified format and timezone.           | `format`: Format of the timestamp. <br> `zone`: Timezone for the timestamp.                                                                   | `format`            | `zone` = "Local"                     |
| epoch     | `"ColumnName": { "type": "epoch", ... }`                               | Generates epoch timestamps in specified unit and timezone.       | `unit`: Time unit (e.g., "millis", "seconds"). <br> `zone`: Timezone for the timestamp.                                                       | -                   | `unit` = "seconds", `zone` = "Local" |
| isId      | `"ColumnName": { "type": "string", "isId": true }`                     | Generates a unique identifier for each row.                      | -                                                                                                                                             | -                   | -                                    |

Sample conf file looks like [this](data.json)

### Field Generators

The Fake Data Producer leverages field generators from the gofakeit library to generate a diverse range of realistic data. These field generators are specialized functions designed to produce random values for specific data types or fields. By utilizing these generators, you can ensure that the generated data closely resembles real-world scenarios and use cases.

#### How Field Generators Work

Field generators operate by applying various algorithms and patterns to produce data that aligns with the selected data type. For example, the gofakeit library includes field generators for names, addresses, numbers, dates, and more. When you define a column in the configuration file, the Fake Data Producer uses the appropriate field generator based on the specified data type to generate the corresponding data.

#### Benefits of Field Generators

Using field generators offers several benefits:

1. **Realism:** Field generators produce data that mimics real-world scenarios, enhancing the authenticity of the generated content.

2. **Diversity:** Field generators ensure diversity in the generated data by introducing variability within defined parameters.

3. **Efficiency:** Field generators eliminate the need for manual data entry, saving time and effort.

#### Using Field Generators in Configuration

To leverage field generators in your configuration, specify the data type for each column in the JSON configuration file. For example, by setting the data type to "string," "int," "float," or other supported types, you automatically enable the relevant field generator for that column.
<br> Supported field generators are [here](../pkg/generators/fieldgenerators.go)