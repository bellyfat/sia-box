# sia-box

## How to use this tool?

1. `go install github.com/jay-dee7/sia-box`
2. `sia-box gen-pwd`
3. `vim $HOME/.sia-box/sia-box.yaml` and enter the path to file, you want to upload
4. `sia-box sync`

After this you'll get a hash, using which you can download the uploaded file.
This file will be encrypted so use:
`sia-box decrypt <downloaded-file`
