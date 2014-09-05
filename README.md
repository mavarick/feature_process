Feature selection by go code
===
given the feature with format `value f1:v1 f2:v2 ...`, return the split result, with processing:
1, split the feature's value
2, add missing value 
3, extend the feature/value to new feature of name `feature_value` and value 1
4, add bias value

then the result could be used in logistic regression with github.com/xlhector/hector

