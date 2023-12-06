# #1 Hands On

Let's Prove Caching's Latency Decrease!

- First of all, import the Postman collection script into your Postman, which consists of these endpoints:
  - GET /item --> Fetches datas directly from DB
  - GET /v2/item --> Caching mechanism implemented
- Run the binary using this command: ```./main```
- Run the Postman's collection, and then select the ```Performance``` tab
- We should see the differences of ```Avg. Resp. Time``` for both endpoints, in which the one with caching mechanism is faster