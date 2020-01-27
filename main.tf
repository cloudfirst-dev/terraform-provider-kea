provider "kea" {
    url     = "http://192.168.16.2:8080"
}

resource "kea_reservation" "my-server" {
    ip_address = "192.0.5.4"
    hw_address = "5a:00:2d:48:43:05"
    subnet_id = 2
}