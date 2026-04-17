go packages & modules:
- packages are directories; collections of go source files compiled together
- modules are collections of related packages that are used together
- repos usually contain one module, define by the go.mod file in the root

go tests:
- just regular functions with some rules
- import the package being tested
- file is in same directory 
- name of test function must start with 'Test'
- Takes one arg of type *testing.T
- table driven testing:
    - Setup struct with input and expected output types (the table)
    - Add in lines of test code as needed connected to map literals
    - iterate through each test struct to see if 'want' and 'got' are equal
    - use 'go-cmp' library with cmp.Diff function 
    - print names from map literals to locate test easier
    - use subtests to keep testing running after a failure
    - using '%#v' makes spotting string errors easier

Using the poke api
-get json data and organize in Jlint
-create struct from data with json-to-go