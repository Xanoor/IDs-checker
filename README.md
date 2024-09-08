
# 🗒️ID Checker / Scanner

Hi, I created this for the Postbook Hacker101 CTF.

This tool scans all IDs of a web page from a given number to another.

If you have any questions (about how it works or anything else), contact me on Discord -> xanoor1

Check my profile for other cool & useful projects.


## Language
![Go](https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white)
## Demo
Here’s what you need to do:

1. Open the file.
2. Change the values of these variables:

```go
var start int = 0                               
var end int = 100                                 
var errorMsg string = "Post not found"            
var url string = "https://example.com/myPage&id=" 
var cookieValue string = "session=XXXXXXXXXXXXXX" 
```
    ⚙️ start ▶️ The starting ID (default: 0)

    ⚙️ end ▶️ The ending ID (default: 100)

    ⚙️ errorMsg ▶️ The message displayed by the web page when an ID is not found (e.g., "Post not found", optional)

    ⚙️ url ▶️ The URL of the web page

    ⚙️ cookieValue ▶️ If a cookie is required to log in before checking all IDs (optional)

3. Run the script using go run ID_Checker.go and enjoy!
## Support

For support, discord -> xanoor1

