# RepeatToken

## 说明

包括支付常用加签、验签算法。

已实现：微信支付

## 微信支付

### 基准测试

| Method      | Count       | ops |
| --------- | ---------- | -------- |
| BenchmarkDigest_Verify-4        | 200000     | 10156 ns/op   | 
| BenchmarkFastDigest_Verify-4         | 1000000     | 2193 ns/op   | 
| BenchmarkDigest_Generate-4        | 1000000     | 1916 ns/op   | 

## TODO

1. 支付宝签名