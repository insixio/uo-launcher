# Insixio UO Launcher Readme

## About

Uses a unique combination of pure htmx for interactivity plus Go templates for creating components and forms, also included:
- Built-in styling using Tailwind and Daisyui.
- Uses HTMX for MPA style interactivity on a single page as per SPA.
- Custom Chi middleware for handling HTMX calls in an easy to maintain routing configuration.
- Built-in version display linked to version variable from main which can be updated on app build for CICD and/or during runtime.
- Scripts configured to use the Bun runtime to launch Vite.
- Also using https://templ.guide/ for components and templates use "templ generate" to update templ files. 

## Initial Setup Instructions
- Install Bun if not already installed
- Run `bun install`
- **For Linux build tag webkit2_40 is required e.g -tags webkit2_40**

## Live Development

To run in live development mode, run `wails dev` in the project directory. In another terminal, go into the `frontend`
directory and run `bun run dev`. The frontend dev server will run on http://localhost:34115. Connect to this in your
browser and connect to your application.

## Building

To build a redistributable, production mode package, use `wails build`.
