# Trampoline
A trampoline allows you to 
https://raganwald.com/2013/03/28/trampolines-in-javascript.html  
Have made certain [assumptions](#ui-assumptions) regarding how the user would interact with the (hypothetical) UI during the life cycle of the game.

## Simulation
The project also contains a simulation for the game that can be run in following ways:

#### Using gradle wrapper
From the root directory
1. Run command `./gradlew build` to build
2. Run command `./gradlew run` to run the simulation

#### Directly running the main method
1. Run the main method in `com.rezilience.runner.SimulateGame` (assuming you have built the code already) 

### Simulator O/p
Each line in o/p indicates a move. Some examples follow:

`1: Eric: 3 --> 8` => (sr. no): (player name): (start) --> (end)  
This means, Eric moved from 3 to 8 by scoring a 5 on the spinner 

`7: Eric: 6 --> 9 --LADDER--> 31`  
Eric moved from 6 to 9 scoring a 3, encountered a ladder and jumped to 31.
Likewise, there can be a CHUTE instead of LADDER. 

`61: Paul: 99 -x-> 105 --> 99`  
Paul scored a 6 but 6 moves not possible from 99 and hence st