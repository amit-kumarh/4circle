import serial
def to_serial(col):
    port = ""
    s = serial.Serial(port)
    s.write(col)
    
to_serial(10)