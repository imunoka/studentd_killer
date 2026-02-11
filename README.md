# Tampering with Apple Classroom remote management software by leveraging process ownership vulnerability
## About studentd_killer.go
This is a small, utility for macOS that **repeatedly terminates the `studentd` daemon**.

`studentd` is the background process Apple uses to implement **Classroom.app remote management / supervision / screen viewing / app locking / remote restriction** features when a Mac is enrolled in an Apple School Manager MDM profile with Classroom or Apple Schoolwork integration enabled.
### How it works
Running `studentd_killer` causes the program to search for a process called `studentd`, take its `PID`, and terminate it if a process matching the query is found. This process is repeated at intervals of 1ms.
## What this vulnerability provides
Killing the `studentd` daemon allows users to use bluetooth and wifi functionality without being monitored. While turning off bluetooth can often be enough to prevent remote system access, it effectively terminates its own functionality, meaning that some external devices like mice, keyboards, and headphones that rely on wireless technology won't be able to connect to the user's computer.
## How this vulnerability could be fixed
Because this issue stems from the way Apple software is designed, it could likely be circumvented by launching the `systemd` process as `root` and not as the user in question (the student, in most cases), such that a regular user of the system would be stripped of the ability to kill the `systemd` process due to lack of privilages. However, promoting `systemd` to a root process could potentially introduce new security risks. At the end of the day, this might not be something that school admins have the ability to fix without changing their remote viewing strategy.
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
