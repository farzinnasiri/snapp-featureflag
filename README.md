# Feature Flag

## Benchmarks

I benchmarked to programme to see how it behaves under high loads of requests. The only important api that we're
concerned about is then one that gets the `user_id` and `version` and returns the list of active feature flags:

```
http://{url}/feature/get-active-features?user_id={userId}&version={version}'
```

Test Machine:
* 8 core intel 11th gen CPU
* 16GB RAM
* localhost 

Test Configs:
* 100 req/s
* 10s duration of each run

Test Criteria:
* n : number of features
* partial feature with coverage of 5%

### Test Results

**For n = 10**
![10](./images/10.png)
**For n = 50**
![50](./images/50.png)
**For n = 100**
![100](./images/100.png)
**For n = 500**
![500](./images/500.png)
**For n = 1000**
![1000](./images/1000.png)
**For n = 2000**
![2000](./images/2000.png)


#### Plots

**For n <= 500**
![1](./images/scatter_plot%20(2).jpeg)

**For n <= 2000**
![1](./images/scatter_plot%20(1).jpeg)


##### Resources
I used [this](https://github.com/nakabonne/ali) tool to get the benchmarks with the following command:

```shell
ali --rate=100 --duration=10s --method=GET 'http://localhost:8080/feature/get-active-features?user_id=1&version=1.0.0'
```

To add features:

```shell
for n in {1..100}; do curl -X POST -H "Content-Type: application/json"  -d '{"name": "'"$n"'", "coverage": 0.05}' localhost:8080/feature/create; done;
```
