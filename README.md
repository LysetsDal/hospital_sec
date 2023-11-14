# hospital_sec

## Running the program

-   open up 4 terminal windows
-   type make proto to create the binaries (if not already made) in one of the windows
-   type make c1 in the first
-   type make c2 in the second
-   type make c3 in the third
-   type make srv in the last

## To run the algorithm, type 'secret'

-   now enter the secret
-   the secret is now split, send to the other peers and the person whom initiated the secret sharing is also the one sending the cumulated data to the hospital.
-   the output of the algorithm is now display on the server. Note that the secret from the other two peers is currently hardcoded to be 20, for demonstration purposes.

## Basics:

-   certs holds all dummy certificates and keys
-   cmd holds the implementation for the server (hospital) and client (Peer2peer network)
-   utils holds the module that calculates and splits the secret and some TLS functions
-   proto holds my service definitions and generated gRPC output files
