# BatchApp (WIP)

This is a WIP local BatchApp — an RMM-style application.

I am creating this to not only learn Windows systems, but also to continue learning Go.

---

##  Tech Stack

The app uses the **[Wails framework](https://wails.io/)**.  
This framework allows you to build the frontend in JavaScript/TypeScript, while writing the backend and core logic in Go.

---

##  Getting Started

To build this on your own:

1. Install dependencies with **npm** in the `frontend` folder
2. Install the Wails CLI: [Installation Docs](https://wails.io/docs/gettingstarted/installation/)

---

##  Commands

- Run the dev build (with terminal output):
  ```bash
  wails dev

- To build into an executable - The executable will be in the src/build/bin folder
  ```bash
  wails build
  
 
## Important note!

This project expects certain PowerShell deployment scripts and software installers (e.g., Sophos, SentinelOne) to be hosted on a local HTTP server.

In my setup:

These files are hosted on a Raspberry Pi running a lightweight Go-based HTTP server

URLs used by the app point to locations like:
http://raspberrypi.local:8080/scripts/... or .../software/...

If you're running this yourself:

You'll need to host your own scripts and .exe files

Update all relevant URLs in the backend logic to match your environment (e.g., use your own IP or hostname)
