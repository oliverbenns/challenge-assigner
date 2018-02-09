# Go Hack

- `/admin`
  - loads a html page with a little javascript that sends a request to `/api/challenge/create` on button click
  - h1: Go Hack Admin
  - Input: Github username
  - Button: Send challenge

- `/start` - loads a html page with a little javascript that sends a request to `/api/challenge/start` on button click
  - Paragraph: You will have X time to complete this challenge. Commits after this period will not be counted towards assessment
  - This will add the brief to the repository under `BRIEF.md`

- POST `/api/challenge/create`
  - Creates new github repository `firstname-lastname-timestamp`
  - Assign user to repository
  - Add initial readme.

- POST `/api/challenge/start`
  - Sends a commit to `firstname-lastname-timestamp` repository with `BRIEF.md`
