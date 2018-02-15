analogReadIndex = 0
analogReadValues = [130.12, 113199.13, 939992.3, 57.5]

tempReadIndex = 0
tempReadValues = [[17.2, 55.0], [18.0, 65.3], [16.1, 33.3], [10.7, 47.5]]


def pinMode(port, value):
    pass


def digitalWrite(port, value):
    pass


def analogRead(port):
    global analogReadIndex
    value = analogReadValues[analogReadIndex % analogReadValues.__len__()]
    analogReadIndex += 1
    return value


def dht(port, module_type):
    global tempReadIndex
    [t, h] = tempReadValues[tempReadIndex % tempReadValues.__len__()]
    tempReadIndex += 1
    return [t, h]
