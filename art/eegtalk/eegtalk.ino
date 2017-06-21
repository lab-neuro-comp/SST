#define LED 13

int inlet;

void setup() {
    pinMode(LED, OUTPUT);
    Serial.begin(9600);
    digitalWrite(LED, LOW);
}

void loop() {
    if (Serial.available() > 0) {
        inlet = Serial.read();
        switch (inlet) {
            case '1':
            digitalWrite(LED, HIGH);
            break;
            case '0':
            digitalWrite(LED, LOW);
            break;
            default:
            Serial.println("What?");
            break;
        }
    }
}
