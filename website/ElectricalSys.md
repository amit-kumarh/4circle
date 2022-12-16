# Electrical System:
Our electrical system consisted of: 
..* Stepper motor driver *(include part number)*
..* Stepper motor
..* 360 degree rotating servo
..* 12V AC/DC converter 
..* Raspberry Pi
..* Arduino Uno
..* 50 micro Farad Capacitor

### Multiple Processing Units
We chose to work with both an arduino and a raspberry pi to allow us to have two different opperation systems. Since our main algorithm was written on Go we figured that it would be best to run it on an a raspberry pi. To prevent slowing down our raspberry pi by making it process data/commands for both our stepper motor and servo we chose decided it would be best to have an arduino handle these commands. 

### Powering the systems
For our processing units the main power we used came from the raspberry pi AC/DC voltage convertor. This powered the raspberry pi at 5 volts which then powered the arduino with 5 volts through the usb connection in between the two. Once the arduino was powered it was able to power both the servo and the stepper motor driver. Both our servos and our stepper motor driver only required 5 volts which was well in the range which the arduino could supply. For our stepper motor since it required 12 volts we have to use an AC/DC voltage converter, to power the motor from a wall socket. In order to prevent any major voltage spikes that occur when the stepper motor runs, which can burnout the stepper motor driver, we added a 50 micro farad capacitor.

## To do:
..* add electrical schematic 
..* add a picture of our electrical system *(unsure if we have a full picture)*