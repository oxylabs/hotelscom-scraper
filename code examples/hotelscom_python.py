import requests
from pprint import pprint

# Structure payload.
payload = {
    'source': 'universal_ecommerce',
    'url': 'https://www.hotels.com/hotel-search?adults=2&children=&d1=2023-12-01&d2=2023-12-05&destination=monaco%2c%20monaco&enddate=2024-02-24&latlong=43.731142%2c7.419758&mapbounds=&pwadialog=&regionid=6050684&rooms=1&semdtl=&sort=recommended&startdate=2024-02-20&theme=&userewards=false&userintent='
}

# Get response.
response = requests.request(
    'POST',
    'https://realtime.oxylabs.io/v1/queries',
    auth=('user', 'pass1'),
    json=payload,
)

# Instead of response with job status and results url, this will return the
# JSON response with the result.
pprint(response.json())