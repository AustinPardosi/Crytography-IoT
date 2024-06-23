#include <Arduino.h>
#include <WiFi.h>
#include <TinyCrypt.h>

// Dummy public key untuk ECC encryption
uint8_t publicKey[64] = { 0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07,
                          0x08, 0x09, 0x0A, 0x0B, 0x0C, 0x0D, 0x0E, 0x0F,
                          0x10, 0x11, 0x12, 0x13, 0x14, 0x15, 0x16, 0x17,
                          0x18, 0x19, 0x1A, 0x1B, 0x1C, 0x1D, 0x1E, 0x1F,
                          0x20, 0x21, 0x22, 0x23, 0x24, 0x25, 0x26, 0x27,
                          0x28, 0x29, 0x2A, 0x2B, 0x2C, 0x2D, 0x2E, 0x2F,
                          0x30, 0x31, 0x32, 0x33, 0x34, 0x35, 0x36, 0x37,
                          0x38, 0x39, 0x3A, 0x3B, 0x3C, 0x3D, 0x3E, 0x3F };

// Fungsi untuk mengenkripsi data menggunakan ECC
String encryptDataECC(String data, uint8_t* publicKey) {
    uint8_t encrypted[128];
    tc_uECC_encrypt((const uint8_t *)data.c_str(), data.length(), publicKey, encrypted);
    return String((char*)encrypted);
}

// Fungsi untuk mengirim data ke database
void sendDataToDatabase(String encryptedData) {
    WiFiClient client;
    const char* server = "dummy.com";

    if (client.connect(server, 80)) {
        client.println("POST /data HTTP/1.1");
        client.println("Host: dummy.com");
        client.println("Content-Type: application/x-www-form-urlencoded");
        client.print("Content-Length: ");
        client.println(encryptedData.length());
        client.println();
        client.println(encryptedData);
    }
}

void setup() {
    // Setup WiFi dan ECC
    WiFi.begin("SSID", "PASSWORD");
    while (WiFi.status() != WL_CONNECTED) {
        delay(1000);
        Serial.println("Connecting to WiFi...");
    }
    Serial.println("Connected to WiFi");

    String data = "95.4";

    // Enkripsi data menggunakan ECC
    String encryptedData = encryptDataECC(data, publicKey);

    // Kirim data terenkripsi ke database
    sendDataToDatabase(encryptedData);
}

void loop() {
    // Kode loop
}
