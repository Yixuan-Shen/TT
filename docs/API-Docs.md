## TT API docs
### Version 0.1

TT API is organized around REST. returns JSON-encoded responses, including errors.

---------------------------------------------

Base URL: http://localhost:10000

---------------------------------------------

#### Get all devices 

**GET** *{Base}*/devices

params:
> none

body:
> none

return:
> all current devices in a json format

status code:
> **200**: success

Example:

```JSON
[
	{
		"ID": "9368582b-fe4f-4543-a66e-8c8c3dc2d982",
		"Name": "Device1",
		"Distance_m": 50
	},
	{
		"ID": "0387670e-4037-413a-9b2d-40c76c95d6ed",
		"Name": "Device2",
		"Distance_m": 30
	},
	{
		"ID": "67f9786e-f354-49f8-bcc7-98076b331650",
		"Name": "self",
		"Distance_m": 0
	}
]
```

---------------------------------------------

#### Get a device by ID

**GET** *{Base}*/devices/{id}

params:
> **id**: string (uuid of the device)

body:
> none

return:
> device in a json format

status code:
> **200**: success
> **404**: device not found

Example:

Success:

```JSON
{
    "ID": "9368582b-fe4f-4543-a66e-8c8c3dc2d982",
    "Name": "Device1",
    "Distance_m": 50
}
```

Error:

```JSON
{
    "ErrorMessage": "device not found"
}
```
---------------------------------------------

#### Add a device

**POST** *{Base}*/device

params:
> none

body:
> device in json format

```JSON
{
    "ID": UUID,
    "Name": Device Name,
    "Distance_m": unsigned int
}
```

return:
> none

status code:
> **200**: success
> **40X**: device already exist
> **40X**: invalid device format

example:

```JSON
{
    "ID": "9368582b-fe4f-4543-a66e-8c8c3dc2d982",
    "Name": "DeviceA",
    "Distance_m": 50
}
```

---------------------------------------------

#### Delete a device

**DELETE** *{Base}*/device/{id}

params:
> **id**: string (uuid of the device)

body:
> none

return:
> none

status code:
> **200**: success
> **404**: device not found

---------------------------------------------

#### Update a device


