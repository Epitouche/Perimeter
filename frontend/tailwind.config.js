/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
    "./components/**/*.{vue,js,ts}",
    "./layouts/**/*.vue",
    "./pages/**/*.vue",
    "./plugins/**/*.{js,ts}",
    "./composables/**/*.{js,ts,mjs}",
    "./utils/**/*.{js,ts,mjs}",
    "./{A,a}pp.{vue,js,jsx,mjs,ts,tsx}",
    "./{E,e}rror.{vue,js,jsx,mjs,ts,tsx}",
    "./nuxt.config.{js,ts}",
    "./app.config.{js,ts,mjs}",
  ],
  theme: {
    extend: {
      colors: {
        custom_color: {
          bg_page: "#F4F4F4", // Light grey
          bg_section: "#FFFFFF", // White
          text: "#000000", // Black
          text_link: "#4187FF", // Blue
          text_other: "#878787", // Dark grey
          border: "#000000", // Black
          input: "#F4F4F4", // Light grey
          google: "#E60000", // Red
          spotify: "#1DC000", // Green
          dropbox: "#001DDA", // Blue
          timer: "#BB00FF", // Purple
          weather: "#946500", // Brown
        },
      },
      fontSize: {
        custom_size_title: '7rem'
      },
      fontWeight: {
        custom_weight_title: '800',
        custom_weight_connection_title: '600'
      },
      borderRadius: {
        custom_border_radius: '3.125rem'
      },
      borderWidth: {
        custom_border_width: '0.125rem'
      },
    },
  },
  plugins: [],
};
