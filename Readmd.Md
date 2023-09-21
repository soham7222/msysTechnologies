# Coding Challenge

## To run tests locally

```go
go mod vendor
go test ./...
```

## Costing Problem Complexity calculation
- Calculating numHouses: This part takes constant time.
- Initializing electricityCosts array: This part takes O(numHouses) time because it iterates over numHouses and performs constant-time operations.
- Setting electricity cost for solar panel houses: This part takes O(len(solarPanelHouses)) time, which is the number of solar panel houses.
- Updating electricity costs based on connections: This part takes O(len(connections)) time, where len(connections) is the number of connections.

Therefore, the **overall time complexity** of the given code is **O(numHouses + len(solarPanelHouses) + len(connections))**. In the **worst-case scenario**, when the number of houses is determined by the larger of (len(connections) + 1) and len(solarPanelHouses), the time complexity simplifies to **O(max(len(connections), len(solarPanelHouses)))**. So, the order of complexity for this code is **O(max(len(connections), len(solarPanelHouses)))** which eventually is **O(n)**.

## Triplet Problem Complexity calculation
- Initialization and Preprocessing: `n := len(arr)` and `(pairSet := make(map[int]bool))` both take **O(1)** time.
- Nested Loops: Outer loop runs `n-1` times and the inner loop runs roughly `(n-i)` times in each iteration of the outer loop. It takes **O(n^2)**
- Operations Inside the Loops: Calculating the sum, complement and looking up the map takes **O(1)** time. 
- Inserting into the map: This also takes **O(1)**.

So, the overall time complexity is dominated by the nested loops, giving us a time complexity of **O(n^2)** for the findTriplet function.


