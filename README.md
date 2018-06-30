# Twitter link shortner bypass server

This server acts like a proxy with t.co shortlinks, permitting you to preserve your privacy.
The URL scheme remains the same for the t.co and this server, so you are able to keep a compatilibity between the Twitter's service and this server.

# THIS IS STILL WORK-IN-PROGRESS

- [x] Twitter's t.co compatible
- [x] Databse caching (uses redis)

TODO:
- [ ] Make Redis optional
- [ ] Make Redis configuration editable in a .toml file
- [ ] Make a reporting system for nocive URLs
- [ ] Make a XSS detector, to avoid risks for the end users
