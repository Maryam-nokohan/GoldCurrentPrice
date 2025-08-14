# Gmail SMTP Email Sender in Go

This  sends emails using Gmail's SMTP server in Go.
It uses a `.env` file to store credentials securely.



## 1. Prerequisites
- A Gmail account
- An App Password (Google no longer allows normal passwords for SMTP)


## 2. How to Get a Gmail App Password
This is how you can get gmail App password for this program. This methode is secure.
### Step 1 — Enable 2-Step Verification
1. Go to your [Google Account Security](https://myaccount.google.com/security) page.
2. Under **"Signing in to Google"**, enable **2-Step Verification**.

### Step 2 — Generate an App Password
1. After enabling 2-Step Verification, go to [App Passwords](https://myaccount.google.com/apppasswords).
2. Sign in again if prompted.
3. Under **Select app**, choose **Mail**.
4. Under **Select device**, choose **Other (Custom name)** and type `Golang`.
5. Click **Generate**.
6. Copy the 16-character password Google shows you (it will look like `abcd efgh ijkl mnop`).

## 3. Add info to .env
```
# Your Gmail address (the sender)
EMAIL_ADDRESS=your-email@gmail.com

# Your Gmail App Password (16 characters)
# How to get this is explained in the README.
EMAIL_KEY="xxxx xxxx xxxx xxxx"

# The recipient's email address
EMAIL_RECEIVER=recipient@example.com
```


