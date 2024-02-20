# Hotelscom Scraper API

[![Oxylabs promo code](https://user-images.githubusercontent.com/129506779/250792357-8289e25e-9c36-4dc0-a5e2-2706db797bb5.png)](https://oxylabs.go2cloud.org/aff_c?offer_id=7&aff_id=877&url_id=112)

Oxylabs' [Hotelscom Scraper](https://oxylabs.io/products/scraper-api/web/hotelscom?utm_source=github&utm_medium=repositories&utm_campaign=product) is a data gathering solution allowing you to extract real-time information from an Hotelscom website effortlessly. This brief guide showcases how Hotelscom Scraper works, along with code examples to help you better understand how to use it hassle-free.

### How it works

You can get Hotelscom results by providing your own URLs to our service. We can return the HTML for any page you like.

#### Python code example

The example below illustrates how you can get HTML of Hotelscom page.

```python
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
```
Find code examples for other programming languages [**here**](https://github.com/oxylabs/hotelscom-scraper/tree/main/code%20examples)

#### Output Example
```json
{
  "results": [
    {
      "content": "<!DOCTYPE html><html data-language=\"en_US\" dir=\"ltr\" lang=\"en\"><head><title></title><meta name=\"robo ... </html>",
      "created_at": "2024-02-20 12:33:34",
      "updated_at": "2024-02-20 12:33:41",
      "page": 1,
      "url": "https://www.hotels.com/hotel-search?adults=2&children=&d1=2023-12-01&d2=2023-12-05&destination=monaco%2c%20monaco&enddate=2024-02-24&latlong=43.731142%2c7.419758&mapbounds=&pwadialog=&regionid=6050684&rooms=1&semdtl=&sort=recommended&startdate=2024-02-20&theme=&userewards=false&userintent=",
      "job_id": "7165684908583591937",
      "status_code": 200
    }
  ]
}
```
With our Hotelscom Scraper, you can painlessly gather valuable public data from any Hotelscom webpage. Procure essential hotel information like room prices, guest reviews, or amenities details to study hospitality trends and outperform your business rivals. If you need assistance or have any inquiries, feel free to reach our support team through live chat or email us at hello@oxylabs.io.