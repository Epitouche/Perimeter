# Remix POC
Remix is a fullstack React framework.

## Description
The official Remix website suggests trying to manually set up a Remix project before using the CLI that does it automatically, to truly understand how it works. This is what was done for this POC.

### Manual
[Manual README](./manual/README.md)

### Automatic
[Automatic README](./automatic/README.md)

### Further information
[Remix Website](https://remix.run/)



MANUAL:
- Must create basic ncessary folders and files manually.
- Must mkdir our own project unlike nuxt.
- Must mkdir our own app folder in the project.
- Documentation not as explicit as nuxt.
- Documentation doesnt seem to be up to date.
- Remix will disappear soon.
- Depends on vulnerable versions of node, react, etc.
- On initialization, many basic files were not created (Readme, .gitignore, etc.) like they were for nuxt.
- Obligatory configurations to have for remix projects are not automatically created or filled, instead one must copy paste them from the website.
- Lots of warnings like: The CJS build of Vite's Node API is deprecated.
- Must change stuff in package.json to even run.
- The Remix app is just one file.
- Must manually enable instant feedback, but can only be done if you have installed & configured another server.
- "behavior is changing in React Router v7" warnings for everything.
npm run dev

AUTOMATIC:
- npm warn deprecated <example>: This module is not supported, and leaks memory.
- All basic files & directories created automatically, including Readme & .gitignore.
- .gitignore is quite empty and contains no comments.
- Instant feedback is enabled.
- Comes with vite.
- Comes with tailwindcss.
- No build folder like in the manual one.
npm run dev
