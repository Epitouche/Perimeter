/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
    "./components/**/*.{vue,js,ts}",
    "./layouts/**/*.vue",
    "./pages/**/*.vue",
    "./plugins/**/*.{js,ts}",
    "./nuxt.config.{js,ts}",
  ],
  theme: {
    extend: {
      colors: {
        custom_color_background: {
          DEFAULT: "#F4F4F4", // Light grey
          section: "#FFFFFF", // White
        },
        custom_color_text: {
          DEFAULT: "#000000", // Black
          link: "#4187FF", // Blue
          other: "#878787", // Dark grey
        },
        custom_color_component: {
          DEFAULT: "#FFFFFF", // White
          border: "#000000", // Black
          input: "#F4F4F4", // Light grey
        },
        custom_color_service: {
          DEFAULT: "#979797", // grey
          google: "#E60000", // Red
          spotify: "#1DC000", // Green
          dropbox: "#001DDA", // Blue
          timer: "#BB00FF", // Purple
          weather: "#946500", // Brown
        },
      },
      fontSize: {
        custom_text_size: {
          DEFAULT: '3rem',
          title: '7rem',
        }
      },
      fontWeight: {
        custom_text_weight: {
          DEFAULT: '700',
          title: '900',
        }
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
