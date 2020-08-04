"""
used car "craig's list":
* make
* model
* year
* price
* phone number

"""

tables:
* users
  * phone number: str<16>
* car_makes
  * id: str (implemented uuid)
  * name: str<64>
  * aliases: list<str<64>>
* car_models
  * id: str (implemented uuid)
  * car_make_id: str<43>
  * name: str<64>
  * start_year: int
  * end_year: nullable<int>
* listings
  * id: str/uuid
  * user_id: str<43>
  * car_model_id: str<43>
  * price: int<32>
  * currency:

// phone number of cheapest Ford

SELECT users.phone_number
FROM users
INNER JOIN listings
ON listings.user_id = users.id
INNER JOIN car_models
ON car_models.id = listings.car_model_id
INNER JOIN car_makes
ON car_makes.id = car_models.car_make_id
WHERE car_makes.name = 'Ford'
ORDER BY listings.price ASC
LIMIT 1
;
