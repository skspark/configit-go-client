# Requirements
support multiple types (github / gitlab / bitbucket)
support multiple data type (yaml / json)
set multiple directory target with name, merge each files to single yaml
sync to config struct
give functionality to get specific field, with given struct type
give functionality to stop sync
give functionality to add callback on sync failure or success
support go version below 18 (generic is optional)

# Requirements by priority
## PO
- support multiple types (github / gitlab / bitbucket)

## P1
- optimize dependency per each types (github / gitlab / bitbucket)

## NOTES
- GO111Module has to be enabled (github.com/go-git/go-git/v5)