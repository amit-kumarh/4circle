
# File to put python function to print to serial port
# Need to debug since I never checked if this works with arduino
# Should only WRITE to serial port and arduino will only READ to serial port

import serial
import time

    
def to_serial(col):
    
    ser = serial.Serial('/dev/ttyACM1', 250000, timeout=1)
    time.sleep(2)
    #ser.reset_input_buffer()

#             col = input()
    #time.sleep(5)
    col = int(col)
    print(col)
    ser.write(str(col).encode('utf-8'))
    ser.flush()
  #  time.sleep(5)
        #break
            
    return col

