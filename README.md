# deaddrop-go

A deaddrop utility written in Go. Put files in a database behind a password to be retrieved at a later date.

This is a part of the University of Wyoming's Secure Software Design Course (Spring 2023). This is the base repository to be forked and updated for various assignments. Alternative language versions are available in:
- [Javascript](https://github.com/andey-robins/deaddrop-js)
- [Rust](https://github.com/andey-robins/deaddrop-rs)

## Versioning

`deaddrop-go` is built with:
- go version go1.19.4 linux/amd64

## Usage

`go run main.go --help` for instructions

Then run `go run main.go -new -user <username here>` and you will be prompted to create the initial password.

## Database

Data gets stored into the local database file dd.db. This file will not by synched to git repos. Delete this file if you don't set up a user properly on the first go

## Logging Strategy

I added a separate package that keeps track of logging. The logging function is pretty simple. It takes username and messages as parameters. And append a log into "logs.txt" detailing the time, user that performed the action, and the result of their action. For example: "2023-02-12 22:29:33.889803 -0700 MST m=+5.916782449	long1	Succesfully received a message". Then, for every important outcome that we want to keep track, we just use that function there with a hard-coded message. Since the scope of this software is small, I figured hard-coded messages were enough. If the software were on a larger scale, we probably can create a database table with messages and maybe an id for each of them, and only reference ids in the code. 

## Mitigation

I worked on this assignment with Morgan Sinclaire since we already started last part together. We each came up with 2 approaches:

- His approach was to programmatically password-protect the database file.
- I personally found his approach to be not sufficiently secured so I wanted to just encrypt/decrypt the messages. Now, even if the attackers have access to the database, they would not be able to straight out read messages. However, since it's encryption/decryption, the key to decrypt is stored in the source code. This creates a secondary risk but from my experience working as a software developer, its much harder to attack the source code than database.

We ended up going with my plan because after hours of trying to follow his approach, it did not seem to work. To achieve this, i also changed message.go, from retrieving array of strings to array of byte slices to feed directly into the decrypt function. The tutorial we referenced to write these encrypt/decrypt functions is below. We ended up having to change the decrypt function a bit since hex.DecodeString could only take hexadecimal strings and I dont know how to convert raw data to string and then to it. This forces us to change how messages are retrieved from the database as mentioned above as well. Reference:
https://www.melvinvivas.com/how-to-encrypt-and-decrypt-data-using-aes
