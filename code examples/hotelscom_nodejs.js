import fetch from 'node-fetch';

const username = 'YOUR_USERNAME';
const password = 'YOUR_PASSWORD';
const body = {
  'source': 'universal_ecommerce',
  'url': 'https://www.hotels.com/hotel-search?adults=2&children=&d1=2023-12-01&d2=2023-12-05&destination=monaco%2c%20monaco&enddate=2024-02-24&latlong=43.731142%2c7.419758&mapbounds=&pwadialog=&regionid=6050684&rooms=1&semdtl=&sort=recommended&startdate=2024-02-20&theme=&userewards=false&userintent=',
};
const response = await fetch('https://realtime.oxylabs.io/v1/queries', {
  method: 'post',
  body: JSON.stringify(body),
  headers: {
    'Content-Type': 'application/json',
    'Authorization': 'Basic ' + Buffer.from(`${username}:${password}`).toString('base64'),
  }
});

console.log(await response.json());