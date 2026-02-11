# Tampering with Apple Classroom remote management software by leveraging process ownership vulnerability
## About studentd_killer.go
This is a small, utility for macOS that **repeatedly terminates the `studentd` daemon**.

`studentd` is the background process Apple uses to implement **Classroom.app remote management / supervision / screen viewing / app locking / remote restriction** features when a Mac is enrolled in an Apple School Manager MDM profile with Classroom or Apple Schoolwork integration enabled.
### How it works
Running `studentd_killer` causes the program to search for a process called `studentd`, take its `PID`, and terminate it if a process matching the query is found. This process is repeated at intervals of 1ms.
## Installation
Open a terminal and run the following:
```bash
git clone https://github.com/imunoka/studentd_killer.git
```
## How to run
From within the repository, run the following:
```bash
chmod +x ./studentd_killer
./studentd_killer
```
