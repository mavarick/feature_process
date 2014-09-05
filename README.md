Feature selection by go code
===
Two Parts:
1. generate splitting points by called IterEqualHeight method.
1. use splitting points to converse the original features to new combined feature, which could be used in large feature processing when using linear regressiong or classfication

==== 1, generate splitting points
see help info

==== 2, generate new features
given the feature with format `value f1:v1 f2:v2 ...`, return the split result, with processing:
1. split the feature's value
1. add missing value 
1. extend the feature/value to new feature of name `feature_value` and value 1
1. add bias value

then the result could be used in logistic regression with github.com/xlhector/hector

==== USAGE
type <./feature_selection -help> to get usage
