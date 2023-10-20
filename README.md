# pwd.go
This a Go cli tool to generate secure passwords. Clone of my own project [pwd.py](https://github.com/lukezielke/pwd.py)
# Features
- create secure passwords
- create passwords in bulk
- output passwords to terminal or to a specified .txt file

# Usage
Create a password with specified length:
```
go run main.go -length <length>
```
Create multiple passwords:
```
go run main.go -amount <amount>
```
Save passwords in a specified .txt file:
```
go run main.go -output <path_to_file>
```