# Eye drop prescription

Generating an eye drop schedule from medical prescription info.

## Input

A JSON file containing a key-value list of medical prescriptions.

```
{
    "medication-brand-name": {
        "interval": integer,
        "interval_size": string,
        "interval_change": integer,
        "interval_mod": integer,
        "type": string,
        "quantity": integer,
        "duration": integer,
        "duration_unit": string,
        "first_medication": timestamp
    },
    ...
}
```

Only `interval_change` and `interval_mod` are optionals, other values must be correctly filled in.

| Field Name | Value | Description |
|-|-|-|
|`interval`| Positive integer | Duration units of `interval_size` between eye drops |
|`interval_size`| `hour` | Time duration for `interval` |
|`interval_change`| Either negative or positive integer | When set defines the delta change on interval regarding to `interval_mod` value |
|`interval_mod`| Positive integer | Describes the day frame between interval change updates|
|`type`| `eye drop` | Medication type |
|`quantity`| Positive integer | Number of applications |
|`duration`| Positive integer | Total `duration_unit` units  of treatment |
|`duration_unit` | `day` | Duration unit |
|`first_medication`| Timestamp | Date and time of the first medication |



