# #3 Hands On

Let's Try Cache Patterns!

Add a ```proxy``` layer for the already-created ```repository``` layer (```internal/repository```).

The ```proxy``` layer will decide which source of data should be fetched from, either from Redis or DB.