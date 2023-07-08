# Go Authentication
Generating random authentication tokens for your users is always a difficult task, you could just generate random strings but what use is that? you might also accidently create duplicate authentication tokens, Go Authentication solves these problems with my own implementation of Discord Tokens. 

## Token Structure:
Authentication tokens contain multiple bits of data together so even if the authentication string is found the account can't actually be accessed unless the person also gets access to the epoc and User ID, This makes this method much more secure.

![App Screenshot](https://raw.githubusercontent.com/NotKatsu/UserToken/main/assets/structure.png)

Full Stops - Full stops are used to seperate bits of information so they can easily be taken apart and certian bits of information can be checked.

User ID - The first bit of information we encode is a User ID, In this case I used my SnowFlake generator (available on my GitHub) to create a Custom ID, this ID is then encoded with base64 and added to the first bit of the authentication token.

Seconds Since Epoc - The second bit of data is the amount of seconds since our epoc which is the 6th of January 2023 this helps us get a creation date of the token later if needed.

Random String - The third and final bit of data stored by the token is a 25 character long randomly generated string, this doesn't contain any information but if the type of string you would normally generate for authentication just with the extra encoded data infront of it, yet again this is encoded in base64.