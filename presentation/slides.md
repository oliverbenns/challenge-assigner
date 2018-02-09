# Can we automate and improve our coding tests for possible hires?

---
# Current Situation

- Candidate goes through initial screening
- Candidate is tracked through Greenhouse *- new!*
- Zip file is emailed by HR to the candidate.
  - Includes supporting files + brief
- On completion, developer will zip up project and upload to Greenhouse. *- new!*
- When notified, someone...
  - Downloads the zip file
  - Extracts to a certain directory in our challenge repository
  - Creates a pull request
  - Assigns a developer to review


---
# Problems

- Manual labour
- No idea of how the code has progressed over time.
  - Thought process / learnings / how they started
- No idea on their git experience outside of code
  - Do they commit regularly?
  - Do they add valuable commit messages?
  - Are their pull requests well laid out and explained?
- No idea of how much time the developer has actually spent on the work
  - Brief recommends 4-5 hours
  - Ask to complete within a week.
  - Ends up being evident in the follow up pair-programming session.

---
# Why?

- Been through it recently
- Inspired by another company
- Small project but with new experiences
  - Go Lang

---
# What I Built

#### Small HR admin dashboard
- List assignees
- Create assignees
  - Creates new **private** github repository
  - Assign user to repository
  - Add an introduction readme to the repo.

#### Webpage for a developer to 'begin' the timer
- Adds the official brief to the repository
- Any commits after a certain period will not be considered.

---
# Showtime!
