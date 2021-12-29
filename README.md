# GoUU
Go UUID Shorten & Revert

# Build
  - go build -v -o gouu ./cmd/... 

# Convert UUID To Shorten UUID based on base58.FlickrAlphabet
  - gouu -l 483ca83c-677d-11ec-949d-42010a006010

# Revert Shorten UUID to Oritinal UUID
  - gouu -s 9VnixksFfQgHYdE5ShW2jG

# References
  - github.com/google/uuid
  - github.com/mr-tron/base58

