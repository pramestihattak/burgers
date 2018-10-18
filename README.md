# burgers!

Penerapan clean architecture dengan depedency injection pada Golang dengan studi kasus membuat aplikasi burger keliling. Aplikasi burger ini merupakan aplikasi POS untuk abang-abang burger keliling yang membutuhkan sebuah aplikasi untuk merekap penjualannya serta merekam jumlah persediaan bumbu-bumbu yang di pakai untuk berjualan burger. Dont take it too seriously :)


# Setup

 1. Buat database dengan nama burgers
 2. Import burgers.sql
 3. Install package-package yang dibutuhkan
	 - Chi router
	 - Mockery
	 - testify
 4. Buat dan sesuaikan pengaturan .env

# Testing

Untuk melakukan testing:

> go test ./... -v

# Run

> go build && ./burgers

# Endpoints

 1. GET | http://localhost:6969/menus
 2. GET | http://localhost:6969/menu/1/receipt
 3. POST | http://localhost:6969/order

# Make an Order

    curl \
	--header "Content-Type: application/json" \
	--request POST \
	--data '{"buyer_name": "Pramesti Hatta K.", "id_discount": "GOPHER", "orders": [{ "id_menu": 1}, {"id_menu": 1}, {"id_menu": 2}]}'\
	http://localhost:6969/order
	
# References
 https://irahardianto.github.io/service-pattern-go/

# License

The MIT License
